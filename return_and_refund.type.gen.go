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
type GetReverseOrdersForSellerRequest struct {
	PageNo                                 int64    `json:"page_no" url:"page_no"`                                                                                   // [Required]
	PageSize                               int64    `json:"page_size" url:"page_size"`                                                                               // [Required]
	RequestTypeList                        []string `json:"request_type_list,omitempty" url:"request_type_list,omitempty"`                                           // [Optional]
	OfcStatusList                          []string `json:"ofc_status_list,omitempty" url:"ofc_status_list,omitempty"`                                               // [Optional]
	ReverseStatusList                      []string `json:"reverse_status_list,omitempty" url:"reverse_status_list,omitempty"`                                       // [Optional]
	ReverseOrderId                         *int64   `json:"reverse_order_id,omitempty" url:"reverse_order_id,omitempty"`                                             // [Optional]
	TradeOrderId                           *int64   `json:"trade_order_id,omitempty" url:"trade_order_id,omitempty"`                                                 // [Optional]
	ReturnToType                           *string  `json:"return_to_type,omitempty" url:"return_to_type,omitempty"`                                                 // [Optional]
	DisputeInProgress                      *bool    `json:"dispute_in_progress,omitempty" url:"dispute_in_progress,omitempty"`                                       // [Optional]
	TradeOrderLineCreatedTimeRangeStart    *int64   `json:"TradeOrderLineCreatedTimeRangeStart,omitempty" url:"TradeOrderLineCreatedTimeRangeStart,omitempty"`       // [Optional]
	TradeOrderLineCreatedTimeRangeEnd      *int64   `json:"TradeOrderLineCreatedTimeRangeEnd,omitempty" url:"TradeOrderLineCreatedTimeRangeEnd,omitempty"`           // [Optional]
	ReverseOrderLineTimeRangeStart         *int64   `json:"ReverseOrderLineTimeRangeStart,omitempty" url:"ReverseOrderLineTimeRangeStart,omitempty"`                 // [Optional]
	ReverseOrderLineTimeRangeEnd           *int64   `json:"ReverseOrderLineTimeRangeEnd,omitempty" url:"ReverseOrderLineTimeRangeEnd,omitempty"`                     // [Optional]
	ReverseOrderLineModifiedTimeRangeStart *int64   `json:"ReverseOrderLineModifiedTimeRangeStart,omitempty" url:"ReverseOrderLineModifiedTimeRangeStart,omitempty"` // [Optional]
	ReverseOrderLineModifiedTimeRangeEnd   *int64   `json:"ReverseOrderLineModifiedTimeRangeEnd,omitempty" url:"ReverseOrderLineModifiedTimeRangeEnd,omitempty"`     // [Optional]
	QCDecision                             *string  `json:"QC_Decision,omitempty" url:"QC_Decision,omitempty"`                                                       // [Optional]
}
type GetReverseOrdersForSellerResponse struct {
	BaseResponse                                       // Common response fields
	Response     GetReverseOrdersForSellerResponseData `json:"result"` // Response data
}
type GetReverseOrdersForSellerResponseData struct {
	Result *GetReverseOrdersForSellerResponseDataResult `json:"result"` // Response data
}
type GetReverseOrdersForSellerResponseDataItems struct {
	ReverseOrderLines []ReverseOrderLines `json:"reverse_order_lines"` // [Required]
	ReverseOrderId    FlexString          `json:"reverse_order_id"`    // [Required]
	RequestType       string              `json:"request_type"`        // [Required]
	IsRtm             FlexString          `json:"is_rtm"`              // [Required]
	ShippingType      string              `json:"shipping_type"`       // [Required]
	TradeOrderId      FlexString          `json:"trade_order_id"`      // [Required]
}
type GetReverseOrdersForSellerResponseDataResult struct {
	Total    FlexInt                                      `json:"total"`     // [Required]
	Success  FlexString                                   `json:"success"`   // [Required]
	PageNo   FlexString                                   `json:"page_no"`   // [Required]
	Items    []GetReverseOrdersForSellerResponseDataItems `json:"items"`     // [Required]
	PageSize FlexInt                                      `json:"page_size"` // [Required]
}
type InitReverseOrderCancelDecideResponse struct {
	BaseResponse // Common response fields
}
type InitReverseOrderCancelRequest struct {
	OrderId      string `json:"order_id"`      // [Required]
	ReasonDetail string `json:"reason_detail"` // [Required]
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
type ReverseOrderLine struct {
	PaidPrice          FlexString `json:"paid_price"`            // [Required]
	IsCancel           string     `json:"is_cancel"`             // [Required]
	ReasonId           string     `json:"reason_id"`             // [Required]
	ReasonSource       string     `json:"reason_source"`         // [Required]
	ReasonDesc         string     `json:"reason_desc"`           // [Required]
	ApplyReason        string     `json:"apply_reason"`          // [Required]
	ReasonType         string     `json:"reason_type"`           // [Required]
	SellerSku          string     `json:"seller_sku"`            // [Required]
	RefundAmount       string     `json:"refund_amount"`         // [Required]
	OrderLineId        string     `json:"order_line_id"`         // [Required]
	ReasonName         string     `json:"reason_name"`           // [Required]
	OrderId            int64      `json:"order_id"`              // [Required]
	ReverseOrderLineId string     `json:"reverse_order_line_id"` // [Required]
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
	ReturnOrderLineGmtCreate   FlexString                `json:"return_order_line_gmt_create"`   // [Required]
	PlatformSkuId              string                    `json:"platform_sku_id"`                // [Required]
	TradeOrderGmtCreate        FlexString                `json:"trade_order_gmt_create"`         // [Required]
	IsNeedRefund               FlexString                `json:"is_need_refund"`                 // [Required]
	ReasonText                 string                    `json:"reason_text"`                    // [Required]
	ItemUnitPrice              FlexString                `json:"item_unit_price"`                // [Required]
	Sla                        FlexString                `json:"sla"`                            // [Required]
	ReturnOrderLineGmtModified FlexString                `json:"return_order_line_gmt_modified"` // [Required]
	TradeOrderLineId           FlexString                `json:"trade_order_line_id"`            // [Required]
	OfcStatus                  string                    `json:"ofc_status"`                     // [Required]
	SellerSkuId                string                    `json:"seller_sku_id"`                  // [Required]
	RefundPaymentMethod        string                    `json:"refund_payment_method"`          // [Required]
	Buyer                      *ReverseOrderLinesBuyer   `json:"buyer"`                          // [Required]
	ReasonCode                 FlexString                `json:"reason_code"`                    // [Required]
	WhqcDecision               string                    `json:"whqc_decision"`                  // [Required]
	ReverseStatus              string                    `json:"reverse_status"`                 // [Required]
	RefundAmount               FlexString                `json:"refund_amount"`                  // [Required]
	TrackingNumber             string                    `json:"tracking_number"`                // [Required]
	ReceiverAddress            string                    `json:"receiver_address"`               // [Required]
	IsDispute                  FlexString                `json:"is_dispute"`                     // [Required]
	ReverseOrderLineId         FlexString                `json:"reverse_order_line_id"`          // [Required]
}
type ReverseOrderLinesBuyer struct {
	BuyerId FlexString `json:"buyer_id"` // [Required]
}
type ReverseOrderLinesProduct struct {
	ProductSku string  `json:"product_sku"` // [Required]
	ProductId  FlexInt `json:"product_id"`  // [Required]
}
type ReverseOrderOnlyRefundDecideRequest struct {
	ReverseOrderId string `json:"reverse_order_id"` // [Required]
	Action         string `json:"action"`           // [Required]
}
type ReverseOrderOnlyRefundDecideResponse struct {
	BaseResponse // Common response fields
}
type ReverseOrderReturnUpdateRequest struct {
	ReverseOrderId string `json:"reverse_order_id"` // [Required]
	Action         string `json:"action"`           // [Required]
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
	TipContent       string             `json:"tip_content"`        // [Required]
	TipType          string             `json:"tip_type"`           // [Required]
}
