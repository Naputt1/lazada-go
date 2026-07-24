package golazada

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_ChoiceCustomized_BatchDeliverJitPurchaseOrder(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_jit_purchase_order_batch_pickup_deliver_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping BatchDeliverJitPurchaseOrder due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/jit/purchase_order/batch_pickup_deliver*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.ChoiceCustomized.BatchDeliverJitPurchaseOrder(ctx)
	if err != nil {
		t.Logf("ChoiceCustomized.BatchDeliverJitPurchaseOrder returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ChoiceCustomized.BatchDeliverJitPurchaseOrder response: %#v", res)
}
func Test_ChoiceCustomized_EditChoiceSkuStock(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_choice_stock_edit_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping EditChoiceSkuStock due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/choice/stock/edit*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.ChoiceCustomized.EditChoiceSkuStock(ctx)
	if err != nil {
		t.Logf("ChoiceCustomized.EditChoiceSkuStock returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ChoiceCustomized.EditChoiceSkuStock response: %#v", res)
}
func Test_ChoiceCustomized_GetChoiceProductItem(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_choice_product_item_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetChoiceProductItem due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/choice/product/item/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.ChoiceCustomized.GetChoiceProductItem(ctx)
	if err != nil {
		t.Logf("ChoiceCustomized.GetChoiceProductItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ChoiceCustomized.GetChoiceProductItem response: %#v", res)
}
func Test_ChoiceCustomized_GetChoiceProducts(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_choice_products_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetChoiceProducts due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/choice/products/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.ChoiceCustomized.GetChoiceProducts(ctx)
	if err != nil {
		t.Logf("ChoiceCustomized.GetChoiceProducts returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ChoiceCustomized.GetChoiceProducts response: %#v", res)
}
func Test_ChoiceCustomized_GetChoiceSeller(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_choice_seller_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetChoiceSeller due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/choice/seller/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.ChoiceCustomized.GetChoiceSeller(ctx)
	if err != nil {
		t.Logf("ChoiceCustomized.GetChoiceSeller returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ChoiceCustomized.GetChoiceSeller response: %#v", res)
}
func Test_ChoiceCustomized_GetChoiceSkuItemRelationBySku(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_choice_sku_item_relation_get_by_sku_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetChoiceSkuItemRelationBySku due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/choice/sku_item_relation/get_by_sku*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.ChoiceCustomized.GetChoiceSkuItemRelationBySku(ctx)
	if err != nil {
		t.Logf("ChoiceCustomized.GetChoiceSkuItemRelationBySku returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ChoiceCustomized.GetChoiceSkuItemRelationBySku response: %#v", res)
}
func Test_ChoiceCustomized_PackageJitPurchaseOrder(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_jit_purchase_order_package_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping PackageJitPurchaseOrder due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/jit/purchase_order/package*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.ChoiceCustomized.PackageJitPurchaseOrder(ctx)
	if err != nil {
		t.Logf("ChoiceCustomized.PackageJitPurchaseOrder returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ChoiceCustomized.PackageJitPurchaseOrder response: %#v", res)
}
func Test_ChoiceCustomized_PrintJitPurchaseOrderAndItem(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_jit_purchase_order_print_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping PrintJitPurchaseOrderAndItem due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/jit/purchase_order/print*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.ChoiceCustomized.PrintJitPurchaseOrderAndItem(ctx)
	if err != nil {
		t.Logf("ChoiceCustomized.PrintJitPurchaseOrderAndItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ChoiceCustomized.PrintJitPurchaseOrderAndItem response: %#v", res)
}
func Test_ChoiceCustomized_PrintPickuoOrder(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_pickup_order_print_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping PrintPickuoOrder due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/pickup_order/print*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.ChoiceCustomized.PrintPickuoOrder(ctx)
	if err != nil {
		t.Logf("ChoiceCustomized.PrintPickuoOrder returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ChoiceCustomized.PrintPickuoOrder response: %#v", res)
}
func Test_ChoiceCustomized_QueryListJitPurchaseOrder(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_jit_purchase_order_query_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping QueryListJitPurchaseOrder due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/jit/purchase_order/query_list*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.ChoiceCustomized.QueryListJitPurchaseOrder(ctx)
	if err != nil {
		t.Logf("ChoiceCustomized.QueryListJitPurchaseOrder returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ChoiceCustomized.QueryListJitPurchaseOrder response: %#v", res)
}
func Test_ChoiceCustomized_QueryListPurchaseItem(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_jit_purchase_order_query_list_purchase_item_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping QueryListPurchaseItem due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/jit/purchase_order/query_list_purchase_item*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.ChoiceCustomized.QueryListPurchaseItem(ctx)
	if err != nil {
		t.Logf("ChoiceCustomized.QueryListPurchaseItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ChoiceCustomized.QueryListPurchaseItem response: %#v", res)
}
func Test_ChoiceCustomized_QueryPickupOrder(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_pickup_order_query_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping QueryPickupOrder due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/pickup_order/query*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.ChoiceCustomized.QueryPickupOrder(ctx)
	if err != nil {
		t.Logf("ChoiceCustomized.QueryPickupOrder returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ChoiceCustomized.QueryPickupOrder response: %#v", res)
}
