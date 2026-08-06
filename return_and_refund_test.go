package golazada

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_ReturnAndRefund_GetReverseOrderDetail(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_order_reverse_return_detail_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetReverseOrderDetail due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/order/reverse/return/detail/list*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.ReturnAndRefund.GetReverseOrderDetail(ctx)
	if err != nil {
		t.Logf("ReturnAndRefund.GetReverseOrderDetail returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ReturnAndRefund.GetReverseOrderDetail response: %#v", res)
}
func Test_ReturnAndRefund_GetReverseOrderHistoryList(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_order_reverse_return_history_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetReverseOrderHistoryList due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/order/reverse/return/history/list*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.ReturnAndRefund.GetReverseOrderHistoryList(ctx)
	if err != nil {
		t.Logf("ReturnAndRefund.GetReverseOrderHistoryList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ReturnAndRefund.GetReverseOrderHistoryList response: %#v", res)
}
func Test_ReturnAndRefund_GetReverseOrderReasonList(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_order_reverse_reason_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetReverseOrderReasonList due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/order/reverse/reason/list*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.ReturnAndRefund.GetReverseOrderReasonList(ctx)
	if err != nil {
		t.Logf("ReturnAndRefund.GetReverseOrderReasonList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ReturnAndRefund.GetReverseOrderReasonList response: %#v", res)
}
func Test_ReturnAndRefund_GetReverseOrdersForSeller(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_reverse_getreverseordersforseller_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetReverseOrdersForSeller due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/reverse/getreverseordersforseller*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	var req GetReverseOrdersForSellerRequest
	ctx := context.Background()
	res, err := client.ReturnAndRefund.GetReverseOrdersForSeller(ctx, req)
	if err != nil {
		t.Logf("ReturnAndRefund.GetReverseOrdersForSeller returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ReturnAndRefund.GetReverseOrdersForSeller response: %#v", res)
}
func Test_ReturnAndRefund_InitReverseOrderCancel(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_order_reverse_cancel_create_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping InitReverseOrderCancel due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/order/reverse/cancel/create*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	var req InitReverseOrderCancelRequest
	ctx := context.Background()
	res, err := client.ReturnAndRefund.InitReverseOrderCancel(ctx, req)
	if err != nil {
		t.Logf("ReturnAndRefund.InitReverseOrderCancel returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ReturnAndRefund.InitReverseOrderCancel response: %#v", res)
}
func Test_ReturnAndRefund_InitReverseOrderCancelDecide(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_order_reverse_cancel_seller_decide_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping InitReverseOrderCancelDecide due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/order/reverse/cancel/seller/decide*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.ReturnAndRefund.InitReverseOrderCancelDecide(ctx)
	if err != nil {
		t.Logf("ReturnAndRefund.InitReverseOrderCancelDecide returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ReturnAndRefund.InitReverseOrderCancelDecide response: %#v", res)
}
func Test_ReturnAndRefund_ReverseOrderOnlyRefundDecide(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_order_reverse_onlyrefund_seller_decide_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping ReverseOrderOnlyRefundDecide due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/order/reverse/onlyrefund/seller/decide*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	var req ReverseOrderOnlyRefundDecideRequest
	ctx := context.Background()
	res, err := client.ReturnAndRefund.ReverseOrderOnlyRefundDecide(ctx, req)
	if err != nil {
		t.Logf("ReturnAndRefund.ReverseOrderOnlyRefundDecide returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ReturnAndRefund.ReverseOrderOnlyRefundDecide response: %#v", res)
}
func Test_ReturnAndRefund_ReverseOrderReturnUpdate(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_order_reverse_return_update_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping ReverseOrderReturnUpdate due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/order/reverse/return/update*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	var req ReverseOrderReturnUpdateRequest
	ctx := context.Background()
	res, err := client.ReturnAndRefund.ReverseOrderReturnUpdate(ctx, req)
	if err != nil {
		t.Logf("ReturnAndRefund.ReverseOrderReturnUpdate returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ReturnAndRefund.ReverseOrderReturnUpdate response: %#v", res)
}
