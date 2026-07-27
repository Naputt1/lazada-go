import { defineProfile, loadTemplate, defaultBuildEndpointStructs } from '@doclient/renderer-go';
import type { IREndpoint, IRParam } from '@doclient/cli';

function param(name: string, type: string, shopeeType: string, children: IRParam[] = [], required = false): IRParam {
  return { name, type, shopeeType, description: '', required, children };
}

const createProductReq: IRParam[] = [
  param('payload', 'string', 'string', [], true),
];

const updateProductReq: IRParam[] = [
  param('payload', 'string', 'string', [], true),
];

const getProductsReq: IRParam[] = [
  param('filter', 'string', 'string', [], false),
  param('created_after', 'string', 'string', [], false),
  param('created_before', 'string', 'string', [], false),
  param('offset', 'int64', 'int64', [], false),
  param('limit', 'int64', 'int64', [], false),
];

const getProductItemReq: IRParam[] = [
  param('item_id', 'int64', 'int64', [], true),
];

const getBrandByPagesReq: IRParam[] = [
  param('startRow', 'int64', 'int64', [], true),
  param('pageSize', 'int64', 'int64', [], true),
];

const getCategoryAttributesReq: IRParam[] = [
  param('primary_category_id', 'int64', 'int64', [], true),
  param('language_code', 'string', 'string', [], true),
];

const manualEndpointTypes: Record<string, { request: IRParam[]; response: IRParam[] }> = {
  GetBrandByPages: { request: getBrandByPagesReq, response: [] },
  GetCategoryAttributes: { request: getCategoryAttributesReq, response: [] },
  CreateProduct: { request: createProductReq, response: [] },
  UpdateProduct: { request: updateProductReq, response: [] },
  GetProducts: { request: getProductsReq, response: [] },
  GetProductItem: { request: getProductItemReq, response: [] },
};

const clientTpl = loadTemplate('./templates/client.go');
const authTpl = loadTemplate('./templates/auth.go');

const profileConfig = {
  responseDataFieldName: 'data' as const,
  commonFields: ['code', 'type', 'message', 'request_id'] as string[],
  commonRequestFields: ['app_key', 'sign_method', 'timestamp', 'partner_id', 'access_token', 'sign'] as string[],
  baseResponseFields: [
    { name: 'Code', type: 'string', jsonTag: 'code', urlTag: '', comment: '' },
    { name: 'Type', type: 'string', jsonTag: 'type', urlTag: '', comment: '' },
    { name: 'Message', type: 'string', jsonTag: 'message', urlTag: '', comment: '' },
    { name: 'RequestID', type: 'string', jsonTag: 'request_id', urlTag: '', comment: '' },
  ],
  name: 'lazada' as const,
};

export const lazadaProfile = defineProfile({
  ...profileConfig,

  buildEndpointStructs: (structGen, moduleName, ep) => {
    const overrides = manualEndpointTypes[ep.name];
    if (overrides) {
      if (overrides.request && overrides.request.length > 0) {
        ep.requestParams = overrides.request;
      }
      if (overrides.response && overrides.response.length > 0) {
        ep.responseParams = overrides.response;
      }
    }

    // Post-process response params to fix known API doc vs reality mismatches
    const fixType = (params: IRParam[]) => {
      for (const p of params) {
        if (p.name === 'marketImages' || p.name === 'MarketImages' || p.name === 'ImageSequence' || p.name === 'imageSequence') {
          p.type = '[]string';
        }
        if (p.type === 'object[]' || p.type === 'object') {
          fixType(p.children);
        }
      }
    };
    fixType(ep.responseParams);

    defaultBuildEndpointStructs(profileConfig as any)(structGen, moduleName, ep);
  },

  renderClientFile: (pkg, services, init) =>
    clientTpl.render({
      PACKAGE_NAME: pkg,
      SERVICES_SECTION: services,
      SERVICES_INIT_SECTION: init,
    }),

  renderAuthFile: (pkg) => authTpl.render({ PACKAGE_NAME: pkg }),

  testSetup: {
    appLiteral: 'App{\n\t\tAppKey:    "test_app_key",\n\t\tAppSecret: "test_app_secret",\n\t}',
    extraSetup: '\tclient.Region = "SG"\n\tclient.Token = "test_access_token"',
  },

  serviceStyle: 'wrapper',

  dependencies: ['github.com/jarcoal/httpmock v1.3.1'],
});
