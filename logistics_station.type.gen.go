package golazada

type CageValidationResponse struct {
	BaseResponse        // Common response fields
	Response     string `json:"data"`                // Response data
	ErrorCode    string `json:"errorCode,omitempty"` //
	ErrorMsg     string `json:"errorMsg,omitempty"`  //
	TraceId      string `json:"traceId,omitempty"`   //
}
type ConfirmInboundResponse struct {
	BaseResponse        // Common response fields
	Response     string `json:"data"`                // Response data
	ErrorCode    string `json:"errorCode,omitempty"` //
	ErrorMsg     string `json:"errorMsg,omitempty"`  //
	TraceId      string `json:"traceId,omitempty"`   //
}
type ConfirmParcelCollectionResponse struct {
	BaseResponse        // Common response fields
	Response     string `json:"data"`                // Response data
	ErrorCode    string `json:"errorCode,omitempty"` //
	ErrorMsg     string `json:"errorMsg,omitempty"`  //
	TraceId      string `json:"traceId,omitempty"`   //
}
type CreateScannedParcelResponse struct {
	BaseResponse                                 // Common response fields
	Response     CreateScannedParcelResponseData `json:"data"`                // Response data
	ErrorCode    string                          `json:"errorCode,omitempty"` //
	ErrorMsg     string                          `json:"errorMsg,omitempty"`  //
	TraceId      string                          `json:"traceId,omitempty"`   //
}
type CreateScannedParcelResponseData struct {
	ServiceType    string `json:"serviceType"`    // [Required]
	CreatedAt      string `json:"createdAt"`      // [Required]
	CageNumber     string `json:"cageNumber"`     // [Required]
	SellerName     string `json:"sellerName"`     // [Required]
	WarningMessage string `json:"warningMessage"` // [Required]
	TrackingNumber string `json:"trackingNumber"` // [Required]
	PickupTplSlug  string `json:"pickupTplSlug"`  // [Required]
	LastmileTpl    string `json:"lastmileTpl"`    // [Required]
}
type DeleteScannedParcelResponse struct {
	BaseResponse        // Common response fields
	Response     string `json:"data"`                // Response data
	ErrorCode    string `json:"errorCode,omitempty"` //
	ErrorMsg     string `json:"errorMsg,omitempty"`  //
	TraceId      string `json:"traceId,omitempty"`   //
}
type DopConfirmInboundResponse struct {
	BaseResponse        // Common response fields
	Response     string `json:"data"`                // Response data
	ErrorCode    string `json:"errorCode,omitempty"` //
	ErrorMsg     string `json:"errorMsg,omitempty"`  //
	TraceId      string `json:"traceId,omitempty"`   //
}
type DopCreateScannedParcelResponse struct {
	BaseResponse                                    // Common response fields
	Response     DopCreateScannedParcelResponseData `json:"data"`                // Response data
	ErrorCode    string                             `json:"errorCode,omitempty"` //
	ErrorMsg     string                             `json:"errorMsg,omitempty"`  //
	TraceId      string                             `json:"traceId,omitempty"`   //
}
type DopCreateScannedParcelResponseData struct {
	StationCode    string `json:"stationCode"`    // [Required]
	CreatedAt      string `json:"createdAt"`      // [Required]
	CageNumber     string `json:"cageNumber"`     // [Required]
	SellerName     string `json:"sellerName"`     // [Required]
	TrackingNumber string `json:"trackingNumber"` // [Required]
	PickupTplSlug  string `json:"pickupTplSlug"`  // [Required]
}
type DopDeleteScannedParcelResponse struct {
	BaseResponse        // Common response fields
	Response     string `json:"data"`                // Response data
	ErrorCode    string `json:"errorCode,omitempty"` //
	ErrorMsg     string `json:"errorMsg,omitempty"`  //
	TraceId      string `json:"traceId,omitempty"`   //
}
type DopGetInboundedParcelResponse struct {
	BaseResponse                                   // Common response fields
	Response     DopGetInboundedParcelResponseData `json:"data"`                // Response data
	ErrorCode    string                            `json:"errorCode,omitempty"` //
	ErrorMsg     string                            `json:"errorMsg,omitempty"`  //
	TraceId      string                            `json:"traceId,omitempty"`   //
}
type DopGetInboundedParcelResponseData struct {
	CageNumber     string `json:"cageNumber"`     // [Required]
	InboundedAt    string `json:"inboundedAt"`    // [Required]
	OutboundedAt   string `json:"outboundedAt"`   // [Required]
	LostAt         string `json:"lostAt"`         // [Required]
	TrackingNumber string `json:"trackingNumber"` // [Required]
	Status         string `json:"status"`         // [Required]
	PickupTplSlug  string `json:"pickupTplSlug"`  // [Required]
}
type DopGetScannedParcelResponse struct {
	BaseResponse                                 // Common response fields
	Response     DopGetScannedParcelResponseData `json:"data"`                // Response data
	ErrorCode    string                          `json:"errorCode,omitempty"` //
	ErrorMsg     string                          `json:"errorMsg,omitempty"`  //
	TraceId      string                          `json:"traceId,omitempty"`   //
}
type DopGetScannedParcelResponseData struct {
	StationCode    string `json:"stationCode"`    // [Required]
	CreatedAt      string `json:"createdAt"`      // [Required]
	CageNumber     string `json:"cageNumber"`     // [Required]
	SellerName     string `json:"sellerName"`     // [Required]
	TrackingNumber string `json:"trackingNumber"` // [Required]
	PickupTplSlug  string `json:"pickupTplSlug"`  // [Required]
}
type GetCpScheduledPuParcelResponse struct {
	BaseResponse                                    // Common response fields
	Response     GetCpScheduledPuParcelResponseData `json:"data"`                // Response data
	ErrorCode    string                             `json:"errorCode,omitempty"` //
	ErrorMsg     string                             `json:"errorMsg,omitempty"`  //
	TraceId      string                             `json:"traceId,omitempty"`   //
}
type GetCpScheduledPuParcelResponseData struct {
	DispatchedAt   string `json:"dispatchedAt"`   // [Required]
	TrackingNumber string `json:"trackingNumber"` // [Required]
}
type GetInboundedParcelResponse struct {
	BaseResponse                                // Common response fields
	Response     GetInboundedParcelResponseData `json:"data"`                // Response data
	ErrorCode    string                         `json:"errorCode,omitempty"` //
	ErrorMsg     string                         `json:"errorMsg,omitempty"`  //
	TraceId      string                         `json:"traceId,omitempty"`   //
}
type GetInboundedParcelResponseData struct {
	ServiceType    string `json:"serviceType"`    // [Required]
	CageNumber     string `json:"cageNumber"`     // [Required]
	InboundedAt    string `json:"inboundedAt"`    // [Required]
	OutboundedAt   string `json:"outboundedAt"`   // [Required]
	WarningMessage string `json:"warningMessage"` // [Required]
	LostAt         string `json:"lostAt"`         // [Required]
	TrackingNumber string `json:"trackingNumber"` // [Required]
	PickupTplSlug  string `json:"pickupTplSlug"`  // [Required]
	LastmileTpl    string `json:"lastmileTpl"`    // [Required]
	Status         string `json:"status"`         // [Required]
}
type GetListAccessStationResponse struct {
	BaseResponse                                  // Common response fields
	Response     GetListAccessStationResponseData `json:"data"`                // Response data
	ErrorCode    string                           `json:"errorCode,omitempty"` //
	ErrorMsg     string                           `json:"errorMsg,omitempty"`  //
	TraceId      string                           `json:"traceId,omitempty"`   //
}
type GetListAccessStationResponseData struct {
	StationCode string `json:"stationCode"` // [Required]
	Active      string `json:"active"`      // [Required]
	StationName string `json:"stationName"` // [Required]
}
type GetMetaDataResponse struct {
	BaseResponse                         // Common response fields
	Response     GetMetaDataResponseData `json:"data"`                // Response data
	ErrorCode    string                  `json:"errorCode,omitempty"` //
	ErrorMsg     string                  `json:"errorMsg,omitempty"`  //
	TraceId      string                  `json:"traceId,omitempty"`   //
}
type GetMetaDataResponseData struct {
	RejectReasons []RejectReasons `json:"rejectReasons"` // [Required]
}
type GetScannedParcelResponse struct {
	BaseResponse                              // Common response fields
	Response     GetScannedParcelResponseData `json:"data"`                // Response data
	ErrorCode    string                       `json:"errorCode,omitempty"` //
	ErrorMsg     string                       `json:"errorMsg,omitempty"`  //
	TraceId      string                       `json:"traceId,omitempty"`   //
}
type GetScannedParcelResponseData struct {
	ServiceType    string `json:"serviceType"`    // [Required]
	CreatedAt      string `json:"createdAt"`      // [Required]
	CageNumber     string `json:"cageNumber"`     // [Required]
	SellerName     string `json:"sellerName"`     // [Required]
	WarningMessage string `json:"warningMessage"` // [Required]
	TrackingNumber string `json:"trackingNumber"` // [Required]
	PickupTplSlug  string `json:"pickupTplSlug"`  // [Required]
	LastmileTpl    string `json:"lastmileTpl"`    // [Required]
}
type RejectReasons struct {
	RejectCode string `json:"rejectCode"` // [Required]
	Text       string `json:"text"`       // [Required]
}
type SearchCustomerReturnParcelResponse struct {
	BaseResponse                                        // Common response fields
	Response     SearchCustomerReturnParcelResponseData `json:"data"`                // Response data
	ErrorCode    string                                 `json:"errorCode,omitempty"` //
	ErrorMsg     string                                 `json:"errorMsg,omitempty"`  //
	TraceId      string                                 `json:"traceId,omitempty"`   //
}
type SearchCustomerReturnParcelResponseData struct {
	MaskedCustomerName string `json:"maskedCustomerName"` // [Required]
	TrackingNumber     string `json:"trackingNumber"`     // [Required]
}
type ValidateCageResponse struct {
	BaseResponse        // Common response fields
	Response     string `json:"data"`                // Response data
	ErrorCode    string `json:"errorCode,omitempty"` //
	ErrorMsg     string `json:"errorMsg,omitempty"`  //
	TraceId      string `json:"traceId,omitempty"`   //
}
type ValidateOTPResponse struct {
	BaseResponse        // Common response fields
	Response     string `json:"data"`                // Response data
	ErrorCode    string `json:"errorCode,omitempty"` //
	ErrorMsg     string `json:"errorMsg,omitempty"`  //
	TraceId      string `json:"traceId,omitempty"`   //
}
