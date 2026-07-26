package golazada

type Chronology struct {
	CalendarType string `json:"calendar_type"` // [Required]
	Id           int64  `json:"id"`            // [Required]
}
type FeeCreationDate struct {
	Offset     int64       `json:"offset"`       // [Required]
	Year       string      `json:"year"`         // [Required]
	DayOfYear  string      `json:"day_of_year"`  // [Required]
	Nano       string      `json:"nano"`         // [Required]
	Chronology *Chronology `json:"chronology"`   // [Required]
	MonthValue string      `json:"month_value"`  // [Required]
	DayOfMonth string      `json:"day_of_month"` // [Required]
	Minute     string      `json:"minute"`       // [Required]
	Second     string      `json:"second"`       // [Required]
	Month      string      `json:"month"`        // [Required]
	Hour       string      `json:"hour"`         // [Required]
	Zone       *Zone       `json:"zone"`         // [Required]
	DayOfWeek  string      `json:"day_of_week"`  // [Required]
}
type GetPayoutStatusResponse struct {
	BaseResponse                             // Common response fields
	Response     GetPayoutStatusResponseData `json:"data"` // Response data
}
type GetPayoutStatusResponseData struct {
	Subtotal2          string `json:"subtotal2"`             // [Required]
	Subtotal1          string `json:"subtotal1"`             // [Required]
	ShipmentFeeCredit  string `json:"shipment_fee_credit"`   // [Required]
	Payout             string `json:"payout"`                // [Required]
	ItemRevenue        string `json:"item_revenue"`          // [Required]
	CreatedAt          string `json:"created_at"`            // [Required]
	OtherRevenueTotal  string `json:"other_revenue_total"`   // [Required]
	FeesTotal          string `json:"fees_total"`            // [Required]
	Refunds            string `json:"refunds"`               // [Required]
	GuaranteeDeposit   string `json:"guarantee_deposit"`     // [Required]
	UpdatedAt          string `json:"updated_at"`            // [Required]
	FeesOnRefundsTotal string `json:"fees_on_refunds_total"` // [Required]
	ClosingBalance     string `json:"closing_balance"`       // [Required]
	Paid               string `json:"paid"`                  // [Required]
	OpeningBalance     string `json:"opening_balance"`       // [Required]
	StatementNumber    string `json:"statement_number"`      // [Required]
	ShipmentFee        string `json:"shipment_fee"`          // [Required]
}
type OrderCreationDate struct {
	Offset     int64  `json:"offset"`       // [Required]
	Year       string `json:"year"`         // [Required]
	DayOfYear  string `json:"day_of_year"`  // [Required]
	Nano       string `json:"nano"`         // [Required]
	Chronology string `json:"chronology"`   // [Required]
	MonthValue string `json:"month_value"`  // [Required]
	DayOfMonth string `json:"day_of_month"` // [Required]
	Minute     string `json:"minute"`       // [Required]
	Second     string `json:"second"`       // [Required]
	Month      string `json:"month"`        // [Required]
	Hour       string `json:"hour"`         // [Required]
	Zone       *Zone  `json:"zone"`         // [Required]
	DayOfWeek  string `json:"day_of_week"`  // [Required]
}
type OrderInfo struct {
	OrderItemStatus   string             `json:"order_item_status"`   // [Required]
	OrderCreationDate *OrderCreationDate `json:"order_creation_date"` // [Required]
}
type PackageInfo struct {
	DeliveryDate            *OrderCreationDate `json:"delivery_date"`             // [Required]
	DestinationAddress      string             `json:"destination_address"`       // [Required]
	OriginAddress           string             `json:"origin_address"`            // [Required]
	TrackingNumber          string             `json:"tracking_number"`           // [Required]
	BillingDate             *OrderCreationDate `json:"billing_date"`              // [Required]
	PackageChargeableWeight string             `json:"package_chargeable_weight"` // [Required]
}
type PageInfo struct {
	TotalCount int64  `json:"total_count"` // [Required]
	TotalPage  int64  `json:"total_page"`  // [Required]
	PageNum    string `json:"page_num"`    // [Required]
	PageSize   int64  `json:"page_size"`   // [Required]
}
type PayeeAccount struct {
	Description string `json:"description"` // [Required]
	Account     string `json:"account"`     // [Required]
}
type QueryAccountTransactionsResponse struct {
	BaseResponse                                      // Common response fields
	Response     QueryAccountTransactionsResponseData `json:"data"`          // Response data
	Msg          string                               `json:"msg,omitempty"` //
}
type QueryAccountTransactionsResponseData struct {
	PageInfo     *PageInfo      `json:"page_info"`    // [Required]
	Transactions []Transactions `json:"transactions"` // [Required]
}
type QueryLogisticsFeeDetailResponse struct {
	BaseResponse                                     // Common response fields
	Response     QueryLogisticsFeeDetailResponseData `json:"data"`             // Response data
	Remark       string                              `json:"remark,omitempty"` //
}
type QueryLogisticsFeeDetailResponseData struct {
	TenantId         string           `json:"tenant_id"`           // [Required]
	Amount           interface{}      `json:"amount"`              // [Required]
	SkuInfo          *SkuInfo         `json:"sku_info"`            // [Required]
	SellerShortCode  string           `json:"seller_short_code"`   // [Required]
	TradeOrderId     string           `json:"trade_order_id"`      // [Required]
	FeeCreationDate  *FeeCreationDate `json:"fee_creation_date"`   // [Required]
	TradeOrderLineId string           `json:"trade_order_line_id"` // [Required]
	StatementId      string           `json:"statement_id"`        // [Required]
	OrderInfo        *OrderInfo       `json:"order_info"`          // [Required]
	FeeName          string           `json:"fee_name"`            // [Required]
	FeeCode          string           `json:"fee_code"`            // [Required]
	Currency         string           `json:"currency"`            // [Required]
	PackageInfo      *PackageInfo     `json:"package_info"`        // [Required]
	TaxInAmount      interface{}      `json:"tax_in_amount"`       // [Required]
	SellerId         int64            `json:"seller_id"`           // [Required]
	StatementPeriod  string           `json:"statement_period"`    // [Required]
}
type QueryTransactionDetailsResponse struct {
	BaseResponse                                     // Common response fields
	Response     QueryTransactionDetailsResponseData `json:"data"` // Response data
}
type QueryTransactionDetailsResponseData struct {
	OrderNo             string `json:"order_no"`               // [Required]
	TransactionDate     string `json:"transaction_date"`       // [Required]
	Amount              string `json:"amount"`                 // [Required]
	PaidStatus          string `json:"paid_status"`            // [Required]
	ShippingProvider    string `json:"shipping_provider"`      // [Required]
	WHTIncludedInAmount string `json:"WHT_included_in_amount"` // [Required]
	PaymentRefId        string `json:"payment_ref_id"`         // [Required]
	LazadaSku           string `json:"lazada_sku"`             // [Required]
	FeeType             string `json:"fee_type"`               // [Required]
	TransactionType     string `json:"transaction_type"`       // [Required]
	OrderItemNo         string `json:"orderItem_no"`           // [Required]
	OrderItemStatus     string `json:"orderItem_status"`       // [Required]
	Reference           string `json:"reference"`              // [Required]
	FeeName             string `json:"fee_name"`               // [Required]
	ShippingSpeed       string `json:"shipping_speed"`         // [Required]
	WHTAmount           string `json:"WHT_amount"`             // [Required]
	TransactionNumber   string `json:"transaction_number"`     // [Required]
	SellerSku           string `json:"seller_sku"`             // [Required]
	Statement           string `json:"statement"`              // [Required]
	Details             string `json:"details"`                // [Required]
	Comment             string `json:"comment"`                // [Required]
	VATInAmount         string `json:"VAT_in_amount"`          // [Required]
	ShipmentType        string `json:"shipment_type"`          // [Required]
}
type Rules struct {
	FixedOffset string `json:"fixed_offset"` // [Required]
}
type SkuInfo struct {
	ItemDetails string `json:"item_details"` // [Required]
	SellerSku   string `json:"seller_sku"`   // [Required]
	LazadaSku   string `json:"lazada_sku"`   // [Required]
}
type Tracking struct {
	UpdateTime string `json:"update_time"` // [Required]
	Name       string `json:"name"`        // [Required]
	Remark     string `json:"remark"`      // [Required]
	Status     string `json:"status"`      // [Required]
}
type Transactions struct {
	PmtReference      string        `json:"pmt_reference"`      // [Required]
	PayeeAccount      *PayeeAccount `json:"payee_account"`      // [Required]
	Amount            string        `json:"amount"`             // [Required]
	SubType           string        `json:"sub_type"`           // [Required]
	TransactionNumber string        `json:"transaction_number"` // [Required]
	TransactionTime   string        `json:"transaction_time"`   // [Required]
	Currency          string        `json:"currency"`           // [Required]
	TrackingList      []Tracking    `json:"tracking_list"`      // [Required]
	Type              string        `json:"type"`               // [Required]
	Remarks           string        `json:"remarks"`            // [Required]
}
type Zone struct {
	Rules *Rules `json:"rules"` // [Required]
	Id    int64  `json:"id"`    // [Required]
}
