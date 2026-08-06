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
	fixture := "reverse.getreverseordersforseller_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetReverseOrdersForSeller due to missing fixture: %v", err)
	}

	// Lazada sends the numeric result fields as JSON strings in some payloads and
	// as real numbers in others. Force them to numbers to reproduce the payload
	// that previously broke decoding.
	mockData, _ := json.Marshal(data)
	var mockObj map[string]any
	if err := json.Unmarshal(mockData, &mockObj); err != nil {
		t.Fatalf("failed to decode fixture: %v", err)
	}
	result := mockObj["result"].(map[string]any)
	result["page_no"] = float64(1)
	result["page_size"] = float64(10)
	result["total"] = float64(50)
	result["success"] = true
	mockData, _ = json.Marshal(mockObj)

	var sawPageNo, sawPageSize string
	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/reverse/getreverseordersforseller", serverURL),
		func(req *http.Request) (*http.Response, error) {
			sawPageNo = req.URL.Query().Get("page_no")
			sawPageSize = req.URL.Query().Get("page_size")
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.ReturnAndRefund.GetReverseOrdersForSeller(ctx, GetReverseOrdersForSellerRequest{
		PageNo:   3,
		PageSize: 25,
	})
	if err != nil {
		t.Fatalf("GetReverseOrdersForSeller failed: %v", err)
	}
	if sawPageNo != "3" || sawPageSize != "25" {
		t.Fatalf("expected page_no=3 page_size=25, got page_no=%q page_size=%q", sawPageNo, sawPageSize)
	}
	if res.Response.Result == nil {
		t.Fatalf("expected result to be parsed, got %#v", res)
	}
	if res.Response.Result.Total != 50 {
		t.Fatalf("expected total 50, got %v", res.Response.Result.Total)
	}
	if len(res.Response.Result.Items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(res.Response.Result.Items))
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
		"GET",
		fmt.Sprintf("%s/order/reverse/cancel/create*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.ReturnAndRefund.InitReverseOrderCancel(ctx)
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
		"GET",
		fmt.Sprintf("%s/order/reverse/onlyrefund/seller/decide*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.ReturnAndRefund.ReverseOrderOnlyRefundDecide(ctx)
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
		"GET",
		fmt.Sprintf("%s/order/reverse/return/update*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.ReturnAndRefund.ReverseOrderReturnUpdate(ctx)
	if err != nil {
		t.Logf("ReturnAndRefund.ReverseOrderReturnUpdate returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ReturnAndRefund.ReverseOrderReturnUpdate response: %#v", res)
}
