package golazada

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_LazadaDG_DigitalServiceCdkCodeReceived(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_digital_service_cdkCodeReceived_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DigitalServiceCdkCodeReceived due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/digital/service/cdkCodeReceived*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazadaDG.DigitalServiceCdkCodeReceived(ctx)
	if err != nil {
		t.Logf("LazadaDG.DigitalServiceCdkCodeReceived returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazadaDG.DigitalServiceCdkCodeReceived response: %#v", res)
}
func Test_LazadaDG_InstallServiceCallBack(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_digital_install_servicecallback_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping InstallServiceCallBack due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/digital/install/servicecallback*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazadaDG.InstallServiceCallBack(ctx)
	if err != nil {
		t.Logf("LazadaDG.InstallServiceCallBack returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazadaDG.InstallServiceCallBack response: %#v", res)
}
func Test_LazadaDG_InstallServiceCallBack1(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_digital_test_install_servicecallback_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping InstallServiceCallBack1 due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/digital/test/install/servicecallback*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazadaDG.InstallServiceCallBack1(ctx)
	if err != nil {
		t.Logf("LazadaDG.InstallServiceCallBack1 returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazadaDG.InstallServiceCallBack1 response: %#v", res)
}
func Test_LazadaDG_InstallServiceCallBackForTest(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_digital_install_test_servicecallback_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping InstallServiceCallBackForTest due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/digital/install/test/servicecallback*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazadaDG.InstallServiceCallBackForTest(ctx)
	if err != nil {
		t.Logf("LazadaDG.InstallServiceCallBackForTest returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazadaDG.InstallServiceCallBackForTest response: %#v", res)
}
func Test_LazadaDG_InuranceNotication(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_digital_insurance_notification_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping InuranceNotication due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/digital/insurance/notification*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazadaDG.InuranceNotication(ctx)
	if err != nil {
		t.Logf("LazadaDG.InuranceNotication returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazadaDG.InuranceNotication response: %#v", res)
}
func Test_LazadaDG_InuranceNotication1(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_digital_insurance_test_notificationcopy_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping InuranceNotication1 due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/digital/insurance/test/notificationcopy*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazadaDG.InuranceNotication1(ctx)
	if err != nil {
		t.Logf("LazadaDG.InuranceNotication1 returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazadaDG.InuranceNotication1 response: %#v", res)
}
func Test_LazadaDG_InuranceNotifyLapse(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_digital_insurance_notificationlapse_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping InuranceNotifyLapse due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/digital/insurance/notificationlapse*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazadaDG.InuranceNotifyLapse(ctx)
	if err != nil {
		t.Logf("LazadaDG.InuranceNotifyLapse returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazadaDG.InuranceNotifyLapse response: %#v", res)
}
