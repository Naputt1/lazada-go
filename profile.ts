import { defineProfile, loadTemplate, defaultBuildEndpointStructs } from '@doclient/renderer-go';
import type { IRParam } from '@doclient/cli';

function param(name: string, type: string, shopeeType: string, children: IRParam[] = [], required = false): IRParam {
  return { name, type, shopeeType, description: '', required, children };
}

const createProductReq: IRParam[] = [
  param('primary_category_id', 'int64', 'int64'),
  param('attributes', 'string', 'string'),
  param('skus', 'string', 'string'),
  param('name', 'string', 'string'),
  param('description', 'string', 'string'),
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

const createProductResp = [
  param('data', 'object', 'object', [
    param('item_id', 'int64', 'int64'),
    param('sku_list', 'object[]', 'object[]', [
      param('seller_sku', 'string', 'string'),
      param('sku_id', 'int64', 'int64'),
    ]),
  ]),
];

const updateProductReq: IRParam[] = [
  param('item_id', 'int64', 'int64'),
  param('attributes', 'string', 'string', [], false),
  param('name', 'string', 'string', [], false),
  param('description', 'string', 'string', [], false),
  param('short_description', 'string', 'string', [], false),
];

const updateProductResp = [
  param('data', 'object', 'object', [
    param('item_id', 'int64', 'int64'),
  ]),
];

const getProductItemReq: IRParam[] = [
  param('item_id', 'int64', 'int64', [], false),
  param('seller_sku', 'string', 'string', [], false),
];

const getProductItemResp = [
  param('data', 'object', 'object', [
    param('item_id', 'int64', 'int64'),
    param('primary_category', 'int64', 'int64'),
    param('name', 'string', 'string'),
    param('description', 'string', 'string'),
    param('short_description', 'string', 'string', [], false),
    param('images', 'string', 'string', [], false),
    param('attributes', 'string', 'string', [], false),
    param('skus', 'object[]', 'object[]', [
      param('seller_sku', 'string', 'string'),
      param('sku_id', 'int64', 'int64'),
      param('quantity', 'int64', 'int64'),
      param('price', 'float64', 'double'),
      param('package_height', 'string', 'string', [], false),
      param('package_length', 'string', 'string', [], false),
      param('package_width', 'string', 'string', [], false),
      param('package_weight', 'string', 'string', [], false),
    ]),
  ]),
];

const getProductsReq: IRParam[] = [
  param('filter', 'string', 'string', [], false),
  param('limit', 'string', 'string', [], false),
  param('offset', 'string', 'string', [], false),
  param('created_after', 'string', 'string', [], false),
  param('created_before', 'string', 'string', [], false),
  param('update_after', 'string', 'string', [], false),
  param('update_before', 'string', 'string', [], false),
  param('search', 'string', 'string', [], false),
];

const getProductsResp = [
  param('data', 'object', 'object', [
    param('total_products', 'int64', 'int64'),
    param('products', 'object[]', 'object[]', [
      param('item_id', 'int64', 'int64'),
      param('primary_category', 'int64', 'int64'),
      param('name', 'string', 'string'),
      param('seller_sku', 'string', 'string'),
    ]),
  ]),
];

const manualEndpointTypes: Record<string, { request: IRParam[]; response: IRParam[] }> = {
  CreateProduct: { request: createProductReq, response: createProductResp },
  UpdateProduct: { request: updateProductReq, response: updateProductResp },
  GetProductItem: { request: getProductItemReq, response: getProductItemResp },
  GetProducts: { request: getProductsReq, response: getProductsResp },
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
      ep.requestParams = overrides.request;
      ep.responseParams = overrides.response;
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
