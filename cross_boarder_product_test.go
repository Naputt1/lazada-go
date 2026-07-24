package golazada

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_CrossBoarderProduct_CreateGlobalProduct(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_product_global_create_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CreateGlobalProduct due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/product/global/create*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.CrossBoarderProduct.CreateGlobalProduct(ctx)
	if err != nil {
		t.Logf("CrossBoarderProduct.CreateGlobalProduct returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("CrossBoarderProduct.CreateGlobalProduct response: %#v", res)
}
func Test_CrossBoarderProduct_DeleteMerchantProduct(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_product_global_delete_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DeleteMerchantProduct due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/product/global/delete*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.CrossBoarderProduct.DeleteMerchantProduct(ctx)
	if err != nil {
		t.Logf("CrossBoarderProduct.DeleteMerchantProduct returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("CrossBoarderProduct.DeleteMerchantProduct response: %#v", res)
}
func Test_CrossBoarderProduct_GetGlobalProductExtension(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_product_global_extension_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetGlobalProductExtension due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/product/global/extension*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.CrossBoarderProduct.GetGlobalProductExtension(ctx)
	if err != nil {
		t.Logf("CrossBoarderProduct.GetGlobalProductExtension returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("CrossBoarderProduct.GetGlobalProductExtension response: %#v", res)
}
func Test_CrossBoarderProduct_GetGlobalProductStatus(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_product_global_status_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetGlobalProductStatus due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/product/global/status/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.CrossBoarderProduct.GetGlobalProductStatus(ctx)
	if err != nil {
		t.Logf("CrossBoarderProduct.GetGlobalProductStatus returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("CrossBoarderProduct.GetGlobalProductStatus response: %#v", res)
}
func Test_CrossBoarderProduct_GetRecommendPrice(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_product_global_semi_recommend_price_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetRecommendPrice due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/product/global/semi/recommend/price/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.CrossBoarderProduct.GetRecommendPrice(ctx)
	if err != nil {
		t.Logf("CrossBoarderProduct.GetRecommendPrice returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("CrossBoarderProduct.GetRecommendPrice response: %#v", res)
}
func Test_CrossBoarderProduct_GetUnfilledAttribute(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_product_global_unfilled_attribute_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetUnfilledAttribute due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/product/global/unfilled/attribute/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.CrossBoarderProduct.GetUnfilledAttribute(ctx)
	if err != nil {
		t.Logf("CrossBoarderProduct.GetUnfilledAttribute returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("CrossBoarderProduct.GetUnfilledAttribute response: %#v", res)
}
func Test_CrossBoarderProduct_GetUpgradableGlobalPlusProductList(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_product_global_semi_avaible_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetUpgradableGlobalPlusProductList due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/product/global/semi/avaible/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.CrossBoarderProduct.GetUpgradableGlobalPlusProductList(ctx)
	if err != nil {
		t.Logf("CrossBoarderProduct.GetUpgradableGlobalPlusProductList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("CrossBoarderProduct.GetUpgradableGlobalPlusProductList response: %#v", res)
}
func Test_CrossBoarderProduct_SemiProductUpdate(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_product_global_semi_update_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SemiProductUpdate due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/product/global/semi/update*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.CrossBoarderProduct.SemiProductUpdate(ctx)
	if err != nil {
		t.Logf("CrossBoarderProduct.SemiProductUpdate returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("CrossBoarderProduct.SemiProductUpdate response: %#v", res)
}
func Test_CrossBoarderProduct_SemiProductUpgrade(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_product_global_semi_upgrade_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SemiProductUpgrade due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/product/global/semi/upgrade*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.CrossBoarderProduct.SemiProductUpgrade(ctx)
	if err != nil {
		t.Logf("CrossBoarderProduct.SemiProductUpgrade returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("CrossBoarderProduct.SemiProductUpgrade response: %#v", res)
}
func Test_CrossBoarderProduct_UpdateGlobalProductAttribute(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_product_global_attribute_update_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateGlobalProductAttribute due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/product/global/attribute/update*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.CrossBoarderProduct.UpdateGlobalProductAttribute(ctx)
	if err != nil {
		t.Logf("CrossBoarderProduct.UpdateGlobalProductAttribute returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("CrossBoarderProduct.UpdateGlobalProductAttribute response: %#v", res)
}
func Test_CrossBoarderProduct_UpdateProductStatus(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_product_global_update_status_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateProductStatus due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/product/global/update/status*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.CrossBoarderProduct.UpdateProductStatus(ctx)
	if err != nil {
		t.Logf("CrossBoarderProduct.UpdateProductStatus returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("CrossBoarderProduct.UpdateProductStatus response: %#v", res)
}
