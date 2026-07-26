package golazada

type Buyer struct {
	UserId string `json:"user_id"` // [Required]
}
type GetReverseOrderDetailResponse struct {
	BaseResponse                                   // Common response fields
	Response     GetReverseOrderDetailResponseData `json:"data"` // Response data
}
type GetReverseOrderDetailResponseData struct {
	ReverseOrderId          string                `json:"reverse_order_id"`        // [Required]
	RequestType             string                `json:"request_type"`            // [Required]
	ReverseOrderLineDTOList []ReverseOrderLineDTO `json:"reverseOrderLineDTOList"` // [Required]
	ShippingType            string                `json:"shipping_type"`           // [Required]
	IsRtm                   string                `json:"is_rtm"`                  // [Required]
	TradeOrderId            string                `json:"trade_order_id"`          // [Required]
}
type GetReverseOrderHistoryListResponse struct {
	BaseResponse                                        // Common response fields
	Response     GetReverseOrderHistoryListResponseData `json:"data"` // Response data
}
type GetReverseOrderHistoryListResponseData struct {
	PageInfo *ResponseDataPageInfo `json:"page_info"` // [Required]
	List     []List                `json:"list"`      // [Required]
}
type GetReverseOrderReasonListResponse struct {
	BaseResponse                                       // Common response fields
	Response     GetReverseOrderReasonListResponseData `json:"data"` // Response data
}
type GetReverseOrderReasonListResponseData struct {
	MutiLanguageText string `json:"muti_language_text"` // [Required]
	Text             string `json:"text"`               // [Required]
	ReasonId         string `json:"reason_id"`          // [Required]
}
type GetReverseOrdersForSellerResponse struct {
	BaseResponse                                              // Common response fields
	Result       *GetReverseOrdersForSellerResponseDataResult `json:"result,omitempty"` //
}
type GetReverseOrdersForSellerResponseDataResult struct {
	Total    int64         `json:"total"`     // [Required]
	Success  bool          `json:"success"`   // [Required]
	PageNo   string        `json:"page_no"`   // [Required]
	Items    []ResultItems `json:"items"`     // [Required]
	PageSize int64         `json:"page_size"` // [Required]
}
type InitReverseOrderCancelDecideResponse struct {
	BaseResponse // Common response fields
}
type InitReverseOrderCancelResponse struct {
	BaseResponse                                    // Common response fields
	Response     InitReverseOrderCancelResponseData `json:"data"` // Response data
}
type InitReverseOrderCancelResponseData struct {
	TipContent string `json:"tip_content"` // [Required]
	TipType    string `json:"tip_type"`    // [Required]
}
type List struct {
	Time     string        `json:"time"`     // [Required]
	Operator string        `json:"operator"` // [Required]
	Picture  []interface{} `json:"picture"`  // [Required]
}
type ProductDTO struct {
	ProductId int64  `json:"product_id"` // [Required]
	Sku       string `json:"sku"`        // [Required]
}
type ReasonOptions struct {
	ReasonName string `json:"reason_name"` // [Required]
	ReasonId   string `json:"reason_id"`   // [Required]
}
type ResponseDataPageInfo struct {
	Total             int64  `json:"total"`               // [Required]
	PageSize          int64  `json:"page_size"`           // [Required]
	CurrentPageNumber string `json:"current_page_number"` // [Required]
}
type ResultItems struct {
	ReverseOrderLines []ReverseOrderLines `json:"reverse_order_lines"` // [Required]
	ReverseOrderId    string              `json:"reverse_order_id"`    // [Required]
	RequestType       string              `json:"request_type"`        // [Required]
	IsRtm             string              `json:"is_rtm"`              // [Required]
	ShippingType      string              `json:"shipping_type"`       // [Required]
	TradeOrderId      string              `json:"trade_order_id"`      // [Required]
}
type ReverseOrderLine struct {
	PaidPrice          string `json:"paid_price"`            // [Required]
	IsCancel           string `json:"is_cancel"`             // [Required]
	ReasonId           string `json:"reason_id"`             // [Required]
	ReasonSource       string `json:"reason_source"`         // [Required]
	ReasonDesc         string `json:"reason_desc"`           // [Required]
	ApplyReason        string `json:"apply_reason"`          // [Required]
	ReasonType         string `json:"reason_type"`           // [Required]
	SellerSku          string `json:"seller_sku"`            // [Required]
	RefundAmount       string `json:"refund_amount"`         // [Required]
	OrderLineId        string `json:"order_line_id"`         // [Required]
	ReasonName         string `json:"reason_name"`           // [Required]
	OrderId            int64  `json:"order_id"`              // [Required]
	ReverseOrderLineId string `json:"reverse_order_line_id"` // [Required]
}
type ReverseOrderLineDTO struct {
	ReturnOrderLineGmtCreate   string      `json:"return_order_line_gmt_create"`   // [Required]
	PlatformSkuId              string      `json:"platform_sku_id"`                // [Required]
	IsNeedRefund               string      `json:"is_need_refund"`                 // [Required]
	TradeOrderGmtCreate        string      `json:"trade_order_gmt_create"`         // [Required]
	ReasonText                 string      `json:"reason_text"`                    // [Required]
	ItemUnitPrice              string      `json:"item_unit_price"`                // [Required]
	Sla                        string      `json:"sla"`                            // [Required]
	TradeOrderLineId           string      `json:"trade_order_line_id"`            // [Required]
	ReturnOrderLineGmtModified string      `json:"return_order_line_gmt_modified"` // [Required]
	OfcStatus                  string      `json:"ofc_status"`                     // [Required]
	SellerSkuId                string      `json:"seller_sku_id"`                  // [Required]
	ProductDTO                 *ProductDTO `json:"productDTO"`                     // [Required]
	RefundPaymentMethod        string      `json:"refund_payment_method"`          // [Required]
	Buyer                      *Buyer      `json:"buyer"`                          // [Required]
	ReasonCode                 string      `json:"reason_code"`                    // [Required]
	WhqcDecision               string      `json:"whqc_decision"`                  // [Required]
	ReverseStatus              string      `json:"reverse_status"`                 // [Required]
	RefundAmount               string      `json:"refund_amount"`                  // [Required]
	TrackingNumber             string      `json:"tracking_number"`                // [Required]
	IsDispute                  string      `json:"is_dispute"`                     // [Required]
	ReverseOrderLineId         string      `json:"reverse_order_line_id"`          // [Required]
}
type ReverseOrderLines struct {
	Product                    *ReverseOrderLinesProduct `json:"product"`                        // [Required]
	ReturnOrderLineGmtCreate   string                    `json:"return_order_line_gmt_create"`   // [Required]
	PlatformSkuId              string                    `json:"platform_sku_id"`                // [Required]
	TradeOrderGmtCreate        string                    `json:"trade_order_gmt_create"`         // [Required]
	IsNeedRefund               string                    `json:"is_need_refund"`                 // [Required]
	ReasonText                 string                    `json:"reason_text"`                    // [Required]
	ItemUnitPrice              string                    `json:"item_unit_price"`                // [Required]
	Sla                        string                    `json:"sla"`                            // [Required]
	ReturnOrderLineGmtModified string                    `json:"return_order_line_gmt_modified"` // [Required]
	TradeOrderLineId           string                    `json:"trade_order_line_id"`            // [Required]
	OfcStatus                  string                    `json:"ofc_status"`                     // [Required]
	SellerSkuId                string                    `json:"seller_sku_id"`                  // [Required]
	RefundPaymentMethod        string                    `json:"refund_payment_method"`          // [Required]
	Buyer                      *ReverseOrderLinesBuyer   `json:"buyer"`                          // [Required]
	ReasonCode                 string                    `json:"reason_code"`                    // [Required]
	WhqcDecision               string                    `json:"whqc_decision"`                  // [Required]
	ReverseStatus              string                    `json:"reverse_status"`                 // [Required]
	RefundAmount               string                    `json:"refund_amount"`                  // [Required]
	TrackingNumber             string                    `json:"tracking_number"`                // [Required]
	ReceiverAddress            string                    `json:"receiver_address"`               // [Required]
	IsDispute                  string                    `json:"is_dispute"`                     // [Required]
	ReverseOrderLineId         string                    `json:"reverse_order_line_id"`          // [Required]
}
type ReverseOrderLinesBuyer struct {
	BuyerId string `json:"buyer_id"` // [Required]
}
type ReverseOrderLinesProduct struct {
	ProductSku string `json:"product_sku"` // [Required]
	ProductId  int64  `json:"product_id"`  // [Required]
}
type ReverseOrderOnlyRefundDecideResponse struct {
	BaseResponse // Common response fields
}
type ReverseOrderReturnUpdateResponse struct {
	BaseResponse                                      // Common response fields
	Response     ReverseOrderReturnUpdateResponseData `json:"data"` // Response data
}
type ReverseOrderReturnUpdateResponseData struct {
	ReasonInfo       []ReasonOptions    `json:"reason_info"`        // [Required]
	ReverseOrderId   string             `json:"reverse_order_id"`   // [Required]
	TotalRefund      string             `json:"total_refund"`       // [Required]
	ReverseOrderLine []ReverseOrderLine `json:"reverse_order_line"` // [Required]
}
