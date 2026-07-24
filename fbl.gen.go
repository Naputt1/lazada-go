package golazada

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
)

type FBLService interface {
	// BuildFulfillmentSkuRelation build the relation between platformSku and fulfillmentSku
	// Path: /fbl/fulfillment_sku_relation/write
	BuildFulfillmentSkuRelation(ctx context.Context) (*BuildFulfillmentSkuRelationResponse, error)
	// CancelFulfillmentOrderForMCL Cancel Fulfillment Order
	// Path: /fbl/fulfillment_order/cancel
	CancelFulfillmentOrderForMCL(ctx context.Context) (*CancelFulfillmentOrderForMCLResponse, error)
	// CancelInboundReservation cancel reservation order
	// Path: /fbl/inbound_reservation/cancel
	CancelInboundReservation(ctx context.Context) (*CancelInboundReservationResponse, error)
	// CancelnBoundOrder Cancel inbound order
	// Path: /fbl/inbound_order/cancel
	CancelnBoundOrder(ctx context.Context) (*CancelnBoundOrderResponse, error)
	// CancelOutboundOrder Cancel outbound order
	// Path: /fbl/outbound_order/cancel
	CancelOutboundOrder(ctx context.Context) (*CancelOutboundOrderResponse, error)
	// CancelVasOrder4FBL 取消增值服务
	// Path: /fbl/vas/cancelVasOrder
	CancelVasOrder4FBL(ctx context.Context) (*CancelVasOrder4FBLResponse, error)
	// CheckInboundReservationSlot Check Available Reservation Slots for Inbound Order
	// Path: /fbl/inbound_reservation/check
	CheckInboundReservationSlot(ctx context.Context) (*CheckInboundReservationSlotResponse, error)
	// CreateFulfillmentOrderForMCL Create Fulfillment Order
	// Path: /fbl/fulfillment_order/create
	CreateFulfillmentOrderForMCL(ctx context.Context) (*CreateFulfillmentOrderForMCLResponse, error)
	// CreateFulfillmentOrderForMCLV2PNF Create Fulfillment Order for MCL2.0 PNF
	// Path: /fbl/fulfillment_order_pnf/create
	CreateFulfillmentOrderForMCLV2PNF(ctx context.Context) (*CreateFulfillmentOrderForMCLV2PNFResponse, error)
	// CreateFulfillmentSkuDecouple create fulfillment sku without product
	// Path: /fbl/fulfillment_sku/create
	CreateFulfillmentSkuDecouple(ctx context.Context) (*CreateFulfillmentSkuDecoupleResponse, error)
	// CreateFulfillmentSkuForFBL create fulfillment sku for specified platform product
	// Path: /fbl/fulfillment_sku_fbl/create
	CreateFulfillmentSkuForFBL(ctx context.Context) (*CreateFulfillmentSkuForFBLResponse, error)
	// CreateInboundOrder Create inbound order
	// Path: /fbl/inbound_order/create
	CreateInboundOrder(ctx context.Context) (*CreateInboundOrderResponse, error)
	// CreateInboundReservation create reservation order
	// Path: /fbl/inbound_reservation/create
	CreateInboundReservation(ctx context.Context) (*CreateInboundReservationResponse, error)
	// CreateOutBoundOrder Create outbound order
	// Path: /fbl/outbound_order/create
	CreateOutBoundOrder(ctx context.Context) (*CreateOutBoundOrderResponse, error)
	// CreateProductReinboundOrderForMCL Create Product Reinbound Order on Failed Delivery for MCL
	// Path: /fbl/product_reinbound/create
	CreateProductReinboundOrderForMCL(ctx context.Context) (*CreateProductReinboundOrderForMCLResponse, error)
	// CreateVasOrder4FBL FBL增值服务创建
	// Path: /fbl/vas/createVasOrder
	CreateVasOrder4FBL(ctx context.Context) (*CreateVasOrder4FBLResponse, error)
	// GetChannelStocksForMCL Query Channel Stocks
	// Path: /fbl/channel_stocks/get
	GetChannelStocksForMCL(ctx context.Context) (*GetChannelStocksForMCLResponse, error)
	// GetFulfillmentProductDetail GET  fulfillment product Detail；Call Get Platform Products for fulfillment_sku first
	// Path: /fbl/fulfillment_products/get
	GetFulfillmentProductDetail(ctx context.Context) (*GetFulfillmentProductDetailResponse, error)
	// GetFulfillmentSkuListForMCL Get Fulfillment SKU List for LAZADA Partner
	// Path: /fbl/fulfillment_sku_list/get
	GetFulfillmentSkuListForMCL(ctx context.Context) (*GetFulfillmentSkuListForMCLResponse, error)
	// GetFulfillmentSkuRelationByScItem get the relation between platformSku and fulfillmentSku by scItem
	// Path: /fbl/fulfillment_sku_relation/get_by_sc_item
	GetFulfillmentSkuRelationByScItem(ctx context.Context) (*GetFulfillmentSkuRelationByScItemResponse, error)
	// GetFulfillmentSkuRelationBySku get the relation between platformSku and fulfillmentSku by sku
	// Path: /fbl/fulfillment_sku_relation/get_by_sku
	GetFulfillmentSkuRelationBySku(ctx context.Context) (*GetFulfillmentSkuRelationBySkuResponse, error)
	// GetFulfillmentSkuRelationsByScItems get fulfillmentSku Relations By ScItems
	// Path: /fbl/fulfillment_sku_relation/get_by_sc_items
	GetFulfillmentSkuRelationsByScItems(ctx context.Context) (*GetFulfillmentSkuRelationsByScItemsResponse, error)
	// GetFulfillmentSkuRelationsBySkus get fulfillmentSku Relations By Skus
	// Path: /fbl/fulfillment_sku_relation/get_by_skus
	GetFulfillmentSkuRelationsBySkus(ctx context.Context) (*GetFulfillmentSkuRelationsBySkusResponse, error)
	// GetIcpOrderFile Get Inbound/Outbound order print PDF file
	// Path: /fbl/icp_order/file
	GetIcpOrderFile(ctx context.Context) (*GetIcpOrderFileResponse, error)
	// GetInboundOrderDetail Use this API to get the Inbound Order Detail
	// Path: /fbl/inbound_order_detail/get
	GetInboundOrderDetail(ctx context.Context) (*GetInboundOrderDetailResponse, error)
	// GetInboundOrderList Use this API to get inbound order list
	// Path: /fbl/inbound_orders/get
	GetInboundOrderList(ctx context.Context) (*GetInboundOrderListResponse, error)
	// GetInboundReservationFile get inbound reservation order file
	// Path: /fbl/inbound_reservation/file
	GetInboundReservationFile(ctx context.Context) (*GetInboundReservationFileResponse, error)
	// GetInventoryChangedSKU Use this API to get SKU list
	// Path: /fbl/inventory_changed_sku/get
	GetInventoryChangedSKU(ctx context.Context) (*GetInventoryChangedSKUResponse, error)
	// GetInventoryOccupyDetails Use this API to get a sku's inventory occupy details
	// Path: /fbl/inventory_occupy_details/get
	GetInventoryOccupyDetails(ctx context.Context) (*GetInventoryOccupyDetailsResponse, error)
	// GetInventoryOperateLog Use this API to get a sku's inventory operate log
	// Path: /fbl/inventory_operate_log/get
	GetInventoryOperateLog(ctx context.Context) (*GetInventoryOperateLogResponse, error)
	// GetOutboundOrderDetail Use this API to Get outbound order detail; shoud call GetOutboundOrderList for outbound_order_no first
	// Path: /fbl/outbound_order_detail/get
	GetOutboundOrderDetail(ctx context.Context) (*GetOutboundOrderDetailResponse, error)
	// GetOutboundOrderList Use this API to get outbound order list
	// Path: /fbl/outbound_orders/get
	GetOutboundOrderList(ctx context.Context) (*GetOutboundOrderListResponse, error)
	// GetPlatformProductsV2 Search products list
	// Path: /fbl/platform_products/get2
	GetPlatformProductsV2(ctx context.Context) (*GetPlatformProductsV2Response, error)
	// GetProductBatchList query product batch list
	// Path: /fbl/product_batch/query
	GetProductBatchList(ctx context.Context) (*GetProductBatchListResponse, error)
	// GetShipperInfo Get Shipper Info for LAZADA Partner
	// Path: /fbl/shipper/get
	GetShipperInfo(ctx context.Context) (*GetShipperInfoResponse, error)
	// GetStockRule Get SKU stock rule by sku and warehouse
	// Path: /fbl/stock_rule/get
	GetStockRule(ctx context.Context) (*GetStockRuleResponse, error)
	// GetVasOrderByNo4FBL get vasOrder by orderNo
	// Path: /fbl/vas/getVasOrderByNo
	GetVasOrderByNo4FBL(ctx context.Context) (*GetVasOrderByNo4FBLResponse, error)
	// GetWarehouseListForMCL Get Warehouse List By Country And Multi-Channel
	// Path: /fbl/warehouses/get
	GetWarehouseListForMCL(ctx context.Context) (*GetWarehouseListForMCLResponse, error)
	// GetWarehouseStock Get SKU list and stock by warehouse code
	// Path: /fbl/stocks/get
	GetWarehouseStock(ctx context.Context) (*GetWarehouseStockResponse, error)
	// GetWarehouseStockV3 Get SKU list and stock by warehouse code, this version separates pending inbound and stock in transit in return json.
	// Path: /fbl/stocks/getV3
	GetWarehouseStockV3(ctx context.Context) (*GetWarehouseStockV3Response, error)
	// ListIcpWarehouse List warehouses for create InboundOrder and outboundOrder
	// Path: /fbl/icp_warehouse/list
	ListIcpWarehouse(ctx context.Context) (*ListIcpWarehouseResponse, error)
	// QueryFulfillmentOrderForMCL Query list of Fulfillment Orders by shipper
	// Path: /fbl/fulfillment_order_list/get
	QueryFulfillmentOrderForMCL(ctx context.Context) (*QueryFulfillmentOrderForMCLResponse, error)
	// QueryInboundBatch query inbound batch
	// Path: /fbl/inbound_batch/query
	QueryInboundBatch(ctx context.Context) (*QueryInboundBatchResponse, error)
	// QueryInboundReservationOrder get inbound reservation order
	// Path: /fbl/inbound_reservation/get
	QueryInboundReservationOrder(ctx context.Context) (*QueryInboundReservationOrderResponse, error)
	// QueryReverseOrderForMCL Query Reverse Order for MCL
	// Path: /fbl/reverse_order/get
	QueryReverseOrderForMCL(ctx context.Context) (*QueryReverseOrderForMCLResponse, error)
	// RemoveFulfillmentSkuRelation remove the relation between platformSku and fulfillmentSku
	// Path: /fbl/fulfillment_sku_relation/remove
	RemoveFulfillmentSkuRelation(ctx context.Context) (*RemoveFulfillmentSkuRelationResponse, error)
	// ReturnCancellation Return order cancellation
	// Path: /fbl/returns/cancel
	ReturnCancellation(ctx context.Context) (*ReturnCancellationResponse, error)
	// ReturnOrderCreation Api to create customer returns
	// Path: /fbl/returns/create
	ReturnOrderCreation(ctx context.Context) (*ReturnOrderCreationResponse, error)
	// SetStockRule set channel ratio by sku and warehouse
	// Path: /fbl/stock_rule/set
	SetStockRule(ctx context.Context) (*SetStockRuleResponse, error)
	// UpdateFulfillmentSkuDecouple update fulfillment sku without product
	// Path: /fbl/fulfillment_sku/update
	UpdateFulfillmentSkuDecouple(ctx context.Context) (*UpdateFulfillmentSkuDecoupleResponse, error)
	// UploadWaybill Use this API to upload a waybill pdf to Lazada site. The maximum size of an pdf file is 1MB.
	// Path: /fbl/waybill/upload
	UploadWaybill(ctx context.Context, filename string, reader io.Reader) (*UploadWaybillResponse, error)
}

type FBLServiceOp[T any] struct {
	client *Client[T]
}

// BuildFulfillmentSkuRelation build the relation between platformSku and fulfillmentSku
// Path: /fbl/fulfillment_sku_relation/write
func (s *FBLServiceOp[T]) BuildFulfillmentSkuRelation(ctx context.Context) (*BuildFulfillmentSkuRelationResponse, error) {
	path := "/fbl/fulfillment_sku_relation/write"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(BuildFulfillmentSkuRelationResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// CancelFulfillmentOrderForMCL Cancel Fulfillment Order
// Path: /fbl/fulfillment_order/cancel
func (s *FBLServiceOp[T]) CancelFulfillmentOrderForMCL(ctx context.Context) (*CancelFulfillmentOrderForMCLResponse, error) {
	path := "/fbl/fulfillment_order/cancel"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(CancelFulfillmentOrderForMCLResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// CancelInboundReservation cancel reservation order
// Path: /fbl/inbound_reservation/cancel
func (s *FBLServiceOp[T]) CancelInboundReservation(ctx context.Context) (*CancelInboundReservationResponse, error) {
	path := "/fbl/inbound_reservation/cancel"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(CancelInboundReservationResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// CancelnBoundOrder Cancel inbound order
// Path: /fbl/inbound_order/cancel
func (s *FBLServiceOp[T]) CancelnBoundOrder(ctx context.Context) (*CancelnBoundOrderResponse, error) {
	path := "/fbl/inbound_order/cancel"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(CancelnBoundOrderResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// CancelOutboundOrder Cancel outbound order
// Path: /fbl/outbound_order/cancel
func (s *FBLServiceOp[T]) CancelOutboundOrder(ctx context.Context) (*CancelOutboundOrderResponse, error) {
	path := "/fbl/outbound_order/cancel"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(CancelOutboundOrderResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// CancelVasOrder4FBL 取消增值服务
// Path: /fbl/vas/cancelVasOrder
func (s *FBLServiceOp[T]) CancelVasOrder4FBL(ctx context.Context) (*CancelVasOrder4FBLResponse, error) {
	path := "/fbl/vas/cancelVasOrder"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(CancelVasOrder4FBLResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// CheckInboundReservationSlot Check Available Reservation Slots for Inbound Order
// Path: /fbl/inbound_reservation/check
func (s *FBLServiceOp[T]) CheckInboundReservationSlot(ctx context.Context) (*CheckInboundReservationSlotResponse, error) {
	path := "/fbl/inbound_reservation/check"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(CheckInboundReservationSlotResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// CreateFulfillmentOrderForMCL Create Fulfillment Order
// Path: /fbl/fulfillment_order/create
func (s *FBLServiceOp[T]) CreateFulfillmentOrderForMCL(ctx context.Context) (*CreateFulfillmentOrderForMCLResponse, error) {
	path := "/fbl/fulfillment_order/create"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(CreateFulfillmentOrderForMCLResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// CreateFulfillmentOrderForMCLV2PNF Create Fulfillment Order for MCL2.0 PNF
// Path: /fbl/fulfillment_order_pnf/create
func (s *FBLServiceOp[T]) CreateFulfillmentOrderForMCLV2PNF(ctx context.Context) (*CreateFulfillmentOrderForMCLV2PNFResponse, error) {
	path := "/fbl/fulfillment_order_pnf/create"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(CreateFulfillmentOrderForMCLV2PNFResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// CreateFulfillmentSkuDecouple create fulfillment sku without product
// Path: /fbl/fulfillment_sku/create
func (s *FBLServiceOp[T]) CreateFulfillmentSkuDecouple(ctx context.Context) (*CreateFulfillmentSkuDecoupleResponse, error) {
	path := "/fbl/fulfillment_sku/create"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(CreateFulfillmentSkuDecoupleResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// CreateFulfillmentSkuForFBL create fulfillment sku for specified platform product
// Path: /fbl/fulfillment_sku_fbl/create
func (s *FBLServiceOp[T]) CreateFulfillmentSkuForFBL(ctx context.Context) (*CreateFulfillmentSkuForFBLResponse, error) {
	path := "/fbl/fulfillment_sku_fbl/create"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(CreateFulfillmentSkuForFBLResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// CreateInboundOrder Create inbound order
// Path: /fbl/inbound_order/create
func (s *FBLServiceOp[T]) CreateInboundOrder(ctx context.Context) (*CreateInboundOrderResponse, error) {
	path := "/fbl/inbound_order/create"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(CreateInboundOrderResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// CreateInboundReservation create reservation order
// Path: /fbl/inbound_reservation/create
func (s *FBLServiceOp[T]) CreateInboundReservation(ctx context.Context) (*CreateInboundReservationResponse, error) {
	path := "/fbl/inbound_reservation/create"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(CreateInboundReservationResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// CreateOutBoundOrder Create outbound order
// Path: /fbl/outbound_order/create
func (s *FBLServiceOp[T]) CreateOutBoundOrder(ctx context.Context) (*CreateOutBoundOrderResponse, error) {
	path := "/fbl/outbound_order/create"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(CreateOutBoundOrderResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// CreateProductReinboundOrderForMCL Create Product Reinbound Order on Failed Delivery for MCL
// Path: /fbl/product_reinbound/create
func (s *FBLServiceOp[T]) CreateProductReinboundOrderForMCL(ctx context.Context) (*CreateProductReinboundOrderForMCLResponse, error) {
	path := "/fbl/product_reinbound/create"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(CreateProductReinboundOrderForMCLResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// CreateVasOrder4FBL FBL增值服务创建
// Path: /fbl/vas/createVasOrder
func (s *FBLServiceOp[T]) CreateVasOrder4FBL(ctx context.Context) (*CreateVasOrder4FBLResponse, error) {
	path := "/fbl/vas/createVasOrder"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(CreateVasOrder4FBLResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetChannelStocksForMCL Query Channel Stocks
// Path: /fbl/channel_stocks/get
func (s *FBLServiceOp[T]) GetChannelStocksForMCL(ctx context.Context) (*GetChannelStocksForMCLResponse, error) {
	path := "/fbl/channel_stocks/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetChannelStocksForMCLResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetFulfillmentProductDetail GET  fulfillment product Detail；Call Get Platform Products for fulfillment_sku first
// Path: /fbl/fulfillment_products/get
func (s *FBLServiceOp[T]) GetFulfillmentProductDetail(ctx context.Context) (*GetFulfillmentProductDetailResponse, error) {
	path := "/fbl/fulfillment_products/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetFulfillmentProductDetailResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetFulfillmentSkuListForMCL Get Fulfillment SKU List for LAZADA Partner
// Path: /fbl/fulfillment_sku_list/get
func (s *FBLServiceOp[T]) GetFulfillmentSkuListForMCL(ctx context.Context) (*GetFulfillmentSkuListForMCLResponse, error) {
	path := "/fbl/fulfillment_sku_list/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetFulfillmentSkuListForMCLResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetFulfillmentSkuRelationByScItem get the relation between platformSku and fulfillmentSku by scItem
// Path: /fbl/fulfillment_sku_relation/get_by_sc_item
func (s *FBLServiceOp[T]) GetFulfillmentSkuRelationByScItem(ctx context.Context) (*GetFulfillmentSkuRelationByScItemResponse, error) {
	path := "/fbl/fulfillment_sku_relation/get_by_sc_item"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetFulfillmentSkuRelationByScItemResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetFulfillmentSkuRelationBySku get the relation between platformSku and fulfillmentSku by sku
// Path: /fbl/fulfillment_sku_relation/get_by_sku
func (s *FBLServiceOp[T]) GetFulfillmentSkuRelationBySku(ctx context.Context) (*GetFulfillmentSkuRelationBySkuResponse, error) {
	path := "/fbl/fulfillment_sku_relation/get_by_sku"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetFulfillmentSkuRelationBySkuResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetFulfillmentSkuRelationsByScItems get fulfillmentSku Relations By ScItems
// Path: /fbl/fulfillment_sku_relation/get_by_sc_items
func (s *FBLServiceOp[T]) GetFulfillmentSkuRelationsByScItems(ctx context.Context) (*GetFulfillmentSkuRelationsByScItemsResponse, error) {
	path := "/fbl/fulfillment_sku_relation/get_by_sc_items"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetFulfillmentSkuRelationsByScItemsResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetFulfillmentSkuRelationsBySkus get fulfillmentSku Relations By Skus
// Path: /fbl/fulfillment_sku_relation/get_by_skus
func (s *FBLServiceOp[T]) GetFulfillmentSkuRelationsBySkus(ctx context.Context) (*GetFulfillmentSkuRelationsBySkusResponse, error) {
	path := "/fbl/fulfillment_sku_relation/get_by_skus"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetFulfillmentSkuRelationsBySkusResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetIcpOrderFile Get Inbound/Outbound order print PDF file
// Path: /fbl/icp_order/file
func (s *FBLServiceOp[T]) GetIcpOrderFile(ctx context.Context) (*GetIcpOrderFileResponse, error) {
	path := "/fbl/icp_order/file"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetIcpOrderFileResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetInboundOrderDetail Use this API to get the Inbound Order Detail
// Path: /fbl/inbound_order_detail/get
func (s *FBLServiceOp[T]) GetInboundOrderDetail(ctx context.Context) (*GetInboundOrderDetailResponse, error) {
	path := "/fbl/inbound_order_detail/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetInboundOrderDetailResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetInboundOrderList Use this API to get inbound order list
// Path: /fbl/inbound_orders/get
func (s *FBLServiceOp[T]) GetInboundOrderList(ctx context.Context) (*GetInboundOrderListResponse, error) {
	path := "/fbl/inbound_orders/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetInboundOrderListResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetInboundReservationFile get inbound reservation order file
// Path: /fbl/inbound_reservation/file
func (s *FBLServiceOp[T]) GetInboundReservationFile(ctx context.Context) (*GetInboundReservationFileResponse, error) {
	path := "/fbl/inbound_reservation/file"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetInboundReservationFileResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetInventoryChangedSKU Use this API to get SKU list
// Path: /fbl/inventory_changed_sku/get
func (s *FBLServiceOp[T]) GetInventoryChangedSKU(ctx context.Context) (*GetInventoryChangedSKUResponse, error) {
	path := "/fbl/inventory_changed_sku/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetInventoryChangedSKUResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetInventoryOccupyDetails Use this API to get a sku's inventory occupy details
// Path: /fbl/inventory_occupy_details/get
func (s *FBLServiceOp[T]) GetInventoryOccupyDetails(ctx context.Context) (*GetInventoryOccupyDetailsResponse, error) {
	path := "/fbl/inventory_occupy_details/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetInventoryOccupyDetailsResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetInventoryOperateLog Use this API to get a sku's inventory operate log
// Path: /fbl/inventory_operate_log/get
func (s *FBLServiceOp[T]) GetInventoryOperateLog(ctx context.Context) (*GetInventoryOperateLogResponse, error) {
	path := "/fbl/inventory_operate_log/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetInventoryOperateLogResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetOutboundOrderDetail Use this API to Get outbound order detail; shoud call GetOutboundOrderList for outbound_order_no first
// Path: /fbl/outbound_order_detail/get
func (s *FBLServiceOp[T]) GetOutboundOrderDetail(ctx context.Context) (*GetOutboundOrderDetailResponse, error) {
	path := "/fbl/outbound_order_detail/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetOutboundOrderDetailResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetOutboundOrderList Use this API to get outbound order list
// Path: /fbl/outbound_orders/get
func (s *FBLServiceOp[T]) GetOutboundOrderList(ctx context.Context) (*GetOutboundOrderListResponse, error) {
	path := "/fbl/outbound_orders/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetOutboundOrderListResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetPlatformProductsV2 Search products list
// Path: /fbl/platform_products/get2
func (s *FBLServiceOp[T]) GetPlatformProductsV2(ctx context.Context) (*GetPlatformProductsV2Response, error) {
	path := "/fbl/platform_products/get2"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetPlatformProductsV2Response)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetProductBatchList query product batch list
// Path: /fbl/product_batch/query
func (s *FBLServiceOp[T]) GetProductBatchList(ctx context.Context) (*GetProductBatchListResponse, error) {
	path := "/fbl/product_batch/query"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetProductBatchListResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetShipperInfo Get Shipper Info for LAZADA Partner
// Path: /fbl/shipper/get
func (s *FBLServiceOp[T]) GetShipperInfo(ctx context.Context) (*GetShipperInfoResponse, error) {
	path := "/fbl/shipper/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetShipperInfoResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetStockRule Get SKU stock rule by sku and warehouse
// Path: /fbl/stock_rule/get
func (s *FBLServiceOp[T]) GetStockRule(ctx context.Context) (*GetStockRuleResponse, error) {
	path := "/fbl/stock_rule/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetStockRuleResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetVasOrderByNo4FBL get vasOrder by orderNo
// Path: /fbl/vas/getVasOrderByNo
func (s *FBLServiceOp[T]) GetVasOrderByNo4FBL(ctx context.Context) (*GetVasOrderByNo4FBLResponse, error) {
	path := "/fbl/vas/getVasOrderByNo"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetVasOrderByNo4FBLResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetWarehouseListForMCL Get Warehouse List By Country And Multi-Channel
// Path: /fbl/warehouses/get
func (s *FBLServiceOp[T]) GetWarehouseListForMCL(ctx context.Context) (*GetWarehouseListForMCLResponse, error) {
	path := "/fbl/warehouses/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetWarehouseListForMCLResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetWarehouseStock Get SKU list and stock by warehouse code
// Path: /fbl/stocks/get
func (s *FBLServiceOp[T]) GetWarehouseStock(ctx context.Context) (*GetWarehouseStockResponse, error) {
	path := "/fbl/stocks/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetWarehouseStockResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetWarehouseStockV3 Get SKU list and stock by warehouse code, this version separates pending inbound and stock in transit in return json.
// Path: /fbl/stocks/getV3
func (s *FBLServiceOp[T]) GetWarehouseStockV3(ctx context.Context) (*GetWarehouseStockV3Response, error) {
	path := "/fbl/stocks/getV3"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetWarehouseStockV3Response)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// ListIcpWarehouse List warehouses for create InboundOrder and outboundOrder
// Path: /fbl/icp_warehouse/list
func (s *FBLServiceOp[T]) ListIcpWarehouse(ctx context.Context) (*ListIcpWarehouseResponse, error) {
	path := "/fbl/icp_warehouse/list"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(ListIcpWarehouseResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// QueryFulfillmentOrderForMCL Query list of Fulfillment Orders by shipper
// Path: /fbl/fulfillment_order_list/get
func (s *FBLServiceOp[T]) QueryFulfillmentOrderForMCL(ctx context.Context) (*QueryFulfillmentOrderForMCLResponse, error) {
	path := "/fbl/fulfillment_order_list/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(QueryFulfillmentOrderForMCLResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// QueryInboundBatch query inbound batch
// Path: /fbl/inbound_batch/query
func (s *FBLServiceOp[T]) QueryInboundBatch(ctx context.Context) (*QueryInboundBatchResponse, error) {
	path := "/fbl/inbound_batch/query"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(QueryInboundBatchResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// QueryInboundReservationOrder get inbound reservation order
// Path: /fbl/inbound_reservation/get
func (s *FBLServiceOp[T]) QueryInboundReservationOrder(ctx context.Context) (*QueryInboundReservationOrderResponse, error) {
	path := "/fbl/inbound_reservation/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(QueryInboundReservationOrderResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// QueryReverseOrderForMCL Query Reverse Order for MCL
// Path: /fbl/reverse_order/get
func (s *FBLServiceOp[T]) QueryReverseOrderForMCL(ctx context.Context) (*QueryReverseOrderForMCLResponse, error) {
	path := "/fbl/reverse_order/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(QueryReverseOrderForMCLResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// RemoveFulfillmentSkuRelation remove the relation between platformSku and fulfillmentSku
// Path: /fbl/fulfillment_sku_relation/remove
func (s *FBLServiceOp[T]) RemoveFulfillmentSkuRelation(ctx context.Context) (*RemoveFulfillmentSkuRelationResponse, error) {
	path := "/fbl/fulfillment_sku_relation/remove"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(RemoveFulfillmentSkuRelationResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// ReturnCancellation Return order cancellation
// Path: /fbl/returns/cancel
func (s *FBLServiceOp[T]) ReturnCancellation(ctx context.Context) (*ReturnCancellationResponse, error) {
	path := "/fbl/returns/cancel"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(ReturnCancellationResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// ReturnOrderCreation Api to create customer returns
// Path: /fbl/returns/create
func (s *FBLServiceOp[T]) ReturnOrderCreation(ctx context.Context) (*ReturnOrderCreationResponse, error) {
	path := "/fbl/returns/create"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(ReturnOrderCreationResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// SetStockRule set channel ratio by sku and warehouse
// Path: /fbl/stock_rule/set
func (s *FBLServiceOp[T]) SetStockRule(ctx context.Context) (*SetStockRuleResponse, error) {
	path := "/fbl/stock_rule/set"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(SetStockRuleResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// UpdateFulfillmentSkuDecouple update fulfillment sku without product
// Path: /fbl/fulfillment_sku/update
func (s *FBLServiceOp[T]) UpdateFulfillmentSkuDecouple(ctx context.Context) (*UpdateFulfillmentSkuDecoupleResponse, error) {
	path := "/fbl/fulfillment_sku/update"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(UpdateFulfillmentSkuDecoupleResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// UploadWaybill Use this API to upload a waybill pdf to Lazada site. The maximum size of an pdf file is 1MB.
// Path: /fbl/waybill/upload
func (s *FBLServiceOp[T]) UploadWaybill(ctx context.Context, filename string, reader io.Reader) (*UploadWaybillResponse, error) {
	path := "/fbl/waybill/upload"
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", filename)
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(part, reader); err != nil {
		return nil, err
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}

	wrapper, err := s.client.execute(ctx, "POST", path, nil, map[string][]byte{"image": {}})
	if err != nil {
		return nil, err
	}
	resp := new(UploadWaybillResponse)
	if err := json.Unmarshal(wrapper.Data, resp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return resp, nil
}
