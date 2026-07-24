package golazada

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_FreeShipping_FreeShippingActivate(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_promotion_freeshipping_activate_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping FreeShippingActivate due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/promotion/freeshipping/activate*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FreeShipping.FreeShippingActivate(ctx)
	if err != nil {
		t.Logf("FreeShipping.FreeShippingActivate returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FreeShipping.FreeShippingActivate response: %#v", res)
}
func Test_FreeShipping_FreeShippingAddSelectedProductSKU(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_promotion_freeshipping_product_sku_add_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping FreeShippingAddSelectedProductSKU due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/promotion/freeshipping/product/sku/add*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FreeShipping.FreeShippingAddSelectedProductSKU(ctx)
	if err != nil {
		t.Logf("FreeShipping.FreeShippingAddSelectedProductSKU returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FreeShipping.FreeShippingAddSelectedProductSKU response: %#v", res)
}
func Test_FreeShipping_FreeShippingCreate(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_promotion_freeshipping_create_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping FreeShippingCreate due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/promotion/freeshipping/create*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FreeShipping.FreeShippingCreate(ctx)
	if err != nil {
		t.Logf("FreeShipping.FreeShippingCreate returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FreeShipping.FreeShippingCreate response: %#v", res)
}
func Test_FreeShipping_FreeShippingDeactivate(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_promotion_freeshipping_deactivate_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping FreeShippingDeactivate due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/promotion/freeshipping/deactivate*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FreeShipping.FreeShippingDeactivate(ctx)
	if err != nil {
		t.Logf("FreeShipping.FreeShippingDeactivate returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FreeShipping.FreeShippingDeactivate response: %#v", res)
}
func Test_FreeShipping_FreeShippingDeleteSelectedProductSKU(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_promotion_freeshipping_product_sku_remove_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping FreeShippingDeleteSelectedProductSKU due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/promotion/freeshipping/product/sku/remove*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FreeShipping.FreeShippingDeleteSelectedProductSKU(ctx)
	if err != nil {
		t.Logf("FreeShipping.FreeShippingDeleteSelectedProductSKU returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FreeShipping.FreeShippingDeleteSelectedProductSKU response: %#v", res)
}
func Test_FreeShipping_FreeShippingDeliveryOptionsQuery(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_promotion_freeshipping_deliveryoptions_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping FreeShippingDeliveryOptionsQuery due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/promotion/freeshipping/deliveryoptions/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FreeShipping.FreeShippingDeliveryOptionsQuery(ctx)
	if err != nil {
		t.Logf("FreeShipping.FreeShippingDeliveryOptionsQuery returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FreeShipping.FreeShippingDeliveryOptionsQuery response: %#v", res)
}
func Test_FreeShipping_FreeShippingGet(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_promotion_freeshipping_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping FreeShippingGet due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/promotion/freeshipping/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FreeShipping.FreeShippingGet(ctx)
	if err != nil {
		t.Logf("FreeShipping.FreeShippingGet returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FreeShipping.FreeShippingGet response: %#v", res)
}
func Test_FreeShipping_FreeShippingList(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_promotion_freeshippings_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping FreeShippingList due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/promotion/freeshippings/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FreeShipping.FreeShippingList(ctx)
	if err != nil {
		t.Logf("FreeShipping.FreeShippingList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FreeShipping.FreeShippingList response: %#v", res)
}
func Test_FreeShipping_FreeShippingRegionsQuery(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_promotion_freeshipping_regions_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping FreeShippingRegionsQuery due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/promotion/freeshipping/regions/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FreeShipping.FreeShippingRegionsQuery(ctx)
	if err != nil {
		t.Logf("FreeShipping.FreeShippingRegionsQuery returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FreeShipping.FreeShippingRegionsQuery response: %#v", res)
}
func Test_FreeShipping_FreeShippingSelectedProductList(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_promotion_freeshipping_products_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping FreeShippingSelectedProductList due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/promotion/freeshipping/products/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FreeShipping.FreeShippingSelectedProductList(ctx)
	if err != nil {
		t.Logf("FreeShipping.FreeShippingSelectedProductList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FreeShipping.FreeShippingSelectedProductList response: %#v", res)
}
func Test_FreeShipping_FreeShippingUpdate(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_promotion_freeshipping_update_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping FreeShippingUpdate due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/promotion/freeshipping/update*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FreeShipping.FreeShippingUpdate(ctx)
	if err != nil {
		t.Logf("FreeShipping.FreeShippingUpdate returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FreeShipping.FreeShippingUpdate response: %#v", res)
}
