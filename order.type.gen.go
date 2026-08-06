package golazada

type Document struct {
	File         FlexString `json:"file"`          // [Required]
	MimeType     FlexString `json:"mime_type"`     // [Required]
	DocumentType FlexString `json:"document_type"` // [Required]
}
type GetDocumentResponse struct {
	BaseResponse                         // Common response fields
	Response     GetDocumentResponseData `json:"data"` // Response data
}
type GetDocumentResponseData struct {
	Document *Document `json:"document"` // [Required]
}
type GetMultipleOrderItemsResponse struct {
	BaseResponse                                     // Common response fields
	Response     []GetMultipleOrderItemsResponseData `json:"data"` // Response data
}
type GetMultipleOrderItemsResponseData struct {
	OrderNumber FlexString   `json:"order_number"` // [Required]
	OrderId     FlexInt      `json:"order_id"`     // [Required]
	OrderItems  []OrderItems `json:"order_items"`  // [Required]
}
type GetOrderItemsResponse struct {
	BaseResponse                           // Common response fields
	Response     GetOrderItemsResponseData `json:"data"` // Response data
}
type GetOrderItemsResponseData struct {
	PickUpStoreInfo               *PickUpStoreInfo `json:"pick_up_store_info"`               // [Required]
	TaxAmount                     FlexString       `json:"tax_amount"`                       // [Required]
	Reason                        FlexString       `json:"reason"`                           // [Required]
	SlaTimeStamp                  FlexString       `json:"sla_time_stamp"`                   // [Required]
	ShowGiftwrappingTag           FlexString       `json:"show_giftwrapping_tag"`            // [Required]
	VoucherSeller                 FlexString       `json:"voucher_seller"`                   // [Required]
	PurchaseOrderId               FlexString       `json:"purchase_order_id"`                // [Required]
	PaymentTime                   FlexString       `json:"payment_time"`                     // [Required]
	VoucherCodeSeller             FlexString       `json:"voucher_code_seller"`              // [Required]
	VoucherCode                   FlexString       `json:"voucher_code"`                     // [Required]
	PackageId                     FlexString       `json:"package_id"`                       // [Required]
	BuyerId                       FlexString       `json:"buyer_id"`                         // [Required]
	Variation                     FlexString       `json:"variation"`                        // [Required]
	IsCancelPending               FlexString       `json:"is_cancel_pending"`                // [Required]
	BizGroup                      FlexString       `json:"biz_group"`                        // [Required]
	ProductId                     FlexInt          `json:"product_id"`                       // [Required]
	VoucherCodePlatform           FlexString       `json:"voucher_code_platform"`            // [Required]
	PurchaseOrderNumber           FlexString       `json:"purchase_order_number"`            // [Required]
	Sku                           FlexString       `json:"sku"`                              // [Required]
	GiftWrapping                  FlexString       `json:"gift_wrapping"`                    // [Required]
	ScheduleDeliveryStartTimeslot FlexString       `json:"schedule_delivery_start_timeslot"` // [Required]
	OrderType                     FlexString       `json:"order_type"`                       // [Required]
	InvoiceNumber                 FlexString       `json:"invoice_number"`                   // [Required]
	ShowPersonalizationTag        FlexString       `json:"show_personalization_tag"`         // [Required]
	CanEscalatePickup             FlexString       `json:"can_escalate_pickup"`              // [Required]
	CancelTriggerTime             FlexString       `json:"cancel_trigger_time"`              // [Required]
	CancelReturnInitiator         FlexString       `json:"cancel_return_initiator"`          // [Required]
	ShopSku                       FlexString       `json:"shop_sku"`                         // [Required]
	IsReroute                     FlexString       `json:"is_reroute"`                       // [Required]
	StagePayStatus                FlexString       `json:"stage_pay_status"`                 // [Required]
	SkuId                         FlexInt          `json:"sku_id"`                           // [Required]
	TrackingCodePre               FlexString       `json:"tracking_code_pre"`                // [Required]
	OrderItemId                   FlexInt          `json:"order_item_id"`                    // [Required]
	ModelQuantityPurchased        FlexInt          `json:"model_quantity_purchased"`         //
	ShopId                        FlexString       `json:"shop_id"`                          // [Required]
	OrderFlag                     FlexString       `json:"order_flag"`                       // [Required]
	IsFbl                         FlexString       `json:"is_fbl"`                           // [Required]
	Name                          FlexString       `json:"name"`                             // [Required]
	DeliveryOptionSof             FlexString       `json:"delivery_option_sof"`              // [Required]
	OrderId                       FlexInt          `json:"order_id"`                         // [Required]
	FulfillmentSla                FlexString       `json:"fulfillment_sla"`                  // [Required]
	NeedCancelConfirm             FlexString       `json:"need_cancel_confirm"`              // [Required]
	Status                        FlexString       `json:"status"`                           // [Required]
	ProductMainImage              FlexString       `json:"product_main_image"`               // [Required]
	VoucherPlatform               FlexString       `json:"voucher_platform"`                 // [Required]
	PaidPrice                     FlexString       `json:"paid_price"`                       // [Required]
	ProductDetailUrl              FlexString       `json:"product_detail_url"`               // [Required]
	WarehouseCode                 FlexString       `json:"warehouse_code"`                   // [Required]
	PromisedShippingTime          FlexString       `json:"promised_shipping_time"`           // [Required]
	ShippingType                  FlexString       `json:"shipping_type"`                    // [Required]
	CreatedAt                     FlexString       `json:"created_at"`                       // [Required]
	SupplyPrice                   FlexString       `json:"supply_price"`                     // [Required]
	Mp3Order                      FlexString       `json:"mp3_order"`                        // [Required]
	VoucherSellerLpi              FlexString       `json:"voucher_seller_lpi"`               // [Required]
	ShippingFeeDiscountPlatform   FlexString       `json:"shipping_fee_discount_platform"`   // [Required]
	Personalization               FlexString       `json:"personalization"`                  // [Required]
	WalletCredits                 FlexString       `json:"wallet_credits"`                   // [Required]
	ReverseOrderId                FlexString       `json:"reverse_order_id"`                 // [Required]
	UpdatedAt                     FlexString       `json:"updated_at"`                       // [Required]
	Currency                      FlexString       `json:"currency"`                         // [Required]
	ShippingProviderType          FlexString       `json:"shipping_provider_type"`           // [Required]
	VoucherPlatformLpi            FlexString       `json:"voucher_platform_lpi"`             // [Required]
	ShippingFeeOriginal           FlexString       `json:"shipping_fee_original"`            // [Required]
	ScheduleDeliveryEndTimeslot   FlexString       `json:"schedule_delivery_end_timeslot"`   // [Required]
	ItemPrice                     FlexString       `json:"item_price"`                       // [Required]
	IsDigital                     FlexString       `json:"is_digital"`                       // [Required]
	ShippingServiceCost           FlexString       `json:"shipping_service_cost"`            // [Required]
	TrackingCode                  FlexString       `json:"tracking_code"`                    // [Required]
	ShippingFeeDiscountSeller     FlexString       `json:"shipping_fee_discount_seller"`     // [Required]
	ShippingAmount                FlexString       `json:"shipping_amount"`                  // [Required]
	ReasonDetail                  FlexString       `json:"reason_detail"`                    // [Required]
	ReturnStatus                  FlexString       `json:"return_status"`                    // [Required]
	SemiManaged                   FlexString       `json:"semi_managed"`                     // [Required]
	ShipmentProvider              FlexString       `json:"shipment_provider"`                // [Required]
	PriorityFulfillmentTag        FlexString       `json:"priority_fulfillment_tag"`         // [Required]
	VoucherAmount                 FlexString       `json:"voucher_amount"`                   // [Required]
	SupplyPriceCurrency           FlexString       `json:"supply_price_currency"`            // [Required]
	DigitalDeliveryInfo           FlexString       `json:"digital_delivery_info"`            // [Required]
	ExtraAttributes               FlexString       `json:"extra_attributes"`                 // [Required]
}
type GetOrderResponse struct {
	BaseResponse                      // Common response fields
	Response     GetOrderResponseData `json:"data"` // Response data
}
type GetOrderResponseData struct {
	Voucher                     FlexString                  `json:"voucher"`                        // [Required]
	WarehouseCode               FlexString                  `json:"warehouse_code"`                 // [Required]
	OrderNumber                 FlexString                  `json:"order_number"`                   // [Required]
	CreatedAt                   FlexString                  `json:"created_at"`                     // [Required]
	VoucherCode                 FlexString                  `json:"voucher_code"`                   // [Required]
	GiftOption                  FlexString                  `json:"gift_option"`                    // [Required]
	IsCancelPending             FlexString                  `json:"is_cancel_pending"`              // [Required]
	ShippingFeeDiscountPlatform FlexString                  `json:"shipping_fee_discount_platform"` // [Required]
	CustomerLastName            FlexString                  `json:"customer_last_name"`             // [Required]
	UpdatedAt                   FlexString                  `json:"updated_at"`                     // [Required]
	PromisedShippingTimes       FlexString                  `json:"promised_shipping_times"`        // [Required]
	Price                       FlexFloat                   `json:"price"`                          // [Required]
	NationalRegistrationNumber  FlexString                  `json:"national_registration_number"`   // [Required]
	ShippingFeeOriginal         FlexString                  `json:"shipping_fee_original"`          // [Required]
	PaymentMethod               FlexString                  `json:"payment_method"`                 // [Required]
	RecipientInfo               *RecipientInfo              `json:"recipient_info"`                 // [Required]
	BuyerNote                   FlexString                  `json:"buyer_note"`                     // [Required]
	CustomerFirstName           FlexString                  `json:"customer_first_name"`            // [Required]
	ShippingFeeDiscountSeller   FlexString                  `json:"shipping_fee_discount_seller"`   // [Required]
	ShippingFee                 FlexString                  `json:"shipping_fee"`                   // [Required]
	BranchNumber                FlexString                  `json:"branch_number"`                  // [Required]
	TaxCode                     FlexString                  `json:"tax_code"`                       // [Required]
	ItemsCount                  FlexString                  `json:"items_count"`                    // [Required]
	DeliveryInfo                FlexString                  `json:"delivery_info"`                  // [Required]
	Statuses                    []interface{}               `json:"statuses"`                       // [Required]
	AddressBilling              *ResponseDataAddressBilling `json:"address_billing"`                // [Required]
	ExtraAttributes             FlexString                  `json:"extra_attributes"`               // [Required]
	OrderId                     FlexInt                     `json:"order_id"`                       // [Required]
	NeedCancelConfirm           FlexString                  `json:"need_cancel_confirm"`            // [Required]
	GiftMessage                 FlexString                  `json:"gift_message"`                   // [Required]
	Remarks                     FlexString                  `json:"remarks"`                        // [Required]
	AddressShipping             *ResponseDataAddressBilling `json:"address_shipping"`               // [Required]
}
type GetOrdersResponse struct {
	BaseResponse                       // Common response fields
	Response     GetOrdersResponseData `json:"data"` // Response data
}
type GetOrdersResponseData struct {
	Count      FlexInt              `json:"count"`      // [Required]
	CountTotal FlexInt              `json:"countTotal"` // [Required]
	Orders     []ResponseDataOrders `json:"orders"`     // [Required]
}
type GetOVOOrdersResponse struct {
	BaseResponse                                 // Common response fields
	Result       *GetOVOOrdersResponseDataResult `json:"result,omitempty"` //
}
type GetOVOOrdersResponseDataResult struct {
	Success     bool          `json:"success"`     // [Required]
	ErrorCode   FlexString    `json:"errorCode"`   // [Required]
	TradeOrders []TradeOrders `json:"tradeOrders"` // [Required]
}
type OrderCancelValidateResponse struct {
	BaseResponse                                 // Common response fields
	Response     OrderCancelValidateResponseData `json:"data"` // Response data
}
type OrderCancelValidateResponseData struct {
	TipContent    FlexString      `json:"tip_content"`    // [Required]
	ReasonOptions []ReasonOptions `json:"reason_options"` // [Required]
	TipType       FlexString      `json:"tip_type"`       // [Required]
}
type OrderItems struct {
	TaxAmount                     FlexString       `json:"tax_amount"`                       // [Required]
	PickUpStoreInfo               *PickUpStoreInfo `json:"pick_up_store_info"`               // [Required]
	Reason                        FlexString       `json:"reason"`                           // [Required]
	SlaTimeStamp                  FlexString       `json:"sla_time_stamp"`                   // [Required]
	PurchaseOrderId               FlexString       `json:"purchase_order_id"`                // [Required]
	VoucherSeller                 FlexString       `json:"voucher_seller"`                   // [Required]
	PaymentTime                   FlexString       `json:"payment_time"`                     // [Required]
	VoucherCodeSeller             FlexString       `json:"voucher_code_seller"`              // [Required]
	VoucherCode                   FlexString       `json:"voucher_code"`                     // [Required]
	PackageId                     FlexString       `json:"package_id"`                       // [Required]
	BuyerId                       FlexString       `json:"buyer_id"`                         // [Required]
	Variation                     FlexString       `json:"variation"`                        // [Required]
	IsCancelPending               FlexString       `json:"is_cancel_pending"`                // [Required]
	BizGroup                      FlexString       `json:"biz_group"`                        // [Required]
	VoucherCodePlatform           FlexString       `json:"voucher_code_platform"`            // [Required]
	PurchaseOrderNumber           FlexString       `json:"purchase_order_number"`            // [Required]
	ShowGiftWrappingTag           FlexString       `json:"show_gift_wrapping_tag"`           // [Required]
	Sku                           FlexString       `json:"sku"`                              // [Required]
	GiftWrapping                  FlexString       `json:"gift_wrapping"`                    // [Required]
	ScheduleDeliveryStartTimeslot FlexString       `json:"schedule_delivery_start_timeslot"` // [Required]
	InvoiceNumber                 FlexString       `json:"invoice_number"`                   // [Required]
	OrderType                     FlexString       `json:"order_type"`                       // [Required]
	ShowPersonalizationTag        FlexString       `json:"show_personalization_tag"`         // [Required]
	CanEscalatePickup             FlexString       `json:"can_escalate_pickup"`              // [Required]
	CancelTriggerTime             FlexString       `json:"cancel_trigger_time"`              // [Required]
	CancelReturnInitiator         FlexString       `json:"cancel_return_initiator"`          // [Required]
	ShopSku                       FlexString       `json:"shop_sku"`                         // [Required]
	IsReroute                     FlexString       `json:"is_reroute"`                       // [Required]
	StagePayStatus                FlexString       `json:"stage_pay_status"`                 // [Required]
	SkuId                         FlexInt          `json:"sku_id"`                           // [Required]
	TrackingCodePre               FlexString       `json:"tracking_code_pre"`                // [Required]
	OrderItemId                   FlexInt          `json:"order_item_id"`                    // [Required]
	ModelQuantityPurchased        FlexInt          `json:"model_quantity_purchased"`         //
	ShopId                        FlexString       `json:"shop_id"`                          // [Required]
	OrderFlag                     FlexString       `json:"order_flag"`                       // [Required]
	IsFbl                         FlexString       `json:"is_fbl"`                           // [Required]
	Name                          FlexString       `json:"name"`                             // [Required]
	DeliveryOptionSof             FlexString       `json:"delivery_option_sof"`              // [Required]
	OrderId                       FlexInt          `json:"order_id"`                         // [Required]
	FulfillmentSla                FlexString       `json:"fulfillment_sla"`                  // [Required]
	NeedCancelConfirm             FlexString       `json:"need_cancel_confirm"`              // [Required]
	Status                        FlexString       `json:"status"`                           // [Required]
	PaidPrice                     FlexString       `json:"paid_price"`                       // [Required]
	ProductMainImage              FlexString       `json:"product_main_image"`               // [Required]
	VoucherPlatform               FlexString       `json:"voucher_platform"`                 // [Required]
	ProductDetailUrl              FlexString       `json:"product_detail_url"`               // [Required]
	PromisedShippingTime          FlexString       `json:"promised_shipping_time"`           // [Required]
	WarehouseCode                 FlexString       `json:"warehouse_code"`                   // [Required]
	ShippingType                  FlexString       `json:"shipping_type"`                    // [Required]
	CreatedAt                     FlexString       `json:"created_at"`                       // [Required]
	SupplyPrice                   FlexString       `json:"supply_price"`                     // [Required]
	Mp3Order                      FlexString       `json:"mp3_order"`                        // [Required]
	VoucherSellerLpi              FlexString       `json:"voucher_seller_lpi"`               // [Required]
	ShippingFeeDiscountPlatform   FlexString       `json:"shipping_fee_discount_platform"`   // [Required]
	Personalization               FlexString       `json:"personalization"`                  // [Required]
	WalletCredits                 FlexString       `json:"wallet_credits"`                   // [Required]
	ReverseOrderId                FlexString       `json:"reverse_order_id"`                 // [Required]
	UpdatedAt                     FlexString       `json:"updated_at"`                       // [Required]
	Currency                      FlexString       `json:"currency"`                         // [Required]
	ShippingProviderType          FlexString       `json:"shipping_provider_type"`           // [Required]
	ShippingFeeOriginal           FlexString       `json:"shipping_fee_original"`            // [Required]
	VoucherPlatformLpi            FlexString       `json:"voucher_platform_lpi"`             // [Required]
	ScheduleDeliveryEndTimeslot   FlexString       `json:"schedule_delivery_end_timeslot"`   // [Required]
	IsDigital                     FlexString       `json:"is_digital"`                       // [Required]
	ItemPrice                     FlexString       `json:"item_price"`                       // [Required]
	ShippingServiceCost           FlexString       `json:"shipping_service_cost"`            // [Required]
	TrackingCode                  FlexString       `json:"tracking_code"`                    // [Required]
	ShippingFeeDiscountSeller     FlexString       `json:"shipping_fee_discount_seller"`     // [Required]
	ShippingAmount                FlexString       `json:"shipping_amount"`                  // [Required]
	ReasonDetail                  FlexString       `json:"reason_detail"`                    // [Required]
	ReturnStatus                  FlexString       `json:"return_status"`                    // [Required]
	SemiManaged                   FlexString       `json:"semi_managed"`                     // [Required]
	ShipmentProvider              FlexString       `json:"shipment_provider"`                // [Required]
	PriorityFulfillmentTag        FlexString       `json:"priority_fulfillment_tag"`         // [Required]
	VoucherAmount                 FlexString       `json:"voucher_amount"`                   // [Required]
	SupplyPriceCurrency           FlexString       `json:"supply_price_currency"`            // [Required]
	DigitalDeliveryInfo           FlexString       `json:"digital_delivery_info"`            // [Required]
	ExtraAttributes               FlexString       `json:"extra_attributes"`                 // [Required]
}
type OrdersAddressBilling struct {
	Country         FlexString `json:"country"`         // [Required]
	Address3        FlexString `json:"address3"`        // [Required]
	Address2        FlexString `json:"address2"`        // [Required]
	City            FlexString `json:"city"`            // [Required]
	Address1        FlexString `json:"address1"`        // [Required]
	Phone2          FlexString `json:"phone2"`          // [Required]
	LastName        FlexString `json:"last_name"`       // [Required]
	AddressDsitrict FlexString `json:"addressDsitrict"` // [Required]
	Phone           FlexString `json:"phone"`           // [Required]
	PostCode        FlexString `json:"post_code"`       // [Required]
	Address5        FlexString `json:"address5"`        // [Required]
	Address4        FlexString `json:"address4"`        // [Required]
	FirstName       FlexString `json:"first_name"`      // [Required]
}
type PickUpStoreInfo struct {
	PickUpStoreAddress  FlexString `json:"pick_up_store_address"`   // [Required]
	PickUpStoreName     FlexString `json:"pick_up_store_name"`      // [Required]
	PickUpStoreOpenHour []string   `json:"pick_up_store_open_hour"` // [Required]
	PickUpStoreCode     FlexString `json:"pick_up_store_code"`      // [Required]
}
type RecipientInfo struct {
	IdentifyNo    FlexString `json:"identify_no"`    // [Required]
	DetailAddress FlexString `json:"detail_address"` // [Required]
	PassportNo    FlexString `json:"passport_no"`    // [Required]
}
type ResponseDataAddressBilling struct {
	Country         FlexString `json:"country"`         // [Required]
	Address3        FlexString `json:"address3"`        // [Required]
	Address2        FlexString `json:"address2"`        // [Required]
	City            FlexString `json:"city"`            // [Required]
	Address1        FlexString `json:"address1"`        // [Required]
	Phone2          FlexString `json:"phone2"`          // [Required]
	LastName        FlexString `json:"last_name"`       // [Required]
	Phone           FlexString `json:"phone"`           // [Required]
	PostCode        FlexString `json:"post_code"`       // [Required]
	Address5        FlexString `json:"address5"`        // [Required]
	Address4        FlexString `json:"address4"`        // [Required]
	AddressDistrict FlexString `json:"addressDistrict"` // [Required]
	FirstName       FlexString `json:"first_name"`      // [Required]
}
type ResponseDataOrders struct {
	VoucherPlatform             FlexString            `json:"voucher_platform"`               // [Required]
	Voucher                     FlexString            `json:"voucher"`                        // [Required]
	WarehouseCode               FlexString            `json:"warehouse_code"`                 // [Required]
	OrderNumber                 FlexString            `json:"order_number"`                   // [Required]
	VoucherSeller               FlexString            `json:"voucher_seller"`                 // [Required]
	CreatedAt                   FlexString            `json:"created_at"`                     // [Required]
	VoucherCode                 FlexString            `json:"voucher_code"`                   // [Required]
	GiftOption                  FlexString            `json:"gift_option"`                    // [Required]
	IsCancelPending             FlexString            `json:"is_cancel_pending"`              // [Required]
	ShippingFeeDiscountPlatform FlexString            `json:"shipping_fee_discount_platform"` // [Required]
	CustomerLastName            FlexString            `json:"customer_last_name"`             // [Required]
	PromisedShippingTimes       FlexString            `json:"promised_shipping_times"`        // [Required]
	UpdatedAt                   FlexString            `json:"updated_at"`                     // [Required]
	Price                       FlexFloat             `json:"price"`                          // [Required]
	NationalRegistrationNumber  FlexString            `json:"national_registration_number"`   // [Required]
	ShippingFeeOriginal         FlexString            `json:"shipping_fee_original"`          // [Required]
	PaymentMethod               FlexString            `json:"payment_method"`                 // [Required]
	AddressUpdatedAt            FlexString            `json:"address_updated_at"`             // [Required]
	RecipientInfo               *RecipientInfo        `json:"recipient_info"`                 // [Required]
	BuyerNote                   FlexString            `json:"buyer_note"`                     // [Required]
	CustomerFirstName           FlexString            `json:"customer_first_name"`            // [Required]
	ShippingFeeDiscountSeller   FlexString            `json:"shipping_fee_discount_seller"`   // [Required]
	ShippingFee                 FlexString            `json:"shipping_fee"`                   // [Required]
	BranchNumber                FlexString            `json:"branch_number"`                  // [Required]
	TaxCode                     FlexString            `json:"tax_code"`                       // [Required]
	ItemsCount                  FlexString            `json:"items_count"`                    // [Required]
	DeliveryInfo                FlexString            `json:"delivery_info"`                  // [Required]
	Statuses                    []interface{}         `json:"statuses"`                       // [Required]
	AddressBilling              *OrdersAddressBilling `json:"address_billing"`                // [Required]
	ExtraAttributes             FlexString            `json:"extra_attributes"`               // [Required]
	OrderId                     FlexInt               `json:"order_id"`                       // [Required]
	NeedCancelConfirm           FlexString            `json:"need_cancel_confirm"`            // [Required]
	Remarks                     FlexString            `json:"remarks"`                        // [Required]
	GiftMessage                 FlexString            `json:"gift_message"`                   // [Required]
	AddressShipping             *OrdersAddressBilling `json:"address_shipping"`               // [Required]
}
type SetInvoiceNumberResponse struct {
	BaseResponse                              // Common response fields
	Response     SetInvoiceNumberResponseData `json:"data"` // Response data
}
type SetInvoiceNumberResponseData struct {
	OrderItemId   int64      `json:"order_item_id"`  // [Required]
	InvoiceNumber FlexString `json:"invoice_number"` // [Required]
}
type TradeOrderLines struct {
	DeliveredTime    FlexString `json:"deliveredTime"`    // [Required]
	TradeOrderLineId FlexString `json:"tradeOrderLineId"` // [Required]
	DeliveryStatus   FlexString `json:"deliveryStatus"`   // [Required]
	ReverseStatus    FlexString `json:"reverseStatus"`    // [Required]
}
type TradeOrders struct {
	TradeOrderId    FlexString        `json:"tradeOrderId"`    // [Required]
	PaymentMethod   FlexString        `json:"paymentMethod"`   // [Required]
	PaidTime        FlexString        `json:"paidTime"`        // [Required]
	TradeOrderLines []TradeOrderLines `json:"tradeOrderLines"` // [Required]
}
