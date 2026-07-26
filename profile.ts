import { defineProfile, loadTemplate, defaultBuildEndpointStructs } from '@doclient/renderer-go';
import type { IRParam } from '@doclient/cli';

function param(name: string, type: string, shopeeType: string, children: IRParam[] = [], required = false): IRParam {
  return { name, type, shopeeType, description: '', required, children };
}

const getBrandByPagesReq: IRParam[] = [
  param('startRow', 'int64', 'int64', [], true),
  param('pageSize', 'int64', 'int64', [], true),
];

const getCategoryAttributesReq: IRParam[] = [
  param('primary_category_id', 'int64', 'int64', [], true),
  param('language_code', 'string', 'string', [], true),
];

const createProductReq: IRParam[] = [
  param('primary_category_id', 'int64', 'int64', [], true),
  param('attributes', 'string', 'string', [], true),
  param('skus', 'string', 'string', [], true),
  param('name', 'string', 'string', [], true),
  param('description', 'string', 'string', [], true),
  param('short_description', 'string', 'string', [], false),
  param('images', 'string', 'string', [], false),
  param('brand', 'int64', 'int64', [], false),
  param('warranty', 'string', 'string', [], false),
  param('warranty_type', 'string', 'string', [], false),
  param('size_guide', 'string', 'string', [], false),
  param('source', 'string', 'string', [], false),
  param('sale_start_date', 'string', 'string', [], false),
  param('sale_end_date', 'string', 'string', [], false),
  param('package_weight', 'string', 'string', [], false),
  param('package_length', 'string', 'string', [], false),
  param('package_width', 'string', 'string', [], false),
  param('package_height', 'string', 'string', [], false),
];

const updateProductReq: IRParam[] = [
  param('item_id', 'int64', 'int64', [], true),
  param('attributes', 'string', 'string', [], false),
  param('name', 'string', 'string', [], false),
  param('description', 'string', 'string', [], false),
  param('short_description', 'string', 'string', [], false),
];

const manualEndpointTypes: Record<string, { request: IRParam[]; response: IRParam[] }> = {
  GetBrandByPages: { request: getBrandByPagesReq, response: [] },
  GetCategoryAttributes: { request: getCategoryAttributesReq, response: [] },
  CreateProduct: { request: createProductReq, response: [] },
  UpdateProduct: { request: updateProductReq, response: [] },
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
