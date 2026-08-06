import { defineProfile, loadTemplate, defaultBuildEndpointStructs } from '@doclient/renderer-go';
import type { StructGenerator } from '@doclient/renderer-go';
import { toPascalCase } from '@doclient/cli';
import type { IREndpoint, IRParam } from '@doclient/cli';
import { structTypeOverrides } from './overrides.js';

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

// /orders/get and the order-items endpoints. Lazada's API doc omits the
// request parameters for these, so the SDK's generated methods historically
// took no arguments and could never send order_ids. They are declared here so
// the generator emits typed request structs.
const getOrdersReq: IRParam[] = [
  param('created_after', 'string', 'string', [], true),
  param('created_before', 'string', 'string', [], false),
  param('update_after', 'string', 'string', [], false),
  param('sort_by', 'string', 'string', [], false),
  param('sort_direction', 'string', 'string', [], false),
  param('offset', 'int64', 'int64', [], false),
  param('limit', 'int64', 'int64', [], false),
];

const getOrderItemsReq: IRParam[] = [
  param('order_ids', 'int64[]', 'int64[]', [], true),
];

const getMultipleOrderItemsReq: IRParam[] = [
  param('order_ids', 'int64[]', 'int64[]', [], true),
];

// /order/package/document/get (PrintAWB) wraps its whole request in a single
// getDocumentReq object. The API doc nests the real fields under it, so model
// it as an object param with children; paramsFromStruct then serializes the
// value as the JSON body Lazada requires.
const printAWBReq: IRParam[] = [
  param('getDocumentReq', 'object', 'object', [
    param('doc_type', 'string', 'string', [], true),
    param('packages', 'object[]', 'object[]', [
      param('package_id', 'string', 'string', [], true),
    ], true),
    param('print_item_list', 'bool', 'boolean', [], false),
  ], true),
];

const getReverseOrdersForSellerReq: IRParam[] = [
  param('page_no', 'int64', 'int64', [], true),
  param('page_size', 'int64', 'int64', [], true),
  param('request_type_list', 'string[]', 'string[]', [], false),
  param('ofc_status_list', 'string[]', 'string[]', [], false),
  param('reverse_status_list', 'string[]', 'string[]', [], false),
  param('reverse_order_id', 'int64', 'int64', [], false),
  param('trade_order_id', 'int64', 'int64', [], false),
  param('return_to_type', 'string', 'string', [], false),
  param('dispute_in_progress', 'bool', 'boolean', [], false),
  param('TradeOrderLineCreatedTimeRangeStart', 'int64', 'int64', [], false),
  param('TradeOrderLineCreatedTimeRangeEnd', 'int64', 'int64', [], false),
  param('ReverseOrderLineTimeRangeStart', 'int64', 'int64', [], false),
  param('ReverseOrderLineTimeRangeEnd', 'int64', 'int64', [], false),
  param('ReverseOrderLineModifiedTimeRangeStart', 'int64', 'int64', [], false),
  param('ReverseOrderLineModifiedTimeRangeEnd', 'int64', 'int64', [], false),
  param('QC_Decision', 'string', 'string', [], false),
];

const manualEndpointTypes: Record<string, { request: IRParam[]; response: IRParam[] }> = {
  GetBrandByPages: { request: getBrandByPagesReq, response: [] },
  GetCategoryAttributes: { request: getCategoryAttributesReq, response: [] },
  CreateProduct: { request: createProductReq, response: [] },
  UpdateProduct: { request: updateProductReq, response: [] },
  GetProducts: { request: getProductsReq, response: [] },
  GetProductItem: { request: getProductItemReq, response: [] },
  GetOrders: { request: getOrdersReq, response: [] },
  GetOrderItems: { request: getOrderItemsReq, response: [] },
  GetMultipleOrderItems: { request: getMultipleOrderItemsReq, response: [] },
  PrintAWB: { request: printAWBReq, response: [] },
  GetReverseOrdersForSeller: { request: getReverseOrdersForSellerReq, response: [] },
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

// The /reverse/getreverseordersforseller endpoint returns its payload under a
// top-level "result" key, and the SDK's default GET method parses wrapper.Data
// straight into &resp.Response. Build the response against the "result" wrapper
// so a `Response` field is emitted, then wrap the generated payload struct in a
// single `Result` field (json tag "result") so that unmarshal maps the key.
// Lazada also encodes numerics as JSON strings, so relax those fields to the
// string-tolerant Flex types.
function buildReverseOrdersResponse(structGen: StructGenerator, moduleName: string, ep: IREndpoint): void {
  defaultBuildEndpointStructs({ ...profileConfig, responseDataFieldName: 'result' } as any)(structGen, moduleName, ep);

  const respDataName = structGen.getNameForChain(moduleName, ep.name, 'ResponseData');
  const payload = structGen.allStructs.get(respDataName);
  if (!payload) return;

  const payloadName = respDataName + 'Result';
  payload.name = payloadName;
  structGen.allStructs.delete(respDataName);
  structGen.allStructs.set(payloadName, payload);

  structGen.allStructs.set(respDataName, {
    name: respDataName,
    fields: [
      { name: 'Result', type: '*' + payloadName, jsonTag: 'result', urlTag: '', comment: 'Response data' },
    ],
    fileName: payload.fileName,
  });

  const flexFields = (s: any, overrides: Record<string, string>) => {
    if (!s) return;
    for (const f of s.fields) {
      if (overrides[f.name]) f.type = overrides[f.name];
    }
  };
  const childStruct = (parent: any, field: string): any => {
    const f = parent?.fields.find((x: any) => x.name === field);
    if (!f) return null;
    const name = f.type.replace(/^(\[\]|\*)/, '');
    return structGen.allStructs.get(name) ?? null;
  };

  // Lazada declares these response fields as Number/Boolean but returns them
  // inconsistently as JSON strings or numbers depending on the payload. Relax
  // every one of them to FlexString so parsing cannot fail.
  flexFields(payload, { Total: 'FlexInt', Success: 'FlexString', PageNo: 'FlexString', PageSize: 'FlexInt' });

  const items = childStruct(payload, 'Items');
  flexFields(items, { ReverseOrderId: 'FlexString', TradeOrderId: 'FlexString', IsRtm: 'FlexString' });

  const lines = childStruct(items, 'ReverseOrderLines');
  flexFields(lines, {
    TradeOrderGmtCreate: 'FlexString',
    RefundAmount: 'FlexString',
    ReasonCode: 'FlexString',
    ReturnOrderLineGmtCreate: 'FlexString',
    ReturnOrderLineGmtModified: 'FlexString',
    ItemUnitPrice: 'FlexString',
    Sla: 'FlexString',
    ReverseOrderLineId: 'FlexString',
    TradeOrderLineId: 'FlexString',
    IsDispute: 'FlexString',
    IsNeedRefund: 'FlexString',
  });

  flexFields(childStruct(lines, 'Product'), { ProductId: 'FlexInt' });
  flexFields(childStruct(lines, 'Buyer'), { BuyerId: 'FlexString' });
}

// The renderer's built-in structTypeOverrides lookup is keyed off the last
// chain segment (e.g. "ResponseData"), which never matches the full struct
// names in doclient.config.ts. Apply the same overrides as a post-processor
// over the generated structs by name instead.
function applyStructTypeOverrides(structGen: StructGenerator): void {
  for (const [structName, fields] of Object.entries(structTypeOverrides)) {
    const s = structGen.allStructs.get(structName);
    if (!s) continue;
    const pascalToType = new Map<string, string>();
    for (const [field, type] of Object.entries(fields)) {
      pascalToType.set(toPascalCase(field), type);
    }
    for (const f of s.fields) {
      const override = pascalToType.get(f.name);
      if (override) f.type = override;
    }
  }
}

// /orders/items/get returns "data" as an array of order batches, but the API
// doc models it as a single object. Make the wrapper field a slice of the
// generated batch struct so the real payload unmarshals.
function applyOrderItemsArrayResponse(structGen: StructGenerator, moduleName: string, ep: IREndpoint): void {
  if (ep.name !== 'GetMultipleOrderItems') return;
  const respName = structGen.getNameForChain(moduleName, ep.name, 'Response');
  const resp = structGen.allStructs.get(respName);
  const dataName = structGen.getNameForChain(moduleName, ep.name, 'ResponseData');
  if (!resp || !structGen.allStructs.has(dataName)) return;
  const field = resp.fields.find((f) => f.name === 'Response');
  if (field) field.type = '[]' + dataName;
}

// PrintAWB (/order/package/document/get) returns its payload under a top-level
// "result" key. The generated method only unmarshals the response when the
// wrapper struct has a field named "Response", and wrapper.Data is the whole
// body ({result:...}), so the Response field must parse that body. Restructure
// like the reverse-orders endpoints: a PrintAWBResponseData wrapper with a
// single Result field (json "result") that captures the payload.
function applyPrintAWBResponse(structGen: StructGenerator, moduleName: string, ep: IREndpoint): void {
  if (ep.name !== 'PrintAWB') return;
  const respName = structGen.getNameForChain(moduleName, ep.name, 'Response');
  const resp = structGen.allStructs.get(respName);
  if (!resp) return;
  const field = resp.fields.find((f) => f.name === 'Result' || f.name === 'Response');
  if (!field) return;
  const payloadName = field.type.replace(/^\*/, '');
  const wrapperName = respName + 'Data';
  structGen.allStructs.set(wrapperName, {
    name: wrapperName,
    fields: [{ name: 'Result', type: '*' + payloadName, jsonTag: 'result', urlTag: '', comment: 'Response data' }],
    fileName: resp.fileName,
  });
  field.name = 'Response';
  field.type = wrapperName;
}

// Tolerances that depend on the generated struct shape rather than field types:
// add the model_quantity_purchased field the API doc omits, and reuse the
// billing address structs for the shipping address so duplicate structs are not
// emitted.
function applyOrderTuning(structGen: StructGenerator): void {
  for (const structName of ['OrderItems', 'GetOrderItemsResponseData']) {
    const s = structGen.allStructs.get(structName);
    if (!s) continue;
    if (!s.fields.some((f) => f.name === 'ModelQuantityPurchased')) {
      s.fields.push({ name: 'ModelQuantityPurchased', type: 'FlexInt', jsonTag: 'model_quantity_purchased', urlTag: '', comment: '' });
    }
  }

  const repoint = (structName: string, fieldName: string, reuse: string, drop: string) => {
    const s = structGen.allStructs.get(structName);
    if (!s) return;
    const f = s.fields.find((x) => x.name === fieldName);
    if (f && f.type === '*' + drop) f.type = '*' + reuse;
    structGen.allStructs.delete(drop);
  };
  repoint('GetOrderResponseData', 'AddressShipping', 'ResponseDataAddressBilling', 'ResponseDataAddressShipping');
  repoint('ResponseDataOrders', 'AddressShipping', 'OrdersAddressBilling', 'OrdersAddressShipping');
}

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

    // GetReverseOrdersForSeller returns its payload under a top-level "result"
    // key and Lazada encodes numerics as JSON strings; see buildReverseOrdersResponse.
    if (ep.name === 'GetReverseOrdersForSeller') {
      buildReverseOrdersResponse(structGen, moduleName, ep);
      applyStructTypeOverrides(structGen);
      return;
    }

    defaultBuildEndpointStructs(profileConfig as any)(structGen, moduleName, ep);
    applyStructTypeOverrides(structGen);
    applyOrderItemsArrayResponse(structGen, moduleName, ep);
    applyPrintAWBResponse(structGen, moduleName, ep);
    applyOrderTuning(structGen);
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
