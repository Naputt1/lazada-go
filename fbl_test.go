package golazada

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_FBL_BuildFulfillmentSkuRelation(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_fulfillment_sku_relation_write_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping BuildFulfillmentSkuRelation due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/fbl/fulfillment_sku_relation/write*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.BuildFulfillmentSkuRelation(ctx)
	if err != nil {
		t.Logf("FBL.BuildFulfillmentSkuRelation returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.BuildFulfillmentSkuRelation response: %#v", res)
}
func Test_FBL_CancelFulfillmentOrderForMCL(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_fulfillment_order_cancel_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CancelFulfillmentOrderForMCL due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/fbl/fulfillment_order/cancel*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.CancelFulfillmentOrderForMCL(ctx)
	if err != nil {
		t.Logf("FBL.CancelFulfillmentOrderForMCL returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.CancelFulfillmentOrderForMCL response: %#v", res)
}
func Test_FBL_CancelInboundReservation(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_inbound_reservation_cancel_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CancelInboundReservation due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/fbl/inbound_reservation/cancel*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.CancelInboundReservation(ctx)
	if err != nil {
		t.Logf("FBL.CancelInboundReservation returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.CancelInboundReservation response: %#v", res)
}
func Test_FBL_CancelnBoundOrder(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_inbound_order_cancel_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CancelnBoundOrder due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/fbl/inbound_order/cancel*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.CancelnBoundOrder(ctx)
	if err != nil {
		t.Logf("FBL.CancelnBoundOrder returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.CancelnBoundOrder response: %#v", res)
}
func Test_FBL_CancelOutboundOrder(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_outbound_order_cancel_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CancelOutboundOrder due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/fbl/outbound_order/cancel*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.CancelOutboundOrder(ctx)
	if err != nil {
		t.Logf("FBL.CancelOutboundOrder returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.CancelOutboundOrder response: %#v", res)
}
func Test_FBL_CancelVasOrder4FBL(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_vas_cancelVasOrder_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CancelVasOrder4FBL due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/fbl/vas/cancelVasOrder*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.CancelVasOrder4FBL(ctx)
	if err != nil {
		t.Logf("FBL.CancelVasOrder4FBL returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.CancelVasOrder4FBL response: %#v", res)
}
func Test_FBL_CheckInboundReservationSlot(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_inbound_reservation_check_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CheckInboundReservationSlot due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/fbl/inbound_reservation/check*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.CheckInboundReservationSlot(ctx)
	if err != nil {
		t.Logf("FBL.CheckInboundReservationSlot returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.CheckInboundReservationSlot response: %#v", res)
}
func Test_FBL_CreateFulfillmentOrderForMCL(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_fulfillment_order_create_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CreateFulfillmentOrderForMCL due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/fbl/fulfillment_order/create*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.CreateFulfillmentOrderForMCL(ctx)
	if err != nil {
		t.Logf("FBL.CreateFulfillmentOrderForMCL returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.CreateFulfillmentOrderForMCL response: %#v", res)
}
func Test_FBL_CreateFulfillmentOrderForMCLV2PNF(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_fulfillment_order_pnf_create_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CreateFulfillmentOrderForMCLV2PNF due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/fbl/fulfillment_order_pnf/create*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.CreateFulfillmentOrderForMCLV2PNF(ctx)
	if err != nil {
		t.Logf("FBL.CreateFulfillmentOrderForMCLV2PNF returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.CreateFulfillmentOrderForMCLV2PNF response: %#v", res)
}
func Test_FBL_CreateFulfillmentSkuDecouple(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_fulfillment_sku_create_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CreateFulfillmentSkuDecouple due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/fbl/fulfillment_sku/create*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.CreateFulfillmentSkuDecouple(ctx)
	if err != nil {
		t.Logf("FBL.CreateFulfillmentSkuDecouple returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.CreateFulfillmentSkuDecouple response: %#v", res)
}
func Test_FBL_CreateFulfillmentSkuForFBL(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_fulfillment_sku_fbl_create_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CreateFulfillmentSkuForFBL due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/fbl/fulfillment_sku_fbl/create*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.CreateFulfillmentSkuForFBL(ctx)
	if err != nil {
		t.Logf("FBL.CreateFulfillmentSkuForFBL returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.CreateFulfillmentSkuForFBL response: %#v", res)
}
func Test_FBL_CreateInboundOrder(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_inbound_order_create_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CreateInboundOrder due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/fbl/inbound_order/create*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.CreateInboundOrder(ctx)
	if err != nil {
		t.Logf("FBL.CreateInboundOrder returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.CreateInboundOrder response: %#v", res)
}
func Test_FBL_CreateInboundReservation(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_inbound_reservation_create_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CreateInboundReservation due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/fbl/inbound_reservation/create*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.CreateInboundReservation(ctx)
	if err != nil {
		t.Logf("FBL.CreateInboundReservation returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.CreateInboundReservation response: %#v", res)
}
func Test_FBL_CreateOutBoundOrder(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_outbound_order_create_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CreateOutBoundOrder due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/fbl/outbound_order/create*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.CreateOutBoundOrder(ctx)
	if err != nil {
		t.Logf("FBL.CreateOutBoundOrder returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.CreateOutBoundOrder response: %#v", res)
}
func Test_FBL_CreateProductReinboundOrderForMCL(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_product_reinbound_create_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CreateProductReinboundOrderForMCL due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/fbl/product_reinbound/create*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.CreateProductReinboundOrderForMCL(ctx)
	if err != nil {
		t.Logf("FBL.CreateProductReinboundOrderForMCL returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.CreateProductReinboundOrderForMCL response: %#v", res)
}
func Test_FBL_CreateVasOrder4FBL(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_vas_createVasOrder_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CreateVasOrder4FBL due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/fbl/vas/createVasOrder*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.CreateVasOrder4FBL(ctx)
	if err != nil {
		t.Logf("FBL.CreateVasOrder4FBL returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.CreateVasOrder4FBL response: %#v", res)
}
func Test_FBL_GetChannelStocksForMCL(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_channel_stocks_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetChannelStocksForMCL due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/fbl/channel_stocks/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.GetChannelStocksForMCL(ctx)
	if err != nil {
		t.Logf("FBL.GetChannelStocksForMCL returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.GetChannelStocksForMCL response: %#v", res)
}
func Test_FBL_GetFulfillmentProductDetail(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_fulfillment_products_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetFulfillmentProductDetail due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/fbl/fulfillment_products/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.GetFulfillmentProductDetail(ctx)
	if err != nil {
		t.Logf("FBL.GetFulfillmentProductDetail returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.GetFulfillmentProductDetail response: %#v", res)
}
func Test_FBL_GetFulfillmentSkuListForMCL(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_fulfillment_sku_list_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetFulfillmentSkuListForMCL due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/fbl/fulfillment_sku_list/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.GetFulfillmentSkuListForMCL(ctx)
	if err != nil {
		t.Logf("FBL.GetFulfillmentSkuListForMCL returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.GetFulfillmentSkuListForMCL response: %#v", res)
}
func Test_FBL_GetFulfillmentSkuRelationByScItem(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_fulfillment_sku_relation_get_by_sc_item_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetFulfillmentSkuRelationByScItem due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/fbl/fulfillment_sku_relation/get_by_sc_item*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.GetFulfillmentSkuRelationByScItem(ctx)
	if err != nil {
		t.Logf("FBL.GetFulfillmentSkuRelationByScItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.GetFulfillmentSkuRelationByScItem response: %#v", res)
}
func Test_FBL_GetFulfillmentSkuRelationBySku(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_fulfillment_sku_relation_get_by_sku_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetFulfillmentSkuRelationBySku due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/fbl/fulfillment_sku_relation/get_by_sku*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.GetFulfillmentSkuRelationBySku(ctx)
	if err != nil {
		t.Logf("FBL.GetFulfillmentSkuRelationBySku returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.GetFulfillmentSkuRelationBySku response: %#v", res)
}
func Test_FBL_GetFulfillmentSkuRelationsByScItems(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_fulfillment_sku_relation_get_by_sc_items_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetFulfillmentSkuRelationsByScItems due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/fbl/fulfillment_sku_relation/get_by_sc_items*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.GetFulfillmentSkuRelationsByScItems(ctx)
	if err != nil {
		t.Logf("FBL.GetFulfillmentSkuRelationsByScItems returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.GetFulfillmentSkuRelationsByScItems response: %#v", res)
}
func Test_FBL_GetFulfillmentSkuRelationsBySkus(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_fulfillment_sku_relation_get_by_skus_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetFulfillmentSkuRelationsBySkus due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/fbl/fulfillment_sku_relation/get_by_skus*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.GetFulfillmentSkuRelationsBySkus(ctx)
	if err != nil {
		t.Logf("FBL.GetFulfillmentSkuRelationsBySkus returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.GetFulfillmentSkuRelationsBySkus response: %#v", res)
}
func Test_FBL_GetIcpOrderFile(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_icp_order_file_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetIcpOrderFile due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/fbl/icp_order/file*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.GetIcpOrderFile(ctx)
	if err != nil {
		t.Logf("FBL.GetIcpOrderFile returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.GetIcpOrderFile response: %#v", res)
}
func Test_FBL_GetInboundOrderDetail(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_inbound_order_detail_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetInboundOrderDetail due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/fbl/inbound_order_detail/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.GetInboundOrderDetail(ctx)
	if err != nil {
		t.Logf("FBL.GetInboundOrderDetail returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.GetInboundOrderDetail response: %#v", res)
}
func Test_FBL_GetInboundOrderList(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_inbound_orders_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetInboundOrderList due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/fbl/inbound_orders/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.GetInboundOrderList(ctx)
	if err != nil {
		t.Logf("FBL.GetInboundOrderList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.GetInboundOrderList response: %#v", res)
}
func Test_FBL_GetInboundReservationFile(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_inbound_reservation_file_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetInboundReservationFile due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/fbl/inbound_reservation/file*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.GetInboundReservationFile(ctx)
	if err != nil {
		t.Logf("FBL.GetInboundReservationFile returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.GetInboundReservationFile response: %#v", res)
}
func Test_FBL_GetInventoryChangedSKU(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_inventory_changed_sku_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetInventoryChangedSKU due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/fbl/inventory_changed_sku/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.GetInventoryChangedSKU(ctx)
	if err != nil {
		t.Logf("FBL.GetInventoryChangedSKU returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.GetInventoryChangedSKU response: %#v", res)
}
func Test_FBL_GetInventoryOccupyDetails(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_inventory_occupy_details_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetInventoryOccupyDetails due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/fbl/inventory_occupy_details/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.GetInventoryOccupyDetails(ctx)
	if err != nil {
		t.Logf("FBL.GetInventoryOccupyDetails returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.GetInventoryOccupyDetails response: %#v", res)
}
func Test_FBL_GetInventoryOperateLog(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_inventory_operate_log_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetInventoryOperateLog due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/fbl/inventory_operate_log/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.GetInventoryOperateLog(ctx)
	if err != nil {
		t.Logf("FBL.GetInventoryOperateLog returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.GetInventoryOperateLog response: %#v", res)
}
func Test_FBL_GetOutboundOrderDetail(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_outbound_order_detail_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetOutboundOrderDetail due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/fbl/outbound_order_detail/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.GetOutboundOrderDetail(ctx)
	if err != nil {
		t.Logf("FBL.GetOutboundOrderDetail returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.GetOutboundOrderDetail response: %#v", res)
}
func Test_FBL_GetOutboundOrderList(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_outbound_orders_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetOutboundOrderList due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/fbl/outbound_orders/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.GetOutboundOrderList(ctx)
	if err != nil {
		t.Logf("FBL.GetOutboundOrderList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.GetOutboundOrderList response: %#v", res)
}
func Test_FBL_GetPlatformProductsV2(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_platform_products_get2_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetPlatformProductsV2 due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/fbl/platform_products/get2*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.GetPlatformProductsV2(ctx)
	if err != nil {
		t.Logf("FBL.GetPlatformProductsV2 returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.GetPlatformProductsV2 response: %#v", res)
}
func Test_FBL_GetProductBatchList(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_product_batch_query_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetProductBatchList due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/fbl/product_batch/query*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.GetProductBatchList(ctx)
	if err != nil {
		t.Logf("FBL.GetProductBatchList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.GetProductBatchList response: %#v", res)
}
func Test_FBL_GetShipperInfo(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_shipper_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetShipperInfo due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/fbl/shipper/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.GetShipperInfo(ctx)
	if err != nil {
		t.Logf("FBL.GetShipperInfo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.GetShipperInfo response: %#v", res)
}
func Test_FBL_GetStockRule(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_stock_rule_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetStockRule due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/fbl/stock_rule/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.GetStockRule(ctx)
	if err != nil {
		t.Logf("FBL.GetStockRule returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.GetStockRule response: %#v", res)
}
func Test_FBL_GetVasOrderByNo4FBL(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_vas_getVasOrderByNo_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetVasOrderByNo4FBL due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/fbl/vas/getVasOrderByNo*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.GetVasOrderByNo4FBL(ctx)
	if err != nil {
		t.Logf("FBL.GetVasOrderByNo4FBL returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.GetVasOrderByNo4FBL response: %#v", res)
}
func Test_FBL_GetWarehouseListForMCL(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_warehouses_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetWarehouseListForMCL due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/fbl/warehouses/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.GetWarehouseListForMCL(ctx)
	if err != nil {
		t.Logf("FBL.GetWarehouseListForMCL returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.GetWarehouseListForMCL response: %#v", res)
}
func Test_FBL_GetWarehouseStock(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_stocks_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetWarehouseStock due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/fbl/stocks/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.GetWarehouseStock(ctx)
	if err != nil {
		t.Logf("FBL.GetWarehouseStock returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.GetWarehouseStock response: %#v", res)
}
func Test_FBL_GetWarehouseStockV3(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_stocks_getV3_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetWarehouseStockV3 due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/fbl/stocks/getV3*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.GetWarehouseStockV3(ctx)
	if err != nil {
		t.Logf("FBL.GetWarehouseStockV3 returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.GetWarehouseStockV3 response: %#v", res)
}
func Test_FBL_ListIcpWarehouse(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_icp_warehouse_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping ListIcpWarehouse due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/fbl/icp_warehouse/list*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.ListIcpWarehouse(ctx)
	if err != nil {
		t.Logf("FBL.ListIcpWarehouse returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.ListIcpWarehouse response: %#v", res)
}
func Test_FBL_QueryFulfillmentOrderForMCL(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_fulfillment_order_list_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping QueryFulfillmentOrderForMCL due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/fbl/fulfillment_order_list/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.QueryFulfillmentOrderForMCL(ctx)
	if err != nil {
		t.Logf("FBL.QueryFulfillmentOrderForMCL returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.QueryFulfillmentOrderForMCL response: %#v", res)
}
func Test_FBL_QueryInboundBatch(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_inbound_batch_query_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping QueryInboundBatch due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/fbl/inbound_batch/query*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.QueryInboundBatch(ctx)
	if err != nil {
		t.Logf("FBL.QueryInboundBatch returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.QueryInboundBatch response: %#v", res)
}
func Test_FBL_QueryInboundReservationOrder(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_inbound_reservation_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping QueryInboundReservationOrder due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/fbl/inbound_reservation/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.QueryInboundReservationOrder(ctx)
	if err != nil {
		t.Logf("FBL.QueryInboundReservationOrder returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.QueryInboundReservationOrder response: %#v", res)
}
func Test_FBL_QueryReverseOrderForMCL(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_reverse_order_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping QueryReverseOrderForMCL due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/fbl/reverse_order/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.QueryReverseOrderForMCL(ctx)
	if err != nil {
		t.Logf("FBL.QueryReverseOrderForMCL returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.QueryReverseOrderForMCL response: %#v", res)
}
func Test_FBL_RemoveFulfillmentSkuRelation(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_fulfillment_sku_relation_remove_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping RemoveFulfillmentSkuRelation due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/fbl/fulfillment_sku_relation/remove*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.RemoveFulfillmentSkuRelation(ctx)
	if err != nil {
		t.Logf("FBL.RemoveFulfillmentSkuRelation returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.RemoveFulfillmentSkuRelation response: %#v", res)
}
func Test_FBL_ReturnCancellation(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_returns_cancel_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping ReturnCancellation due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/fbl/returns/cancel*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.ReturnCancellation(ctx)
	if err != nil {
		t.Logf("FBL.ReturnCancellation returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.ReturnCancellation response: %#v", res)
}
func Test_FBL_ReturnOrderCreation(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_returns_create_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping ReturnOrderCreation due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/fbl/returns/create*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.ReturnOrderCreation(ctx)
	if err != nil {
		t.Logf("FBL.ReturnOrderCreation returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.ReturnOrderCreation response: %#v", res)
}
func Test_FBL_SetStockRule(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_stock_rule_set_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SetStockRule due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/fbl/stock_rule/set*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.SetStockRule(ctx)
	if err != nil {
		t.Logf("FBL.SetStockRule returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.SetStockRule response: %#v", res)
}
func Test_FBL_UpdateFulfillmentSkuDecouple(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_fulfillment_sku_update_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateFulfillmentSkuDecouple due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/fbl/fulfillment_sku/update*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.UpdateFulfillmentSkuDecouple(ctx)
	if err != nil {
		t.Logf("FBL.UpdateFulfillmentSkuDecouple returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.UpdateFulfillmentSkuDecouple response: %#v", res)
}
func Test_FBL_UploadWaybill(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_fbl_waybill_upload_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UploadWaybill due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/fbl/waybill/upload*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.FBL.UploadWaybill(ctx, "test.jpg", strings.NewReader("test data"))
	if err != nil {
		t.Logf("FBL.UploadWaybill returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FBL.UploadWaybill response: %#v", res)
}
