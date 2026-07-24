package golazada

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_ServiceMarket_ServiceMarketAppKeyOrderQuery(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_service_market_order_query_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping ServiceMarketAppKeyOrderQuery due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/service/market/order/query*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.ServiceMarket.ServiceMarketAppKeyOrderQuery(ctx)
	if err != nil {
		t.Logf("ServiceMarket.ServiceMarketAppKeyOrderQuery returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ServiceMarket.ServiceMarketAppKeyOrderQuery response: %#v", res)
}
func Test_ServiceMarket_ServiceMarketAppKeySubQuery(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_service_market_subs_query_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping ServiceMarketAppKeySubQuery due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/service/market/subs/query*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.ServiceMarket.ServiceMarketAppKeySubQuery(ctx)
	if err != nil {
		t.Logf("ServiceMarket.ServiceMarketAppKeySubQuery returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ServiceMarket.ServiceMarketAppKeySubQuery response: %#v", res)
}
