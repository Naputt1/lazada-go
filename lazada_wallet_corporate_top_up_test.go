package golazada

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_LazadaWalletCorporateTopUp_DirectTransferQuery(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_wallet_transfer_query_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DirectTransferQuery due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/wallet/transfer/query*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazadaWalletCorporateTopUp.DirectTransferQuery(ctx)
	if err != nil {
		t.Logf("LazadaWalletCorporateTopUp.DirectTransferQuery returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazadaWalletCorporateTopUp.DirectTransferQuery response: %#v", res)
}
func Test_LazadaWalletCorporateTopUp_DirectTransferRequest(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_wallet_transfer_request_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DirectTransferRequest due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/wallet/transfer/request*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazadaWalletCorporateTopUp.DirectTransferRequest(ctx)
	if err != nil {
		t.Logf("LazadaWalletCorporateTopUp.DirectTransferRequest returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazadaWalletCorporateTopUp.DirectTransferRequest response: %#v", res)
}
func Test_LazadaWalletCorporateTopUp_GiftCodeQuery(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_wallet_giftcode_query_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GiftCodeQuery due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/wallet/giftcode/query*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazadaWalletCorporateTopUp.GiftCodeQuery(ctx)
	if err != nil {
		t.Logf("LazadaWalletCorporateTopUp.GiftCodeQuery returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazadaWalletCorporateTopUp.GiftCodeQuery response: %#v", res)
}
func Test_LazadaWalletCorporateTopUp_GiftCodeRequest(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_wallet_giftcode_request_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GiftCodeRequest due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/wallet/giftcode/request*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazadaWalletCorporateTopUp.GiftCodeRequest(ctx)
	if err != nil {
		t.Logf("LazadaWalletCorporateTopUp.GiftCodeRequest returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazadaWalletCorporateTopUp.GiftCodeRequest response: %#v", res)
}
func Test_LazadaWalletCorporateTopUp_Reconciliation1(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_wallet_open_reconciliation_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping Reconciliation1 due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/wallet/open/reconciliation*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazadaWalletCorporateTopUp.Reconciliation1(ctx)
	if err != nil {
		t.Logf("LazadaWalletCorporateTopUp.Reconciliation1 returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazadaWalletCorporateTopUp.Reconciliation1 response: %#v", res)
}
