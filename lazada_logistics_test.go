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

func Test_LazadaLogistics_CreateCustomerAccountRelationshipByOTP(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_epis_customers_external_relationships_bundle_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CreateCustomerAccountRelationshipByOTP due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/epis/customers/external_relationships_bundle*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazadaLogistics.CreateCustomerAccountRelationshipByOTP(ctx)
	if err != nil {
		t.Logf("LazadaLogistics.CreateCustomerAccountRelationshipByOTP returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazadaLogistics.CreateCustomerAccountRelationshipByOTP response: %#v", res)
}
func Test_LazadaLogistics_CreateCustomerAccountRelationshipForExternal(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_epis_customers_external_relationships_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CreateCustomerAccountRelationshipForExternal due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/epis/customers/external_relationships*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazadaLogistics.CreateCustomerAccountRelationshipForExternal(ctx)
	if err != nil {
		t.Logf("LazadaLogistics.CreateCustomerAccountRelationshipForExternal returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazadaLogistics.CreateCustomerAccountRelationshipForExternal response: %#v", res)
}
func Test_LazadaLogistics_CreateOrUpdateCustomerWarehouse(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_epis_customers_warehouses_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CreateOrUpdateCustomerWarehouse due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/epis/customers/warehouses*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazadaLogistics.CreateOrUpdateCustomerWarehouse(ctx)
	if err != nil {
		t.Logf("LazadaLogistics.CreateOrUpdateCustomerWarehouse returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazadaLogistics.CreateOrUpdateCustomerWarehouse response: %#v", res)
}
func Test_LazadaLogistics_EpisGetDeliveryOptions(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_epis_service_delivery_options_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping EpisGetDeliveryOptions due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/logistics/epis/service/delivery_options*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazadaLogistics.EpisGetDeliveryOptions(ctx)
	if err != nil {
		t.Logf("LazadaLogistics.EpisGetDeliveryOptions returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazadaLogistics.EpisGetDeliveryOptions response: %#v", res)
}
func Test_LazadaLogistics_EpisPackageCancellation(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_epis_packages_cancel_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping EpisPackageCancellation due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/epis/packages/cancel*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazadaLogistics.EpisPackageCancellation(ctx)
	if err != nil {
		t.Logf("LazadaLogistics.EpisPackageCancellation returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazadaLogistics.EpisPackageCancellation response: %#v", res)
}
func Test_LazadaLogistics_EpisPackageCancellationV3(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_epis_packages_cancel_v3_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping EpisPackageCancellationV3 due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/epis/packages/cancel/v3*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazadaLogistics.EpisPackageCancellationV3(ctx)
	if err != nil {
		t.Logf("LazadaLogistics.EpisPackageCancellationV3 returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazadaLogistics.EpisPackageCancellationV3 response: %#v", res)
}
func Test_LazadaLogistics_EpisPackageConsignment(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_epis_packages_consign_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping EpisPackageConsignment due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/epis/packages/consign*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazadaLogistics.EpisPackageConsignment(ctx)
	if err != nil {
		t.Logf("LazadaLogistics.EpisPackageConsignment returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazadaLogistics.EpisPackageConsignment response: %#v", res)
}
func Test_LazadaLogistics_EpisPackageConsignmentV2(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_epis_packages_consign_v2_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping EpisPackageConsignmentV2 due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/epis/packages/consign/v2*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazadaLogistics.EpisPackageConsignmentV2(ctx)
	if err != nil {
		t.Logf("LazadaLogistics.EpisPackageConsignmentV2 returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazadaLogistics.EpisPackageConsignmentV2 response: %#v", res)
}
func Test_LazadaLogistics_EpisPackageCreation(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_epis_packages_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping EpisPackageCreation due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/epis/packages*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazadaLogistics.EpisPackageCreation(ctx)
	if err != nil {
		t.Logf("LazadaLogistics.EpisPackageCreation returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazadaLogistics.EpisPackageCreation response: %#v", res)
}
func Test_LazadaLogistics_EpisPackageInfoUpdate(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_epis_packages_update_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping EpisPackageInfoUpdate due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/epis/packages/update*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazadaLogistics.EpisPackageInfoUpdate(ctx)
	if err != nil {
		t.Logf("LazadaLogistics.EpisPackageInfoUpdate returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazadaLogistics.EpisPackageInfoUpdate response: %#v", res)
}
func Test_LazadaLogistics_EpisPackagePrintAwb(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_epis_packages_awb_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping EpisPackagePrintAwb due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/logistics/epis/packages/awb*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazadaLogistics.EpisPackagePrintAwb(ctx)
	if err != nil {
		t.Logf("LazadaLogistics.EpisPackagePrintAwb returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazadaLogistics.EpisPackagePrintAwb response: %#v", res)
}
func Test_LazadaLogistics_EpisPackageReadyToBeShipped(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_epis_packages_rts_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping EpisPackageReadyToBeShipped due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/epis/packages/rts*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazadaLogistics.EpisPackageReadyToBeShipped(ctx)
	if err != nil {
		t.Logf("LazadaLogistics.EpisPackageReadyToBeShipped returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazadaLogistics.EpisPackageReadyToBeShipped response: %#v", res)
}
func Test_LazadaLogistics_EpisPackageReAttempt(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_epis_packages_reattempt_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping EpisPackageReAttempt due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/epis/packages/reattempt*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazadaLogistics.EpisPackageReAttempt(ctx)
	if err != nil {
		t.Logf("LazadaLogistics.EpisPackageReAttempt returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazadaLogistics.EpisPackageReAttempt response: %#v", res)
}
func Test_LazadaLogistics_EpisUploadAwbFulfillment(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_epis_fulfillment_upload_awb_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping EpisUploadAwbFulfillment due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/epis/fulfillment/upload_awb*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazadaLogistics.EpisUploadAwbFulfillment(ctx, "test.jpg", strings.NewReader("test data"))
	if err != nil {
		t.Logf("LazadaLogistics.EpisUploadAwbFulfillment returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazadaLogistics.EpisUploadAwbFulfillment response: %#v", res)
}
func Test_LazadaLogistics_EpisXspaceCreate(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_epis_xspace_create_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping EpisXspaceCreate due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/epis/xspace/create*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazadaLogistics.EpisXspaceCreate(ctx)
	if err != nil {
		t.Logf("LazadaLogistics.EpisXspaceCreate returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazadaLogistics.EpisXspaceCreate response: %#v", res)
}
func Test_LazadaLogistics_EpisXspaceGetDetail(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_epis_xspace_detail_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping EpisXspaceGetDetail due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/epis/xspace/detail*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazadaLogistics.EpisXspaceGetDetail(ctx)
	if err != nil {
		t.Logf("LazadaLogistics.EpisXspaceGetDetail returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazadaLogistics.EpisXspaceGetDetail response: %#v", res)
}
func Test_LazadaLogistics_EpisXspaceQuery(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_epis_xspace_query_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping EpisXspaceQuery due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/epis/xspace/query*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazadaLogistics.EpisXspaceQuery(ctx)
	if err != nil {
		t.Logf("LazadaLogistics.EpisXspaceQuery returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazadaLogistics.EpisXspaceQuery response: %#v", res)
}
func Test_LazadaLogistics_EpisXspaceRateTicket(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_epis_xspace_rate_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping EpisXspaceRateTicket due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/epis/xspace/rate*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazadaLogistics.EpisXspaceRateTicket(ctx)
	if err != nil {
		t.Logf("LazadaLogistics.EpisXspaceRateTicket returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazadaLogistics.EpisXspaceRateTicket response: %#v", res)
}
func Test_LazadaLogistics_EstimateShippingFee(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_epis_estimate_shipping_fee_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping EstimateShippingFee due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/epis/estimate_shipping_fee*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazadaLogistics.EstimateShippingFee(ctx)
	if err != nil {
		t.Logf("LazadaLogistics.EstimateShippingFee returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazadaLogistics.EstimateShippingFee response: %#v", res)
}
func Test_LazadaLogistics_GetShippingFee(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_epis_get_shipping_fee_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetShippingFee due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/logistics/epis/get_shipping_fee*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazadaLogistics.GetShippingFee(ctx)
	if err != nil {
		t.Logf("LazadaLogistics.GetShippingFee returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazadaLogistics.GetShippingFee response: %#v", res)
}
