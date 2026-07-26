import type {
  Config,
  IR,
  IREndpoint,
  IRModule,
  IRError,
  IRParam,
  SourceAdapter,
} from '@doclient/cli';
import { toPascalCase, createCachedFetcher } from '@doclient/cli';

const CATEGORY_API = 'https://isvconsole.lazada.com/handler/share/apidoc/getApiCategoryMixed.json';
const DETAIL_API = 'https://isvconsole.lazada.com/handler/share/apidoc/getApi.json';

const AUTH_PATHS = new Set([
  '/auth/token/create',
  '/auth/token/createWithOpenId',
  '/auth/token/refresh',
]);

const COMMON_PARAMS = new Set([
  'app_key',
  'sign_method',
  'timestamp',
  'partner_id',
  'access_token',
  'sign',
]);

interface Category {
  name: string;
  description?: string;
  id: number;
  apiList: ApiSummary[];
}

interface ApiSummary {
  title: string;
  description?: string;
  path: string;
  methodType: string;
}

interface ApiDetailData {
  desc?: string;
  authType?: number;
  id?: number;
  parameters?: { data: Param[] };
  outputParameters?: { data: Param[] };
  apiList?: ApiDetailItem[];
}

interface ApiDetailItem {
  title: string;
  path: string;
  method?: string;
  errorCodes?: { data: ErrorCodeEntry[] };
  examples?: { data: ExampleEntry[] };
}

interface ErrorCodeEntry {
  code: string;
  codeDesc?: string;
  solution?: string;
}

interface ExampleEntry {
  responseCode?: string;
}

interface Param {
  name: string;
  type: string;
  required?: boolean;
  desc?: string;
  children?: Param[];
}

function isAuthPath(path: string): boolean {
  return AUTH_PATHS.has(path);
}

function isCommonParam(name: string): boolean {
  return COMMON_PARAMS.has(name);
}

function normalizeType(lazadaType: string): string {
  switch (lazadaType) {
    case 'String':
    case 'Date':
    case 'Byte':
    case 'json':
      return 'string';
    case 'Number':
      return 'number';
    case 'Boolean':
      return 'boolean';
    case 'Object':
      return 'object';
    case 'Object[]':
      return 'object[]';
    case 'String[]':
      return 'string[]';
    case 'Number[]':
      return 'number[]';
    default:
      return lazadaType;
  }
}

function mapParam(p: Param, typeOverrides?: Record<string, string>): IRParam {
  const irType = typeOverrides?.[p.name] ?? normalizeType(p.type);
  return {
    name: p.name,
    type: irType,
    shopeeType: p.type,
    description: p.desc ?? '',
    required: p.required ?? false,
    children: (p.children ?? []).map((c) => mapParam(c, typeOverrides)),
  };
}

const SKIP_FIXTURE_FIELDS = new Set(['code', 'type', 'message', 'request_id', 'success', 'error_msg', 'error_code']);

function inferParamsFromFixture(fixture: any, typeOverrides?: Record<string, string>): IRParam[] {
  if (typeof fixture !== 'object' || fixture === null) return [];
  return Object.entries(fixture)
    .filter(([k]) => !SKIP_FIXTURE_FIELDS.has(k))
    .map(([k, v]) => paramFromValue(k, v, typeOverrides));
}

function paramFromValue(name: string, value: any, typeOverrides?: Record<string, string>): IRParam {
  const override = typeOverrides?.[name];
  if (override) {
    return { name, type: override, shopeeType: override, description: '', required: true, children: [] };
  }
  if (Array.isArray(value)) {
    if (value.length > 0 && typeof value[0] === 'object' && value[0] !== null) {
      return {
        name,
        type: 'object[]',
        shopeeType: 'object[]',
        description: '',
        required: true,
        children: Object.entries(value[0]).map(([k, v]) => paramFromValue(k, v, typeOverrides)),
      };
    }
    if (value.length > 0) {
      return { name, type: 'string[]', shopeeType: 'string[]', description: '', required: true, children: [] };
    }
    return { name, type: 'object[]', shopeeType: 'object[]', description: '', required: true, children: [] };
  }
  if (typeof value === 'object' && value !== null) {
    return {
      name,
      type: 'object',
      shopeeType: 'object',
      description: '',
      required: true,
      children: Object.entries(value).map(([k, v]) => paramFromValue(k, v, typeOverrides)),
    };
  }
  if (typeof value === 'number') {
    return { name, type: 'int64', shopeeType: 'int64', description: '', required: true, children: [] };
  }
  if (typeof value === 'boolean') {
    return { name, type: 'boolean', shopeeType: 'boolean', description: '', required: true, children: [] };
  }
  return { name, type: 'string', shopeeType: 'string', description: '', required: true, children: [] };
}

function determineMethod(methodType: string, title: string): 'GET' | 'POST' {
  switch (methodType) {
    case 'GET':
      return 'GET';
    case 'POST':
      return 'POST';
    default:
      const getPrefixes = ['Get', 'List', 'Query', 'Search', 'Fetch', 'View'];
      for (const prefix of getPrefixes) {
        if (title.startsWith(prefix)) {
          return 'GET';
        }
      }
      return 'POST';
  }
}

function cleanFixtureName(name: string): string {
  let cleaned = '';
  for (const ch of name) {
    if (/[a-zA-Z0-9._]/.test(ch)) cleaned += ch;
    else cleaned += '_';
  }
  while (cleaned.includes('__')) cleaned = cleaned.replace(/__/g, '_');
  return cleaned;
}

function toGoErrorName(code: string): string {
  const withUnderscore = code.replace(/[-.]/g, '_');
  const parts = withUnderscore.split('_').filter(Boolean);
  const cameled = parts.map((p) => p.charAt(0).toUpperCase() + p.slice(1)).join('');
  return 'Err' + cameled;
}

function sleep(ms: number): Promise<void> {
  return new Promise((r) => setTimeout(r, ms));
}

function createFetcher(config: Config) {
  if (config.cacheDir) {
    return createCachedFetcher(config.cacheDir).fetchJSON;
  }
  return async <T = unknown>(url: string): Promise<T> => {
    const res = await fetch(url);
    if (!res.ok) throw new Error(`fetch failed: ${url} (${res.status})`);
    return res.json() as Promise<T>;
  };
}

function moduleNameFromCategory(categoryName: string): string {
  let name = categoryName;
  if (name.endsWith(' API')) {
    name = name.slice(0, -4);
  }
  return toPascalCase(name.trim());
}

function makeUniqueName(name: string, used: Map<string, number>): string {
  const count = used.get(name) ?? 0;
  used.set(name, count + 1);
  if (count > 0) {
    return `${name}${count}`;
  }
  return name;
}

export const lazadaSource: SourceAdapter = {
  name: 'lazada',

  async execute(config: Config): Promise<IR> {
    const typeOverrides = config.mappings?.typeOverrides;
    const staticModules = config.mappings?.staticModules;
    const fetchJSON = createFetcher(config);

    const catResp = await fetchJSON<{ data: Category[]; success: boolean }>(CATEGORY_API);
    const categories = catResp.data ?? [];

    const irModules: IRModule[] = [];
    const allErrors = new Map<string, IRError>();
    const fixtures: Array<{ filename: string; content: string }> = [];
    const allMethodNames = new Map<string, number>();

    for (const cat of categories) {
      const moduleName = moduleNameFromCategory(cat.name);

      if (staticModules?.some((m) => m.toLowerCase() === moduleName.toLowerCase())) continue;

      const irEndpoints: IREndpoint[] = [];

      for (const api of cat.apiList ?? []) {
        if (isAuthPath(api.path)) continue;

        await sleep(15);

        const detailUrl = `${DETAIL_API}?path=${encodeURIComponent(api.path)}`;
        let detail: ApiDetailData;
        try {
          const resp = await fetchJSON<{ data: ApiDetailData; success: boolean }>(detailUrl);
          if (!resp.success) continue;
          detail = resp.data;
        } catch {
          continue;
        }

        const rawParams = detail.parameters?.data ?? [];
        const requestParams = rawParams
          .filter((p) => !isCommonParam(p.name))
          .map((p) => mapParam(p, typeOverrides));

		const rawOutput = detail.outputParameters?.data ?? [];
		let responseParams = rawOutput.map((p) => mapParam(p, typeOverrides));

        const method = determineMethod(api.methodType, api.title);

        const endpointName = toPascalCase(api.title);
        const uniqueName = makeUniqueName(endpointName, allMethodNames);

        const item = detail.apiList?.[0];
        const errorList = item?.errorCodes?.data ?? [];
        for (const e of errorList) {
          const rawCode = e.code?.trim() ?? '';
          const code = rawCode.replace(/^"|"$/g, '');
          if (code) {
            const key = toGoErrorName(code);
            if (!allErrors.has(key)) {
              allErrors.set(key, { code, description: e.codeDesc ?? '' });
            }
          }
        }

        const fixtureName = cleanFixtureName(api.path.replace(/^\//, '').replace(/\//g, '.')) + '_resp.json';

        let fixtureContent: string | undefined;
        const example = item?.examples?.data?.[0];
        if (example?.responseCode) {
          try {
            const parsed = JSON.parse(example.responseCode);
            fixtureContent = JSON.stringify(parsed, null, 2);
          } catch {
            fixtureContent = example.responseCode;
          }
        }
        if (fixtureContent) {
          fixtures.push({ filename: fixtureName, content: fixtureContent });
        }

		if (responseParams.length === 0 && fixtureContent) {
			try {
				const parsed = JSON.parse(fixtureContent);
				const inferred = inferParamsFromFixture(parsed, typeOverrides);
				if (inferred.length > 0) {
					responseParams = inferred;
				}
			} catch {
				// fixture not parseable as JSON, keep empty responseParams
			}
		}

        const irEp: IREndpoint = {
          name: uniqueName,
          method,
          path: api.path,
          fullPath: api.path,
          description: detail.desc ?? api.description ?? '',
          docUrl: `https://open.lazada.com/apps/doc/api?path=${encodeURIComponent(api.path)}`,
          apiType: 'Shop',
          isUpload: api.path.includes('upload') || api.path.includes('image'),
          fullApiName: api.path,
          requestParams: requestParams.sort((a, b) => a.name.localeCompare(b.name)),
          responseParams: responseParams.sort((a, b) => a.name.localeCompare(b.name)),
          errors: errorList
            .map((e) => ({ code: (e.code?.trim() ?? '').replace(/^"|"$/g, ''), description: e.codeDesc ?? '' }))
            .filter((e) => e.code),
        };

        irEndpoints.push(irEp);
      }

      if (irEndpoints.length > 0) {
        irEndpoints.sort((a, b) => a.name.localeCompare(b.name));
        irModules.push({ name: moduleName, moduleId: cat.id, endpoints: irEndpoints });
      }
    }

    irModules.sort((a, b) => a.name.localeCompare(b.name));

    return {
      name: config.name,
      modules: irModules,
      constants: [],
      errors: Array.from(allErrors.values()).sort((a, b) => a.code.localeCompare(b.code)),
      fixtures: fixtures.sort((a, b) => a.filename.localeCompare(b.filename)),
    };
  },
};
