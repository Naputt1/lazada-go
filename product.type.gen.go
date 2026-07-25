package golazada

type AdjustSellableQuantityResponse struct {
	BaseResponse // Common response fields
}
type BatchUpdateSizeChartResponse struct {
	BaseResponse // Common response fields
}
type CreateProductRequest struct {
	PrimaryCategoryId *int64  `json:"primary_category_id,omitempty"` // [Optional]
	Attributes        *string `json:"attributes,omitempty"`          // [Optional]
	Skus              *string `json:"skus,omitempty"`                // [Optional]
	Name              *string `json:"name,omitempty"`                // [Optional]
	Description       *string `json:"description,omitempty"`         // [Optional]
	ShortDescription  *string `json:"short_description,omitempty"`   // [Optional]
	Images            *string `json:"images,omitempty"`              // [Optional]
	Brand             *int64  `json:"brand,omitempty"`               // [Optional]
	Warranty          *string `json:"warranty,omitempty"`            // [Optional]
	WarrantyType      *string `json:"warranty_type,omitempty"`       // [Optional]
	SizeGuide         *string `json:"size_guide,omitempty"`          // [Optional]
	Source            *string `json:"source,omitempty"`              // [Optional]
	SaleStartDate     *string `json:"sale_start_date,omitempty"`     // [Optional]
	SaleEndDate       *string `json:"sale_end_date,omitempty"`       // [Optional]
	PackageWeight     *string `json:"package_weight,omitempty"`      // [Optional]
	PackageLength     *string `json:"package_length,omitempty"`      // [Optional]
	PackageWidth      *string `json:"package_width,omitempty"`       // [Optional]
	PackageHeight     *string `json:"package_height,omitempty"`      // [Optional]
}
type CreateProductResponse struct {
	BaseResponse                           // Common response fields
	Response     CreateProductResponseData `json:"data"` // Response data
}
type CreateProductResponseData struct {
	ItemId  *int64 `json:"item_id,omitempty"`  // [Optional]
	SkuList []Sku  `json:"sku_list,omitempty"` // [Optional]
}
type DeactivateProductResponse struct {
	BaseResponse // Common response fields
}
type GetBrandByPagesResponse struct {
	BaseResponse // Common response fields
}
type GetCategoryAttributesResponse struct {
	BaseResponse // Common response fields
}
type GetCategorySuggestionResponse struct {
	BaseResponse // Common response fields
}
type GetCategoryTreeResponse struct {
	BaseResponse // Common response fields
}
type GetNextCascadePropResponse struct {
	BaseResponse // Common response fields
}
type GetPreQcRulesResponse struct {
	BaseResponse // Common response fields
}
type GetProductContentScoreResponse struct {
	BaseResponse // Common response fields
}
type GetProductItemRequest struct {
	ItemId    *int64  `json:"item_id,omitempty" url:"item_id,omitempty"`       // [Optional]
	SellerSku *string `json:"seller_sku,omitempty" url:"seller_sku,omitempty"` // [Optional]
}
type GetProductItemResponse struct {
	BaseResponse                            // Common response fields
	Response     GetProductItemResponseData `json:"data"` // Response data
}
type GetProductItemResponseData struct {
	ItemId           *int64  `json:"item_id,omitempty"`           // [Optional]
	PrimaryCategory  *int64  `json:"primary_category,omitempty"`  // [Optional]
	Name             *string `json:"name,omitempty"`              // [Optional]
	Description      *string `json:"description,omitempty"`       // [Optional]
	ShortDescription *string `json:"short_description,omitempty"` // [Optional]
	Images           *string `json:"images,omitempty"`            // [Optional]
	Attributes       *string `json:"attributes,omitempty"`        // [Optional]
	Skus             []Skus  `json:"skus,omitempty"`              // [Optional]
}
type GetProductsRequest struct {
	Filter        *string `json:"filter,omitempty" url:"filter,omitempty"`                 // [Optional]
	Limit         *int64  `json:"limit,omitempty" url:"limit,omitempty"`                   // [Optional]
	Offset        *int64  `json:"offset,omitempty" url:"offset,omitempty"`                 // [Optional]
	CreatedAfter  *string `json:"created_after,omitempty" url:"created_after,omitempty"`   // [Optional]
	CreatedBefore *string `json:"created_before,omitempty" url:"created_before,omitempty"` // [Optional]
	UpdateAfter   *string `json:"update_after,omitempty" url:"update_after,omitempty"`     // [Optional]
	UpdateBefore  *string `json:"update_before,omitempty" url:"update_before,omitempty"`   // [Optional]
	Search        *string `json:"search,omitempty" url:"search,omitempty"`                 // [Optional]
}
type GetProductsResponse struct {
	BaseResponse                         // Common response fields
	Response     GetProductsResponseData `json:"data"` // Response data
}
type GetProductsResponseData struct {
	TotalProducts *int64     `json:"total_products,omitempty"` // [Optional]
	Products      []Products `json:"products,omitempty"`       // [Optional]
}
type GetQCAlertProductsResponse struct {
	BaseResponse // Common response fields
}
type GetResponseResponse struct {
	BaseResponse // Common response fields
}
type GetSellerItemLimitResponse struct {
	BaseResponse // Common response fields
}
type GetSizeChartTemplateResponse struct {
	BaseResponse // Common response fields
}
type GetUnfilledAttributeItemResponse struct {
	BaseResponse // Common response fields
}
type MigrateImageResponse struct {
	BaseResponse // Common response fields
}
type MigrateImagesResponse struct {
	BaseResponse // Common response fields
}
type ProductCheckResponse struct {
	BaseResponse // Common response fields
}
type Products struct {
	ItemId          *int64  `json:"item_id,omitempty"`          // [Optional]
	PrimaryCategory *int64  `json:"primary_category,omitempty"` // [Optional]
	Name            *string `json:"name,omitempty"`             // [Optional]
	SellerSku       *string `json:"seller_sku,omitempty"`       // [Optional]
}
type RemoveProductResponse struct {
	BaseResponse // Common response fields
}
type RemoveSkuResponse struct {
	BaseResponse // Common response fields
}
type SetImagesResponse struct {
	BaseResponse // Common response fields
}
type Sku struct {
	SellerSku *string `json:"seller_sku,omitempty"` // [Optional]
	SkuId     *int64  `json:"sku_id,omitempty"`     // [Optional]
}
type Skus struct {
	SellerSku     *string  `json:"seller_sku,omitempty"`     // [Optional]
	SkuId         *int64   `json:"sku_id,omitempty"`         // [Optional]
	Quantity      *int64   `json:"quantity,omitempty"`       // [Optional]
	Price         *float64 `json:"price,omitempty"`          // [Optional]
	PackageHeight *string  `json:"package_height,omitempty"` // [Optional]
	PackageLength *string  `json:"package_length,omitempty"` // [Optional]
	PackageWidth  *string  `json:"package_width,omitempty"`  // [Optional]
	PackageWeight *string  `json:"package_weight,omitempty"` // [Optional]
}
type UpdatePriceQuantityResponse struct {
	BaseResponse // Common response fields
}
type UpdateProductRequest struct {
	ItemId           *int64  `json:"item_id,omitempty"`           // [Optional]
	Attributes       *string `json:"attributes,omitempty"`        // [Optional]
	Name             *string `json:"name,omitempty"`              // [Optional]
	Description      *string `json:"description,omitempty"`       // [Optional]
	ShortDescription *string `json:"short_description,omitempty"` // [Optional]
}
type UpdateProductResponse struct {
	BaseResponse                           // Common response fields
	Response     UpdateProductResponseData `json:"data"` // Response data
}
type UpdateProductResponseData struct {
	ItemId *int64 `json:"item_id,omitempty"` // [Optional]
}
type UpdateSellableQuantityResponse struct {
	BaseResponse // Common response fields
}
type UploadImageResponse struct {
	BaseResponse // Common response fields
}
