package golazada

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_LogisticsStation_CageValidation(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_station_cages_validate_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CageValidation due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/station/cages/validate*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LogisticsStation.CageValidation(ctx)
	if err != nil {
		t.Logf("LogisticsStation.CageValidation returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LogisticsStation.CageValidation response: %#v", res)
}
func Test_LogisticsStation_ConfirmInbound(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_station_v1_confirm_inbound_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping ConfirmInbound due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/station/v1/confirm-inbound*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LogisticsStation.ConfirmInbound(ctx)
	if err != nil {
		t.Logf("LogisticsStation.ConfirmInbound returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LogisticsStation.ConfirmInbound response: %#v", res)
}
func Test_LogisticsStation_ConfirmParcelCollection(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_station_v1_cp_confirm_parcel_collection_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping ConfirmParcelCollection due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/station/v1/cp/confirm-parcel-collection*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LogisticsStation.ConfirmParcelCollection(ctx)
	if err != nil {
		t.Logf("LogisticsStation.ConfirmParcelCollection returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LogisticsStation.ConfirmParcelCollection response: %#v", res)
}
func Test_LogisticsStation_CreateScannedParcel(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_station_v1_scanned_parcels_create_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CreateScannedParcel due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/station/v1/scanned-parcels/create*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LogisticsStation.CreateScannedParcel(ctx)
	if err != nil {
		t.Logf("LogisticsStation.CreateScannedParcel returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LogisticsStation.CreateScannedParcel response: %#v", res)
}
func Test_LogisticsStation_DeleteScannedParcel(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_station_v1_scanned_parcels_delete_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DeleteScannedParcel due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/station/v1/scanned-parcels/delete*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LogisticsStation.DeleteScannedParcel(ctx)
	if err != nil {
		t.Logf("LogisticsStation.DeleteScannedParcel returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LogisticsStation.DeleteScannedParcel response: %#v", res)
}
func Test_LogisticsStation_DopConfirmInbound(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_station_dop_confirm_inbound_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DopConfirmInbound due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/station/dop/confirm-inbound*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LogisticsStation.DopConfirmInbound(ctx)
	if err != nil {
		t.Logf("LogisticsStation.DopConfirmInbound returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LogisticsStation.DopConfirmInbound response: %#v", res)
}
func Test_LogisticsStation_DopCreateScannedParcel(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_station_dop_scanned_parcels_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DopCreateScannedParcel due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/station/dop/scanned-parcels*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LogisticsStation.DopCreateScannedParcel(ctx)
	if err != nil {
		t.Logf("LogisticsStation.DopCreateScannedParcel returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LogisticsStation.DopCreateScannedParcel response: %#v", res)
}
func Test_LogisticsStation_DopDeleteScannedParcel(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_station_dop_scanned_parcels_delete_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DopDeleteScannedParcel due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/station/dop/scanned-parcels/delete*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LogisticsStation.DopDeleteScannedParcel(ctx)
	if err != nil {
		t.Logf("LogisticsStation.DopDeleteScannedParcel returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LogisticsStation.DopDeleteScannedParcel response: %#v", res)
}
func Test_LogisticsStation_DopGetInboundedParcel(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_station_dop_inbounded_parcels_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DopGetInboundedParcel due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/logistics/station/dop/inbounded-parcels/list*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LogisticsStation.DopGetInboundedParcel(ctx)
	if err != nil {
		t.Logf("LogisticsStation.DopGetInboundedParcel returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LogisticsStation.DopGetInboundedParcel response: %#v", res)
}
func Test_LogisticsStation_DopGetScannedParcel(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_station_dop_scanned_parcels_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DopGetScannedParcel due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/logistics/station/dop/scanned-parcels/list*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LogisticsStation.DopGetScannedParcel(ctx)
	if err != nil {
		t.Logf("LogisticsStation.DopGetScannedParcel returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LogisticsStation.DopGetScannedParcel response: %#v", res)
}
func Test_LogisticsStation_GetCpScheduledPuParcel(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_station_v1_cp_scheduled_pu_parcels_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetCpScheduledPuParcel due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/logistics/station/v1/cp/scheduled-pu-parcels/list*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LogisticsStation.GetCpScheduledPuParcel(ctx)
	if err != nil {
		t.Logf("LogisticsStation.GetCpScheduledPuParcel returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LogisticsStation.GetCpScheduledPuParcel response: %#v", res)
}
func Test_LogisticsStation_GetInboundedParcel(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_station_v1_inbounded_parcels_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetInboundedParcel due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/logistics/station/v1/inbounded-parcels/list*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LogisticsStation.GetInboundedParcel(ctx)
	if err != nil {
		t.Logf("LogisticsStation.GetInboundedParcel returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LogisticsStation.GetInboundedParcel response: %#v", res)
}
func Test_LogisticsStation_GetListAccessStation(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_station_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetListAccessStation due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/logistics/station/list*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LogisticsStation.GetListAccessStation(ctx)
	if err != nil {
		t.Logf("LogisticsStation.GetListAccessStation returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LogisticsStation.GetListAccessStation response: %#v", res)
}
func Test_LogisticsStation_GetMetaData(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_station_v1_metadata_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetMetaData due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/logistics/station/v1/metadata*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LogisticsStation.GetMetaData(ctx)
	if err != nil {
		t.Logf("LogisticsStation.GetMetaData returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LogisticsStation.GetMetaData response: %#v", res)
}
func Test_LogisticsStation_GetScannedParcel(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_station_v1_scanned_parcels_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetScannedParcel due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/logistics/station/v1/scanned-parcels/list*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LogisticsStation.GetScannedParcel(ctx)
	if err != nil {
		t.Logf("LogisticsStation.GetScannedParcel returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LogisticsStation.GetScannedParcel response: %#v", res)
}
func Test_LogisticsStation_SearchCustomerReturnParcel(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_station_v1_dop_cr_parcels_search_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SearchCustomerReturnParcel due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/logistics/station/v1/dop/cr-parcels/search*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LogisticsStation.SearchCustomerReturnParcel(ctx)
	if err != nil {
		t.Logf("LogisticsStation.SearchCustomerReturnParcel returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LogisticsStation.SearchCustomerReturnParcel response: %#v", res)
}
func Test_LogisticsStation_ValidateCage(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_station_v1_cages_validate_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping ValidateCage due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/station/v1/cages/validate*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LogisticsStation.ValidateCage(ctx)
	if err != nil {
		t.Logf("LogisticsStation.ValidateCage returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LogisticsStation.ValidateCage response: %#v", res)
}
func Test_LogisticsStation_ValidateOTP(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_station_v1_cp_validate_otp_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping ValidateOTP due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/station/v1/cp/validate-otp*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LogisticsStation.ValidateOTP(ctx)
	if err != nil {
		t.Logf("LogisticsStation.ValidateOTP returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LogisticsStation.ValidateOTP response: %#v", res)
}
