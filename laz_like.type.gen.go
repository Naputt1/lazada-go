package golazada

type ApiResult struct {
	Result       string `json:"result"`       // [Required]
	Success      bool   `json:"success"`      // [Required]
	ErrorMessage string `json:"errorMessage"` // [Required]
	ErrorCode    string `json:"errorCode"`    // [Required]
}
type McnContentCancelSchedulePublishResponse struct {
	BaseResponse            // Common response fields
	ApiResult    *ApiResult `json:"api_result,omitempty"` //
}
type McnContentCompleteCreateVideoResponse struct {
	BaseResponse                                                  // Common response fields
	Result       *McnContentCompleteCreateVideoResponseDataResult `json:"result,omitempty"` //
}
type McnContentCompleteCreateVideoResponseDataResult struct {
	ResultMessage string `json:"result_message"` // [Required]
	Success       bool   `json:"success"`        // [Required]
	VideoId       string `json:"videoId"`        // [Required]
	ResultCode    string `json:"result_code"`    // [Required]
}
type McnContentCreateResponse struct {
	BaseResponse                                     // Common response fields
	Result       *McnContentCreateResponseDataResult `json:"result,omitempty"` //
}
type McnContentCreateResponseDataResult struct {
	ResultMessage string `json:"result_message"` // [Required]
	Success       bool   `json:"success"`        // [Required]
	ContentId     string `json:"contentId"`      // [Required]
	ResultCode    string `json:"result_code"`    // [Required]
}
type McnContentInitCreateVideoResponse struct {
	BaseResponse                                              // Common response fields
	Result       *McnContentInitCreateVideoResponseDataResult `json:"result,omitempty"` //
}
type McnContentInitCreateVideoResponseDataResult struct {
	UploadId      string `json:"upload_id"`      // [Required]
	ResultMessage string `json:"result_message"` // [Required]
	Success       bool   `json:"success"`        // [Required]
	ResultCode    string `json:"result_code"`    // [Required]
}
type McnContentListCategoryResponse struct {
	BaseResponse                                           // Common response fields
	Result       *McnContentListCategoryResponseDataResult `json:"result,omitempty"` //
}
type McnContentListCategoryResponseDataResult struct {
	ResultMessage string        `json:"result_message"` // [Required]
	Success       bool          `json:"success"`        // [Required]
	CategoryList  []interface{} `json:"categoryList"`   // [Required]
	ResultCode    string        `json:"result_code"`    // [Required]
}
type McnContentPropertyTagListResponse struct {
	BaseResponse                // Common response fields
	ResultCode    string        `json:"resultCode,omitempty"`    //
	ResultMessage string        `json:"resultMessage,omitempty"` //
	TagList       []interface{} `json:"tagList,omitempty"`       //
}
type McnContentReplySchedulePublishResponse struct {
	BaseResponse            // Common response fields
	ApiResult    *ApiResult `json:"api_result,omitempty"` //
}
type McnContentUploadImageResponse struct {
	BaseResponse                                          // Common response fields
	Result       *McnContentUploadImageResponseDataResult `json:"result,omitempty"` //
}
type McnContentUploadImageResponseDataResult struct {
	ResultMessage string `json:"result_message"` // [Required]
	Success       bool   `json:"success"`        // [Required]
	ResultCode    string `json:"result_code"`    // [Required]
	Url           string `json:"url"`            // [Required]
}
type McnContentUploadVideoBlockResponse struct {
	BaseResponse                                               // Common response fields
	Result       *McnContentUploadVideoBlockResponseDataResult `json:"result,omitempty"` //
}
type McnContentUploadVideoBlockResponseDataResult struct {
	ResultMessage string `json:"result_message"` // [Required]
	Success       bool   `json:"success"`        // [Required]
	ETag          string `json:"eTag"`           // [Required]
	ResultCode    string `json:"result_code"`    // [Required]
}
type McnProductValidatorResponse struct {
	BaseResponse                                        // Common response fields
	Result       *McnProductValidatorResponseDataResult `json:"result,omitempty"` //
}
type McnProductValidatorResponseDataResult struct {
	ResultMessage    string   `json:"result_message"`   // [Required]
	Success          bool     `json:"success"`          // [Required]
	HighRiskItemList []string `json:"highRiskItemList"` // [Required]
	NormalItemList   []string `json:"normalItemList"`   // [Required]
	ResultCode       string   `json:"result_code"`      // [Required]
}
type MCNQueryTagInfoByNameResponse struct {
	BaseResponse                        // Common response fields
	ApiResult    *ResponseDataApiResult `json:"api_result,omitempty"` //
}
type McnSimilarProductSearchResponse struct {
	BaseResponse                       // Common response fields
	ConfidentialityStatement string    `json:"confidentialityStatement,omitempty"` //
	ProductList              []Product `json:"productList,omitempty"`              //
	ResultCode               string    `json:"result_code,omitempty"`              //
	ResultMessage            string    `json:"result_message,omitempty"`           //
}
type Product struct {
	ProductId                int64  `json:"productId"`                // [Required]
	ImageUrl                 string `json:"imageUrl"`                 // [Required]
	ProductLink              string `json:"productLink"`              // [Required]
	MainPicture              string `json:"mainPicture"`              // [Required]
	ConfidentialityStatement string `json:"confidentialityStatement"` // [Required]
	SkuId                    int64  `json:"skuId"`                    // [Required]
}
type QueryContentReviewRecordsResponse struct {
	BaseResponse                                              // Common response fields
	Result       *QueryContentReviewRecordsResponseDataResult `json:"result,omitempty"` //
}
type QueryContentReviewRecordsResponseDataResult struct {
	Success       bool            `json:"success"`       // [Required]
	ResultCode    string          `json:"resultCode"`    // [Required]
	ResultMessage string          `json:"resultMessage"` // [Required]
	ReviewRecords []ReviewRecords `json:"reviewRecords"` // [Required]
}
type ResponseDataApiResult struct {
	Success       bool     `json:"success"`       // [Required]
	ResultCode    string   `json:"resultCode"`    // [Required]
	TagDTOList    []TagDTO `json:"tagDTOList"`    // [Required]
	ResultMessage string   `json:"resultMessage"` // [Required]
}
type ReviewRecords struct {
	ReviewedType            string `json:"reviewedType"`            // [Required]
	Reason                  string `json:"reason"`                  // [Required]
	ReviewedTime            string `json:"reviewedTime"`            // [Required]
	ContentId               string `json:"contentId"`               // [Required]
	CurrentContentBaseState string `json:"currentContentBaseState"` // [Required]
}
type TagDTO struct {
	Owner             string `json:"owner"`             // [Required]
	GmtModified       int64  `json:"gmtModified"`       // [Required]
	Creator           string `json:"creator"`           // [Required]
	TagCode           string `json:"tagCode"`           // [Required]
	Modifier          string `json:"modifier"`          // [Required]
	Description       string `json:"description"`       // [Required]
	GmtCreate         int64  `json:"gmtCreate"`         // [Required]
	TagName           string `json:"tagName"`           // [Required]
	ParentTagId       int64  `json:"parentTagId"`       // [Required]
	IsDeleted         string `json:"isDeleted"`         // [Required]
	TagPath           string `json:"tagPath"`           // [Required]
	Id                int64  `json:"id"`                // [Required]
	IsSetDeadline     string `json:"isSetDeadline"`     // [Required]
	Class             string `json:"class"`             // [Required]
	ParentTagCode     string `json:"parentTagCode"`     // [Required]
	TagCategoryCode   string `json:"tagCategoryCode"`   // [Required]
	EntityAttrVersion string `json:"entityAttrVersion"` // [Required]
}
