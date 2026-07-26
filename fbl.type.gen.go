package golazada

type Batch struct {
	OutboundBanDate   string `json:"outbound_ban_date"`  // [Required]
	Quantity          string `json:"quantity"`           // [Required]
	FulfillmentSkuId  string `json:"fulfillment_sku_id"` // [Required]
	ExpiryDate        string `json:"expiry_date"`        // [Required]
	ManufacturingDate string `json:"manufacturing_date"` // [Required]
	InventoryStatus   string `json:"inventory_status"`   // [Required]
	ProductBatch      string `json:"product_batch"`      // [Required]
}
type BuildFulfillmentSkuRelationResponse struct {
	BaseResponse                                                // Common response fields
	Result       *BuildFulfillmentSkuRelationResponseDataResult `json:"result,omitempty"` //
}
type BuildFulfillmentSkuRelationResponseDataResult struct {
	ErrorMsg  string `json:"error_msg"`  // [Required]
	Success   bool   `json:"success"`    // [Required]
	Failure   string `json:"failure"`    // [Required]
	ErrorCode string `json:"error_code"` // [Required]
}
type CancelFulfillmentOrderForMCLResponse struct {
	BaseResponse        // Common response fields
	ErrorMessage string `json:"error_message,omitempty"` //
}
type CancelInboundReservationResponse struct {
	BaseResponse        // Common response fields
	ErrorMessage string `json:"error_message,omitempty"` //
}
type CancelnBoundOrderResponse struct {
	BaseResponse        // Common response fields
	ErrorMessage string `json:"error_message,omitempty"` //
}
type CancelOutboundOrderResponse struct {
	BaseResponse        // Common response fields
	ErrorMessage string `json:"error_message,omitempty"` //
}
type CancelVasOrder4FBLResponse struct {
	BaseResponse        // Common response fields
	Response     string `json:"data"` // Response data
}
type ChannelRatio struct {
	ChannelCode string `json:"channel_code"` // [Required]
	Ratio       string `json:"ratio"`        // [Required]
}
type ChannelStocks struct {
	Quantity string `json:"quantity"` // [Required]
	Channel  string `json:"channel"`  // [Required]
}
type CheckInboundReservationSlotResponse struct {
	BaseResponse                                         // Common response fields
	Response     CheckInboundReservationSlotResponseData `json:"data"`                    // Response data
	ErrorMessage string                                  `json:"error_message,omitempty"` //
}
type CheckInboundReservationSlotResponseData struct {
	Slots []string `json:"slots"` // [Required]
}
type CreateFulfillmentOrderForMCLResponse struct {
	BaseResponse        // Common response fields
	ErrorMessage string `json:"error_message,omitempty"` //
}
type CreateFulfillmentOrderForMCLV2PNFResponse struct {
	BaseResponse        // Common response fields
	ErrorMessage string `json:"error_message,omitempty"` //
}
type CreateFulfillmentSkuDecoupleResponse struct {
	BaseResponse                                          // Common response fields
	Response     CreateFulfillmentSkuDecoupleResponseData `json:"data"`                    // Response data
	ErrorMessage string                                   `json:"error_message,omitempty"` //
}
type CreateFulfillmentSkuDecoupleResponseData struct {
	FulfillmentSkuId   string `json:"fulfillment_sku_id"`   // [Required]
	FulfillmentSkuCode string `json:"fulfillment_sku_code"` // [Required]
}
type CreateFulfillmentSkuForFBLResponse struct {
	BaseResponse                                        // Common response fields
	Response     CreateFulfillmentSkuForFBLResponseData `json:"data"`                    // Response data
	ErrorMessage string                                 `json:"error_message,omitempty"` //
}
type CreateFulfillmentSkuForFBLResponseData struct {
	FulfillmentSkuId   string `json:"fulfillment_sku_id"`   // [Required]
	FulfillmentSkuCode string `json:"fulfillment_sku_code"` // [Required]
}
type CreateInboundOrderResponse struct {
	BaseResponse          // Common response fields
	ErrorMessage   string `json:"error_message,omitempty"`    //
	InboundOrderNo string `json:"inbound_order_no,omitempty"` //
}
type CreateInboundReservationResponse struct {
	BaseResponse                                      // Common response fields
	Response     CreateInboundReservationResponseData `json:"data"`                    // Response data
	ErrorMessage string                               `json:"error_message,omitempty"` //
}
type CreateInboundReservationResponseData struct {
	ReservationOrder string `json:"reservation_order"` // [Required]
}
type CreateOutBoundOrderResponse struct {
	BaseResponse           // Common response fields
	ErrorMessage    string `json:"error_message,omitempty"`     //
	OutboundOrderNo string `json:"outbound_order_no,omitempty"` //
}
type CreateProductReinboundOrderForMCLResponse struct {
	BaseResponse        // Common response fields
	ErrorMessage string `json:"error_message,omitempty"` //
}
type CreateVasOrder4FBLResponse struct {
	BaseResponse        // Common response fields
	Response     string `json:"data"` // Response data
}
type DataBatch struct {
	Quantity         string `json:"quantity"`           // [Required]
	FulfillmentSkuId string `json:"fulfillment_sku_id"` // [Required]
	InventoryStatus  string `json:"inventory_status"`   // [Required]
	ProductBatch     string `json:"product_batch"`      // [Required]
}
type GetChannelStocksForMCLResponse struct {
	BaseResponse                                    // Common response fields
	Response     GetChannelStocksForMCLResponseData `json:"data"`                    // Response data
	ErrorMessage string                             `json:"error_message,omitempty"` //
}
type GetChannelStocksForMCLResponseData struct {
	FulfillmentSkuId string   `json:"fulfillment_sku_id"` // [Required]
	Stocks           []Stocks `json:"stocks"`             // [Required]
}
type GetFulfillmentProductDetailResponse struct {
	BaseResponse                                         // Common response fields
	Response     GetFulfillmentProductDetailResponseData `json:"data"` // Response data
}
type GetFulfillmentProductDetailResponseData struct {
	ShelfLifeDays          string     `json:"shelf_life_days"`         // [Required]
	Precious               string     `json:"precious"`                // [Required]
	Color                  string     `json:"color"`                   // [Required]
	FulfillmentSku         string     `json:"fulfillment_sku"`         // [Required]
	SerialNumberFlag       string     `json:"serial_number_flag"`      // [Required]
	Length                 string     `json:"length"`                  // [Required]
	OfflineShelfLive       string     `json:"offline_shelf_live"`      // [Required]
	Barcodes               string     `json:"barcodes"`                // [Required]
	NetWeight              string     `json:"net_weight"`              // [Required]
	AlertShelfLive         string     `json:"alert_shelf_live"`        // [Required]
	ShelfLifeFlag          string     `json:"shelf_life_flag"`         // [Required]
	RejectShelfLive        string     `json:"reject_shelf_live"`       // [Required]
	ProductType            string     `json:"product_type"`            // [Required]
	SellerSkus             []string   `json:"seller_skus"`             // [Required]
	SnSampleList           []SnSample `json:"sn_sample_list"`          // [Required]
	Width                  string     `json:"width"`                   // [Required]
	TemperatureRequirement string     `json:"temperature_requirement"` // [Required]
	ShipperId              string     `json:"shipper_id"`              // [Required]
	SerialNumberMode       string     `json:"serial_number_mode"`      // [Required]
	Hygroscopic            string     `json:"hygroscopic"`             // [Required]
	FulfillmentSkuName     string     `json:"fulfillment_sku_name"`    // [Required]
	GrossWeight            string     `json:"gross_weight"`            // [Required]
	Height                 string     `json:"height"`                  // [Required]
}
type GetFulfillmentSkuListForMCLResponse struct {
	BaseResponse                                         // Common response fields
	Response     GetFulfillmentSkuListForMCLResponseData `json:"data"`                    // Response data
	ErrorMessage string                                  `json:"error_message,omitempty"` //
	Page         string                                  `json:"page,omitempty"`          //
	PerPage      string                                  `json:"per_page,omitempty"`      //
	TotalCount   int64                                   `json:"total_count,omitempty"`   //
}
type GetFulfillmentSkuListForMCLResponseData struct {
	HasStock           string `json:"has_stock"`            // [Required]
	FulfillmentSkuId   string `json:"fulfillment_sku_id"`   // [Required]
	SerialNumFlag      string `json:"serial_num_flag"`      // [Required]
	OwnerId            string `json:"owner_id"`             // [Required]
	MinStockAlert      string `json:"min_stock_alert"`      // [Required]
	PicUrls            string `json:"pic_urls"`             // [Required]
	Barcodes           string `json:"barcodes"`             // [Required]
	SalePrice          string `json:"sale_price"`           // [Required]
	ShelfLifeFlag      string `json:"shelf_life_flag"`      // [Required]
	SellerSkus         string `json:"seller_skus"`          // [Required]
	PlatformName       string `json:"platform_name"`        // [Required]
	Currency           string `json:"currency"`             // [Required]
	FulfillmentSkuName string `json:"fulfillment_sku_name"` // [Required]
	PlatformSkuStatus  string `json:"platform_sku_status"`  // [Required]
	FulfillmentSkuCode string `json:"fulfillment_sku_code"` // [Required]
	SellerId           int64  `json:"seller_id"`            // [Required]
}
type GetFulfillmentSkuRelationByScItemResponse struct {
	BaseResponse                                                      // Common response fields
	Result       *GetFulfillmentSkuRelationByScItemResponseDataResult `json:"result,omitempty"` //
}
type GetFulfillmentSkuRelationByScItemResponseDataResult struct {
	ErrorMsg  string                                                    `json:"error_msg"`  // [Required]
	Data      []GetFulfillmentSkuRelationByScItemResponseDataResultData `json:"data"`       // [Required]
	Failure   string                                                    `json:"failure"`    // [Required]
	Success   bool                                                      `json:"success"`    // [Required]
	ErrorCode string                                                    `json:"error_code"` // [Required]
}
type GetFulfillmentSkuRelationByScItemResponseDataResultData struct {
	Site           string `json:"site"`            // [Required]
	ItemId         int64  `json:"item_id"`         // [Required]
	FulfillmentSku string `json:"fulfillment_sku"` // [Required]
	ScItemUserId   string `json:"sc_item_user_id"` // [Required]
	SkuId          int64  `json:"sku_id"`          // [Required]
	Source         string `json:"source"`          // [Required]
	SellerId       int64  `json:"seller_id"`       // [Required]
	ScItemId       string `json:"sc_item_id"`      // [Required]
}
type GetFulfillmentSkuRelationBySkuResponse struct {
	BaseResponse                                                   // Common response fields
	Result       *GetFulfillmentSkuRelationBySkuResponseDataResult `json:"result,omitempty"` //
}
type GetFulfillmentSkuRelationBySkuResponseDataResult struct {
	ErrorMsg  string                                                   `json:"error_msg"`  // [Required]
	Data      *GetFulfillmentSkuRelationByScItemResponseDataResultData `json:"data"`       // [Required]
	Failure   string                                                   `json:"failure"`    // [Required]
	Success   bool                                                     `json:"success"`    // [Required]
	ErrorCode string                                                   `json:"error_code"` // [Required]
}
type GetFulfillmentSkuRelationsByScItemsResponse struct {
	BaseResponse                                                      // Common response fields
	Result       *GetFulfillmentSkuRelationByScItemResponseDataResult `json:"result,omitempty"` //
}
type GetFulfillmentSkuRelationsBySkusResponse struct {
	BaseResponse                                                      // Common response fields
	Result       *GetFulfillmentSkuRelationByScItemResponseDataResult `json:"result,omitempty"` //
}
type GetIcpOrderFileResponse struct {
	BaseResponse                             // Common response fields
	Response     GetIcpOrderFileResponseData `json:"data"`                    // Response data
	ErrorMessage string                      `json:"error_message,omitempty"` //
}
type GetIcpOrderFileResponseData struct {
	Url string `json:"url"` // [Required]
}
type GetInboundOrderDetailResponse struct {
	BaseResponse                                   // Common response fields
	Response     GetInboundOrderDetailResponseData `json:"data"` // Response data
}
type GetInboundOrderDetailResponseData struct {
	InboundWarehouse       string                                  `json:"inbound_warehouse"`        // [Required]
	Skus                   []GetInboundOrderDetailResponseDataSkus `json:"skus"`                     // [Required]
	InboundTime            string                                  `json:"inbound_time"`             // [Required]
	InboundWarehouseCode   string                                  `json:"inbound_warehouse_code"`   // [Required]
	CreatedAt              string                                  `json:"created_at"`               // [Required]
	SellerMobile           string                                  `json:"seller_mobile"`            // [Required]
	SellerCountry          string                                  `json:"seller_country"`           // [Required]
	FulfillmentOrderNumber string                                  `json:"fulfillment_order_number"` // [Required]
	NeedReservation        string                                  `json:"need_reservation"`         // [Required]
	SellerPostcode         string                                  `json:"seller_postcode"`          // [Required]
	SellerWarehouseName    string                                  `json:"seller_warehouse_name"`    // [Required]
	UpdatedAt              string                                  `json:"updated_at"`               // [Required]
	EstimateTime           string                                  `json:"estimate_time"`            // [Required]
	DeliveryType           string                                  `json:"delivery_type"`            // [Required]
	SellerContact          string                                  `json:"seller_contact"`           // [Required]
	IoStatus               string                                  `json:"io_status"`                // [Required]
	Comments               string                                  `json:"comments"`                 // [Required]
	Marketplace            string                                  `json:"marketplace"`              // [Required]
	WarehouseAddress       string                                  `json:"warehouse_address"`        // [Required]
	ReservationOrder       string                                  `json:"reservation_order"`        // [Required]
	ShopName               string                                  `json:"shop_name"`                // [Required]
	ReferenceNumber        string                                  `json:"reference_number"`         // [Required]
	SellerAddress          string                                  `json:"seller_address"`           // [Required]
	SellerCity             string                                  `json:"seller_city"`              // [Required]
	ReservationStatus      string                                  `json:"reservation_status"`       // [Required]
	WarehouseName          string                                  `json:"warehouse_name"`           // [Required]
	IoType                 string                                  `json:"io_type"`                  // [Required]
	IoNumber               string                                  `json:"io_number"`                // [Required]
}
type GetInboundOrderDetailResponseDataSkus struct {
	ShelfLifeFlag        string   `json:"shelf_life_flag"`        // [Required]
	Comments             string   `json:"comments"`               // [Required]
	ItemInboundedDamaged string   `json:"item_inbounded_damaged"` // [Required]
	RequestedQuantity    string   `json:"requested_quantity"`     // [Required]
	SerialNumberFlag     string   `json:"serial_number_flag"`     // [Required]
	FulfillmentSku       string   `json:"fulfillment_sku"`        // [Required]
	SellerSku            []string `json:"seller_sku"`             // [Required]
	ItemInboundedExpired string   `json:"item_inbounded_expired"` // [Required]
	ItemInboundedGood    string   `json:"item_inbounded_good"`    // [Required]
	SkuStatus            string   `json:"sku_status"`             // [Required]
	FulfillmentSkuName   string   `json:"fulfillment_sku_name"`   // [Required]
	Barcodes             []string `json:"barcodes"`               // [Required]
}
type GetInboundOrderListResponse struct {
	BaseResponse                                        // Common response fields
	Result       *GetInboundOrderListResponseDataResult `json:"result,omitempty"` //
}
type GetInboundOrderListResponseDataResult struct {
	PerPage    string                                      `json:"per_page"`    // [Required]
	Data       []GetInboundOrderListResponseDataResultData `json:"data"`        // [Required]
	TotalCount int64                                       `json:"total_count"` // [Required]
	Page       string                                      `json:"page"`        // [Required]
}
type GetInboundOrderListResponseDataResultData struct {
	InboundWarehouse     string `json:"inbound_warehouse"`      // [Required]
	InboundTime          string `json:"inbound_time"`           // [Required]
	Marketplace          string `json:"marketplace"`            // [Required]
	ItemInboundedDamaged string `json:"item_inbounded_damaged"` // [Required]
	SkuApproved          string `json:"sku_approved"`           // [Required]
	ReservationOrder     string `json:"reservation_order"`      // [Required]
	ItemRequested        string `json:"item_requested"`         // [Required]
	InboundWarehouseCode string `json:"inbound_warehouse_code"` // [Required]
	CreatedAt            string `json:"created_at"`             // [Required]
	ItemInboundedExpired string `json:"item_inbounded_expired"` // [Required]
	ShopName             string `json:"shop_name"`              // [Required]
	ReferenceNumber      string `json:"reference_number"`       // [Required]
	NeedReservation      string `json:"need_reservation"`       // [Required]
	SkuInbounded         string `json:"sku_inbounded"`          // [Required]
	SkuRequested         string `json:"sku_requested"`          // [Required]
	ReservationStatus    string `json:"reservation_status"`     // [Required]
	UpdatedAt            string `json:"updated_at"`             // [Required]
	EstimateTime         string `json:"estimate_time"`          // [Required]
	DeliveryType         string `json:"delivery_type"`          // [Required]
	IoType               string `json:"io_type"`                // [Required]
	ItemInboundedGood    string `json:"item_inbounded_good"`    // [Required]
	IoNumber             string `json:"io_number"`              // [Required]
	Status               string `json:"status"`                 // [Required]
}
type GetInboundReservationFileResponse struct {
	BaseResponse                                       // Common response fields
	Response     GetInboundReservationFileResponseData `json:"data"`                    // Response data
	ErrorMessage string                                `json:"error_message,omitempty"` //
}
type GetInboundReservationFileResponseData struct {
	Url string `json:"url"` // [Required]
}
type GetInventoryChangedSKUResponse struct {
	BaseResponse                   // Common response fields
	ErrCode      string            `json:"errCode,omitempty"`     //
	ErrMessage   string            `json:"errMessage,omitempty"`  //
	Page         string            `json:"page,omitempty"`        //
	PerPage      string            `json:"per_page,omitempty"`    //
	SkuList      []ResponseDataSku `json:"sku_list,omitempty"`    //
	TotalCount   int64             `json:"total_count,omitempty"` //
}
type GetInventoryOccupyDetailsResponse struct {
	BaseResponse                                    // Common response fields
	InventoryOccupyDetails []InventoryOccupyDetails `json:"inventoryOccupyDetails,omitempty"` //
}
type GetInventoryOperateLogResponse struct {
	BaseResponse                              // Common response fields
	ErrCode             string                `json:"errCode,omitempty"`               //
	ErrMessage          string                `json:"errMessage,omitempty"`            //
	InventoryOperateLog []InventoryOperateLog `json:"inventory_operate_log,omitempty"` //
	Page                string                `json:"page,omitempty"`                  //
	PerPage             string                `json:"per_page,omitempty"`              //
	TotalCount          int64                 `json:"total_count,omitempty"`           //
}
type GetOutboundOrderDetailResponse struct {
	BaseResponse                                    // Common response fields
	Response     GetOutboundOrderDetailResponseData `json:"data"` // Response data
}
type GetOutboundOrderDetailResponseData struct {
	Skus                   []GetOutboundOrderDetailResponseDataSkus `json:"skus"`                     // [Required]
	CreatedAt              string                                   `json:"created_at"`               // [Required]
	SellerMobile           string                                   `json:"seller_mobile"`            // [Required]
	SellerCountry          string                                   `json:"seller_country"`           // [Required]
	FulfillmentOrderNumber string                                   `json:"fulfillment_order_number"` // [Required]
	SellerPostcode         string                                   `json:"seller_postcode"`          // [Required]
	OutboundOrderNo        string                                   `json:"outbound_order_no"`        // [Required]
	SellerWarehouseName    string                                   `json:"seller_warehouse_name"`    // [Required]
	UpdatedAt              string                                   `json:"updated_at"`               // [Required]
	EstimateTime           string                                   `json:"estimate_time"`            // [Required]
	OutboundWarehouse      string                                   `json:"outbound_warehouse"`       // [Required]
	DeliveryType           string                                   `json:"delivery_type"`            // [Required]
	SellerContact          string                                   `json:"seller_contact"`           // [Required]
	OutboundWarehouseCode  string                                   `json:"outbound_warehouse_code"`  // [Required]
	OutboundReason         string                                   `json:"outbound_reason"`          // [Required]
	Comments               string                                   `json:"comments"`                 // [Required]
	Marketplace            string                                   `json:"marketplace"`              // [Required]
	WarehouseAddress       string                                   `json:"warehouse_address"`        // [Required]
	OutboundTime           string                                   `json:"outbound_time"`            // [Required]
	ShopName               string                                   `json:"shop_name"`                // [Required]
	ReferenceNumber        string                                   `json:"reference_number"`         // [Required]
	CreatedBy              string                                   `json:"created_by"`               // [Required]
	SellerAddress          string                                   `json:"seller_address"`           // [Required]
	SellerCity             string                                   `json:"seller_city"`              // [Required]
	ItemOutbounded         string                                   `json:"item_outbounded"`          // [Required]
	WarehouseName          string                                   `json:"warehouse_name"`           // [Required]
	InventoryType          string                                   `json:"inventory_type"`           // [Required]
	Status                 string                                   `json:"status"`                   // [Required]
}
type GetOutboundOrderDetailResponseDataSkus struct {
	ItemOutbounded     string   `json:"item_outbounded"`      // [Required]
	ShelfLifeFlag      string   `json:"shelf_life_flag"`      // [Required]
	Comments           string   `json:"comments"`             // [Required]
	RequestedQuantity  string   `json:"requested_quantity"`   // [Required]
	SerialNumberFlag   string   `json:"serial_number_flag"`   // [Required]
	FulfillmentSku     string   `json:"fulfillment_sku"`      // [Required]
	SellerSku          []string `json:"seller_sku"`           // [Required]
	SkuStatus          string   `json:"sku_status"`           // [Required]
	FulfillmentSkuName string   `json:"fulfillment_sku_name"` // [Required]
	Barcodes           []string `json:"barcodes"`             // [Required]
}
type GetOutboundOrderListResponse struct {
	BaseResponse                                         // Common response fields
	Result       *GetOutboundOrderListResponseDataResult `json:"result,omitempty"` //
}
type GetOutboundOrderListResponseDataResult struct {
	PerPage    string                                       `json:"per_page"`    // [Required]
	Data       []GetOutboundOrderListResponseDataResultData `json:"data"`        // [Required]
	TotalCount int64                                        `json:"total_count"` // [Required]
	Page       string                                       `json:"page"`        // [Required]
}
type GetOutboundOrderListResponseDataResultData struct {
	Marketplace            string `json:"marketplace"`              // [Required]
	SkuApproved            string `json:"sku_approved"`             // [Required]
	ItemRequested          string `json:"item_requested"`           // [Required]
	CreatedAt              string `json:"created_at"`               // [Required]
	OutboundTime           string `json:"outbound_time"`            // [Required]
	ShopName               string `json:"shop_name"`                // [Required]
	FulfillmentOrderNumber string `json:"fulfillment_order_number"` // [Required]
	ReferenceNumber        string `json:"reference_number"`         // [Required]
	CreatedBy              string `json:"created_by"`               // [Required]
	ItemOutbounded         string `json:"item_outbounded"`          // [Required]
	SkuRequested           string `json:"sku_requested"`            // [Required]
	UpdatedAt              string `json:"updated_at"`               // [Required]
	EstimateTime           string `json:"estimate_time"`            // [Required]
	DeliveryType           string `json:"delivery_type"`            // [Required]
	OutboundWarehouse      string `json:"outbound_warehouse"`       // [Required]
	OutboundWarehouseCode  string `json:"outbound_warehouse_code"`  // [Required]
	SkuOutbounded          string `json:"sku_outbounded"`           // [Required]
	OutboundReason         string `json:"outbound_reason"`          // [Required]
	OoNumber               string `json:"oo_number"`                // [Required]
	Status                 string `json:"status"`                   // [Required]
}
type GetPlatformProductsV2Response struct {
	BaseResponse                                   // Common response fields
	Response     GetPlatformProductsV2ResponseData `json:"data"` // Response data
}
type GetPlatformProductsV2ResponseData struct {
	Skus            []GetPlatformProductsV2ResponseDataSkus `json:"skus"`              // [Required]
	Marketplace     string                                  `json:"marketplace"`       // [Required]
	ProductId       int64                                   `json:"product_id"`        // [Required]
	PlatformSkuName string                                  `json:"platform_sku_name"` // [Required]
	Source          string                                  `json:"source"`            // [Required]
	Status          string                                  `json:"status"`            // [Required]
}
type GetPlatformProductsV2ResponseDataSkus struct {
	FulfillmentSku     string `json:"fulfillment_sku"`      // [Required]
	SellerSku          string `json:"seller_sku"`           // [Required]
	ExtendFields       string `json:"extend_fields"`        // [Required]
	SkuStatus          string `json:"sku_status"`           // [Required]
	PlatformSku        string `json:"platform_sku"`         // [Required]
	FulfillmentSkuName string `json:"fulfillment_sku_name"` // [Required]
}
type GetProductBatchListResponse struct {
	BaseResponse                                        // Common response fields
	Result       *GetProductBatchListResponseDataResult `json:"result,omitempty"` //
}
type GetProductBatchListResponseDataResult struct {
	ErrorMessage string                                     `json:"error_message"` // [Required]
	Data         *GetProductBatchListResponseDataResultData `json:"data"`          // [Required]
	Success      bool                                       `json:"success"`       // [Required]
	ErrorCode    string                                     `json:"error_code"`    // [Required]
}
type GetProductBatchListResponseDataResultData struct {
	StoreCode string  `json:"store_code"` // [Required]
	BatchList []Batch `json:"batch_list"` // [Required]
	PageNo    string  `json:"page_no"`    // [Required]
	PageSize  int64   `json:"page_size"`  // [Required]
}
type GetShipperInfoResponse struct {
	BaseResponse                            // Common response fields
	Response     GetShipperInfoResponseData `json:"data"`                    // Response data
	ErrorMessage string                     `json:"error_message,omitempty"` //
}
type GetShipperInfoResponseData struct {
	MainSellerSite string `json:"main_seller_site"` // [Required]
	MainShipperId  string `json:"main_shipper_id"`  // [Required]
	PartnerName    string `json:"partner_name"`     // [Required]
	IsCb           string `json:"is_cb"`            // [Required]
	MainSellerId   string `json:"main_seller_id"`   // [Required]
	ShipperId      string `json:"shipper_id"`       // [Required]
	IsMcl          string `json:"is_mcl"`           // [Required]
}
type GetStockRuleResponse struct {
	BaseResponse                          // Common response fields
	Response     GetStockRuleResponseData `json:"data"`                    // Response data
	ErrorMessage string                   `json:"error_message,omitempty"` //
	Page         string                   `json:"page,omitempty"`          //
	PerPage      string                   `json:"per_page,omitempty"`      //
	TotalCount   int64                    `json:"total_count,omitempty"`   //
}
type GetStockRuleResponseData struct {
	StoreCode        string         `json:"store_code"`         // [Required]
	ChannelRatio     []ChannelRatio `json:"channel_ratio"`      // [Required]
	FulfillmentSkuId string         `json:"fulfillment_sku_id"` // [Required]
	AutoBalancing    string         `json:"auto_balancing"`     // [Required]
}
type GetVasOrderByNo4FBLResponse struct {
	BaseResponse        // Common response fields
	Response     string `json:"data"` // Response data
}
type GetWarehouseListForMCLResponse struct {
	BaseResponse                                    // Common response fields
	Response     GetWarehouseListForMCLResponseData `json:"data"`                    // Response data
	ErrorMessage string                             `json:"error_message,omitempty"` //
	Page         string                             `json:"page,omitempty"`          //
	PerPage      string                             `json:"per_page,omitempty"`      //
	TotalCount   int64                              `json:"total_count,omitempty"`   //
	TotalPage    int64                              `json:"total_page,omitempty"`    //
}
type GetWarehouseListForMCLResponseData struct {
	CountryCode   string `json:"country_code"`   // [Required]
	TownCode      string `json:"town_code"`      // [Required]
	WarehouseName string `json:"warehouse_name"` // [Required]
	MultiChannel  string `json:"multi_channel"`  // [Required]
	WarehouseCode string `json:"warehouse_code"` // [Required]
	AreaCode      string `json:"area_code"`      // [Required]
	Latitude      string `json:"latitude"`       // [Required]
	PlatformName  string `json:"platform_name"`  // [Required]
	CityCode      string `json:"city_code"`      // [Required]
	DivisionId    string `json:"division_id"`    // [Required]
	ZipCode       string `json:"zip_code"`       // [Required]
	Longitude     string `json:"longitude"`      // [Required]
}
type GetWarehouseStockResponse struct {
	BaseResponse                               // Common response fields
	Response     GetWarehouseStockResponseData `json:"data"` // Response data
}
type GetWarehouseStockResponseData struct {
	FulfilmentSku string        `json:"fulfilment_sku"` // [Required]
	StoreStocks   []StoreStocks `json:"store_stocks"`   // [Required]
}
type GetWarehouseStockV3Response struct {
	BaseResponse                                 // Common response fields
	Response     GetWarehouseStockV3ResponseData `json:"data"` // Response data
}
type GetWarehouseStockV3ResponseData struct {
	FulfilmentSku string                    `json:"fulfilment_sku"` // [Required]
	StoreStocks   []ResponseDataStoreStocks `json:"store_stocks"`   // [Required]
}
type InventoryOccupyDetails struct {
	OrderType     string `json:"orderType"`     // [Required]
	InventoryType string `json:"inventoryType"` // [Required]
	Quantity      string `json:"quantity"`      // [Required]
	OrderCode     string `json:"orderCode"`     // [Required]
}
type InventoryOperateLog struct {
	OrderTypeCode    string         `json:"order_type_code"`    // [Required]
	RefOrderCode     []RefOrderCode `json:"ref_order_code"`     // [Required]
	WarehouseName    string         `json:"warehouse_name"`     // [Required]
	ChangeQuantity   string         `json:"change_quantity"`    // [Required]
	FulfillmentSkuId string         `json:"fulfillment_sku_id"` // [Required]
	WarehouseCode    string         `json:"warehouse_code"`     // [Required]
	CustomerOrder    string         `json:"customer_order"`     // [Required]
	InventoryType    string         `json:"inventory_type"`     // [Required]
	OrderType        string         `json:"order_type"`         // [Required]
	ResultQuantity   string         `json:"result_quantity"`    // [Required]
	OperateTime      string         `json:"operate_time"`       // [Required]
}
type Items struct {
	FulfillmentSkuId string `json:"fulfillment_sku_id"` // [Required]
	PlatformItemId   string `json:"platform_item_id"`   // [Required]
	Status           string `json:"status"`             // [Required]
}
type ListIcpWarehouseResponse struct {
	BaseResponse                              // Common response fields
	Response     ListIcpWarehouseResponseData `json:"data"`                    // Response data
	ErrorMessage string                       `json:"error_message,omitempty"` //
}
type ListIcpWarehouseResponseData struct {
	WarehouseName string `json:"warehouse_name"` // [Required]
	WarehouseCode string `json:"warehouse_code"` // [Required]
}
type Pending struct {
	Reserved  string `json:"reserved"`  // [Required]
	Available string `json:"available"` // [Required]
}
type QueryFulfillmentOrderForMCLResponse struct {
	BaseResponse                                         // Common response fields
	Response     QueryFulfillmentOrderForMCLResponseData `json:"data"`                    // Response data
	ErrorMessage string                                  `json:"error_message,omitempty"` //
	Page         string                                  `json:"page,omitempty"`          //
	PerPage      string                                  `json:"per_page,omitempty"`      //
	TotalCount   int64                                   `json:"total_count,omitempty"`   //
}
type QueryFulfillmentOrderForMCLResponseData struct {
	SalesOrderNumber string  `json:"sales_order_number"` // [Required]
	PlatformOrderId  string  `json:"platform_order_id"`  // [Required]
	CreateTime       string  `json:"create_time"`        // [Required]
	Items            []Items `json:"items"`              // [Required]
}
type QueryInboundBatchResponse struct {
	BaseResponse                                      // Common response fields
	Result       *QueryInboundBatchResponseDataResult `json:"result,omitempty"` //
}
type QueryInboundBatchResponseDataResult struct {
	ErrorMessage string                                   `json:"error_message"` // [Required]
	Data         *QueryInboundBatchResponseDataResultData `json:"data"`          // [Required]
	Success      bool                                     `json:"success"`       // [Required]
	ErrorCode    string                                   `json:"error_code"`    // [Required]
}
type QueryInboundBatchResponseDataResultData struct {
	StoreCode    string      `json:"store_code"`    // [Required]
	BatchList    []DataBatch `json:"batch_list"`    // [Required]
	InboundOrder string      `json:"inbound_order"` // [Required]
}
type QueryInboundReservationOrderResponse struct {
	BaseResponse                                          // Common response fields
	Response     QueryInboundReservationOrderResponseData `json:"data"`                    // Response data
	ErrorMessage string                                   `json:"error_message,omitempty"` //
}
type QueryInboundReservationOrderResponseData struct {
	ReservationOrder string   `json:"reservation_order"` // [Required]
	InboundOrders    []string `json:"inbound_orders"`    // [Required]
	Slot             string   `json:"slot"`              // [Required]
	Status           string   `json:"status"`            // [Required]
}
type QueryReverseOrderForMCLResponse struct {
	BaseResponse                                     // Common response fields
	Response     QueryReverseOrderForMCLResponseData `json:"data"`                    // Response data
	ErrorMessage string                              `json:"error_message,omitempty"` //
}
type QueryReverseOrderForMCLResponseData struct {
	SalesOrderNumber string              `json:"sales_order_number"` // [Required]
	CreateTime       string              `json:"create_time"`        // [Required]
	Type             string              `json:"type"`               // [Required]
	Items            []ResponseDataItems `json:"items"`              // [Required]
	Status           string              `json:"status"`             // [Required]
}
type RefOrderCode struct {
	OrderCode string `json:"order_code"` // [Required]
	Type      string `json:"type"`       // [Required]
}
type RemoveFulfillmentSkuRelationResponse struct {
	BaseResponse                                                // Common response fields
	Result       *BuildFulfillmentSkuRelationResponseDataResult `json:"result,omitempty"` //
}
type ResponseDataItems struct {
	Quantity           string `json:"quantity"`             // [Required]
	FulfillmentSkuId   string `json:"fulfillment_sku_id"`   // [Required]
	FulfillmentSkuCode string `json:"fulfillment_sku_code"` // [Required]
}
type ResponseDataSku struct {
	FulfillmentSkuId string `json:"fulfillment_sku_id"` // [Required]
	OperateLogCount  string `json:"operate_log_count"`  // [Required]
}
type ResponseDataStoreStocks struct {
	StoreCode string                         `json:"store_code"` // [Required]
	Stocks    *ResponseDataStoreStocksStocks `json:"stocks"`     // [Required]
}
type ResponseDataStoreStocksStocks struct {
	DamagedUnsellable *Pending `json:"damagedUnsellable"` // [Required]
	Transfer          *Pending `json:"transfer"`          // [Required]
	Pending           *Pending `json:"pending"`           // [Required]
	Unsellable        *Pending `json:"unsellable"`        // [Required]
	ExpiredUnsellable *Pending `json:"expiredUnsellable"` // [Required]
	Sellable          *Pending `json:"sellable"`          // [Required]
}
type ReturnCancellationResponse struct {
	BaseResponse        // Common response fields
	ErrorMessage string `json:"error_message,omitempty"` //
}
type ReturnOrderCreationResponse struct {
	BaseResponse                                 // Common response fields
	Response     ReturnOrderCreationResponseData `json:"data"`                    // Response data
	ErrorMessage string                          `json:"error_message,omitempty"` //
}
type ReturnOrderCreationResponseData struct {
	ReturnId string `json:"return_id"` // [Required]
}
type SampleRule struct {
	RuleRegularExpression string `json:"rule_regular_expression"` // [Required]
	RuleDesc              string `json:"rule_desc"`               // [Required]
	RuleImgUrl            string `json:"rule_img_url"`            // [Required]
	RuleSample            string `json:"rule_sample"`             // [Required]
}
type SetStockRuleResponse struct {
	BaseResponse        // Common response fields
	ErrorMessage string `json:"error_message,omitempty"` //
}
type SnSample struct {
	SampleSeq      string       `json:"sample_seq"`       // [Required]
	SampleDesc     string       `json:"sample_desc"`      // [Required]
	SampleRuleList []SampleRule `json:"sample_rule_list"` // [Required]
}
type Stocks struct {
	WarehouseCode string          `json:"warehouse_code"` // [Required]
	ChannelStocks []ChannelStocks `json:"channel_stocks"` // [Required]
}
type StoreStocks struct {
	StoreCode string             `json:"store_code"` // [Required]
	Stocks    *StoreStocksStocks `json:"stocks"`     // [Required]
}
type StoreStocksStocks struct {
	Pending    *Pending `json:"pending"`    // [Required]
	Unsellable *Pending `json:"unsellable"` // [Required]
	Sellable   *Pending `json:"sellable"`   // [Required]
}
type UpdateFulfillmentSkuDecoupleResponse struct {
	BaseResponse        // Common response fields
	Response     string `json:"data"`                    // Response data
	ErrorMessage string `json:"error_message,omitempty"` //
}
type UploadWaybillResponse struct {
	BaseResponse        // Common response fields
	ErrorMessage string `json:"error_message,omitempty"` //
}
