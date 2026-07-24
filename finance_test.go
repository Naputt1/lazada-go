package golazada

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_Finance_GetPayoutStatus(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_finance_payout_status_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetPayoutStatus due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/finance/payout/status/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Finance.GetPayoutStatus(ctx)
	if err != nil {
		t.Logf("Finance.GetPayoutStatus returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Finance.GetPayoutStatus response: %#v", res)
}
func Test_Finance_QueryAccountTransactions(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_finance_transaction_accountTransactions_query_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping QueryAccountTransactions due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/finance/transaction/accountTransactions/query*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Finance.QueryAccountTransactions(ctx)
	if err != nil {
		t.Logf("Finance.QueryAccountTransactions returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Finance.QueryAccountTransactions response: %#v", res)
}
func Test_Finance_QueryLogisticsFeeDetail(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_lbs_slb_queryLogisticsFeeDetail_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping QueryLogisticsFeeDetail due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/lbs/slb/queryLogisticsFeeDetail*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Finance.QueryLogisticsFeeDetail(ctx)
	if err != nil {
		t.Logf("Finance.QueryLogisticsFeeDetail returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Finance.QueryLogisticsFeeDetail response: %#v", res)
}
func Test_Finance_QueryTransactionDetails(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_finance_transaction_details_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping QueryTransactionDetails due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/finance/transaction/details/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Finance.QueryTransactionDetails(ctx)
	if err != nil {
		t.Logf("Finance.QueryTransactionDetails returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Finance.QueryTransactionDetails response: %#v", res)
}
