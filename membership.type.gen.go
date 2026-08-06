package golazada

type AddressBilling struct {
	Country   FlexString `json:"country"`    // [Required]
	Address3  FlexString `json:"address3"`   // [Required]
	Phone     FlexString `json:"phone"`      // [Required]
	Address2  FlexString `json:"address2"`   // [Required]
	City      FlexString `json:"city"`       // [Required]
	Address1  FlexString `json:"address1"`   // [Required]
	PostCode  FlexString `json:"post_code"`  // [Required]
	Phone2    FlexString `json:"phone2"`     // [Required]
	LastName  FlexString `json:"last_name"`  // [Required]
	Address5  FlexString `json:"address5"`   // [Required]
	Address4  FlexString `json:"address4"`   // [Required]
	FirstName FlexString `json:"first_name"` // [Required]
}
type AddressShipping struct {
	Country   string `json:"country"`    // [Required]
	Address3  string `json:"address3"`   // [Required]
	Phone     string `json:"phone"`      // [Required]
	Address2  string `json:"address2"`   // [Required]
	City      string `json:"city"`       // [Required]
	Address1  string `json:"address1"`   // [Required]
	PostCode  string `json:"post_code"`  // [Required]
	Phone2    string `json:"phone2"`     // [Required]
	LastName  string `json:"last_name"`  // [Required]
	Address5  string `json:"address5"`   // [Required]
	Address4  string `json:"address4"`   // [Required]
	FirstName string `json:"first_name"` // [Required]
}
type GetLinkMember1Response struct {
	BaseResponse                                  // Common response fields
	Result       *GetLinkMemberResponseDataResult `json:"result,omitempty"` //
}
type GetLinkMemberList1Response struct {
	BaseResponse                                      // Common response fields
	Result       *GetLinkMemberListResponseDataResult `json:"result,omitempty"` //
}
type GetLinkMemberListResponse struct {
	BaseResponse                                      // Common response fields
	Result       *GetLinkMemberListResponseDataResult `json:"result,omitempty"` //
}
type GetLinkMemberListResponseDataResult struct {
	ModelList  []ResponseDataResultModule `json:"model_list"`  // [Required]
	TotalCount int64                      `json:"total_count"` // [Required]
}
type GetLinkMemberResponse struct {
	BaseResponse                                  // Common response fields
	Result       *GetLinkMemberResponseDataResult `json:"result,omitempty"` //
}
type GetLinkMemberResponseDataResult struct {
	Module *ResponseDataResultModule `json:"module"` // [Required]
}
type LinkMembershipResponse struct {
	BaseResponse                                   // Common response fields
	Result       *LinkMembershipResponseDataResult `json:"result,omitempty"` //
}
type LinkMembershipResponseDataResult struct {
	Data      interface{}      `json:"data"`       // [Required]
	Success   bool             `json:"success"`    // [Required]
	ErrorCode *ResultErrorCode `json:"error_code"` // [Required]
}
type MemberSubOrder struct {
	PickUpStoreInfo             *PickUpStoreInfo `json:"pick_up_store_info"`             // [Required]
	TaxAmount                   FlexString       `json:"tax_amount"`                     // [Required]
	Reason                      FlexString       `json:"reason"`                         // [Required]
	SlaTimeStamp                FlexString       `json:"sla_time_stamp"`                 // [Required]
	VoucherSeller               FlexString       `json:"voucher_seller"`                 // [Required]
	PurchaseOrderId             FlexString       `json:"purchase_order_id"`              // [Required]
	VoucherCodeSeller           FlexString       `json:"voucher_code_seller"`            // [Required]
	VoucherCode                 FlexString       `json:"voucher_code"`                   // [Required]
	PackageId                   FlexString       `json:"package_id"`                     // [Required]
	BuyerId                     FlexString       `json:"buyer_id"`                       // [Required]
	Variation                   FlexString       `json:"variation"`                      // [Required]
	ProductId                   int64            `json:"product_id"`                     // [Required]
	VoucherCodePlatform         FlexString       `json:"voucher_code_platform"`          // [Required]
	PurchaseOrderNumber         FlexString       `json:"purchase_order_number"`          // [Required]
	Sku                         FlexString       `json:"sku"`                            // [Required]
	OrderType                   FlexString       `json:"order_type"`                     // [Required]
	InvoiceNumber               FlexString       `json:"invoice_number"`                 // [Required]
	SellerId                    int64            `json:"seller_id"`                      // [Required]
	CancelReturnInitiator       FlexString       `json:"cancel_return_initiator"`        // [Required]
	ShopSku                     FlexString       `json:"shop_sku"`                       // [Required]
	IsReroute                   FlexString       `json:"is_reroute"`                     // [Required]
	StagePayStatus              FlexString       `json:"stage_pay_status"`               // [Required]
	SkuId                       int64            `json:"sku_id"`                         // [Required]
	TrackingCodePre             FlexString       `json:"tracking_code_pre"`              // [Required]
	OrderItemId                 int64            `json:"order_item_id"`                  // [Required]
	ShopId                      int64            `json:"shop_id"`                        // [Required]
	OrderFlag                   FlexString       `json:"order_flag"`                     // [Required]
	IsFbl                       FlexString       `json:"is_fbl"`                         // [Required]
	Name                        FlexString       `json:"name"`                           // [Required]
	OrderId                     FlexInt          `json:"order_id"`                       // [Required]
	Status                      FlexString       `json:"status"`                         // [Required]
	ProductMainImage            FlexString       `json:"product_main_image"`             // [Required]
	VoucherPlatform             FlexString       `json:"voucher_platform"`               // [Required]
	PaidPrice                   FlexString       `json:"paid_price"`                     // [Required]
	ProductDetailUrl            FlexString       `json:"product_detail_url"`             // [Required]
	WarehouseCode               FlexString       `json:"warehouse_code"`                 // [Required]
	PromisedShippingTime        FlexString       `json:"promised_shipping_time"`         // [Required]
	ShippingType                FlexString       `json:"shipping_type"`                  // [Required]
	CreatedAt                   FlexString       `json:"created_at"`                     // [Required]
	VoucherSellerLpi            FlexString       `json:"voucher_seller_lpi"`             // [Required]
	ShippingFeeDiscountPlatform FlexString       `json:"shipping_fee_discount_platform"` // [Required]
	WalletCredits               FlexString       `json:"wallet_credits"`                 // [Required]
	UpdatedAt                   FlexString       `json:"updated_at"`                     // [Required]
	Currency                    FlexString       `json:"currency"`                       // [Required]
	ShippingProviderType        FlexString       `json:"shipping_provider_type"`         // [Required]
	VoucherPlatformLpi          FlexString       `json:"voucher_platform_lpi"`           // [Required]
	ShippingFeeOriginal         FlexString       `json:"shipping_fee_original"`          // [Required]
	ItemPrice                   FlexString       `json:"item_price"`                     // [Required]
	IsDigital                   FlexString       `json:"is_digital"`                     // [Required]
	ShippingServiceCost         FlexString       `json:"shipping_service_cost"`          // [Required]
	TrackingCode                FlexString       `json:"tracking_code"`                  // [Required]
	ShippingFeeDiscountSeller   FlexString       `json:"shipping_fee_discount_seller"`   // [Required]
	ShippingAmount              FlexString       `json:"shipping_amount"`                // [Required]
	ReasonDetail                FlexString       `json:"reason_detail"`                  // [Required]
	ReturnStatus                FlexString       `json:"return_status"`                  // [Required]
	PartnerUserId               FlexString       `json:"partner_user_id"`                // [Required]
	ShipmentProvider            FlexString       `json:"shipment_provider"`              // [Required]
	VoucherAmount               FlexString       `json:"voucher_amount"`                 // [Required]
	DigitalDeliveryInfo         FlexString       `json:"digital_delivery_info"`          // [Required]
	ExtraAttributes             FlexString       `json:"extra_attributes"`               // [Required]
}
type Model struct {
	VoucherPlatform             FlexString       `json:"voucher_platform"`               // [Required]
	Voucher                     FlexString       `json:"voucher"`                        // [Required]
	WarehouseCode               FlexString       `json:"warehouse_code"`                 // [Required]
	OrderNumber                 FlexString       `json:"order_number"`                   // [Required]
	VoucherSeller               FlexString       `json:"voucher_seller"`                 // [Required]
	CreatedAt                   FlexString       `json:"created_at"`                     // [Required]
	VoucherCode                 FlexString       `json:"voucher_code"`                   // [Required]
	GiftOption                  FlexString       `json:"gift_option"`                    // [Required]
	ShippingFeeDiscountPlatform FlexString       `json:"shipping_fee_discount_platform"` // [Required]
	CustomerLastName            FlexString       `json:"customer_last_name"`             // [Required]
	UpdatedAt                   FlexString       `json:"updated_at"`                     // [Required]
	PromisedShippingTimes       FlexString       `json:"promised_shipping_times"`        // [Required]
	Price                       FlexFloat        `json:"price"`                          // [Required]
	NationalRegistrationNumber  FlexString       `json:"national_registration_number"`   // [Required]
	ShippingFeeOriginal         FlexString       `json:"shipping_fee_original"`          // [Required]
	PaymentMethod               FlexString       `json:"payment_method"`                 // [Required]
	AddressUpdatedAt            FlexString       `json:"address_updated_at"`             // [Required]
	CustomerFirstName           FlexString       `json:"customer_first_name"`            // [Required]
	MemberSubOrderList          []MemberSubOrder `json:"member_sub_order_list"`          // [Required]
	ShippingFeeDiscountSeller   FlexString       `json:"shipping_fee_discount_seller"`   // [Required]
	ShippingFee                 FlexString       `json:"shipping_fee"`                   // [Required]
	BranchNumber                FlexString       `json:"branch_number"`                  // [Required]
	TaxCode                     FlexString       `json:"tax_code"`                       // [Required]
	ItemsCount                  FlexString       `json:"items_count"`                    // [Required]
	DeliveryInfo                FlexString       `json:"delivery_info"`                  // [Required]
	Statuses                    []string         `json:"statuses"`                       // [Required]
	AddressBilling              *AddressBilling  `json:"address_billing"`                // [Required]
	ExtraAttributes             FlexString       `json:"extra_attributes"`               // [Required]
	OrderId                     FlexInt          `json:"order_id"`                       // [Required]
	GiftMessage                 FlexString       `json:"gift_message"`                   // [Required]
	Remarks                     FlexString       `json:"remarks"`                        // [Required]
	AddressShipping             *AddressShipping `json:"address_shipping"`               // [Required]
}
type PartnerLinkResponse struct {
	BaseResponse                                // Common response fields
	Result       *PartnerLinkResponseDataResult `json:"result,omitempty"` //
}
type PartnerLinkResponseDataResult struct {
	Success   bool                                 `json:"success"`   // [Required]
	Module    *PartnerLinkResponseDataResultModule `json:"module"`    // [Required]
	ErrorCode *ResponseDataResultErrorCode         `json:"errorCode"` // [Required]
}
type PartnerLinkResponseDataResultModule struct {
	PartnerUid string     `json:"partnerUid"` // [Required]
	Status     FlexString `json:"status"`     // [Required]
}
type PartnerTransactionResponse struct {
	BaseResponse                                       // Common response fields
	Result       *PartnerTransactionResponseDataResult `json:"result,omitempty"` //
}
type PartnerTransactionResponseDataResult struct {
	ModelList  []Model    `json:"model_list"`  // [Required]
	TotalCount int64      `json:"total_count"` // [Required]
	PageNo     FlexString `json:"page_no"`     // [Required]
	PageSize   int64      `json:"page_size"`   // [Required]
}
type PartnerUnlinkResponse struct {
	BaseResponse                                   // Common response fields
	Result       *LinkMembershipResponseDataResult `json:"result,omitempty"` //
}
type PartnerUpdateResponse struct {
	BaseResponse                                   // Common response fields
	Result       *LinkMembershipResponseDataResult `json:"result,omitempty"` //
}
type ResponseDataResultErrorCode struct {
	DisplayMessage string     `json:"displayMessage"` // [Required]
	Key            FlexString `json:"key"`            // [Required]
}
type ResponseDataResultModule struct {
	BuyerId       FlexString `json:"buyer_id"`       // [Required]
	SellerId      int64      `json:"seller_id"`      // [Required]
	PartneruserId FlexString `json:"partneruser_id"` // [Required]
}
type ResultErrorCode struct {
	DisplayMessage FlexString `json:"display_message"` // [Required]
	Key            FlexString `json:"key"`             // [Required]
}
type UpdatePartnerUserIdResponse struct {
	BaseResponse                                        // Common response fields
	Result       *UpdatePartnerUserIdResponseDataResult `json:"result,omitempty"` //
}
type UpdatePartnerUserIdResponseDataResult struct {
	Success   bool                         `json:"success"`   // [Required]
	Module    interface{}                  `json:"module"`    // [Required]
	ErrorCode *ResponseDataResultErrorCode `json:"errorCode"` // [Required]
}
