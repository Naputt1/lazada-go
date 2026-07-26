package golazada

type Certificate struct {
	CertificateCode string      `json:"certificate_code"` // [Required]
	InitialNum      string      `json:"initial_num"`      // [Required]
	BizType         string      `json:"biz_type"`         // [Required]
	EndTime         string      `json:"end_time"`         // [Required]
	OuterId         string      `json:"outer_id"`         // [Required]
	QrCodeUrl       string      `json:"qr_code_url"`      // [Required]
	LockedNum       string      `json:"locked_num"`       // [Required]
	StartTime       string      `json:"start_time"`       // [Required]
	AvailableNum    string      `json:"available_num"`    // [Required]
	UsedNum         string      `json:"used_num"`         // [Required]
	Attributes      interface{} `json:"attributes"`       // [Required]
	ConsumeStatus   string      `json:"consume_status"`   // [Required]
	CodeStatus      string      `json:"code_status"`      // [Required]
}
type GetOrderItemsFromBarCodeResponse struct {
	BaseResponse                                      // Common response fields
	Response     GetOrderItemsFromBarCodeResponseData `json:"data"` // Response data
}
type GetOrderItemsFromBarCodeResponseData struct {
	StrartTime      string `json:"strart_time"`      // [Required]
	CertificateCode string `json:"certificate_code"` // [Required]
	ItemList        []Item `json:"item_list"`        // [Required]
	BizType         string `json:"biz_type"`         // [Required]
	EndTime         string `json:"end_time"`         // [Required]
	TradeOrderId    string `json:"trade_order_id"`   // [Required]
	CodeStatus      string `json:"code_status"`      // [Required]
	OuterId         string `json:"outer_id"`         // [Required]
	SerialNum       string `json:"serial_num"`       // [Required]
}
type GlobalEticketMerchantMaAvailableResponse struct {
	BaseResponse           // Common response fields
	RespBody     *RespBody `json:"resp_body,omitempty"` //
	RetCode      string    `json:"ret_code,omitempty"`  //
	RetMsg       string    `json:"ret_msg,omitempty"`   //
}
type GlobalEticketMerchantMaConsumeResponse struct {
	BaseResponse           // Common response fields
	RespBody     *RespBody `json:"resp_body,omitempty"` //
	RetCode      string    `json:"ret_code,omitempty"`  //
	RetMsg       string    `json:"ret_msg,omitempty"`   //
}
type GlobalEticketMerchantMaFailsendResponse struct {
	BaseResponse             // Common response fields
	RespBody     interface{} `json:"resp_body,omitempty"` //
	RetCode      string      `json:"ret_code,omitempty"`  //
	RetMsg       string      `json:"ret_msg,omitempty"`   //
}
type GlobalEticketMerchantMaQueryResponse struct {
	BaseResponse                       // Common response fields
	RespBody     *ResponseDataRespBody `json:"resp_body,omitempty"` //
	RetCode      string                `json:"ret_code,omitempty"`  //
	RetMsg       string                `json:"ret_msg,omitempty"`   //
}
type GlobalEticketMerchantMaQueryTbMaResponse struct {
	BaseResponse             // Common response fields
	RespBody     interface{} `json:"resp_body,omitempty"` //
	RetCode      string      `json:"ret_code,omitempty"`  //
	RetMsg       string      `json:"ret_msg,omitempty"`   //
}
type GlobalEticketMerchantMaSendResponse struct {
	BaseResponse             // Common response fields
	RespBody     interface{} `json:"resp_body,omitempty"` //
	RetCode      string      `json:"ret_code,omitempty"`  //
	RetMsg       string      `json:"ret_msg,omitempty"`   //
}
type Item struct {
	ItemImg           string `json:"item_img"`            // [Required]
	ItemId            int64  `json:"item_id"`             // [Required]
	ActualFee         string `json:"actual_fee"`          // [Required]
	ActualFeeCurrency string `json:"actual_fee_currency"` // [Required]
	UnitFee           string `json:"unit_fee"`            // [Required]
	ItemName          string `json:"item_name"`           // [Required]
	UnitFeeCurrency   string `json:"unit_fee_currency"`   // [Required]
}
type RedeemOrderItemsResponse struct {
	BaseResponse                              // Common response fields
	Response     RedeemOrderItemsResponseData `json:"data"` // Response data
}
type RedeemOrderItemsResponseData struct {
	LeftNum   string `json:"left_num"`   // [Required]
	OuterId   string `json:"outer_id"`   // [Required]
	SerialNum string `json:"serial_num"` // [Required]
}
type RespBody struct {
	AttributeMap interface{} `json:"attribute_map"` // [Required]
}
type ResponseDataRespBody struct {
	Certificate *Certificate `json:"certificate"` // [Required]
}
