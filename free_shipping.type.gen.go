package golazada

type FreeShippingActivateResponse struct {
	BaseResponse // Common response fields
}
type FreeShippingAddSelectedProductSKUResponse struct {
	BaseResponse                                               // Common response fields
	Response     FreeShippingAddSelectedProductSKUResponseData `json:"data"` // Response data
}
type FreeShippingAddSelectedProductSKUResponseData struct {
	SkuId string `json:"sku id"` // [Required]
}
type FreeShippingCreateResponse struct {
	BaseResponse        // Common response fields
	Response     string `json:"data"` // Response data
}
type FreeShippingDeactivateResponse struct {
	BaseResponse // Common response fields
}
type FreeShippingDeleteSelectedProductSKUResponse struct {
	BaseResponse // Common response fields
}
type FreeShippingDeliveryOptionsQueryResponse struct {
	BaseResponse                                              // Common response fields
	Response     FreeShippingDeliveryOptionsQueryResponseData `json:"data"` // Response data
}
type FreeShippingDeliveryOptionsQueryResponseData struct {
	Name  string `json:"name"`  // [Required]
	Value string `json:"value"` // [Required]
}
type FreeShippingGetResponse struct {
	BaseResponse                             // Common response fields
	Response     FreeShippingGetResponseData `json:"data"` // Response data
}
type FreeShippingGetResponseData struct {
	PeriodEndTime   string        `json:"period_end_time"`   // [Required]
	CategoryName    string        `json:"category_name"`     // [Required]
	Apply           string        `json:"apply"`             // [Required]
	BudgetValue     string        `json:"budget_value"`      // [Required]
	CampaignTag     string        `json:"campaign_tag"`      // [Required]
	RegionType      string        `json:"region_type"`       // [Required]
	RegionValue     []interface{} `json:"region_value"`      // [Required]
	PromoTier       *PromoTier    `json:"promo_tier"`        // [Required]
	TemplateCode    string        `json:"template_code"`     // [Required]
	PeriodStartTime string        `json:"period_start_time"` // [Required]
	PromotionName   string        `json:"promotion_name"`    // [Required]
	UsedBudgetValue string        `json:"used_budget_value"` // [Required]
	PlatformChannel string        `json:"platform_channel"`  // [Required]
	TemplateType    string        `json:"template_type"`     // [Required]
	Currency        string        `json:"currency"`          // [Required]
	Id              int64         `json:"id"`                // [Required]
	BudgetType      string        `json:"budget_type"`       // [Required]
	PeriodType      string        `json:"period_type"`       // [Required]
	DeliveryOption  string        `json:"delivery_option"`   // [Required]
	Status          string        `json:"status"`            // [Required]
}
type FreeShippingListResponse struct {
	BaseResponse                              // Common response fields
	Response     FreeShippingListResponseData `json:"data"` // Response data
}
type FreeShippingListResponseData struct {
	DataList []FreeShippingListResponseDataData `json:"data_list"` // [Required]
	Total    int64                              `json:"total"`     // [Required]
	Current  string                             `json:"current"`   // [Required]
	PageSize int64                              `json:"page_size"` // [Required]
}
type FreeShippingListResponseDataData struct {
	PeriodEndTime   string        `json:"period_end_time"`   // [Required]
	CategoryName    string        `json:"category_name"`     // [Required]
	Apply           string        `json:"apply"`             // [Required]
	BudgetValue     string        `json:"budget_value"`      // [Required]
	CampaignTag     string        `json:"campaign_tag"`      // [Required]
	RegionType      string        `json:"region_type"`       // [Required]
	RegionValue     []interface{} `json:"region_value"`      // [Required]
	PromoTier       *PromoTier    `json:"promo_tier"`        // [Required]
	TemplateCode    string        `json:"template_code"`     // [Required]
	PeriodStartTime string        `json:"period_start_time"` // [Required]
	PromotionName   string        `json:"promotion_name"`    // [Required]
	UsedBudgetValue string        `json:"used_budget_value"` // [Required]
	PlatformChannel string        `json:"platform_channel"`  // [Required]
	TemplateType    string        `json:"template_type"`     // [Required]
	Currency        string        `json:"currency"`          // [Required]
	Id              int64         `json:"id"`                // [Required]
	BudgetType      string        `json:"budget_type"`       // [Required]
	PeriodType      string        `json:"period_type"`       // [Required]
	DeliveryOption  string        `json:"delivery_option"`   // [Required]
	Status          string        `json:"status"`            // [Required]
}
type FreeShippingRegionsQueryResponse struct {
	BaseResponse                                      // Common response fields
	Response     FreeShippingRegionsQueryResponseData `json:"data"` // Response data
}
type FreeShippingRegionsQueryResponseData struct {
	Name  string `json:"name"`  // [Required]
	Value string `json:"value"` // [Required]
}
type FreeShippingSelectedProductListResponse struct {
	BaseResponse                                             // Common response fields
	Response     FreeShippingSelectedProductListResponseData `json:"data"` // Response data
}
type FreeShippingSelectedProductListResponseData struct {
	DataList []FreeShippingSelectedProductListResponseDataData `json:"data_list"` // [Required]
	Total    int64                                             `json:"total"`     // [Required]
	Current  string                                            `json:"current"`   // [Required]
	PageSize int64                                             `json:"page_size"` // [Required]
}
type FreeShippingUpdateResponse struct {
	BaseResponse        // Common response fields
	Response     string `json:"data"` // Response data
}
type PromoTier struct {
	Tiers        []Tiers `json:"tiers"`         // [Required]
	DealCriteria string  `json:"deal_criteria"` // [Required]
	DiscountType string  `json:"discount_type"` // [Required]
}
type Tiers struct {
	Filter string `json:"filter"` // [Required]
	Result string `json:"result"` // [Required]
}
