package golazada

type AddressBilling struct {
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
	TaxAmount                   string           `json:"tax_amount"`                     // [Required]
	Reason                      string           `json:"reason"`                         // [Required]
	SlaTimeStamp                string           `json:"sla_time_stamp"`                 // [Required]
	VoucherSeller               string           `json:"voucher_seller"`                 // [Required]
	PurchaseOrderId             string           `json:"purchase_order_id"`              // [Required]
	VoucherCodeSeller           string           `json:"voucher_code_seller"`            // [Required]
	VoucherCode                 string           `json:"voucher_code"`                   // [Required]
	PackageId                   string           `json:"package_id"`                     // [Required]
	BuyerId                     string           `json:"buyer_id"`                       // [Required]
	Variation                   string           `json:"variation"`                      // [Required]
	ProductId                   int64            `json:"product_id"`                     // [Required]
	VoucherCodePlatform         string           `json:"voucher_code_platform"`          // [Required]
	PurchaseOrderNumber         string           `json:"purchase_order_number"`          // [Required]
	Sku                         string           `json:"sku"`                            // [Required]
	OrderType                   string           `json:"order_type"`                     // [Required]
	InvoiceNumber               string           `json:"invoice_number"`                 // [Required]
	SellerId                    int64            `json:"seller_id"`                      // [Required]
	CancelReturnInitiator       string           `json:"cancel_return_initiator"`        // [Required]
	ShopSku                     string           `json:"shop_sku"`                       // [Required]
	IsReroute                   string           `json:"is_reroute"`                     // [Required]
	StagePayStatus              string           `json:"stage_pay_status"`               // [Required]
	SkuId                       int64            `json:"sku_id"`                         // [Required]
	TrackingCodePre             string           `json:"tracking_code_pre"`              // [Required]
	OrderItemId                 int64            `json:"order_item_id"`                  // [Required]
	ShopId                      int64            `json:"shop_id"`                        // [Required]
	OrderFlag                   string           `json:"order_flag"`                     // [Required]
	IsFbl                       string           `json:"is_fbl"`                         // [Required]
	Name                        string           `json:"name"`                           // [Required]
	OrderId                     int64            `json:"order_id"`                       // [Required]
	Status                      string           `json:"status"`                         // [Required]
	ProductMainImage            string           `json:"product_main_image"`             // [Required]
	VoucherPlatform             FlexString       `json:"voucher_platform"`               // [Required]
	PaidPrice                   string           `json:"paid_price"`                     // [Required]
	ProductDetailUrl            string           `json:"product_detail_url"`             // [Required]
	WarehouseCode               string           `json:"warehouse_code"`                 // [Required]
	PromisedShippingTime        string           `json:"promised_shipping_time"`         // [Required]
	ShippingType                string           `json:"shipping_type"`                  // [Required]
	CreatedAt                   string           `json:"created_at"`                     // [Required]
	VoucherSellerLpi            string           `json:"voucher_seller_lpi"`             // [Required]
	ShippingFeeDiscountPlatform string           `json:"shipping_fee_discount_platform"` // [Required]
	WalletCredits               string           `json:"wallet_credits"`                 // [Required]
	UpdatedAt                   string           `json:"updated_at"`                     // [Required]
	Currency                    string           `json:"currency"`                       // [Required]
	ShippingProviderType        string           `json:"shipping_provider_type"`         // [Required]
	VoucherPlatformLpi          FlexString       `json:"voucher_platform_lpi"`           // [Required]
	ShippingFeeOriginal         string           `json:"shipping_fee_original"`          // [Required]
	ItemPrice                   string           `json:"item_price"`                     // [Required]
	IsDigital                   string           `json:"is_digital"`                     // [Required]
	ShippingServiceCost         string           `json:"shipping_service_cost"`          // [Required]
	TrackingCode                string           `json:"tracking_code"`                  // [Required]
	ShippingFeeDiscountSeller   string           `json:"shipping_fee_discount_seller"`   // [Required]
	ShippingAmount              string           `json:"shipping_amount"`                // [Required]
	ReasonDetail                string           `json:"reason_detail"`                  // [Required]
	ReturnStatus                string           `json:"return_status"`                  // [Required]
	PartnerUserId               string           `json:"partner_user_id"`                // [Required]
	ShipmentProvider            string           `json:"shipment_provider"`              // [Required]
	VoucherAmount               string           `json:"voucher_amount"`                 // [Required]
	DigitalDeliveryInfo         string           `json:"digital_delivery_info"`          // [Required]
	ExtraAttributes             string           `json:"extra_attributes"`               // [Required]
}
type Model struct {
	VoucherPlatform             FlexString       `json:"voucher_platform"`               // [Required]
	Voucher                     string           `json:"voucher"`                        // [Required]
	WarehouseCode               string           `json:"warehouse_code"`                 // [Required]
	OrderNumber                 string           `json:"order_number"`                   // [Required]
	VoucherSeller               string           `json:"voucher_seller"`                 // [Required]
	CreatedAt                   string           `json:"created_at"`                     // [Required]
	VoucherCode                 string           `json:"voucher_code"`                   // [Required]
	GiftOption                  string           `json:"gift_option"`                    // [Required]
	ShippingFeeDiscountPlatform string           `json:"shipping_fee_discount_platform"` // [Required]
	CustomerLastName            string           `json:"customer_last_name"`             // [Required]
	UpdatedAt                   string           `json:"updated_at"`                     // [Required]
	PromisedShippingTimes       string           `json:"promised_shipping_times"`        // [Required]
	Price                       float64          `json:"price"`                          // [Required]
	NationalRegistrationNumber  string           `json:"national_registration_number"`   // [Required]
	ShippingFeeOriginal         string           `json:"shipping_fee_original"`          // [Required]
	PaymentMethod               string           `json:"payment_method"`                 // [Required]
	AddressUpdatedAt            string           `json:"address_updated_at"`             // [Required]
	CustomerFirstName           string           `json:"customer_first_name"`            // [Required]
	MemberSubOrderList          []MemberSubOrder `json:"member_sub_order_list"`          // [Required]
	ShippingFeeDiscountSeller   string           `json:"shipping_fee_discount_seller"`   // [Required]
	ShippingFee                 string           `json:"shipping_fee"`                   // [Required]
	BranchNumber                string           `json:"branch_number"`                  // [Required]
	TaxCode                     string           `json:"tax_code"`                       // [Required]
	ItemsCount                  string           `json:"items_count"`                    // [Required]
	DeliveryInfo                string           `json:"delivery_info"`                  // [Required]
	Statuses                    []string         `json:"statuses"`                       // [Required]
	AddressBilling              *AddressBilling  `json:"address_billing"`                // [Required]
	ExtraAttributes             string           `json:"extra_attributes"`               // [Required]
	OrderId                     int64            `json:"order_id"`                       // [Required]
	GiftMessage                 string           `json:"gift_message"`                   // [Required]
	Remarks                     string           `json:"remarks"`                        // [Required]
	AddressShipping             *AddressBilling  `json:"address_shipping"`               // [Required]
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
	PartnerUid string `json:"partnerUid"` // [Required]
	Status     string `json:"status"`     // [Required]
}
type PartnerTransactionResponse struct {
	BaseResponse                                       // Common response fields
	Result       *PartnerTransactionResponseDataResult `json:"result,omitempty"` //
}
type PartnerTransactionResponseDataResult struct {
	ModelList  []Model `json:"model_list"`  // [Required]
	TotalCount int64   `json:"total_count"` // [Required]
	PageNo     string  `json:"page_no"`     // [Required]
	PageSize   int64   `json:"page_size"`   // [Required]
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
	DisplayMessage string `json:"displayMessage"` // [Required]
	Key            string `json:"key"`            // [Required]
}
type ResponseDataResultModule struct {
	BuyerId       string `json:"buyer_id"`       // [Required]
	SellerId      int64  `json:"seller_id"`      // [Required]
	PartneruserId string `json:"partneruser_id"` // [Required]
}
type ResultErrorCode struct {
	DisplayMessage string `json:"display_message"` // [Required]
	Key            string `json:"key"`             // [Required]
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
