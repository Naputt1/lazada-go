package golazada

type BaseInfo struct {
	RegisterCountry string `json:"registerCountry"` // [Required]
	Phone           string `json:"phone"`           // [Required]
	ShopName        string `json:"shopName"`        // [Required]
	ReqNo           string `json:"reqNo"`           // [Required]
	Email           string `json:"email"`           // [Required]
	Status          string `json:"status"`          // [Required]
}
type BatchQueryFollowStatusResponse struct {
	BaseResponse                                           // Common response fields
	Result       *BatchQueryFollowStatusResponseDataResult `json:"result,omitempty"` //
}
type BatchQueryFollowStatusResponseDataResult struct {
	Result  []interface{} `json:"result"`  // [Required]
	Success bool          `json:"success"` // [Required]
	Error   interface{}   `json:"error"`   // [Required]
}
type DataPageInfo struct {
	Current  string `json:"current"`  // [Required]
	Total    int64  `json:"total"`    // [Required]
	PageSize string `json:"pageSize"` // [Required]
}
type DataSource struct {
	MessageContent *MessageContent `json:"message_content"` // [Required]
	Id             int64           `json:"id"`              // [Required]
	Time           string          `json:"time"`            // [Required]
}
type GetCountryInfoResponse struct {
	BaseResponse                            // Common response fields
	Response     GetCountryInfoResponseData `json:"data"` // Response data
}
type GetCountryInfoResponseData struct {
	Label string `json:"label"` // [Required]
	Value string `json:"value"` // [Required]
}
type GetPickUpStoreListResponse struct {
	BaseResponse                                       // Common response fields
	Result       *GetPickUpStoreListResponseDataResult `json:"result,omitempty"` //
}
type GetSellerMetricsByIdResponse struct {
	BaseResponse                                  // Common response fields
	Response     GetSellerMetricsByIdResponseData `json:"data"` // Response data
}
type GetSellerMetricsByIdResponseData struct {
	MainCategoryName     string `json:"main_category_name"`     // [Required]
	ShipOnTime           string `json:"ship_on_time"`           // [Required]
	PositiveSellerRating string `json:"positive_seller_rating"` // [Required]
	ResponseTime         string `json:"response_time"`          // [Required]
	SellerId             int64  `json:"seller_id"`              // [Required]
	ResponseRate         string `json:"response_rate"`          // [Required]
	MainCategoryId       string `json:"main_category_id"`       // [Required]
}
type GetSellerPerformanceResponse struct {
	BaseResponse                                  // Common response fields
	Response     GetSellerPerformanceResponseData `json:"data"` // Response data
}
type GetSellerPerformanceResponseData struct {
	MainCategoryName string                   `json:"main_category_name"` // [Required]
	Indicators       []ResponseDataIndicators `json:"indicators"`         // [Required]
	SellerId         int64                    `json:"seller_id"`          // [Required]
	MainCategoryId   string                   `json:"main_category_id"`   // [Required]
}
type GetSellerRegisterInfoResponse struct {
	BaseResponse                                   // Common response fields
	Response     GetSellerRegisterInfoResponseData `json:"data"` // Response data
}
type GetSellerRegisterInfoResponseData struct {
	BaseInfoList  []BaseInfo `json:"baseInfoList"`  // [Required]
	CompanyName   string     `json:"companyName"`   // [Required]
	LicenseNumber string     `json:"licenseNumber"` // [Required]
}
type GetSellerResponse struct {
	BaseResponse                       // Common response fields
	Response     GetSellerResponseData `json:"data"` // Response data
}
type GetSellerResponseData struct {
	NameCompany         string `json:"name_company"`        // [Required]
	LogoUrl             string `json:"logo_url"`            // [Required]
	Name                string `json:"name"`                // [Required]
	Verified            string `json:"verified"`            // [Required]
	Location            string `json:"location"`            // [Required]
	MarketplaceEaseMode string `json:"marketplaceEaseMode"` // [Required]
	SellerId            int64  `json:"seller_id"`           // [Required]
	Email               string `json:"email"`               // [Required]
	ShortCode           string `json:"short_code"`          // [Required]
	Cb                  string `json:"cb"`                  // [Required]
	Status              string `json:"status"`              // [Required]
}
type GetSubAddressResponse struct {
	BaseResponse                           // Common response fields
	Response     GetSubAddressResponseData `json:"data"` // Response data
}
type GetSubAddressResponseData struct {
	Label string `json:"label"` // [Required]
	Value string `json:"value"` // [Required]
}
type GetWarehouseBySellerIdResponse struct {
	BaseResponse                                           // Common response fields
	Result       *GetWarehouseBySellerIdResponseDataResult `json:"result,omitempty"` //
}
type GetWarehouseBySellerIdResponseDataResult struct {
	NotSuccess string      `json:"not_success"` // [Required]
	Success    bool        `json:"success"`     // [Required]
	Module     interface{} `json:"module"`      // [Required]
	ErrorCode  string      `json:"error_code"`  // [Required]
	Repeated   string      `json:"repeated"`    // [Required]
	Retry      string      `json:"retry"`       // [Required]
}
type MessageContent struct {
	AppLink      string `json:"appLink"`      // [Required]
	WebLink      string `json:"webLink"`      // [Required]
	Description  string `json:"description"`  // [Required]
	Title        string `json:"title"`        // [Required]
	CategoryName string `json:"categoryName"` // [Required]
	Picture      string `json:"picture"`      // [Required]
}
type PaymentBindingResponse struct {
	BaseResponse                            // Common response fields
	Response     PaymentBindingResponseData `json:"data"` // Response data
}
type PaymentBindingResponseData struct {
	Result    string `json:"result"`    // [Required]
	Reason    string `json:"reason"`    // [Required]
	ShortCode string `json:"shortCode"` // [Required]
}
type QueryBuyboxHuntingInfoResponse struct {
	BaseResponse                                           // Common response fields
	Result       *QueryBuyboxHuntingInfoResponseDataResult `json:"result,omitempty"` //
}
type QueryBuyboxHuntingInfoResponseDataResult struct {
	Data       *QueryBuyboxHuntingInfoResponseDataResultData `json:"data"`       // [Required]
	RetSuccess string                                        `json:"retSuccess"` // [Required]
}
type QueryBuyboxHuntingInfoResponseDataResultData struct {
	ItemId    string `json:"itemId"`    // [Required]
	IsValid   string `json:"isValid"`   // [Required]
	Venture   string `json:"venture"`   // [Required]
	SkuId     string `json:"skuId"`     // [Required]
	PriceRank string `json:"priceRank"` // [Required]
}
type QueryWarehouseDetailInfoBySellerIdResponse struct {
	BaseResponse                                                       // Common response fields
	Result       *QueryWarehouseDetailInfoBySellerIdResponseDataResult `json:"result,omitempty"` //
}
type QueryWarehouseDetailInfoBySellerIdResponseDataResult struct {
	NotSuccess string                                                      `json:"not_success"` // [Required]
	Success    bool                                                        `json:"success"`     // [Required]
	Module     *QueryWarehouseDetailInfoBySellerIdResponseDataResultModule `json:"module"`      // [Required]
	ErrorCode  string                                                      `json:"error_code"`  // [Required]
	Repeated   string                                                      `json:"repeated"`    // [Required]
	ClassName  string                                                      `json:"class_name"`  // [Required]
	Retry      string                                                      `json:"retry"`       // [Required]
}
type QueryWarehouseDetailInfoBySellerIdResponseDataResultModule struct {
	Country        string `json:"country"`         // [Required]
	DefaultAddress string `json:"default_address"` // [Required]
	Province       string `json:"province"`        // [Required]
	City           string `json:"city"`            // [Required]
	DetailAddress  string `json:"detail_address"`  // [Required]
	WarehouseCode  string `json:"warehouse_code"`  // [Required]
	District       string `json:"district"`        // [Required]
	PostCode       string `json:"post_code"`       // [Required]
	Name           string `json:"name"`            // [Required]
	Status         string `json:"status"`          // [Required]
}
type ResponseDataIndicators struct {
	ActionUrl       string `json:"action_url"`       // [Required]
	Score           string `json:"score"`            // [Required]
	ScoreFormat     string `json:"score_format"`     // [Required]
	FormattedScore  string `json:"formatted_score"`  // [Required]
	Name            string `json:"name"`             // [Required]
	Tip             string `json:"tip"`              // [Required]
	Type            string `json:"type"`             // [Required]
	FormattedTarget string `json:"formatted_target"` // [Required]
	Target          string `json:"target"`           // [Required]
	TargetFormat    string `json:"target_format"`    // [Required]
	TargetRespected string `json:"target_respected"` // [Required]
}
type ResponseDataModel struct {
	Uid string `json:"uid"` // [Required]
}
type SaveSellerWarehouseInfoResponse struct {
	BaseResponse                                            // Common response fields
	Result       *SaveSellerWarehouseInfoResponseDataResult `json:"result,omitempty"` //
}
type SaveSellerWarehouseInfoResponseDataResult struct {
	NotSuccess string `json:"not_success"` // [Required]
	Success    bool   `json:"success"`     // [Required]
	Module     string `json:"module"`      // [Required]
	Repeated   string `json:"repeated"`    // [Required]
	Retry      string `json:"retry"`       // [Required]
}
type SellerCenterMsgListResponse struct {
	BaseResponse                                        // Common response fields
	Result       *SellerCenterMsgListResponseDataResult `json:"result,omitempty"` //
}
type SellerCenterMsgListResponseDataResult struct {
	Data      *SellerCenterMsgListResponseDataResultData `json:"data"`      // [Required]
	Success   bool                                       `json:"success"`   // [Required]
	ErrorCode string                                     `json:"errorCode"` // [Required]
	Type      string                                     `json:"type"`      // [Required]
	Error     string                                     `json:"error"`     // [Required]
}
type SellerCenterMsgListResponseDataResultData struct {
	PageInfo   *DataPageInfo `json:"pageInfo"`   // [Required]
	DataSource []DataSource  `json:"dataSource"` // [Required]
}
type SellerFieldVerifyResponse struct {
	BaseResponse                               // Common response fields
	Response     SellerFieldVerifyResponseData `json:"data"` // Response data
}
type SellerFieldVerifyResponseData struct {
	Result   string `json:"result"`    // [Required]
	ErrorMsg string `json:"error_msg"` // [Required]
	Name     string `json:"name"`      // [Required]
	ErrCode  string `json:"err_code"`  // [Required]
}
type SellerPolicyFetchResponse struct {
	BaseResponse        // Common response fields
	Response     string `json:"data"` // Response data
}
type SynchronizeSellerItemArConfigResponse struct {
	BaseResponse                    // Common response fields
	ErrorCode    string             `json:"errorCode,omitempty"` //
	ErrorMsg     string             `json:"errorMsg,omitempty"`  //
	Model        *ResponseDataModel `json:"model,omitempty"`     //
}
