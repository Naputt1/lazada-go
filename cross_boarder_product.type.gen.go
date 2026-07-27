package golazada

type CountryInfo struct {
	Market       string  `json:"market"`        // [Required]
	Quantity     string  `json:"quantity"`      // [Required]
	Abs          string  `json:"abs"`           // [Required]
	SpecialPrice string  `json:"special_price"` // [Required]
	ItemId       int64   `json:"item_id"`       // [Required]
	Price        float64 `json:"price"`         // [Required]
	Currency     string  `json:"currency"`      // [Required]
	SkuId        int64   `json:"sku_id"`        // [Required]
}
type CountryPrice struct {
	Market         string `json:"market"`           // [Required]
	NoPostagePrice string `json:"no_postage_price"` // [Required]
	Currency       string `json:"currency"`         // [Required]
}
type CreateGlobalProductResponse struct {
	BaseResponse                                 // Common response fields
	Response     CreateGlobalProductResponseData `json:"data"` // Response data
}
type CreateGlobalProductResponseData struct {
	SkuList []Sku `json:"sku_list"` // [Required]
}
type DeleteIcProductFailResult struct {
	Market       string `json:"market"`       // [Required]
	ProductId    string `json:"productId"`    // [Required]
	UpdateMsg    string `json:"updateMsg"`    // [Required]
	UpdateResult string `json:"updateResult"` // [Required]
}
type DeleteMerchantProductResponse struct {
	BaseResponse                                   // Common response fields
	Response     DeleteMerchantProductResponseData `json:"data"` // Response data
}
type DeleteMerchantProductResponseData struct {
	DeleteICProductResult         string                      `json:"deleteICProductResult"`         // [Required]
	DeleteIcProductFailResultList []DeleteIcProductFailResult `json:"deleteIcProductFailResultList"` // [Required]
	DeleteGspProductResult        string                      `json:"deleteGspProductResult"`        // [Required]
}
type GetGlobalProductExtensionResponse struct {
	BaseResponse                                       // Common response fields
	Response     GetGlobalProductExtensionResponseData `json:"data"` // Response data
}
type GetGlobalProductExtensionResponseData struct {
	GlobalItemId string                 `json:"global_item_id"` // [Required]
	ItemId       int64                  `json:"item_id"`        // [Required]
	Products     []ResponseDataProducts `json:"products"`       // [Required]
}
type GetGlobalProductStatusResponse struct {
	BaseResponse        // Common response fields
	Response     string `json:"data"` // Response data
}
type GetRecommendPriceResponse struct {
	BaseResponse                               // Common response fields
	Response     GetRecommendPriceResponseData `json:"data"` // Response data
}
type GetRecommendPriceResponseData struct {
	GlobalItemId string             `json:"global_item_id"` // [Required]
	Skus         []ResponseDataSkus `json:"skus"`           // [Required]
	ItemId       int64              `json:"item_id"`        // [Required]
}
type GetUnfilledAttributeResponse struct {
	BaseResponse                                  // Common response fields
	Response     GetUnfilledAttributeResponseData `json:"data"`                   // Response data
	ErrorDetail  string                           `json:"error_detail,omitempty"` //
	Errors       string                           `json:"errors,omitempty"`       //
}
type GetUnfilledAttributeResponseData struct {
	TotalProducts int64                                      `json:"total_products"` // [Required]
	Products      []GetUnfilledAttributeResponseDataProducts `json:"products"`       // [Required]
}
type GetUnfilledAttributeResponseDataProducts struct {
	ItemId          int64                `json:"item_id"`          // [Required]
	PrimaryCategory int64                `json:"primary_category"` // [Required]
	SellerSku       string               `json:"seller_sku"`       // [Required]
	Attributes      []ProductsAttributes `json:"attributes"`       // [Required]
}
type GetUpgradableGlobalPlusProductListResponse struct {
	BaseResponse                                                // Common response fields
	Response     GetUpgradableGlobalPlusProductListResponseData `json:"data"` // Response data
}
type GetUpgradableGlobalPlusProductListResponseData struct {
	Type          string                                                   `json:"type"`           // [Required]
	TotalProducts int64                                                    `json:"total_products"` // [Required]
	CurrentPage   string                                                   `json:"current_page"`   // [Required]
	PageSize      int64                                                    `json:"page_size"`      // [Required]
	CurrentIndex  string                                                   `json:"current_index"`  // [Required]
	Products      []GetUpgradableGlobalPlusProductListResponseDataProducts `json:"products"`       // [Required]
}
type GetUpgradableGlobalPlusProductListResponseDataProducts struct {
	GlobalItemId string                     `json:"global_item_id"` // [Required]
	Skus         []ResponseDataProductsSkus `json:"skus"`           // [Required]
	ItemId       int64                      `json:"item_id"`        // [Required]
}
type ProductsAttributes struct {
	Advanced      *Advanced     `json:"advanced"`       // [Required]
	InputType     string        `json:"input_type"`     // [Required]
	Options       []interface{} `json:"options"`        // [Required]
	Name          string        `json:"name"`           // [Required]
	IsMandatory   int64         `json:"is_mandatory"`   // [Required]
	AttributeType string        `json:"attribute_type"` // [Required]
	Label         string        `json:"label"`          // [Required]
}
type ProductsSkus struct {
	SpecialPrice *SpecialPrice `json:"special_price"`  // [Required]
	Price        float64       `json:"price"`          // [Required]
	SellerSku    string        `json:"seller_sku"`     // [Required]
	NoPostageFee *SpecialPrice `json:"no_postage_fee"` // [Required]
	SkuId        int64         `json:"sku_id"`         // [Required]
}
type ResponseDataProducts struct {
	Market     string         `json:"market"`      // [Required]
	SemiStatus string         `json:"semi_status"` // [Required]
	Abs        string         `json:"abs"`         // [Required]
	Skus       []ProductsSkus `json:"skus"`        // [Required]
	ItemId     int64          `json:"item_id"`     // [Required]
}
type ResponseDataProductsSkus struct {
	PackageWidth  string        `json:"package_width"`  // [Required]
	PackageHeight string        `json:"package_height"` // [Required]
	ItemId        int64         `json:"item_id"`        // [Required]
	PackageLength string        `json:"package_length"` // [Required]
	SellerSku     string        `json:"seller_sku"`     // [Required]
	PackageWeight string        `json:"package_weight"` // [Required]
	SkuId         int64         `json:"sku_id"`         // [Required]
	CountryInfo   []CountryInfo `json:"country_info"`   // [Required]
}
type ResponseDataSkus struct {
	SellerSku    string         `json:"seller_sku"`    // [Required]
	SkuId        int64          `json:"sku_id"`        // [Required]
	CountryPrice []CountryPrice `json:"country_price"` // [Required]
}
type SemiProductUpdateResponse struct {
	BaseResponse                               // Common response fields
	Response     SemiProductUpdateResponseData `json:"data"` // Response data
}
type SemiProductUpdateResponseData struct {
	ProductId int64 `json:"product_id"` // [Required]
}
type SemiProductUpgradeResponse struct {
	BaseResponse                                // Common response fields
	Response     SemiProductUpgradeResponseData `json:"data"` // Response data
}
type SemiProductUpgradeResponseData struct {
	ProductId int64 `json:"product_id"` // [Required]
}
type Sku struct {
	SellerSku string `json:"seller_sku"` // [Required]
}
type SpecialPrice struct {
	Amount   int64  `json:"amount"`   // [Required]
	Currency string `json:"currency"` // [Required]
}
type UpdateGlobalProductAttributeResponse struct {
	BaseResponse        // Common response fields
	ErrorDetail  string `json:"error_detail,omitempty"` //
	Errors       string `json:"errors,omitempty"`       //
}
type UpdateIcProductFailResult struct {
	Market       string `json:"market"`        // [Required]
	ProductId    int64  `json:"product_id"`    // [Required]
	UpdateResult string `json:"update_result"` // [Required]
	UpdateMsg    string `json:"update_msg"`    // [Required]
}
type UpdateProductStatusResponse struct {
	BaseResponse                                 // Common response fields
	Response     UpdateProductStatusResponseData `json:"data"` // Response data
}
type UpdateProductStatusResponseData struct {
	UpdateIcProductResult         string                      `json:"update_ic_product_result"`           // [Required]
	UpdateGspProductResult        string                      `json:"update_gsp_product_result"`          // [Required]
	UpdateIcProductFailResultList []UpdateIcProductFailResult `json:"update_ic_product_fail_result_list"` // [Required]
}
