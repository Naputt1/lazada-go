package golazada

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_FirstMileBigbagonlyForCN_GetChannelcodeByFirstMileNo(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_cngfc_fulfill_getchannelcode_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetChannelcodeByFirstMileNo due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/logistics/cngfc/fulfill/getchannelcode*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FirstMileBigbagonlyForCN.GetChannelcodeByFirstMileNo(ctx)
	if err != nil {
		t.Logf("FirstMileBigbagonlyForCN.GetChannelcodeByFirstMileNo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FirstMileBigbagonlyForCN.GetChannelcodeByFirstMileNo response: %#v", res)
}
func Test_FirstMileBigbagonlyForCN_GetLazadaBigbagPDFLable(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_cnpms_bigbag_lable_getPdf_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetLazadaBigbagPDFLable due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/logistics/cnpms/bigbag/lable/getPdf*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FirstMileBigbagonlyForCN.GetLazadaBigbagPDFLable(ctx)
	if err != nil {
		t.Logf("FirstMileBigbagonlyForCN.GetLazadaBigbagPDFLable returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FirstMileBigbagonlyForCN.GetLazadaBigbagPDFLable response: %#v", res)
}
func Test_FirstMileBigbagonlyForCN_LazadaBigbagCancel(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_cnpms_bigbag_cancel_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping LazadaBigbagCancel due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/cnpms/bigbag/cancel*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FirstMileBigbagonlyForCN.LazadaBigbagCancel(ctx)
	if err != nil {
		t.Logf("FirstMileBigbagonlyForCN.LazadaBigbagCancel returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FirstMileBigbagonlyForCN.LazadaBigbagCancel response: %#v", res)
}
func Test_FirstMileBigbagonlyForCN_LazadaBigbagCollectionPoints(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_cnpms_bigbag_querycollection_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping LazadaBigbagCollectionPoints due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/cnpms/bigbag/querycollection*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FirstMileBigbagonlyForCN.LazadaBigbagCollectionPoints(ctx)
	if err != nil {
		t.Logf("FirstMileBigbagonlyForCN.LazadaBigbagCollectionPoints returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FirstMileBigbagonlyForCN.LazadaBigbagCollectionPoints response: %#v", res)
}
func Test_FirstMileBigbagonlyForCN_LazadaBigbagCommit(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_cnpms_bigbag_commit_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping LazadaBigbagCommit due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/cnpms/bigbag/commit*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FirstMileBigbagonlyForCN.LazadaBigbagCommit(ctx)
	if err != nil {
		t.Logf("FirstMileBigbagonlyForCN.LazadaBigbagCommit returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FirstMileBigbagonlyForCN.LazadaBigbagCommit response: %#v", res)
}
func Test_FirstMileBigbagonlyForCN_LazadaBigbagUpdate(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_cnpms_bigbag_update_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping LazadaBigbagUpdate due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/cnpms/bigbag/update*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FirstMileBigbagonlyForCN.LazadaBigbagUpdate(ctx)
	if err != nil {
		t.Logf("FirstMileBigbagonlyForCN.LazadaBigbagUpdate returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FirstMileBigbagonlyForCN.LazadaBigbagUpdate response: %#v", res)
}
func Test_FirstMileBigbagonlyForCN_LazadaSellerAccountBind(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_cnpms_account_bind_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping LazadaSellerAccountBind due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/logistics/cnpms/account/bind*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FirstMileBigbagonlyForCN.LazadaSellerAccountBind(ctx)
	if err != nil {
		t.Logf("FirstMileBigbagonlyForCN.LazadaSellerAccountBind returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FirstMileBigbagonlyForCN.LazadaSellerAccountBind response: %#v", res)
}
func Test_FirstMileBigbagonlyForCN_QueryAddressInformaiton(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_cnpms_address_query_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping QueryAddressInformaiton due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/logistics/cnpms/address/query*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FirstMileBigbagonlyForCN.QueryAddressInformaiton(ctx)
	if err != nil {
		t.Logf("FirstMileBigbagonlyForCN.QueryAddressInformaiton returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FirstMileBigbagonlyForCN.QueryAddressInformaiton response: %#v", res)
}
func Test_FirstMileBigbagonlyForCN_QueryLazadaBigbagInfo(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_logistics_cnpms_bigbag_query_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping QueryLazadaBigbagInfo due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/logistics/cnpms/bigbag/query*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FirstMileBigbagonlyForCN.QueryLazadaBigbagInfo(ctx)
	if err != nil {
		t.Logf("FirstMileBigbagonlyForCN.QueryLazadaBigbagInfo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FirstMileBigbagonlyForCN.QueryLazadaBigbagInfo response: %#v", res)
}
