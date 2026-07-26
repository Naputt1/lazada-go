package golazada

type Attributes struct {
	ShortDescription string `json:"short_description"` // [Required]
	Name             string `json:"name"`              // [Required]
	Description      string `json:"description"`       // [Required]
	NameEngravement  string `json:"name_engravement"`  // [Required]
	WarrantyType     string `json:"warranty_type"`     // [Required]
	GiftWrapping     string `json:"gift_wrapping"`     // [Required]
	Brand            string `json:"brand"`             // [Required]
}
type BatchDeliverJitPurchaseOrderResponse struct {
	BaseResponse         // Common response fields
	Result       *Result `json:"result,omitempty"` //
}
type BizSupplement struct {
	ItemType int64 `json:"item_type"` // [Required]
}
type Data struct {
	ErrorMessage    string        `json:"error_message"`     // [Required]
	PickupNo        string        `json:"pickup_no"`         // [Required]
	AllowDateRange  []interface{} `json:"allow_date_range"`  // [Required]
	PurchaseOrderNo string        `json:"purchase_order_no"` // [Required]
	Status          string        `json:"status"`            // [Required]
}
type EditChoiceSkuStockResponse struct {
	BaseResponse                                // Common response fields
	Response     EditChoiceSkuStockResponseData `json:"data"` // Response data
}
type EditChoiceSkuStockResponseData struct {
	SuccessSku []string      `json:"success_sku"` // [Required]
	FailedSku  []interface{} `json:"failed_sku"`  // [Required]
}
type GetChoiceProductItemResponse struct {
	BaseResponse                                  // Common response fields
	Response     GetChoiceProductItemResponseData `json:"data"` // Response data
}
type GetChoiceProductItemResponseData struct {
	CreatedTime     string         `json:"created_time"`     // [Required]
	UpdatedTime     string         `json:"updated_time"`     // [Required]
	Images          string         `json:"images"`           // [Required]
	Skus            []Skus         `json:"skus"`             // [Required]
	ItemId          int64          `json:"item_id"`          // [Required]
	HiddenStatus    string         `json:"hiddenStatus"`     // [Required]
	BizSupplement   *BizSupplement `json:"bizSupplement"`    // [Required]
	SuspendedSkus   []interface{}  `json:"suspendedSkus"`    // [Required]
	SubStatus       string         `json:"subStatus"`        // [Required]
	Variation       *Variation     `json:"variation"`        // [Required]
	TrialProduct    string         `json:"trialProduct"`     // [Required]
	RejectReason    []RejectReason `json:"rejectReason"`     // [Required]
	PrimaryCategory int64          `json:"primary_category"` // [Required]
	MarketImages    string         `json:"marketImages"`     // [Required]
	Attributes      *Attributes    `json:"attributes"`       // [Required]
	HiddenReason    string         `json:"hiddenReason"`     // [Required]
	Status          string         `json:"status"`           // [Required]
}
type GetChoiceProductsResponse struct {
	BaseResponse                               // Common response fields
	Response     GetChoiceProductsResponseData `json:"data"` // Response data
}
type GetChoiceProductsResponseData struct {
	TotalProducts int64      `json:"total_products"` // [Required]
	Products      []Products `json:"products"`       // [Required]
}
type GetChoiceSellerResponse struct {
	BaseResponse                             // Common response fields
	Response     GetChoiceSellerResponseData `json:"data"` // Response data
}
type GetChoiceSellerResponseData struct {
	NameCompany string `json:"name_company"` // [Required]
	Name        string `json:"name"`         // [Required]
	Verified    string `json:"verified"`     // [Required]
	Location    string `json:"location"`     // [Required]
	SellerId    int64  `json:"seller_id"`    // [Required]
	Email       string `json:"email"`        // [Required]
	ShortCode   string `json:"short_code"`   // [Required]
	Cb          string `json:"cb"`           // [Required]
	Status      string `json:"status"`       // [Required]
}
type GetChoiceSkuItemRelationBySkuResponse struct {
	BaseResponse                                           // Common response fields
	Response     GetChoiceSkuItemRelationBySkuResponseData `json:"data"` // Response data
}
type GetChoiceSkuItemRelationBySkuResponseData struct {
	Site         string `json:"site"`            // [Required]
	ItemId       int64  `json:"item_id"`         // [Required]
	ScItemUserId string `json:"sc_item_user_id"` // [Required]
	SkuId        int64  `json:"sku_id"`          // [Required]
	Source       string `json:"source"`          // [Required]
	Barcode      string `json:"barcode"`         // [Required]
	SellerId     int64  `json:"seller_id"`       // [Required]
	ScItemId     string `json:"sc_item_id"`      // [Required]
}
type MultiWarehouseInventories struct {
	Quantity      int64  `json:"quantity"`      // [Required]
	WarehouseCode string `json:"warehouseCode"` // [Required]
}
type PackageJitPurchaseOrderResponse struct {
	BaseResponse                     // Common response fields
	Result       *ResponseDataResult `json:"result,omitempty"` //
}
type PrintJitPurchaseOrderAndItemResponse struct {
	BaseResponse                                                 // Common response fields
	Result       *PrintJitPurchaseOrderAndItemResponseDataResult `json:"result,omitempty"` //
}
type PrintJitPurchaseOrderAndItemResponseDataResult struct {
	ErrorMessage string                  `json:"error_message"` // [Required]
	Data         *ResponseDataResultData `json:"data"`          // [Required]
	Success      bool                    `json:"success"`       // [Required]
	ErrorCode    string                  `json:"error_code"`    // [Required]
}
type PrintPickuoOrderResponse struct {
	BaseResponse                                                 // Common response fields
	Result       *PrintJitPurchaseOrderAndItemResponseDataResult `json:"result,omitempty"` //
}
type Products struct {
	CreatedTime     string         `json:"created_time"`     // [Required]
	UpdatedTime     string         `json:"updated_time"`     // [Required]
	Images          string         `json:"images"`           // [Required]
	Skus            []Skus         `json:"skus"`             // [Required]
	ItemId          int64          `json:"item_id"`          // [Required]
	HiddenStatus    string         `json:"hiddenStatus"`     // [Required]
	BizSupplement   *BizSupplement `json:"bizSupplement"`    // [Required]
	SuspendedSkus   []interface{}  `json:"suspendedSkus"`    // [Required]
	SubStatus       string         `json:"subStatus"`        // [Required]
	TrialProduct    string         `json:"trialProduct"`     // [Required]
	RejectReason    []RejectReason `json:"rejectReason"`     // [Required]
	PrimaryCategory int64          `json:"primary_category"` // [Required]
	MarketImages    string         `json:"marketImages"`     // [Required]
	Attributes      *Attributes    `json:"attributes"`       // [Required]
	HiddenReason    string         `json:"hiddenReason"`     // [Required]
	Status          string         `json:"status"`           // [Required]
}
type QueryListJitPurchaseOrderResponse struct {
	BaseResponse                                              // Common response fields
	Result       *QueryListJitPurchaseOrderResponseDataResult `json:"result,omitempty"` //
}
type QueryListJitPurchaseOrderResponseDataResult struct {
	ErrorMessage string                                            `json:"error_message"` // [Required]
	Data         []QueryListJitPurchaseOrderResponseDataResultData `json:"data"`          // [Required]
	Success      bool                                              `json:"success"`       // [Required]
	TotalCount   int64                                             `json:"total_count"`   // [Required]
	PageIndex    int64                                             `json:"page_index"`    // [Required]
	TotalPage    int64                                             `json:"total_page"`    // [Required]
	ErrorCode    string                                            `json:"error_code"`    // [Required]
	PageSize     int64                                             `json:"page_size"`     // [Required]
}
type QueryListJitPurchaseOrderResponseDataResultData struct {
	GmtCreate               string        `json:"gmt_create"`                // [Required]
	StoreAddress            string        `json:"store_address"`             // [Required]
	GmtModified             string        `json:"gmt_modified"`              // [Required]
	FulfillmentCancelStatus string        `json:"fulfillment_cancel_status"` // [Required]
	TradeOrderIdList        []interface{} `json:"trade_order_id_list"`       // [Required]
	StoreContactName        string        `json:"store_contact_name"`        // [Required]
	DeliveryMethod          string        `json:"delivery_method"`           // [Required]
	GmtArriveTime           string        `json:"gmt_arrive_time"`           // [Required]
	TotalQuantity           string        `json:"total_quantity"`            // [Required]
	StoreName               string        `json:"store_name"`                // [Required]
	StoreContactPhone       string        `json:"store_contact_phone"`       // [Required]
	SupplierName            string        `json:"supplier_name"`             // [Required]
	ExtFields               string        `json:"ext_fields"`                // [Required]
	SellerId                int64         `json:"seller_id"`                 // [Required]
	StoreCode               string        `json:"store_code"`                // [Required]
	Creator                 string        `json:"creator"`                   // [Required]
	BizStatus               string        `json:"biz_status"`                // [Required]
	ConsignOrderNoList      string        `json:"consign_order_no_list"`     // [Required]
	TotalSkuCount           string        `json:"total_sku_count"`           // [Required]
	GmtExceptArriveTime     string        `json:"gmt_except_arrive_time"`    // [Required]
	PickupOrderNo           string        `json:"pickup_order_no"`           // [Required]
	SiteId                  string        `json:"site_id"`                   // [Required]
	LogisticsNoList         string        `json:"logistics_no_list"`         // [Required]
	SupplierId              string        `json:"supplier_id"`               // [Required]
	SupplierCode            string        `json:"supplier_code"`             // [Required]
	PurchaseOrderNo         string        `json:"purchase_order_no"`         // [Required]
}
type QueryListPurchaseItemResponse struct {
	BaseResponse                                          // Common response fields
	Result       *QueryListPurchaseItemResponseDataResult `json:"result,omitempty"` //
}
type QueryListPurchaseItemResponseDataResult struct {
	ErrorMessage string                                        `json:"error_message"` // [Required]
	Data         []QueryListPurchaseItemResponseDataResultData `json:"data"`          // [Required]
	Success      bool                                          `json:"success"`       // [Required]
	TotalCount   int64                                         `json:"total_count"`   // [Required]
	PageIndex    int64                                         `json:"page_index"`    // [Required]
	TotalPage    int64                                         `json:"total_page"`    // [Required]
	ErrorCode    string                                        `json:"error_code"`    // [Required]
	PageSize     int64                                         `json:"page_size"`     // [Required]
}
type QueryListPurchaseItemResponseDataResultData struct {
	ReceivedDefectiveQty string        `json:"received_defective_qty"` // [Required]
	SkuId                int64         `json:"sku_id"`                 // [Required]
	Barcodes             []interface{} `json:"barcodes"`               // [Required]
	ProductTitle         string        `json:"product_title"`          // [Required]
	BuyerQty             string        `json:"buyer_qty"`              // [Required]
	ScItemCode           string        `json:"sc_item_code"`           // [Required]
	ImgUrl               string        `json:"img_url"`                // [Required]
	ScItemName           string        `json:"sc_item_name"`           // [Required]
	ProductId            int64         `json:"product_id"`             // [Required]
	SellerSku            string        `json:"seller_sku"`             // [Required]
	ReceivedNormalQty    string        `json:"received_normal_qty"`    // [Required]
	PurchaseOrderNo      string        `json:"purchase_order_no"`      // [Required]
	ScItemId             string        `json:"sc_item_id"`             // [Required]
}
type QueryPickupOrderResponse struct {
	BaseResponse                                     // Common response fields
	Result       *QueryPickupOrderResponseDataResult `json:"result,omitempty"` //
}
type QueryPickupOrderResponseDataResult struct {
	ErrorMessage string                                  `json:"error_message"` // [Required]
	Data         *QueryPickupOrderResponseDataResultData `json:"data"`          // [Required]
	Success      bool                                    `json:"success"`       // [Required]
	ErrorCode    string                                  `json:"error_code"`    // [Required]
}
type QueryPickupOrderResponseDataResultData struct {
	ActualPickupTime      string   `json:"actual_pickup_time"`       // [Required]
	Reason                string   `json:"reason"`                   // [Required]
	EstimatedVolume       string   `json:"estimated_volume"`         // [Required]
	PurchaseOrderNoList   []string `json:"purchase_order_no_list"`   // [Required]
	CreateTime            string   `json:"create_time"`              // [Required]
	CarDriverPhone        string   `json:"car_driver_phone"`         // [Required]
	ActualArriveTime      string   `json:"actual_arrive_time"`       // [Required]
	ShipperPhone          string   `json:"shipper_phone"`            // [Required]
	CarDriverName         string   `json:"car_driver_name"`          // [Required]
	EstimatedPickupTime   string   `json:"estimated_pickup_time"`    // [Required]
	ActualWeight          string   `json:"actual_weight"`            // [Required]
	UpdateTime            string   `json:"update_time"`              // [Required]
	ReceiveStoreCode      string   `json:"receive_store_code"`       // [Required]
	PickupOrderNo         string   `json:"pickup_order_no"`          // [Required]
	ReceiveStoreAddress   string   `json:"receive_store_address"`    // [Required]
	CarNumber             string   `json:"car_number"`               // [Required]
	EstimatedWeight       string   `json:"estimated_weight"`         // [Required]
	LogisticsNoList       []string `json:"logistics_no_list"`        // [Required]
	ActualLogisticsNoList []string `json:"actual_logistics_no_list"` // [Required]
	EstimatedBoxNumber    string   `json:"estimated_box_number"`     // [Required]
	ShipperName           string   `json:"shipper_name"`             // [Required]
	ShipperAddress        string   `json:"shipper_address"`          // [Required]
	Status                string   `json:"status"`                   // [Required]
}
type ResponseDataResult struct {
	ErrorMessage string      `json:"error_message"` // [Required]
	Data         *ResultData `json:"data"`          // [Required]
	Success      bool        `json:"success"`       // [Required]
	ErrorCode    string      `json:"error_code"`    // [Required]
}
type ResponseDataResultData struct {
	File string `json:"file"` // [Required]
}
type Result struct {
	ErrorMessage string `json:"error_message"` // [Required]
	Data         []Data `json:"data"`          // [Required]
	Success      bool   `json:"success"`       // [Required]
	ErrorCode    string `json:"error_code"`    // [Required]
}
type ResultData struct {
	Status string `json:"status"` // [Required]
}
type Skus struct {
	Status                    string                      `json:"Status"`                    // [Required]
	Quantity                  int64                       `json:"quantity"`                  // [Required]
	ProductWeight             string                      `json:"product_weight"`            // [Required]
	Images                    []string                    `json:"Images"`                    // [Required]
	SellerSku                 string                      `json:"SellerSku"`                 // [Required]
	ShopSku                   string                      `json:"ShopSku"`                   // [Required]
	CurrencyUnit              string                      `json:"currency_unit"`             // [Required]
	MultiWarehouseInventories []MultiWarehouseInventories `json:"multiWarehouseInventories"` // [Required]
	SkuSupplyPrice            int64                       `json:"sku_supply_price"`          // [Required]
	PackageWidth              string                      `json:"package_width"`             // [Required]
	SpecialToTime             string                      `json:"special_to_time"`           // [Required]
	SpecialFromTime           string                      `json:"special_from_time"`         // [Required]
	PackageHeight             string                      `json:"package_height"`            // [Required]
	PackageLength             string                      `json:"package_length"`            // [Required]
	PackageWeight             string                      `json:"package_weight"`            // [Required]
	Available                 int64                       `json:"Available"`                 // [Required]
	SkuId                     int64                       `json:"SkuId"`                     // [Required]
	SpecialToDate             string                      `json:"special_to_date"`           // [Required]
}
