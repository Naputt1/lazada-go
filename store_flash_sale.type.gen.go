package golazada

type ActivateStoreFlashSaleResponse struct {
	BaseResponse // Common response fields
}
type CreateStoreFlashSaleResponse struct {
	BaseResponse        // Common response fields
	Response     string `json:"data"` // Response data
}
type DeactivateStoreFlashSaleResponse struct {
	BaseResponse // Common response fields
}
type GetStoreFlashSaleResponse struct {
	BaseResponse                               // Common response fields
	Response     GetStoreFlashSaleResponseData `json:"data"` // Response data
}
type GetStoreFlashSaleResponseData struct {
	PeriodStartTime string     `json:"period_start_time"` // [Required]
	PeriodEndTime   string     `json:"period_end_time"`   // [Required]
	Id              int64      `json:"id"`                // [Required]
	SkuModelList    []SkuModel `json:"sku_model_list"`    // [Required]
	Status          string     `json:"status"`            // [Required]
}
type ListStoreFlashSaleResponse struct {
	BaseResponse                                // Common response fields
	Response     ListStoreFlashSaleResponseData `json:"data"` // Response data
}
type ListStoreFlashSaleResponseData struct {
	DataList []ListStoreFlashSaleResponseDataData `json:"data_list"` // [Required]
	Total    int64                                `json:"total"`     // [Required]
	Current  string                               `json:"current"`   // [Required]
	PageSize int64                                `json:"page_size"` // [Required]
}
type ListStoreFlashSaleResponseDataData struct {
	PeriodStartTime string     `json:"period_start_time"` // [Required]
	PeriodEndTime   string     `json:"period_end_time"`   // [Required]
	Id              int64      `json:"id"`                // [Required]
	SkuModelList    []SkuModel `json:"sku_model_list"`    // [Required]
	Status          string     `json:"status"`            // [Required]
}
type SkuModel struct {
	ItemId      int64   `json:"item_id"`      // [Required]
	Price       float64 `json:"price"`        // [Required]
	SkuId       int64   `json:"sku_id"`       // [Required]
	SkuQuantity string  `json:"sku_quantity"` // [Required]
}
type UpdateStoreFlashSaleResponse struct {
	BaseResponse // Common response fields
}
