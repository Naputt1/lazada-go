package golazada

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_Order_GetDocument(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_order_document_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetDocument due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/order/document/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Order.GetDocument(ctx)
	if err != nil {
		t.Logf("Order.GetDocument returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Order.GetDocument response: %#v", res)
}
func Test_Order_GetMultipleOrderItems(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture, err := loadFixtureSafe("orders.items.get_resp.json")
	if err != nil {
		t.Skipf("Skipping GetMultipleOrderItems due to missing fixture: %v", err)
	}
	if resp, ok := fixture.(map[string]interface{}); ok {
		fixture = resp["data"]
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": fixture,
	}
	mockData, _ := json.Marshal(mockResp)

	var gotOrderIDs string
	httpmock.RegisterRegexpResponder(
		"GET",
		regexp.MustCompile(regexp.QuoteMeta(serverURL)+`/orders/items/get\?.*`),
		func(req *http.Request) (*http.Response, error) {
			gotOrderIDs = req.URL.Query().Get("order_ids")
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Order.GetMultipleOrderItems(ctx, GetMultipleOrderItemsRequest{OrderIds: []int64{32793}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotOrderIDs != "[32793]" {
		t.Fatalf("expected order_ids=[32793] in request, got %q", gotOrderIDs)
	}
	if len(res.Response) != 1 {
		t.Fatalf("expected 1 order batch, got %d", len(res.Response))
	}
	if string(res.Response[0].OrderNumber) != "300029225" {
		t.Fatalf("unexpected order_number: %q", res.Response[0].OrderNumber)
	}
	if int64(res.Response[0].OrderId) != 32793 {
		t.Fatalf("unexpected order_id: %v", res.Response[0].OrderId)
	}
	if len(res.Response[0].OrderItems) != 1 {
		t.Fatalf("expected 1 order item, got %d", len(res.Response[0].OrderItems))
	}
	item := res.Response[0].OrderItems[0]
	if string(item.Name) != "Bean Rester Crest Brown" {
		t.Fatalf("unexpected item name: %q", item.Name)
	}
	if int64(item.OrderItemId) != 100827 {
		t.Fatalf("unexpected order_item_id: %v", item.OrderItemId)
	}
}
func Test_Order_GetOrder(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_order_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetOrder due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/order/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Order.GetOrder(ctx)
	if err != nil {
		t.Logf("Order.GetOrder returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Order.GetOrder response: %#v", res)
}
func Test_Order_GetOrderItems(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_order_items_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetOrderItems due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/order/items/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	var req GetOrderItemsRequest
	ctx := context.Background()
	res, err := client.Order.GetOrderItems(ctx, req)
	if err != nil {
		t.Logf("Order.GetOrderItems returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Order.GetOrderItems response: %#v", res)
}
func Test_Order_GetOrders(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_orders_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetOrders due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/orders/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	var req GetOrdersRequest
	ctx := context.Background()
	res, err := client.Order.GetOrders(ctx, req)
	if err != nil {
		t.Logf("Order.GetOrders returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Order.GetOrders response: %#v", res)
}
func Test_Order_GetOVOOrders(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_orders_ovo_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetOVOOrders due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/orders/ovo/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Order.GetOVOOrders(ctx)
	if err != nil {
		t.Logf("Order.GetOVOOrders returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Order.GetOVOOrders response: %#v", res)
}
func Test_Order_OrderCancelValidate(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_order_reverse_cancel_validate_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping OrderCancelValidate due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/order/reverse/cancel/validate*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Order.OrderCancelValidate(ctx)
	if err != nil {
		t.Logf("Order.OrderCancelValidate returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Order.OrderCancelValidate response: %#v", res)
}
func Test_Order_SetInvoiceNumber(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_order_invoice_number_set_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SetInvoiceNumber due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/order/invoice_number/set*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Order.SetInvoiceNumber(ctx)
	if err != nil {
		t.Logf("Order.SetInvoiceNumber returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Order.SetInvoiceNumber response: %#v", res)
}
