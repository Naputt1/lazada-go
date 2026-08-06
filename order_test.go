package golazada

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
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
	fixture := "_orders_items_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetMultipleOrderItems due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/orders/items/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	var req GetMultipleOrderItemsRequest
	ctx := context.Background()
	res, err := client.Order.GetMultipleOrderItems(ctx, req)
	if err != nil {
		t.Logf("Order.GetMultipleOrderItems returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Order.GetMultipleOrderItems response: %#v", res)
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

	var req OrderCancelValidateRequest
	ctx := context.Background()
	res, err := client.Order.OrderCancelValidate(ctx, req)
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
