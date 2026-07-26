package golazada

type AddOrUpdatePickupStopResponse struct {
	BaseResponse                      // Common response fields
	ErrorCode    string               `json:"errorCode,omitempty"`    //
	ErrorMessage string               `json:"errorMessage,omitempty"` //
	Errors       []ResponseDataErrors `json:"errors,omitempty"`       //
	Retryable    string               `json:"retryable,omitempty"`    //
}
type Create3PLStationResponse struct {
	BaseResponse                      // Common response fields
	ErrorCode    string               `json:"errorCode,omitempty"`    //
	ErrorMessage string               `json:"errorMessage,omitempty"` //
	Errors       []ResponseDataErrors `json:"errors,omitempty"`       //
	Retryable    string               `json:"retryable,omitempty"`    //
}
type CreateConsolidationServiceResponse struct {
	BaseResponse        // Common response fields
	Response     string `json:"data"`                // Response data
	ErrorCode    string `json:"errorCode,omitempty"` //
	ErrorMsg     string `json:"errorMsg,omitempty"`  //
}
type Error struct {
	ErrorCode string `json:"errorCode"` // [Required]
}
type GetOrderTraceResponse struct {
	BaseResponse                                  // Common response fields
	Result       *GetOrderTraceResponseDataResult `json:"result,omitempty"` //
}
type GetOrderTraceResponseDataResult struct {
	NotSuccess string                 `json:"not_success"` // [Required]
	Success    bool                   `json:"success"`     // [Required]
	Module     []ResultModule         `json:"module"`      // [Required]
	ErrorCode  *ResponseDataErrorCode `json:"error_code"`  // [Required]
	Repeated   string                 `json:"repeated"`    // [Required]
	Retry      string                 `json:"retry"`       // [Required]
}
type LogisticDetailInfo struct {
	PackageLocationName string        `json:"package_location_name"` // [Required]
	StatusCode          string        `json:"status_code"`           // [Required]
	ProofImages         []interface{} `json:"proof_images"`          // [Required]
	DetailType          string        `json:"detail_type"`           // [Required]
	EventDate           string        `json:"event_date"`            // [Required]
	ReceiveTime         string        `json:"receive_time"`          // [Required]
	Icon                string        `json:"icon"`                  // [Required]
	Description         string        `json:"description"`           // [Required]
	Title               string        `json:"title"`                 // [Required]
	EventTime           string        `json:"event_time"`            // [Required]
}
type PackageDetailInfo struct {
	OrderLineInfoList      string               `json:"order_line_info_list"`      // [Required]
	OfcPackageId           string               `json:"ofc_package_id"`            // [Required]
	TrackingNumber         string               `json:"tracking_number"`           // [Required]
	LogisticDetailInfoList []LogisticDetailInfo `json:"logistic_detail_info_list"` // [Required]
}
type ResponseDataErrorCode struct {
	DisplayMessage string `json:"displayMessage"` // [Required]
}
type ResponseDataErrors struct {
	Field        string `json:"field"`        // [Required]
	ErrorMessage string `json:"errorMessage"` // [Required]
	ErrorCode    string `json:"errorCode"`    // [Required]
}
type ResultModule struct {
	WarehouseDetailInfo   string              `json:"warehouse_detail_info"`    // [Required]
	OfcOrderId            string              `json:"ofc_order_id"`             // [Required]
	PackageDetailInfoList []PackageDetailInfo `json:"package_detail_info_list"` // [Required]
}
type ScanParcelResponse struct {
	BaseResponse          // Common response fields
	TrackingNumber string `json:"trackingNumber,omitempty"` //
}
type StationDopScanResponse struct {
	BaseResponse                            // Common response fields
	Response     StationDopScanResponseData `json:"data"`            // Response data
	Error        *Error                     `json:"error,omitempty"` //
}
type StationDopScanResponseData struct {
	TrackingNumber string `json:"trackingNumber"` // [Required]
}
type Update3PLStationResponse struct {
	BaseResponse                      // Common response fields
	ErrorCode    string               `json:"errorCode,omitempty"`    //
	ErrorMessage string               `json:"errorMessage,omitempty"` //
	Errors       []ResponseDataErrors `json:"errors,omitempty"`       //
	Retryable    string               `json:"retryable,omitempty"`    //
}
type UpdateLastMileResponse struct {
	BaseResponse        // Common response fields
	Response     string `json:"data"`                // Response data
	ErrorCode    string `json:"errorCode,omitempty"` //
	ErrorMsg     string `json:"errorMsg,omitempty"`  //
}
type UpdatePickupTimeSlotResponse struct {
	BaseResponse                      // Common response fields
	ErrorCode    string               `json:"errorCode,omitempty"`    //
	ErrorMessage string               `json:"errorMessage,omitempty"` //
	Errors       []ResponseDataErrors `json:"errors,omitempty"`       //
	Retryable    string               `json:"retryable,omitempty"`    //
}
