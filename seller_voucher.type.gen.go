package golazada

type FreeShippingSelectedProductListResponseDataData struct {
	SkuIds    []interface{} `json:"sku_ids"`    // [Required]
	ProductId int64         `json:"product_id"` // [Required]
}
type SellerVoucheDeleteSelectedProductSKUResponse struct {
	BaseResponse // Common response fields
}
type SellerVoucherActivateResponse struct {
	BaseResponse // Common response fields
}
type SellerVoucherAddSelectedProductSKUResponse struct {
	BaseResponse                                                // Common response fields
	Response     SellerVoucherAddSelectedProductSKUResponseData `json:"data"` // Response data
}
type SellerVoucherAddSelectedProductSKUResponseData struct {
	SkuId string `json:"sku id"` // [Required]
}
type SellerVoucherCreateResponse struct {
	BaseResponse        // Common response fields
	Response     string `json:"data"` // Response data
}
type SellerVoucherDeactivateResponse struct {
	BaseResponse // Common response fields
}
type SellerVoucherDetailQueryResponse struct {
	BaseResponse                                      // Common response fields
	Response     SellerVoucherDetailQueryResponseData `json:"data"` // Response data
}
type SellerVoucherDetailQueryResponseData struct {
	PeriodEndTime                 string `json:"period_end_time"`                   // [Required]
	MaxDiscountOfferingMoneyValue string `json:"max_discount_offering_money_value"` // [Required]
	CriteriaOverMoney             string `json:"criteria_over_money"`               // [Required]
	Apply                         string `json:"apply"`                             // [Required]
	VoucherName                   string `json:"voucher_name"`                      // [Required]
	VoucherCode                   string `json:"voucher_code"`                      // [Required]
	OfferingMoneyValueOff         string `json:"offering_money_value_off"`          // [Required]
	OrderUsedBudget               string `json:"order_used_budget"`                 // [Required]
	OfferingPercentageDiscountOff string `json:"offering_percentage_discount_off"`  // [Required]
	PeriodStartTime               string `json:"period_start_time"`                 // [Required]
	DisplayArea                   string `json:"display_area"`                      // [Required]
	VoucherType                   string `json:"voucher_type"`                      // [Required]
	Limit                         int64  `json:"limit"`                             // [Required]
	CollectStart                  string `json:"collect_start"`                     // [Required]
	VoucherDiscountType           string `json:"voucher_discount_type"`             // [Required]
	Currency                      string `json:"currency"`                          // [Required]
	Id                            int64  `json:"id"`                                // [Required]
	Issued                        string `json:"issued"`                            // [Required]
	Status                        string `json:"status"`                            // [Required]
}
type SellerVoucherListResponse struct {
	BaseResponse                               // Common response fields
	Response     SellerVoucherListResponseData `json:"data"` // Response data
}
type SellerVoucherListResponseData struct {
	DataList []SellerVoucherListResponseDataData `json:"data_list"` // [Required]
	Total    int64                               `json:"total"`     // [Required]
	Current  string                              `json:"current"`   // [Required]
	PageSize int64                               `json:"page_size"` // [Required]
}
type SellerVoucherListResponseDataData struct {
	PeriodEndTime                 string `json:"period_end_time"`                   // [Required]
	MaxDiscountOfferingMoneyValue string `json:"max_discount_offering_money_value"` // [Required]
	CriteriaOverMoney             string `json:"criteria_over_money"`               // [Required]
	Apply                         string `json:"apply"`                             // [Required]
	VoucherName                   string `json:"voucher_name"`                      // [Required]
	VoucherCode                   string `json:"voucher_code"`                      // [Required]
	OfferingMoneyValueOff         string `json:"offering_money_value_off"`          // [Required]
	OrderUsedBudget               string `json:"order_used_budget"`                 // [Required]
	OfferingPercentageDiscountOff string `json:"offering_percentage_discount_off"`  // [Required]
	PeriodStartTime               string `json:"period_start_time"`                 // [Required]
	DisplayArea                   string `json:"display_area"`                      // [Required]
	VoucherType                   string `json:"voucher_type"`                      // [Required]
	Limit                         int64  `json:"limit"`                             // [Required]
	CollectStart                  string `json:"collect_start"`                     // [Required]
	VoucherDiscountType           string `json:"voucher_discount_type"`             // [Required]
	Currency                      string `json:"currency"`                          // [Required]
	Id                            int64  `json:"id"`                                // [Required]
	Issued                        string `json:"issued"`                            // [Required]
	Status                        string `json:"status"`                            // [Required]
}
type SellerVoucherSelectedProductListResponse struct {
	BaseResponse                                              // Common response fields
	Response     SellerVoucherSelectedProductListResponseData `json:"data"` // Response data
}
type SellerVoucherSelectedProductListResponseData struct {
	DataList []FreeShippingSelectedProductListResponseDataData `json:"data_list"` // [Required]
	Total    int64                                             `json:"total"`     // [Required]
	Current  string                                            `json:"current"`   // [Required]
	PageSize int64                                             `json:"page_size"` // [Required]
}
type SellerVoucherUpdateResponse struct {
	BaseResponse        // Common response fields
	Response     string `json:"data"` // Response data
}
