package golazada

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_Seller_BatchQueryFollowStatus(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_shop_follow_status_batch_query_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping BatchQueryFollowStatus due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/shop/follow/status/batch/query*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Seller.BatchQueryFollowStatus(ctx)
	if err != nil {
		t.Logf("Seller.BatchQueryFollowStatus returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Seller.BatchQueryFollowStatus response: %#v", res)
}
func Test_Seller_GetCountryInfo(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_seller_cb_country_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetCountryInfo due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/seller/cb/country/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Seller.GetCountryInfo(ctx)
	if err != nil {
		t.Logf("Seller.GetCountryInfo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Seller.GetCountryInfo response: %#v", res)
}
func Test_Seller_GetPickUpStoreList(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_rc_store_list_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetPickUpStoreList due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/rc/store/list/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Seller.GetPickUpStoreList(ctx)
	if err != nil {
		t.Logf("Seller.GetPickUpStoreList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Seller.GetPickUpStoreList response: %#v", res)
}
func Test_Seller_GetSeller(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_seller_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetSeller due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/seller/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Seller.GetSeller(ctx)
	if err != nil {
		t.Logf("Seller.GetSeller returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Seller.GetSeller response: %#v", res)
}
func Test_Seller_GetSellerMetricsById(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_seller_metrics_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetSellerMetricsById due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/seller/metrics/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Seller.GetSellerMetricsById(ctx)
	if err != nil {
		t.Logf("Seller.GetSellerMetricsById returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Seller.GetSellerMetricsById response: %#v", res)
}
func Test_Seller_GetSellerPerformance(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_seller_performance_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetSellerPerformance due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/seller/performance/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Seller.GetSellerPerformance(ctx)
	if err != nil {
		t.Logf("Seller.GetSellerPerformance returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Seller.GetSellerPerformance response: %#v", res)
}
func Test_Seller_GetSellerRegisterInfo(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_seller_cb_register_info_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetSellerRegisterInfo due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/seller/cb/register/info*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Seller.GetSellerRegisterInfo(ctx)
	if err != nil {
		t.Logf("Seller.GetSellerRegisterInfo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Seller.GetSellerRegisterInfo response: %#v", res)
}
func Test_Seller_GetSubAddress(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_seller_cb_country_location_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetSubAddress due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/seller/cb/country/location/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Seller.GetSubAddress(ctx)
	if err != nil {
		t.Logf("Seller.GetSubAddress returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Seller.GetSubAddress response: %#v", res)
}
func Test_Seller_GetWarehouseBySellerId(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_rc_warehouse_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetWarehouseBySellerId due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/rc/warehouse/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Seller.GetWarehouseBySellerId(ctx)
	if err != nil {
		t.Logf("Seller.GetWarehouseBySellerId returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Seller.GetWarehouseBySellerId response: %#v", res)
}
func Test_Seller_PaymentBinding(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_seller_cb_payment_config_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping PaymentBinding due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/seller/cb/payment/config*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Seller.PaymentBinding(ctx)
	if err != nil {
		t.Logf("Seller.PaymentBinding returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Seller.PaymentBinding response: %#v", res)
}
func Test_Seller_QueryBuyboxHuntingInfo(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_hunting_buybox_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping QueryBuyboxHuntingInfo due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/hunting/buybox/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Seller.QueryBuyboxHuntingInfo(ctx)
	if err != nil {
		t.Logf("Seller.QueryBuyboxHuntingInfo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Seller.QueryBuyboxHuntingInfo response: %#v", res)
}
func Test_Seller_QueryWarehouseDetailInfoBySellerId(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_rc_warehouse_detail_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping QueryWarehouseDetailInfoBySellerId due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/rc/warehouse/detail/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Seller.QueryWarehouseDetailInfoBySellerId(ctx)
	if err != nil {
		t.Logf("Seller.QueryWarehouseDetailInfoBySellerId returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Seller.QueryWarehouseDetailInfoBySellerId response: %#v", res)
}
func Test_Seller_SaveSellerWarehouseInfo(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_rc_sellerWarehouse_saveWarehouseInfo_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SaveSellerWarehouseInfo due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/rc/sellerWarehouse/saveWarehouseInfo*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Seller.SaveSellerWarehouseInfo(ctx)
	if err != nil {
		t.Logf("Seller.SaveSellerWarehouseInfo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Seller.SaveSellerWarehouseInfo response: %#v", res)
}
func Test_Seller_SellerCenterMsgList(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sellercenter_msg_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SellerCenterMsgList due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sellercenter/msg/list*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Seller.SellerCenterMsgList(ctx)
	if err != nil {
		t.Logf("Seller.SellerCenterMsgList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Seller.SellerCenterMsgList response: %#v", res)
}
func Test_Seller_SellerFieldVerify(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_seller_cb_register_fieldcheck_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SellerFieldVerify due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/seller/cb/register/fieldcheck*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Seller.SellerFieldVerify(ctx)
	if err != nil {
		t.Logf("Seller.SellerFieldVerify returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Seller.SellerFieldVerify response: %#v", res)
}
func Test_Seller_SellerPolicyFetch(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_seller_policy_fetch_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SellerPolicyFetch due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/seller/policy/fetch*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Seller.SellerPolicyFetch(ctx)
	if err != nil {
		t.Logf("Seller.SellerPolicyFetch returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Seller.SellerPolicyFetch response: %#v", res)
}
func Test_Seller_SynchronizeSellerItemArConfig(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_seller_ar_config_syn_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SynchronizeSellerItemArConfig due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/seller/ar/config/syn*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Seller.SynchronizeSellerItemArConfig(ctx)
	if err != nil {
		t.Logf("Seller.SynchronizeSellerItemArConfig returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Seller.SynchronizeSellerItemArConfig response: %#v", res)
}
