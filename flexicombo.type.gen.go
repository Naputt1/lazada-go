package golazada

type ActivateFlexiComboResponse struct {
	BaseResponse // Common response fields
}
type AddFlexiComboProductsResponse struct {
	BaseResponse                                   // Common response fields
	Response     AddFlexiComboProductsResponseData `json:"data"` // Response data
}
type AddFlexiComboProductsResponseData struct {
	SkuId string `json:"sku id"` // [Required]
}
type CreateFlexiComboResponse struct {
	BaseResponse        // Common response fields
	Response     string `json:"data"` // Response data
}
type DataGiftSkus struct {
	ProductId int64 `json:"product_id"` // [Required]
	SkuId     int64 `json:"sku_id"`     // [Required]
}
type DeactivateFlexiComboResponse struct {
	BaseResponse // Common response fields
}
type DeleteFlexiComboProductsResponse struct {
	BaseResponse // Common response fields
}
type GetFlexiComboDetailsResponse struct {
	BaseResponse                                  // Common response fields
	Response     GetFlexiComboDetailsResponseData `json:"data"` // Response data
}
type GetFlexiComboDetailsResponseData struct {
	Stackable         string        `json:"stackable"`            // [Required]
	GiftBuyLimitValue []string      `json:"gift_buy_limit_value"` // [Required]
	Apply             string        `json:"apply"`                // [Required]
	GiftSkus          []GiftSkus    `json:"gift_skus"`            // [Required]
	EndTime           string        `json:"end_time"`             // [Required]
	SampleSkus        []GiftSkus    `json:"sample_skus"`          // [Required]
	DiscountValue     []interface{} `json:"discount_value"`       // [Required]
	Type              string        `json:"type"`                 // [Required]
	DiscountType      string        `json:"discount_type"`        // [Required]
	OrderUsedNumbers  string        `json:"order_used_numbers"`   // [Required]
	StartTime         string        `json:"start_time"`           // [Required]
	Name              string        `json:"name"`                 // [Required]
	PlatformChannel   string        `json:"platform_channel"`     // [Required]
	Id                int64         `json:"id"`                   // [Required]
	CriteriaType      string        `json:"criteria_type"`        // [Required]
	CriteriaValue     []string      `json:"criteria_value"`       // [Required]
	OrderNumbers      string        `json:"order_numbers"`        // [Required]
	Status            string        `json:"status"`               // [Required]
}
type GiftSkus struct {
	Tier      string `json:"tier"`       // [Required]
	ProductId int64  `json:"product_id"` // [Required]
	SkuId     int64  `json:"sku_id"`     // [Required]
}
type ListFlexiComboProductsResponse struct {
	BaseResponse                                    // Common response fields
	Response     ListFlexiComboProductsResponseData `json:"data"` // Response data
}
type ListFlexiComboProductsResponseData struct {
	DataList []interface{} `json:"data_list"` // [Required]
	Total    int64         `json:"total"`     // [Required]
	Current  string        `json:"current"`   // [Required]
	PageSize int64         `json:"page_size"` // [Required]
}
type ListFlexiComboResponse struct {
	BaseResponse                            // Common response fields
	Response     ListFlexiComboResponseData `json:"data"` // Response data
}
type ListFlexiComboResponseData struct {
	DataList []ResponseDataData `json:"data_list"` // [Required]
	Total    int64              `json:"total"`     // [Required]
	Current  string             `json:"current"`   // [Required]
	PageSize int64              `json:"page_size"` // [Required]
}
type ResponseDataData struct {
	Stackable        string         `json:"stackable"`          // [Required]
	Apply            string         `json:"apply"`              // [Required]
	GiftSkus         []DataGiftSkus `json:"gift_skus"`          // [Required]
	EndTime          string         `json:"end_time"`           // [Required]
	DiscountValue    []interface{}  `json:"discount_value"`     // [Required]
	SampleSkus       []DataGiftSkus `json:"sample_skus"`        // [Required]
	DiscountType     string         `json:"discount_type"`      // [Required]
	Type             string         `json:"type"`               // [Required]
	StartTime        string         `json:"start_time"`         // [Required]
	OrderUsedNumbers string         `json:"order_used_numbers"` // [Required]
	Name             string         `json:"name"`               // [Required]
	PlatformChannel  string         `json:"platform_channel"`   // [Required]
	Id               int64          `json:"id"`                 // [Required]
	CriteriaType     string         `json:"criteria_type"`      // [Required]
	CriteriaValue    []interface{}  `json:"criteria_value"`     // [Required]
	OrderNumbers     string         `json:"order_numbers"`      // [Required]
	Status           string         `json:"status"`             // [Required]
}
type UpdateFlexiComboResponse struct {
	BaseResponse // Common response fields
}
