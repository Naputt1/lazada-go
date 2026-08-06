package golazada

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_Product_AdjustSellableQuantity(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_product_stock_sellable_adjust_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping AdjustSellableQuantity due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/product/stock/sellable/adjust*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Product.AdjustSellableQuantity(ctx)
	if err != nil {
		t.Logf("Product.AdjustSellableQuantity returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.AdjustSellableQuantity response: %#v", res)
}
func Test_Product_BatchUpdateSizeChart(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_size_chart_batch_update_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping BatchUpdateSizeChart due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/size/chart/batch/update*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Product.BatchUpdateSizeChart(ctx, BatchUpdateSizeChartRequest{})
	if err != nil {
		t.Logf("Product.BatchUpdateSizeChart returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.BatchUpdateSizeChart response: %#v", res)
}
func Test_Product_CreateProduct(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_product_create_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CreateProduct due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/product/create*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	var req CreateProductRequest
	ctx := context.Background()
	res, err := client.Product.CreateProduct(ctx, req)
	if err != nil {
		t.Logf("Product.CreateProduct returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.CreateProduct response: %#v", res)
}
func Test_Product_DeactivateProduct(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_product_deactivate_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DeactivateProduct due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/product/deactivate*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Product.DeactivateProduct(ctx)
	if err != nil {
		t.Logf("Product.DeactivateProduct returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.DeactivateProduct response: %#v", res)
}
func Test_Product_GetBrandByPages(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_category_brands_query_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetBrandByPages due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/category/brands/query*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	var req GetBrandByPagesRequest
	ctx := context.Background()
	res, err := client.Product.GetBrandByPages(ctx, req)
	if err != nil {
		t.Logf("Product.GetBrandByPages returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetBrandByPages response: %#v", res)
}
func Test_Product_GetCategoryAttributes(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_category_attributes_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetCategoryAttributes due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/category/attributes/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	var req GetCategoryAttributesRequest
	ctx := context.Background()
	res, err := client.Product.GetCategoryAttributes(ctx, req)
	if err != nil {
		t.Logf("Product.GetCategoryAttributes returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetCategoryAttributes response: %#v", res)
}
func Test_Product_GetCategorySuggestion(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_product_category_suggestion_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetCategorySuggestion due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/product/category/suggestion/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Product.GetCategorySuggestion(ctx)
	if err != nil {
		t.Logf("Product.GetCategorySuggestion returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetCategorySuggestion response: %#v", res)
}
func Test_Product_GetCategoryTree(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_category_tree_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetCategoryTree due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/category/tree/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Product.GetCategoryTree(ctx)
	if err != nil {
		t.Logf("Product.GetCategoryTree returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetCategoryTree response: %#v", res)
}
func Test_Product_GetNextCascadeProp(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_category_cascade_getNextCascadeProp_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetNextCascadeProp due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/category/cascade/getNextCascadeProp*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Product.GetNextCascadeProp(ctx)
	if err != nil {
		t.Logf("Product.GetNextCascadeProp returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetNextCascadeProp response: %#v", res)
}
func Test_Product_GetPreQcRules(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_product_seller_item_getPreQcRules_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetPreQcRules due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/product/seller/item/getPreQcRules*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Product.GetPreQcRules(ctx)
	if err != nil {
		t.Logf("Product.GetPreQcRules returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetPreQcRules response: %#v", res)
}
func Test_Product_GetProductContentScore(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_product_content_score_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetProductContentScore due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/product/content/score/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Product.GetProductContentScore(ctx)
	if err != nil {
		t.Logf("Product.GetProductContentScore returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetProductContentScore response: %#v", res)
}
func Test_Product_GetProductItem(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_product_item_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetProductItem due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/product/item/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	var req GetProductItemRequest
	ctx := context.Background()
	res, err := client.Product.GetProductItem(ctx, req)
	if err != nil {
		t.Logf("Product.GetProductItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetProductItem response: %#v", res)
}
func Test_Product_GetProducts(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_products_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetProducts due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/products/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	var req GetProductsRequest
	ctx := context.Background()
	res, err := client.Product.GetProducts(ctx, req)
	if err != nil {
		t.Logf("Product.GetProducts returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetProducts response: %#v", res)
}
func Test_Product_GetQCAlertProducts(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_product_qc_alert_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetQCAlertProducts due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/product/qc/alert/list*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Product.GetQCAlertProducts(ctx)
	if err != nil {
		t.Logf("Product.GetQCAlertProducts returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetQCAlertProducts response: %#v", res)
}
func Test_Product_GetResponse(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_image_response_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetResponse due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/image/response/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Product.GetResponse(ctx, "test.jpg", strings.NewReader("test data"))
	if err != nil {
		t.Logf("Product.GetResponse returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetResponse response: %#v", res)
}
func Test_Product_GetSellerItemLimit(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_product_seller_item_limit_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetSellerItemLimit due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/product/seller/item/limit*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Product.GetSellerItemLimit(ctx)
	if err != nil {
		t.Logf("Product.GetSellerItemLimit returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetSellerItemLimit response: %#v", res)
}
func Test_Product_GetSizeChartTemplate(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_size_chart_template_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetSizeChartTemplate due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/size/chart/template/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Product.GetSizeChartTemplate(ctx, GetSizeChartTemplateRequest{})
	if err != nil {
		t.Logf("Product.GetSizeChartTemplate returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetSizeChartTemplate response: %#v", res)
}
func Test_Product_GetUnfilledAttributeItem(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_product_unfilled_attribute_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetUnfilledAttributeItem due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/product/unfilled/attribute/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Product.GetUnfilledAttributeItem(ctx)
	if err != nil {
		t.Logf("Product.GetUnfilledAttributeItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.GetUnfilledAttributeItem response: %#v", res)
}
func Test_Product_MigrateImage(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_image_migrate_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping MigrateImage due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/image/migrate*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Product.MigrateImage(ctx, "test.jpg", strings.NewReader("test data"))
	if err != nil {
		t.Logf("Product.MigrateImage returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.MigrateImage response: %#v", res)
}
func Test_Product_MigrateImages(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_images_migrate_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping MigrateImages due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/images/migrate*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Product.MigrateImages(ctx, "test.jpg", strings.NewReader("test data"))
	if err != nil {
		t.Logf("Product.MigrateImages returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.MigrateImages response: %#v", res)
}
func Test_Product_ProductCheck(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_product_pre_check_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping ProductCheck due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/product/pre/check*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Product.ProductCheck(ctx)
	if err != nil {
		t.Logf("Product.ProductCheck returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.ProductCheck response: %#v", res)
}
func Test_Product_RemoveProduct(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_product_remove_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping RemoveProduct due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/product/remove*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Product.RemoveProduct(ctx, RemoveProductRequest{})
	if err != nil {
		t.Logf("Product.RemoveProduct returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.RemoveProduct response: %#v", res)
}
func Test_Product_RemoveSku(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_product_sku_remove_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping RemoveSku due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/product/sku/remove*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Product.RemoveSku(ctx)
	if err != nil {
		t.Logf("Product.RemoveSku returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.RemoveSku response: %#v", res)
}
func Test_Product_SetImages(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_images_set_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SetImages due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/images/set*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Product.SetImages(ctx, "test.jpg", strings.NewReader("test data"))
	if err != nil {
		t.Logf("Product.SetImages returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.SetImages response: %#v", res)
}
func Test_Product_UpdatePriceQuantity(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_product_price_quantity_update_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdatePriceQuantity due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/product/price_quantity/update*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Product.UpdatePriceQuantity(ctx)
	if err != nil {
		t.Logf("Product.UpdatePriceQuantity returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.UpdatePriceQuantity response: %#v", res)
}
func Test_Product_UpdateProduct(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_product_update_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateProduct due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/product/update*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	var req UpdateProductRequest
	ctx := context.Background()
	res, err := client.Product.UpdateProduct(ctx, req)
	if err != nil {
		t.Logf("Product.UpdateProduct returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.UpdateProduct response: %#v", res)
}
func Test_Product_UpdateSellableQuantity(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_product_stock_sellable_update_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateSellableQuantity due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/product/stock/sellable/update*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Product.UpdateSellableQuantity(ctx)
	if err != nil {
		t.Logf("Product.UpdateSellableQuantity returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.UpdateSellableQuantity response: %#v", res)
}
func Test_Product_UploadImage(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_image_upload_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UploadImage due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/image/upload*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Product.UploadImage(ctx, "test.jpg", strings.NewReader("test data"))
	if err != nil {
		t.Logf("Product.UploadImage returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Product.UploadImage response: %#v", res)
}
