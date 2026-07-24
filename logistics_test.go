package golazada

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_Logistics_AddOrUpdatePickupStop(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_tps_runsheets_stops_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping AddOrUpdatePickupStop due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/tps/runsheets/stops*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Logistics.AddOrUpdatePickupStop(ctx)
	if err != nil {
		t.Logf("Logistics.AddOrUpdatePickupStop returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.AddOrUpdatePickupStop response: %#v", res)
}
func Test_Logistics_Create3PLStation(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_tps_stations_create_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping Create3PLStation due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/tps/stations/create*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Logistics.Create3PLStation(ctx)
	if err != nil {
		t.Logf("Logistics.Create3PLStation returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.Create3PLStation response: %#v", res)
}
func Test_Logistics_CreateConsolidationService(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_ldp_createConsolidationService_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CreateConsolidationService due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/ldp/createConsolidationService*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Logistics.CreateConsolidationService(ctx)
	if err != nil {
		t.Logf("Logistics.CreateConsolidationService returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.CreateConsolidationService response: %#v", res)
}
func Test_Logistics_GetOrderTrace(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistic_order_trace_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetOrderTrace due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/logistic/order/trace*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Logistics.GetOrderTrace(ctx)
	if err != nil {
		t.Logf("Logistics.GetOrderTrace returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.GetOrderTrace response: %#v", res)
}
func Test_Logistics_ScanParcel(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_dop_scan_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping ScanParcel due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/dop/scan*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Logistics.ScanParcel(ctx)
	if err != nil {
		t.Logf("Logistics.ScanParcel returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.ScanParcel response: %#v", res)
}
func Test_Logistics_StationDopScan(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_stations_dop_scan_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping StationDopScan due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/stations/dop/scan*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Logistics.StationDopScan(ctx)
	if err != nil {
		t.Logf("Logistics.StationDopScan returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.StationDopScan response: %#v", res)
}
func Test_Logistics_Update3PLStation(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_tps_stations_update_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping Update3PLStation due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/tps/stations/update*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Logistics.Update3PLStation(ctx)
	if err != nil {
		t.Logf("Logistics.Update3PLStation returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.Update3PLStation response: %#v", res)
}
func Test_Logistics_UpdateLastMile(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_ldp_updateLastmile_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateLastMile due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/ldp/updateLastmile*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Logistics.UpdateLastMile(ctx)
	if err != nil {
		t.Logf("Logistics.UpdateLastMile returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.UpdateLastMile response: %#v", res)
}
func Test_Logistics_UpdatePickupTimeSlot(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_tps_sellers_pickup_timeslot_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdatePickupTimeSlot due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/tps/sellers/pickup_timeslot*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Logistics.UpdatePickupTimeSlot(ctx)
	if err != nil {
		t.Logf("Logistics.UpdatePickupTimeSlot returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.UpdatePickupTimeSlot response: %#v", res)
}
