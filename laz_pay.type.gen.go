package golazada

type CollectBenefitResponse struct {
	BaseResponse         // Common response fields
	Response      string `json:"data"`                    // Response data
	ResultCode    string `json:"resultCode,omitempty"`    //
	ResultMessage string `json:"resultMessage,omitempty"` //
	TraceId       string `json:"trace_id,omitempty"`      //
}
type ConsultPaymentResponse struct {
	BaseResponse                 // Common response fields
	AdditionalInfo  string       `json:"additionalInfo,omitempty"`  //
	ErrorCode       string       `json:"errorCode,omitempty"`       //
	PayOptions      []PayOptions `json:"payOptions,omitempty"`      //
	ResponseCode    string       `json:"responseCode,omitempty"`    //
	ResponseMessage string       `json:"responseMessage,omitempty"` //
}
type CreateSubscriptionToFusionResponse struct {
	BaseResponse              // Common response fields
	SubscribeTime      string `json:"subscribeTime,omitempty"`      //
	SubscriptionStatus string `json:"subscriptionStatus,omitempty"` //
	UnsubscribeTime    string `json:"unsubscribeTime,omitempty"`    //
}
type DGUtiityPreCreateOrderResponse struct {
	BaseResponse        // Common response fields
	ResultCode   string `json:"resultCode,omitempty"` //
	ResultMsg    string `json:"resultMsg,omitempty"`  //
	TradeNo      string `json:"tradeNo,omitempty"`    //
}
type DGUtilityPreGetPaymentStatusResponse struct {
	BaseResponse        // Common response fields
	ResultCode   string `json:"resultCode,omitempty"` //
	ResultMsg    string `json:"resultMsg,omitempty"`  //
}
type DGUtilityPreUpdateFulfillemtStatusResponse struct {
	BaseResponse        // Common response fields
	ResultCode   string `json:"resultCode,omitempty"` //
	ResultMsg    string `json:"resultMsg,omitempty"`  //
}
type DigitalAlterOrderStatusResponse struct {
	BaseResponse         // Common response fields
	OrderStatus   string `json:"orderStatus,omitempty"`   //
	PaymentStatus string `json:"paymentStatus,omitempty"` //
	ResultCode    string `json:"resultCode,omitempty"`    //
	TraceId       string `json:"traceId,omitempty"`       //
	TransactionId string `json:"transactionId,omitempty"` //
}
type DigitalCreateOrderResponse struct {
	BaseResponse            // Common response fields
	PaymentLink      string `json:"paymentLink,omitempty"`      //
	ResultCode       string `json:"resultCode,omitempty"`       //
	TraceId          string `json:"traceId,omitempty"`          //
	TradeOrderLineId string `json:"tradeOrderLineId,omitempty"` //
	TransactionId    string `json:"transactionId,omitempty"`    //
}
type DigitalQueryOrderResponse struct {
	BaseResponse         // Common response fields
	OrderStatus   string `json:"orderStatus,omitempty"`   //
	PaymentStatus string `json:"paymentStatus,omitempty"` //
	ResultCode    string `json:"resultCode,omitempty"`    //
	TraceId       string `json:"traceId,omitempty"`       //
	TransactionId string `json:"transactionId,omitempty"` //
}
type GetSubscriptionToFusionResponse struct {
	BaseResponse              // Common response fields
	SubscribeTime      string `json:"subscribeTime,omitempty"`      //
	SubscriptionStatus string `json:"subscriptionStatus,omitempty"` //
	UnsubscribeTime    string `json:"unsubscribeTime,omitempty"`    //
}
type InsuranceAlterOrderStatusResponse struct {
	BaseResponse         // Common response fields
	OrderStatus   string `json:"orderStatus,omitempty"`   //
	PaymentStatus string `json:"paymentStatus,omitempty"` //
	ResultCode    string `json:"resultCode,omitempty"`    //
	TraceId       string `json:"traceId,omitempty"`       //
	TransactionId string `json:"transactionId,omitempty"` //
}
type InsuranceCreateOrderResponse struct {
	BaseResponse               // Common response fields
	ItemPrice           string `json:"itemPrice,omitempty"`           //
	PaymentLink         string `json:"paymentLink,omitempty"`         //
	ResultCode          string `json:"resultCode,omitempty"`          //
	SubItemPrice        string `json:"subItemPrice,omitempty"`        //
	SubTradeOrderLineId string `json:"subTradeOrderLineId,omitempty"` //
	TraceId             string `json:"traceId,omitempty"`             //
	TradeOrderLineId    string `json:"tradeOrderLineId,omitempty"`    //
	TransactionId       string `json:"transactionId,omitempty"`       //
}
type InsuranceGetPromotionsResponse struct {
	BaseResponse         // Common response fields
	Response      string `json:"data"`                    // Response data
	ResultCode    string `json:"resultCode,omitempty"`    //
	ResultMessage string `json:"resultMessage,omitempty"` //
	TraceId       string `json:"traceId,omitempty"`       //
}
type InsuranceQueryOrderResponse struct {
	BaseResponse         // Common response fields
	OrderStatus   string `json:"orderStatus,omitempty"`   //
	PaymentStatus string `json:"paymentStatus,omitempty"` //
	ResultCode    string `json:"resultCode,omitempty"`    //
	TraceId       string `json:"traceId,omitempty"`       //
	TransactionId string `json:"transactionId,omitempty"` //
}
type InsuranceRealTimeCDPResponse struct {
	BaseResponse         // Common response fields
	Response      string `json:"data"`                    // Response data
	RedirectUrl   string `json:"redirectUrl,omitempty"`   //
	ResultCode    string `json:"resultCode,omitempty"`    //
	ResultMessage string `json:"resultMessage,omitempty"` //
}
type LazadaCFOInvoiceRpaCallbackResponse struct {
	BaseResponse        // Common response fields
	Content      string `json:"content,omitempty"`    //
	IsSuccess    string `json:"is_success,omitempty"` //
	ResCode      string `json:"res_code,omitempty"`   //
	ResMsg       string `json:"res_msg,omitempty"`    //
}
type OpenServiceBalanceQueryResponse struct {
	BaseResponse               // Common response fields
	AvailableAmount     string `json:"available_amount,omitempty"`      //
	AvailableAmountCent string `json:"available_amount_cent,omitempty"` //
	Currency            string `json:"currency,omitempty"`              //
	DateTime            string `json:"date_time,omitempty"`             //
}
type OpenServiceKycQueryResponse struct {
	BaseResponse          // Common response fields
	Birthday       string `json:"birthday,omitempty"`         //
	CertFrontImage string `json:"cert_front_image,omitempty"` //
	CertType       string `json:"cert_type,omitempty"`        //
	ExtendInfo     string `json:"extend_info,omitempty"`      //
	FullKycStatus  string `json:"full_kyc_status,omitempty"`  //
	FullName       string `json:"full_name,omitempty"`        //
	KycJumpUrl     string `json:"kyc_jump_url,omitempty"`     //
	Phone          string `json:"phone,omitempty"`            //
	Prefix         string `json:"prefix,omitempty"`           //
	UserId         string `json:"userId,omitempty"`           //
}
type OpenServiceWithdrawApplyResponse struct {
	BaseResponse             // Common response fields
	Currency          string `json:"currency,omitempty"`            //
	PartnerDeposit    string `json:"partner_deposit,omitempty"`     //
	WithdrawAmount    string `json:"withdraw_amount,omitempty"`     //
	WithdrawId        string `json:"withdraw_id,omitempty"`         //
	WithdrawRequestId string `json:"withdraw_request_id,omitempty"` //
	Withdrawable      string `json:"withdrawable,omitempty"`        //
}
type OpenServiceWithdrawQueryResponse struct {
	BaseResponse             // Common response fields
	Currency          string `json:"currency,omitempty"`            //
	PartnerDeposit    string `json:"partner_deposit,omitempty"`     //
	WithdrawAmount    string `json:"withdraw_amount,omitempty"`     //
	WithdrawId        string `json:"withdraw_id,omitempty"`         //
	WithdrawRequestId string `json:"withdraw_request_id,omitempty"` //
	Withdrawable      string `json:"withdrawable,omitempty"`        //
}
type Order struct {
	Premium         string `json:"premium"`         // [Required]
	ExpireTime      string `json:"expireTime"`      // [Required]
	OrderDetailLink string `json:"orderDetailLink"` // [Required]
	EffectiveTime   string `json:"effectiveTime"`   // [Required]
	InsuranceName   string `json:"insuranceName"`   // [Required]
	OrderStatus     string `json:"orderStatus"`     // [Required]
	ZoneId          string `json:"zoneId"`          // [Required]
	PolicyLink      string `json:"policyLink"`      // [Required]
	InsuredName     string `json:"insuredName"`     // [Required]
	PaidPremium     string `json:"paidPremium"`     // [Required]
	TransactionId   string `json:"transactionId"`   // [Required]
	ProductName     string `json:"productName"`     // [Required]
}
type PayAssetDetails struct {
	BankAccount     interface{} `json:"bankAccount"`     // [Required]
	Coupon          interface{} `json:"coupon"`          // [Required]
	Rebate          interface{} `json:"rebate"`          // [Required]
	AdditionalInfo  string      `json:"additionalInfo"`  // [Required]
	ExternalAccount interface{} `json:"externalAccount"` // [Required]
	Discount        interface{} `json:"discount"`        // [Required]
	PayAssetType    string      `json:"payAssetType"`    // [Required]
	StoreValue      interface{} `json:"storeValue"`      // [Required]
	Card            interface{} `json:"card"`            // [Required]
}
type PayOptions struct {
	DisableReasonCode   string            `json:"disableReasonCode"`   // [Required]
	DisableReasonDesc   string            `json:"disableReasonDesc"`   // [Required]
	AmountLimitMap      interface{}       `json:"amountLimitMap"`      // [Required]
	PayOptionInfo       interface{}       `json:"payOptionInfo"`       // [Required]
	Enabled             string            `json:"enabled"`             // [Required]
	SupportedCurrencies []interface{}     `json:"supportedCurrencies"` // [Required]
	PayCategory         string            `json:"payCategory"`         // [Required]
	PayMethod           string            `json:"payMethod"`           // [Required]
	AdditionalInfo      string            `json:"additionalInfo"`      // [Required]
	PayOption           string            `json:"payOption"`           // [Required]
	Rank                string            `json:"rank"`                // [Required]
	PayAssetDetails     []PayAssetDetails `json:"payAssetDetails"`     // [Required]
	Preferred           string            `json:"preferred"`           // [Required]
}
type QueryAddonOrderResponse struct {
	BaseResponse                              // Common response fields
	Response      QueryAddonOrderResponseData `json:"data"`                    // Response data
	RedirectUrl   string                      `json:"redirectUrl,omitempty"`   //
	ResultCode    string                      `json:"resultCode,omitempty"`    //
	ResultMessage string                      `json:"resultMessage,omitempty"` //
}
type QueryAddonOrderResponseData struct {
	TraceId    string  `json:"traceId"`    // [Required]
	Total      int64   `json:"total"`      // [Required]
	TotalPages string  `json:"totalPages"` // [Required]
	PageSize   string  `json:"pageSize"`   // [Required]
	OrderList  []Order `json:"orderList"`  // [Required]
	PageNum    string  `json:"pageNum"`    // [Required]
}
type QueryBenefitResponse struct {
	BaseResponse         // Common response fields
	Response      string `json:"data"`                    // Response data
	ResultCode    string `json:"resultCode,omitempty"`    //
	ResultMessage string `json:"resultMessage,omitempty"` //
	TraceId       string `json:"trace_id,omitempty"`      //
}
type ReconciliationResponse struct {
	BaseResponse        // Common response fields
	Res          string `json:"res,omitempty"` //
}
type RedeemMpVoucherResponse struct {
	BaseResponse             // Common response fields
	BrokerName        string `json:"brokerName,omitempty"`        //
	ResultCode        string `json:"resultCode,omitempty"`        //
	ResultMessage     string `json:"resultMessage,omitempty"`     //
	TraceId           string `json:"traceId,omitempty"`           //
	VoucherTemplateId string `json:"voucherTemplateId,omitempty"` //
}
