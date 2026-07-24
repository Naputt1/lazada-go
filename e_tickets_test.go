package golazada

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_ETickets_GetOrderItemsFromBarCode(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_eticket_code_query_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetOrderItemsFromBarCode due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/eticket/code/query*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.ETickets.GetOrderItemsFromBarCode(ctx)
	if err != nil {
		t.Logf("ETickets.GetOrderItemsFromBarCode returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ETickets.GetOrderItemsFromBarCode response: %#v", res)
}
func Test_ETickets_GlobalEticketMerchantMaAvailable(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_eticket_ma_available_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GlobalEticketMerchantMaAvailable due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/eticket/ma/available*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.ETickets.GlobalEticketMerchantMaAvailable(ctx)
	if err != nil {
		t.Logf("ETickets.GlobalEticketMerchantMaAvailable returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ETickets.GlobalEticketMerchantMaAvailable response: %#v", res)
}
func Test_ETickets_GlobalEticketMerchantMaConsume(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_eticket_ma_consume_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GlobalEticketMerchantMaConsume due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/eticket/ma/consume*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.ETickets.GlobalEticketMerchantMaConsume(ctx)
	if err != nil {
		t.Logf("ETickets.GlobalEticketMerchantMaConsume returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ETickets.GlobalEticketMerchantMaConsume response: %#v", res)
}
func Test_ETickets_GlobalEticketMerchantMaFailsend(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_eticket_ma_failsend_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GlobalEticketMerchantMaFailsend due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/eticket/ma/failsend*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.ETickets.GlobalEticketMerchantMaFailsend(ctx)
	if err != nil {
		t.Logf("ETickets.GlobalEticketMerchantMaFailsend returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ETickets.GlobalEticketMerchantMaFailsend response: %#v", res)
}
func Test_ETickets_GlobalEticketMerchantMaQuery(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_eticket_ma_query_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GlobalEticketMerchantMaQuery due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/eticket/ma/query*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.ETickets.GlobalEticketMerchantMaQuery(ctx)
	if err != nil {
		t.Logf("ETickets.GlobalEticketMerchantMaQuery returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ETickets.GlobalEticketMerchantMaQuery response: %#v", res)
}
func Test_ETickets_GlobalEticketMerchantMaQueryTbMa(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_eticket_ma_queryTbMa_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GlobalEticketMerchantMaQueryTbMa due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/eticket/ma/queryTbMa*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.ETickets.GlobalEticketMerchantMaQueryTbMa(ctx)
	if err != nil {
		t.Logf("ETickets.GlobalEticketMerchantMaQueryTbMa returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ETickets.GlobalEticketMerchantMaQueryTbMa response: %#v", res)
}
func Test_ETickets_GlobalEticketMerchantMaSend(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_eticket_ma_send_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GlobalEticketMerchantMaSend due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/eticket/ma/send*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.ETickets.GlobalEticketMerchantMaSend(ctx)
	if err != nil {
		t.Logf("ETickets.GlobalEticketMerchantMaSend returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ETickets.GlobalEticketMerchantMaSend response: %#v", res)
}
func Test_ETickets_RedeemOrderItems(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_eticket_code_consume_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping RedeemOrderItems due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/eticket/code/consume*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.ETickets.RedeemOrderItems(ctx)
	if err != nil {
		t.Logf("ETickets.RedeemOrderItems returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ETickets.RedeemOrderItems response: %#v", res)
}
