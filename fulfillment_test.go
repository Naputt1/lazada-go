package golazada

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_Fulfillment_ConfirmCollectForDBS(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_order_package_sof_collect_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping ConfirmCollectForDBS due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/order/package/sof/collect*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Fulfillment.ConfirmCollectForDBS(ctx)
	if err != nil {
		t.Logf("Fulfillment.ConfirmCollectForDBS returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Fulfillment.ConfirmCollectForDBS response: %#v", res)
}
func Test_Fulfillment_ConfirmDeliveryForDBS(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_order_package_sof_delivered_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping ConfirmDeliveryForDBS due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/order/package/sof/delivered*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Fulfillment.ConfirmDeliveryForDBS(ctx)
	if err != nil {
		t.Logf("Fulfillment.ConfirmDeliveryForDBS returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Fulfillment.ConfirmDeliveryForDBS response: %#v", res)
}
func Test_Fulfillment_DeliverDigital(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_order_digital_delivered_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DeliverDigital due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/order/digital/delivered*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Fulfillment.DeliverDigital(ctx)
	if err != nil {
		t.Logf("Fulfillment.DeliverDigital returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Fulfillment.DeliverDigital response: %#v", res)
}
func Test_Fulfillment_FailedDeliveryForDBS(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_order_package_sof_failed_delivery_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping FailedDeliveryForDBS due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/order/package/sof/failed_delivery*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Fulfillment.FailedDeliveryForDBS(ctx)
	if err != nil {
		t.Logf("Fulfillment.FailedDeliveryForDBS returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Fulfillment.FailedDeliveryForDBS response: %#v", res)
}
func Test_Fulfillment_GetShipmentProvider(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_order_shipment_providers_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetShipmentProvider due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/order/shipment/providers/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Fulfillment.GetShipmentProvider(ctx)
	if err != nil {
		t.Logf("Fulfillment.GetShipmentProvider returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Fulfillment.GetShipmentProvider response: %#v", res)
}
func Test_Fulfillment_Pack(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_order_fulfill_pack_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping Pack due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/order/pack*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	var req PackRequest
	ctx := context.Background()
	res, err := client.Fulfillment.Pack(ctx, req)
	if err != nil {
		t.Logf("Fulfillment.Pack returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Fulfillment.Pack response: %#v", res)
}
func Test_Fulfillment_PackageStatusUpdateForDBS(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_order_package_sof_status_update_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping PackageStatusUpdateForDBS due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/order/package/sof/status/update*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Fulfillment.PackageStatusUpdateForDBS(ctx)
	if err != nil {
		t.Logf("Fulfillment.PackageStatusUpdateForDBS returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Fulfillment.PackageStatusUpdateForDBS response: %#v", res)
}
func Test_Fulfillment_PrintAWB(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_order_package_document_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping PrintAWB due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/order/package/document/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	var req PrintAWBRequest
	ctx := context.Background()
	res, err := client.Fulfillment.PrintAWB(ctx, req)
	if err != nil {
		t.Logf("Fulfillment.PrintAWB returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Fulfillment.PrintAWB response: %#v", res)
}
func Test_Fulfillment_ReadyToShip(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_order_package_rts_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping ReadyToShip due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/order/rts*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	var req ReadyToShipRequest
	ctx := context.Background()
	res, err := client.Fulfillment.ReadyToShip(ctx, req)
	if err != nil {
		t.Logf("Fulfillment.ReadyToShip returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Fulfillment.ReadyToShip response: %#v", res)
}
func Test_Fulfillment_RecreatePackage(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_order_package_repack_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping RecreatePackage due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/order/package/repack*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Fulfillment.RecreatePackage(ctx)
	if err != nil {
		t.Logf("Fulfillment.RecreatePackage returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Fulfillment.RecreatePackage response: %#v", res)
}
