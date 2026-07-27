package golazada

type AdjustSellableQuantityResponse struct {
	BaseResponse // Common response fields
}
type Advanced struct {
	IsKeyProp int64 `json:"is_key_prop"` // [Required]
}
type AttributesOptions struct {
	Name string `json:"name"` // [Required]
}
type BatchUpdateSizeChartResponse struct {
	BaseResponse // Common response fields
}
type CategorySuggestions struct {
	CategoryPath string `json:"categoryPath"` // [Required]
	CategoryName string `json:"categoryName"` // [Required]
	CategoryId   string `json:"categoryId"`   // [Required]
}
type Children struct {
	CategoryId int64  `json:"category_id"` // [Required]
	Var        bool   `json:"var"`         // [Required]
	Name       string `json:"name"`        // [Required]
	Leaf       bool   `json:"leaf"`        // [Required]
}
type CreateProductRequest struct {
	Payload string `json:"payload"` // [Required]
}
type CreateProductResponse struct {
	BaseResponse                           // Common response fields
	Response     CreateProductResponseData `json:"data"` // Response data
}
type CreateProductResponseData struct {
	ItemId     int64                          `json:"item_id"`     // [Required]
	SkuList    []CreateProductResponseDataSku `json:"sku_list"`    // [Required]
	ItemStatus string                         `json:"item_status"` // [Required]
}
type CreateProductResponseDataSku struct {
	ShopSku   string `json:"shop_sku"`   // [Required]
	SellerSku string `json:"seller_sku"` // [Required]
	SkuId     int64  `json:"sku_id"`     // [Required]
}
type DataItems struct {
	Score      string       `json:"score"`      // [Required]
	Total      int64        `json:"total"`      // [Required]
	ItemTitle  string       `json:"itemTitle"`  // [Required]
	Label      string       `json:"label"`      // [Required]
	Indicators []Indicators `json:"indicators"` // [Required]
	ImageList  []Image      `json:"imageList"`  // [Required]
	Key        string       `json:"key"`        // [Required]
	Group      string       `json:"group"`      // [Required]
	Latest     string       `json:"latest"`     // [Required]
}
type DeactivateProductResponse struct {
	BaseResponse // Common response fields
}
type GetBrandByPagesRequest struct {
	StartRow int64 `json:"startRow" url:"startRow"` // [Required]
	PageSize int64 `json:"pageSize" url:"pageSize"` // [Required]
}
type GetBrandByPagesResponse struct {
	BaseResponse                             // Common response fields
	Response     GetBrandByPagesResponseData `json:"data"` // Response data
}
type GetBrandByPagesResponseData struct {
	EnableTotal bool                 `json:"enable_total"` // [Required]
	StartRow    int64                `json:"start_row"`    // [Required]
	PageIndex   int64                `json:"page_index"`   // [Required]
	Module      []ResponseDataModule `json:"module"`       // [Required]
	TotalPage   int64                `json:"total_page"`   // [Required]
	PageSize    int64                `json:"page_size"`    // [Required]
	TotalRecord int64                `json:"total_record"` // [Required]
}
type GetCategoryAttributesRequest struct {
	PrimaryCategoryId int64  `json:"primary_category_id" url:"primary_category_id"` // [Required]
	LanguageCode      string `json:"language_code" url:"language_code"`             // [Required]
}
type GetCategoryAttributesResponse struct {
	BaseResponse                                   // Common response fields
	Response     GetCategoryAttributesResponseData `json:"data"` // Response data
}
type GetCategoryAttributesResponseData struct {
	Unit          *Unit                 `json:"unit"`           // [Required]
	Advanced      *Advanced             `json:"advanced"`       // [Required]
	IsSaleProp    int64                 `json:"is_sale_prop"`   // [Required]
	Name          string                `json:"name"`           // [Required]
	InputType     string                `json:"input_type"`     // [Required]
	Options       []ResponseDataOptions `json:"options"`        // [Required]
	IsMandatory   int64                 `json:"is_mandatory"`   // [Required]
	AttributeType string                `json:"attribute_type"` // [Required]
	Label         string                `json:"label"`          // [Required]
	Id            int64                 `json:"id"`             // [Required]
}
type GetCategorySuggestionResponse struct {
	BaseResponse                                   // Common response fields
	Response     GetCategorySuggestionResponseData `json:"data"` // Response data
}
type GetCategorySuggestionResponseData struct {
	CategorySuggestions []CategorySuggestions `json:"categorySuggestions"` // [Required]
}
type GetCategoryTreeResponse struct {
	BaseResponse                             // Common response fields
	Response     GetCategoryTreeResponseData `json:"data"` // Response data
}
type GetCategoryTreeResponseData struct {
	CategoryId int64      `json:"category_id"` // [Required]
	Children   []Children `json:"children"`    // [Required]
	Var        bool       `json:"var"`         // [Required]
	Name       string     `json:"name"`        // [Required]
	Leaf       bool       `json:"leaf"`        // [Required]
}
type GetNextCascadePropResponse struct {
	BaseResponse                                // Common response fields
	Response     GetNextCascadePropResponseData `json:"data"` // Response data
}
type GetNextCascadePropResponseData struct {
	Prop      *Prop       `json:"prop"`      // [Required]
	PropValue []PropValue `json:"propValue"` // [Required]
}
type GetPreQcRulesResponse struct {
	BaseResponse         // Common response fields
	Values       *Values `json:"values,omitempty"` //
}
type GetProductContentScoreResponse struct {
	BaseResponse                                           // Common response fields
	Result       *GetProductContentScoreResponseDataResult `json:"result,omitempty"` //
}
type GetProductContentScoreResponseDataResult struct {
	Data *GetProductContentScoreResponseDataResultData `json:"data"` // [Required]
}
type GetProductContentScoreResponseDataResultData struct {
	ProductTitle string      `json:"productTitle"` // [Required]
	Score        string      `json:"score"`        // [Required]
	Image        string      `json:"image"`        // [Required]
	Total        int64       `json:"total"`        // [Required]
	ProductId    string      `json:"productId"`    // [Required]
	Items        []DataItems `json:"items"`        // [Required]
}
type GetProductItemRequest struct {
	ItemId int64 `json:"item_id" url:"item_id"` // [Required]
}
type GetProductItemResponse struct {
	BaseResponse                            // Common response fields
	Response     GetProductItemResponseData `json:"data"` // Response data
}
type GetProductItemResponseData struct {
	CreatedTime     string                           `json:"created_time"`     // [Required]
	UpdatedTime     string                           `json:"updated_time"`     // [Required]
	Images          []string                           `json:"images"`           // [Required]
	Skus            []GetProductItemResponseDataSkus `json:"skus"`             // [Required]
	ImageSequence   *ImageSequence                   `json:"imageSequence"`    // [Required]
	ItemId          int64                            `json:"item_id"`          // [Required]
	HiddenStatus    string                           `json:"hiddenStatus"`     // [Required]
	BizSupplement   *ResponseDataBizSupplement       `json:"bizSupplement"`    // [Required]
	SuspendedSkus   []interface{}                    `json:"suspendedSkus"`    // [Required]
	SubStatus       string                           `json:"subStatus"`        // [Required]
	Variation       *Variation                       `json:"variation"`        // [Required]
	TrialProduct    string                           `json:"trialProduct"`     // [Required]
	RejectReason    []RejectReason                   `json:"rejectReason"`     // [Required]
	PrimaryCategory int64                            `json:"primary_category"` // [Required]
	MarketImages    string                           `json:"marketImages"`     // [Required]
	Attributes      *ResponseDataAttributes          `json:"attributes"`       // [Required]
	HiddenReason    string                           `json:"hiddenReason"`     // [Required]
	Status          string                           `json:"status"`           // [Required]
}
type GetProductItemResponseDataSkus struct {
	Status          string         `json:"Status"`            // [Required]
	Quantity        int64          `json:"quantity"`          // [Required]
	ImageSequence   *ImageSequence `json:"ImageSequence"`     // [Required]
	ProductWeight   string         `json:"product_weight"`    // [Required]
	Images          []string       `json:"Images"`            // [Required]
	SellerSku       string         `json:"SellerSku"`         // [Required]
	ShopSku         string         `json:"ShopSku"`           // [Required]
	Url             string         `json:"Url"`               // [Required]
	ComingSoon      string         `json:"coming_soon"`       // [Required]
	PackageWidth    string         `json:"package_width"`     // [Required]
	SpecialToTime   string         `json:"special_to_time"`   // [Required]
	SpecialFromTime string         `json:"special_from_time"` // [Required]
	PackageHeight   string         `json:"package_height"`    // [Required]
	SpecialPrice    int64          `json:"special_price"`     // [Required]
	Price           float64          `json:"price"`             // [Required]
	PackageLength   string         `json:"package_length"`    // [Required]
	PackageWeight   string         `json:"package_weight"`    // [Required]
	Available       int64          `json:"Available"`         // [Required]
	SkuId           int64          `json:"SkuId"`             // [Required]
	SpecialToDate   string         `json:"special_to_date"`   // [Required]
}
type GetProductsRequest struct {
	Filter        *string `json:"filter,omitempty" url:"filter,omitempty"`                 // [Optional]
	CreatedAfter  *string `json:"created_after,omitempty" url:"created_after,omitempty"`   // [Optional]
	CreatedBefore *string `json:"created_before,omitempty" url:"created_before,omitempty"` // [Optional]
	Offset        *int64  `json:"offset,omitempty" url:"offset,omitempty"`                 // [Optional]
	Limit         *int64  `json:"limit,omitempty" url:"limit,omitempty"`                   // [Optional]
}
type GetProductsResponse struct {
	BaseResponse                         // Common response fields
	Response     GetProductsResponseData `json:"data"` // Response data
}
type GetProductsResponseData struct {
	TotalProducts int64                             `json:"total_products"` // [Required]
	Products      []GetProductsResponseDataProducts `json:"products"`       // [Required]
}
type GetProductsResponseDataProducts struct {
	CreatedTime     string                                `json:"created_time"`     // [Required]
	UpdatedTime     string                                `json:"updated_time"`     // [Required]
	Images          []string                           `json:"images"`           // [Required]
	Skus            []GetProductsResponseDataProductsSkus `json:"skus"`             // [Required]
	ItemId          int64                                 `json:"item_id"`          // [Required]
	HiddenStatus    string                                `json:"hiddenStatus"`     // [Required]
	SuspendedSkus   []interface{}                         `json:"suspendedSkus"`    // [Required]
	SubStatus       string                                `json:"subStatus"`        // [Required]
	TrialProduct    string                                `json:"trialProduct"`     // [Required]
	RejectReason    []RejectReason                        `json:"rejectReason"`     // [Required]
	PrimaryCategory int64                                 `json:"primary_category"` // [Required]
	MarketImages    string                                `json:"marketImages"`     // [Required]
	Attributes      *ResponseDataAttributes               `json:"attributes"`       // [Required]
	HiddenReason    string                                `json:"hiddenReason"`     // [Required]
	Status          string                                `json:"status"`           // [Required]
}
type GetProductsResponseDataProductsSkus struct {
	Status          string   `json:"Status"`            // [Required]
	Quantity        int64    `json:"quantity"`          // [Required]
	ProductWeight   string   `json:"product_weight"`    // [Required]
	Images          []string `json:"Images"`            // [Required]
	SellerSku       string   `json:"SellerSku"`         // [Required]
	ShopSku         string   `json:"ShopSku"`           // [Required]
	Url             string   `json:"Url"`               // [Required]
	PackageWidth    string   `json:"package_width"`     // [Required]
	SpecialToTime   string   `json:"special_to_time"`   // [Required]
	SpecialFromTime string   `json:"special_from_time"` // [Required]
	PackageHeight   string   `json:"package_height"`    // [Required]
	SpecialPrice    int64    `json:"special_price"`     // [Required]
	Price           float64    `json:"price"`             // [Required]
	PackageLength   string   `json:"package_length"`    // [Required]
	PackageWeight   string   `json:"package_weight"`    // [Required]
	Available       int64    `json:"Available"`         // [Required]
	SkuId           int64    `json:"SkuId"`             // [Required]
	SpecialToDate   string   `json:"special_to_date"`   // [Required]
}
type GetQCAlertProductsResponse struct {
	BaseResponse                                // Common response fields
	Response     GetQCAlertProductsResponseData `json:"data"` // Response data
}
type GetQCAlertProductsResponseData struct {
	ProductId            string   `json:"productId"`            // [Required]
	SuggestionCategories []string `json:"suggestionCategories"` // [Required]
	CategoryId           string   `json:"categoryId"`           // [Required]
	DeactivationTime     string   `json:"deactivationTime"`     // [Required]
}
type GetResponseResponse struct {
	BaseResponse                         // Common response fields
	Response     GetResponseResponseData `json:"data"` // Response data
}
type GetResponseResponseData struct {
	Images []Images                        `json:"images"` // [Required]
	Errors []GetResponseResponseDataErrors `json:"errors"` // [Required]
}
type GetResponseResponseDataErrors struct {
	Msg         string `json:"msg"`          // [Required]
	Field       string `json:"field"`        // [Required]
	OriginalUrl string `json:"original_url"` // [Required]
}
type GetSellerItemLimitResponse struct {
	BaseResponse                                // Common response fields
	Response     GetSellerItemLimitResponseData `json:"data"`                 // Response data
	ErrorCodes   []interface{}                  `json:"errorCodes,omitempty"` //
	ErrorMsgs    []interface{}                  `json:"errorMsgs,omitempty"`  //
}
type GetSellerItemLimitResponseData struct {
	PayByrCnt       string `json:"payByrCnt"`       // [Required]
	PayItemCnt      string `json:"payItemCnt"`      // [Required]
	ItemLimit       string `json:"itemLimit"`       // [Required]
	OnlineItemCount string `json:"onlineItemCount"` // [Required]
}
type GetSizeChartTemplateResponse struct {
	BaseResponse                                  // Common response fields
	Response     GetSizeChartTemplateResponseData `json:"data"` // Response data
}
type GetSizeChartTemplateResponseData struct {
	Total              int64         `json:"total"`              // [Required]
	PageNo             string        `json:"pageNo"`             // [Required]
	TotalPage          string        `json:"totalPage"`          // [Required]
	PageSize           string        `json:"pageSize"`           // [Required]
	SizeChartResponses []interface{} `json:"sizeChartResponses"` // [Required]
}
type GetUnfilledAttributeItemResponse struct {
	BaseResponse                                                 // Common response fields
	Products      []GetUnfilledAttributeItemResponseDataProducts `json:"products,omitempty"`       //
	TotalProducts int64                                          `json:"total_products,omitempty"` //
}
type GetUnfilledAttributeItemResponseDataProducts struct {
	ItemId          int64                            `json:"item_id"`          // [Required]
	PrimaryCategory int64                            `json:"primary_category"` // [Required]
	Attributes      []ResponseDataProductsAttributes `json:"attributes"`       // [Required]
	SellerSkuId     string                           `json:"seller_sku_id"`    // [Required]
}
type Image struct {
	Score      string            `json:"score"`      // [Required]
	ImageUrl   string            `json:"imageUrl"`   // [Required]
	Text       string            `json:"text"`       // [Required]
	Type       string            `json:"type"`       // [Required]
	Indicators []ImageIndicators `json:"indicators"` // [Required]
	ImageType  string            `json:"imageType"`  // [Required]
}
type ImageIndicators struct {
	Text string `json:"text"` // [Required]
	Key  string `json:"key"`  // [Required]
}
type Images struct {
	HashCode string `json:"hash_code"` // [Required]
	Url      string `json:"url"`       // [Required]
}
type ImageSequence struct {
	Score       string   `json:"score"`       // [Required]
	NeedSuggest bool     `json:"needSuggest"` // [Required]
	IsDistinct  bool     `json:"isDistinct"`  // [Required]
	Url         []string `json:"url"`         // [Required]
}
type Indicators struct {
	Critical string `json:"critical"` // [Required]
	Text     string `json:"text"`     // [Required]
	Key      string `json:"key"`      // [Required]
}
type MigrateImageResponse struct {
	BaseResponse                          // Common response fields
	Response     MigrateImageResponseData `json:"data"` // Response data
}
type MigrateImageResponseData struct {
	Image *Images `json:"image"` // [Required]
}
type MigrateImagesResponse struct {
	BaseResponse        // Common response fields
	BatchId      string `json:"batch_id,omitempty"` //
}
type ProductCheckResponse struct {
	BaseResponse // Common response fields
}
type Prop struct {
	Name     string `json:"name"`     // [Required]
	Id       int64  `json:"id"`       // [Required]
	Required string `json:"required"` // [Required]
}
type PropValue struct {
	Name string `json:"name"` // [Required]
	Id   int64  `json:"id"`   // [Required]
	Leaf bool   `json:"leaf"` // [Required]
}
type RejectReason struct {
	Suggestion      string `json:"suggestion"`      // [Required]
	ViolationDetail string `json:"violationDetail"` // [Required]
}
type RemoveProductResponse struct {
	BaseResponse // Common response fields
}
type RemoveSkuResponse struct {
	BaseResponse // Common response fields
}
type ResponseDataAttributes struct {
	ShortDescription string `json:"short_description"` // [Required]
	Name             string `json:"name"`              // [Required]
	Description      string `json:"description"`       // [Required]
	NameEngravement  string `json:"name_engravement"`  // [Required]
	WarrantyType     string `json:"warranty_type"`     // [Required]
	GiftWrapping     string `json:"gift_wrapping"`     // [Required]
	PreorderDays     int64  `json:"preorder_days"`     // [Required]
	Brand            string `json:"brand"`             // [Required]
	Preorder         string `json:"preorder"`          // [Required]
}
type ResponseDataBizSupplement struct {
	GlobalPlusProductStatus int64 `json:"globalPlusProductStatus"` // [Required]
}
type ResponseDataModule struct {
	Name             string `json:"name"`              // [Required]
	GlobalIdentifier string `json:"global_identifier"` // [Required]
	NameEn           string `json:"name_en"`           // [Required]
	BrandId          int64  `json:"brand_id"`          // [Required]
}
type ResponseDataOptions struct {
	Name   string `json:"name"`    // [Required]
	EnName string `json:"en_name"` // [Required]
	Id     int64  `json:"id"`      // [Required]
}
type ResponseDataProductsAttributes struct {
	Advanced      *Advanced           `json:"advanced"`       // [Required]
	Name          string              `json:"name"`           // [Required]
	InputType     string              `json:"input_type"`     // [Required]
	Options       []AttributesOptions `json:"options"`        // [Required]
	IsMandatory   int64               `json:"is_mandatory"`   // [Required]
	AttributeType string              `json:"attribute_type"` // [Required]
	Label         string              `json:"label"`          // [Required]
}
type ResponseDataVariation struct {
	Variation1 *Variation1 `json:"Variation1"` // [Required]
	Variation2 *Variation1 `json:"Variation2"` // [Required]
	Variation3 *Variation1 `json:"Variation3"` // [Required]
	Variation4 *Variation1 `json:"Variation4"` // [Required]
}
type SetImagesResponse struct {
	BaseResponse // Common response fields
}
type Unit struct {
	Precision  string        `json:"precision"`  // [Required]
	Type       []interface{} `json:"type"`       // [Required]
	NumericMin string        `json:"numericMin"` // [Required]
	NumericMax string        `json:"numericMax"` // [Required]
}
type UpdatePriceQuantityResponse struct {
	BaseResponse // Common response fields
}
type UpdateProductRequest struct {
	Payload string `json:"payload"` // [Required]
}
type UpdateProductResponse struct {
	BaseResponse                           // Common response fields
	Response     UpdateProductResponseData `json:"data"` // Response data
}
type UpdateProductResponseData struct {
	ItemStatus string                 `json:"item_status"` // [Required]
	Variation  *ResponseDataVariation `json:"variation"`   // [Required]
}
type UpdateSellableQuantityResponse struct {
	BaseResponse // Common response fields
}
type UploadImageResponse struct {
	BaseResponse                         // Common response fields
	Response     UploadImageResponseData `json:"data"` // Response data
}
type UploadImageResponseData struct {
	Image *Images `json:"image"` // [Required]
}
type Values struct {
	ItemLimit         string        `json:"item_limit"`          // [Required]
	ItemCount         string        `json:"item_count"`          // [Required]
	RestrictedCateIds []interface{} `json:"restricted_cate_ids"` // [Required]
}
type Variation struct {
	Variation3 *Variation3 `json:"variation3"` // [Required]
	Variation4 *Variation3 `json:"variation4"` // [Required]
	Variation1 *Variation3 `json:"variation1"` // [Required]
	Variation2 *Variation3 `json:"variation2"` // [Required]
}
type Variation1 struct {
	HasImage  string        `json:"has_image"` // [Required]
	Name      string        `json:"name"`      // [Required]
	Options   []interface{} `json:"options"`   // [Required]
	Customize string        `json:"customize"` // [Required]
}
type Variation3 struct {
	HasImage  string        `json:"has_image"` // [Required]
	Name      string        `json:"name"`      // [Required]
	Options   []interface{} `json:"options"`   // [Required]
	Label     string        `json:"label"`     // [Required]
	Customize string        `json:"customize"` // [Required]
}
