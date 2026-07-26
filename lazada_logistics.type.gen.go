package golazada

type AppliedVas struct {
	VasExchangeOrderOption        string `json:"vasExchangeOrderOption"`        // [Required]
	OpenBox                       string `json:"openBox"`                       // [Required]
	VasFdCollectShippingFeeOption string `json:"vasFdCollectShippingFeeOption"` // [Required]
	VasFdStorageOption            string `json:"vasFdStorageOption"`            // [Required]
	VasFdCallOption               string `json:"vasFdCallOption"`               // [Required]
	VasPartialDeliveryOption      string `json:"vasPartialDeliveryOption"`      // [Required]
}
type Content struct {
	GmtModified    string   `json:"gmtModified"`    // [Required]
	Attachments    string   `json:"attachments"`    // [Required]
	OrderId        string   `json:"orderId"`        // [Required]
	Subject        string   `json:"subject"`        // [Required]
	ContactName    string   `json:"contactName"`    // [Required]
	BuyerEmail     string   `json:"buyerEmail"`     // [Required]
	SellerName     string   `json:"sellerName"`     // [Required]
	Description    string   `json:"description"`    // [Required]
	BuyerName      string   `json:"buyerName"`      // [Required]
	GmtCreate      string   `json:"gmtCreate"`      // [Required]
	SellerEmail    string   `json:"sellerEmail"`    // [Required]
	RatingStar     string   `json:"ratingStar"`     // [Required]
	GmtDeleted     string   `json:"gmtDeleted"`     // [Required]
	RatingRemark   string   `json:"ratingRemark"`   // [Required]
	MerchantId     string   `json:"merchantId"`     // [Required]
	CaseId         string   `json:"caseId"`         // [Required]
	CaseTemplateId string   `json:"caseTemplateId"` // [Required]
	SellerPhoneNo  string   `json:"sellerPhoneNo"`  // [Required]
	Attributes     string   `json:"attributes"`     // [Required]
	Id             int64    `json:"id"`             // [Required]
	TrackingNumber string   `json:"trackingNumber"` // [Required]
	RatingReasons  []string `json:"ratingReasons"`  // [Required]
	CategoryId     string   `json:"categoryId"`     // [Required]
	Status         string   `json:"status"`         // [Required]
}
type ConvertedAddress struct {
	Details string `json:"details"` // [Required]
	Id      int64  `json:"id"`      // [Required]
}
type CreateCustomerAccountRelationshipByOTPResponse struct {
	BaseResponse          // Common response fields
	ErrorCode    string   `json:"errorCode,omitempty"`    //
	ErrorMessage string   `json:"errorMessage,omitempty"` //
	Errors       []Errors `json:"errors,omitempty"`       //
	Retryable    string   `json:"retryable,omitempty"`    //
	TraceId      string   `json:"traceId,omitempty"`      //
}
type CreateCustomerAccountRelationshipForExternalResponse struct {
	BaseResponse          // Common response fields
	ErrorCode    string   `json:"errorCode,omitempty"`    //
	ErrorMessage string   `json:"errorMessage,omitempty"` //
	Errors       []Errors `json:"errors,omitempty"`       //
	Retryable    string   `json:"retryable,omitempty"`    //
	TraceId      string   `json:"traceId,omitempty"`      //
}
type CreateOrUpdateCustomerWarehouseResponse struct {
	BaseResponse                                             // Common response fields
	Response     CreateOrUpdateCustomerWarehouseResponseData `json:"data"`                   // Response data
	ErrorCode    string                                      `json:"errorCode,omitempty"`    //
	ErrorMessage string                                      `json:"errorMessage,omitempty"` //
	Errors       []Errors                                    `json:"errors,omitempty"`       //
	Retryable    string                                      `json:"retryable,omitempty"`    //
	TraceId      string                                      `json:"traceId,omitempty"`      //
}
type CreateOrUpdateCustomerWarehouseResponseData struct {
	ConvertedAddress *ConvertedAddress `json:"convertedAddress"` // [Required]
}
type EpisGetDeliveryOptionsResponse struct {
	BaseResponse                                    // Common response fields
	Response     EpisGetDeliveryOptionsResponseData `json:"data"`                   // Response data
	ErrorCode    string                             `json:"errorCode,omitempty"`    //
	ErrorMessage string                             `json:"errorMessage,omitempty"` //
	Errors       []Errors                           `json:"errors,omitempty"`       //
	Retryable    string                             `json:"retryable,omitempty"`    //
	TraceId      string                             `json:"traceId,omitempty"`      //
}
type EpisGetDeliveryOptionsResponseData struct {
	LastMileShippingProvider      string `json:"lastMileShippingProvider"`      // [Required]
	FirstMileShippingProviderSlug string `json:"firstMileShippingProviderSlug"` // [Required]
	FirstMileDeliveryType         string `json:"firstMileDeliveryType"`         // [Required]
	FirstMileShippingProvider     string `json:"firstMileShippingProvider"`     // [Required]
	DeliveryOption                string `json:"deliveryOption"`                // [Required]
	PickupTargetCutoffTime        string `json:"pickupTargetCutoffTime"`        // [Required]
	LastMileShippingProviderSlug  string `json:"lastMileShippingProviderSlug"`  // [Required]
}
type EpisPackageCancellationResponse struct {
	BaseResponse          // Common response fields
	ErrorCode    string   `json:"errorCode,omitempty"`    //
	ErrorMessage string   `json:"errorMessage,omitempty"` //
	Errors       []Errors `json:"errors,omitempty"`       //
	Retryable    string   `json:"retryable,omitempty"`    //
	TraceId      string   `json:"traceId,omitempty"`      //
}
type EpisPackageCancellationV3Response struct {
	BaseResponse          // Common response fields
	ErrorCode    string   `json:"errorCode,omitempty"`    //
	ErrorMessage string   `json:"errorMessage,omitempty"` //
	Errors       []Errors `json:"errors,omitempty"`       //
	Retryable    string   `json:"retryable,omitempty"`    //
	TraceId      string   `json:"traceId,omitempty"`      //
}
type EpisPackageConsignmentResponse struct {
	BaseResponse                                    // Common response fields
	Response     EpisPackageConsignmentResponseData `json:"data"`                   // Response data
	ErrorCode    string                             `json:"errorCode,omitempty"`    //
	ErrorMessage string                             `json:"errorMessage,omitempty"` //
	Errors       []Errors                           `json:"errors,omitempty"`       //
	Retryable    string                             `json:"retryable,omitempty"`    //
	TraceId      string                             `json:"traceId,omitempty"`      //
}
type EpisPackageConsignmentResponseData struct {
	RouteCode                 string                    `json:"routeCode"`                 // [Required]
	LastMileShippingProvider  *LastMileShippingProvider `json:"lastMileShippingProvider"`  // [Required]
	AoiName                   string                    `json:"aoiName"`                   // [Required]
	Origin                    *ConvertedAddress         `json:"origin"`                    // [Required]
	Options                   *Options                  `json:"options"`                   // [Required]
	AppliedVas                *AppliedVas               `json:"appliedVas"`                // [Required]
	Destination               *ConvertedAddress         `json:"destination"`               // [Required]
	FirstMileShippingProvider *LastMileShippingProvider `json:"firstMileShippingProvider"` // [Required]
	PortCode                  string                    `json:"portCode"`                  // [Required]
	TrackingNumber            string                    `json:"trackingNumber"`            // [Required]
}
type EpisPackageConsignmentV2Response struct {
	BaseResponse                                      // Common response fields
	Response     EpisPackageConsignmentV2ResponseData `json:"data"`                   // Response data
	ErrorCode    string                               `json:"errorCode,omitempty"`    //
	ErrorMessage string                               `json:"errorMessage,omitempty"` //
	Errors       []Errors                             `json:"errors,omitempty"`       //
	Retryable    string                               `json:"retryable,omitempty"`    //
	TraceId      string                               `json:"traceId,omitempty"`      //
}
type EpisPackageConsignmentV2ResponseData struct {
	RouteCode                 string                    `json:"routeCode"`                 // [Required]
	LogisticsOrderId          string                    `json:"logisticsOrderId"`          // [Required]
	LastMileShippingProvider  *LastMileShippingProvider `json:"lastMileShippingProvider"`  // [Required]
	AoiName                   string                    `json:"aoiName"`                   // [Required]
	Origin                    *ConvertedAddress         `json:"origin"`                    // [Required]
	Options                   *Options                  `json:"options"`                   // [Required]
	AppliedVas                *AppliedVas               `json:"appliedVas"`                // [Required]
	Destination               *ConvertedAddress         `json:"destination"`               // [Required]
	FirstMileShippingProvider *LastMileShippingProvider `json:"firstMileShippingProvider"` // [Required]
	PortCode                  string                    `json:"portCode"`                  // [Required]
	TrackingNumber            string                    `json:"trackingNumber"`            // [Required]
}
type EpisPackageCreationResponse struct {
	BaseResponse                                 // Common response fields
	Response     EpisPackageCreationResponseData `json:"data"`                   // Response data
	ErrorCode    string                          `json:"errorCode,omitempty"`    //
	ErrorMessage string                          `json:"errorMessage,omitempty"` //
	Errors       []Errors                        `json:"errors,omitempty"`       //
	Retryable    string                          `json:"retryable,omitempty"`    //
	TraceId      string                          `json:"traceId,omitempty"`      //
}
type EpisPackageCreationResponseData struct {
	RouteCode                 string                    `json:"routeCode"`                 // [Required]
	MaxEta                    string                    `json:"maxEta"`                    // [Required]
	LastMileShippingProvider  *LastMileShippingProvider `json:"lastMileShippingProvider"`  // [Required]
	Origin                    *ConvertedAddress         `json:"origin"`                    // [Required]
	Destination               *ConvertedAddress         `json:"destination"`               // [Required]
	FirstMileShippingProvider *LastMileShippingProvider `json:"firstMileShippingProvider"` // [Required]
	PortCode                  string                    `json:"portCode"`                  // [Required]
	AoiName                   string                    `json:"aoiName"`                   // [Required]
	PackageCode               string                    `json:"packageCode"`               // [Required]
	Options                   *Options                  `json:"options"`                   // [Required]
	AppliedVas                *AppliedVas               `json:"appliedVas"`                // [Required]
	MinEta                    string                    `json:"minEta"`                    // [Required]
	TrackingNumber            string                    `json:"trackingNumber"`            // [Required]
}
type EpisPackageInfoUpdateResponse struct {
	BaseResponse                                   // Common response fields
	Response     EpisPackageInfoUpdateResponseData `json:"data"`                   // Response data
	ErrorCode    string                            `json:"errorCode,omitempty"`    //
	ErrorMessage string                            `json:"errorMessage,omitempty"` //
	Errors       []ResponseDataErrors              `json:"errors,omitempty"`       //
	Retryable    string                            `json:"retryable,omitempty"`    //
	TraceId      string                            `json:"traceId,omitempty"`      //
}
type EpisPackageInfoUpdateResponseData struct {
	ConvertedAddress *ResponseDataConvertedAddress `json:"convertedAddress"` // [Required]
}
type EpisPackagePrintAwbResponse struct {
	BaseResponse                                 // Common response fields
	Response     EpisPackagePrintAwbResponseData `json:"data"`                   // Response data
	ErrorCode    string                          `json:"errorCode,omitempty"`    //
	ErrorMessage string                          `json:"errorMessage,omitempty"` //
	Errors       []Errors                        `json:"errors,omitempty"`       //
	Retryable    string                          `json:"retryable,omitempty"`    //
	TraceId      string                          `json:"traceId,omitempty"`      //
}
type EpisPackagePrintAwbResponseData struct {
	Url string `json:"url"` // [Required]
}
type EpisPackageReadyToBeShippedResponse struct {
	BaseResponse                                                 // Common response fields
	Response     EpisPackageReadyToBeShippedResponseData         `json:"data"`                   // Response data
	ErrorCode    string                                          `json:"errorCode,omitempty"`    //
	ErrorMessage string                                          `json:"errorMessage,omitempty"` //
	Errors       []EpisPackageReadyToBeShippedResponseDataErrors `json:"errors,omitempty"`       //
	Retryable    string                                          `json:"retryable,omitempty"`    //
	TraceId      string                                          `json:"traceId,omitempty"`      //
}
type EpisPackageReadyToBeShippedResponseData struct {
	RouteCode                 string                    `json:"routeCode"`                 // [Required]
	MaxEta                    string                    `json:"maxEta"`                    // [Required]
	LastMileShippingProvider  interface{}               `json:"lastMileShippingProvider"`  // [Required]
	PackageCode               string                    `json:"packageCode"`               // [Required]
	Options                   *Options                  `json:"options"`                   // [Required]
	AppliedVas                *AppliedVas               `json:"appliedVas"`                // [Required]
	FirstMileShippingProvider *LastMileShippingProvider `json:"firstMileShippingProvider"` // [Required]
	MinEta                    string                    `json:"minEta"`                    // [Required]
	PortCode                  string                    `json:"portCode"`                  // [Required]
	TrackingNumber            string                    `json:"trackingNumber"`            // [Required]
}
type EpisPackageReadyToBeShippedResponseDataErrors struct {
	Field string `json:"field"` // [Required]
}
type EpisPackageReAttemptResponse struct {
	BaseResponse        // Common response fields
	ErrorCode    string `json:"errorCode,omitempty"`    //
	ErrorMessage string `json:"errorMessage,omitempty"` //
	Retryable    string `json:"retryable,omitempty"`    //
	TraceId      string `json:"traceId,omitempty"`      //
}
type EpisUploadAwbFulfillmentResponse struct {
	BaseResponse          // Common response fields
	ErrorCode    string   `json:"errorCode,omitempty"`    //
	ErrorMessage string   `json:"errorMessage,omitempty"` //
	Errors       []Errors `json:"errors,omitempty"`       //
	Retryable    string   `json:"retryable,omitempty"`    //
	TraceId      string   `json:"traceId,omitempty"`      //
}
type EpisXspaceCreateResponse struct {
	BaseResponse                              // Common response fields
	Response     EpisXspaceCreateResponseData `json:"data"`                   // Response data
	ErrorCode    string                       `json:"errorCode,omitempty"`    //
	ErrorMessage string                       `json:"errorMessage,omitempty"` //
	Retryable    string                       `json:"retryable,omitempty"`    //
	TraceId      string                       `json:"traceId,omitempty"`      //
}
type EpisXspaceCreateResponseData struct {
	CaseId string `json:"caseId"` // [Required]
}
type EpisXspaceGetDetailResponse struct {
	BaseResponse                                 // Common response fields
	Response     EpisXspaceGetDetailResponseData `json:"data"`                   // Response data
	ErrorCode    string                          `json:"errorCode,omitempty"`    //
	ErrorMessage string                          `json:"errorMessage,omitempty"` //
	Retryable    string                          `json:"retryable,omitempty"`    //
	TraceId      string                          `json:"traceId,omitempty"`      //
}
type EpisXspaceGetDetailResponseData struct {
	GmtModified    string        `json:"gmtModified"`    // [Required]
	Attachments    string        `json:"attachments"`    // [Required]
	OrderId        string        `json:"orderId"`        // [Required]
	Subject        string        `json:"subject"`        // [Required]
	ContactName    string        `json:"contactName"`    // [Required]
	BuyerEmail     string        `json:"buyerEmail"`     // [Required]
	SellerName     string        `json:"sellerName"`     // [Required]
	Description    string        `json:"description"`    // [Required]
	BuyerName      string        `json:"buyerName"`      // [Required]
	GmtCreate      string        `json:"gmtCreate"`      // [Required]
	RatingStar     string        `json:"ratingStar"`     // [Required]
	GmtDeleted     string        `json:"gmtDeleted"`     // [Required]
	Mails          []interface{} `json:"mails"`          // [Required]
	RatingRemark   string        `json:"ratingRemark"`   // [Required]
	MerchantId     string        `json:"merchantId"`     // [Required]
	CaseId         string        `json:"caseId"`         // [Required]
	CaseTemplateId string        `json:"caseTemplateId"` // [Required]
	SellerPhoneNo  string        `json:"sellerPhoneNo"`  // [Required]
	Attributes     string        `json:"attributes"`     // [Required]
	Actions        []interface{} `json:"actions"`        // [Required]
	RatingReasons  []string      `json:"ratingReasons"`  // [Required]
	TrackingNumber string        `json:"trackingNumber"` // [Required]
	CategoryId     string        `json:"categoryId"`     // [Required]
	Status         string        `json:"status"`         // [Required]
}
type EpisXspaceQueryResponse struct {
	BaseResponse                             // Common response fields
	Response     EpisXspaceQueryResponseData `json:"data"`                   // Response data
	ErrorCode    string                      `json:"errorCode,omitempty"`    //
	ErrorMessage string                      `json:"errorMessage,omitempty"` //
	Retryable    string                      `json:"retryable,omitempty"`    //
	TraceId      string                      `json:"traceId,omitempty"`      //
}
type EpisXspaceQueryResponseData struct {
	Page    *Page     `json:"page"`    // [Required]
	Content []Content `json:"content"` // [Required]
}
type EpisXspaceRateTicketResponse struct {
	BaseResponse        // Common response fields
	ErrorCode    string `json:"errorCode,omitempty"`    //
	ErrorMessage string `json:"errorMessage,omitempty"` //
	Retryable    string `json:"retryable,omitempty"`    //
	TraceId      string `json:"traceId,omitempty"`      //
}
type Errors struct {
	Field        string `json:"field"`        // [Required]
	ErrorMessage string `json:"errorMessage"` // [Required]
}
type EstimateShippingFeeResponse struct {
	BaseResponse                                 // Common response fields
	Response     EstimateShippingFeeResponseData `json:"data"`                   // Response data
	ErrorCode    string                          `json:"errorCode,omitempty"`    //
	ErrorMessage string                          `json:"errorMessage,omitempty"` //
	Errors       []Errors                        `json:"errors,omitempty"`       //
	Retryable    string                          `json:"retryable,omitempty"`    //
	TraceId      string                          `json:"traceId,omitempty"`      //
}
type EstimateShippingFeeResponseData struct {
	TransactionType string `json:"transactionType"` // [Required]
	Amount          string `json:"amount"`          // [Required]
	Currency        string `json:"currency"`        // [Required]
	TransactionName string `json:"transactionName"` // [Required]
	TaxAmount       string `json:"taxAmount"`       // [Required]
	TransactionId   string `json:"transactionId"`   // [Required]
}
type GetShippingFeeResponse struct {
	BaseResponse                            // Common response fields
	Response     GetShippingFeeResponseData `json:"data"`                   // Response data
	ErrorCode    string                     `json:"errorCode,omitempty"`    //
	ErrorMessage string                     `json:"errorMessage,omitempty"` //
	Errors       []Errors                   `json:"errors,omitempty"`       //
	Retryable    string                     `json:"retryable,omitempty"`    //
	TraceId      string                     `json:"traceId,omitempty"`      //
}
type GetShippingFeeResponseData struct {
	OriginEstimatedShippingFee string `json:"originEstimatedShippingFee"` // [Required]
	ActualShippingFee          string `json:"actualShippingFee"`          // [Required]
	EstimatedShippingFee       string `json:"estimatedShippingFee"`       // [Required]
	Currency                   string `json:"currency"`                   // [Required]
}
type LastMileShippingProvider struct {
	TplSlug string `json:"tplSlug"` // [Required]
	TplName string `json:"tplName"` // [Required]
	TplCode string `json:"tplCode"` // [Required]
}
type Options struct {
	PromotionCode                        string `json:"promotionCode"`                        // [Required]
	VasPartialDeliveryOptionNotAvailable string `json:"vasPartialDeliveryOptionNotAvailable"` // [Required]
}
type Page struct {
	TotalRecords string `json:"totalRecords"` // [Required]
	PageNo       string `json:"pageNo"`       // [Required]
	PageSize     string `json:"pageSize"`     // [Required]
}
type ResponseDataConvertedAddress struct {
	Details string `json:"details"` // [Required]
	Id      int64  `json:"id"`      // [Required]
	Type    string `json:"type"`    // [Required]
}
