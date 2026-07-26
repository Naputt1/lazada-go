package golazada

type CreateEarlyBirdActivityV2Response struct {
	BaseResponse                                              // Common response fields
	Result       *CreateEarlyBirdActivityV2ResponseDataResult `json:"result,omitempty"` //
}
type CreateEarlyBirdActivityV2ResponseDataResult struct {
	Success   bool        `json:"success"`    // [Required]
	Module    interface{} `json:"module"`     // [Required]
	ErrorCode *ErrorCode  `json:"error_code"` // [Required]
	Repeated  string      `json:"repeated"`   // [Required]
	Retry     string      `json:"retry"`      // [Required]
}
type EarlyBirdActivityAddSkusV2Response struct {
	BaseResponse                                               // Common response fields
	Result       *EarlyBirdActivityAddSkusV2ResponseDataResult `json:"result,omitempty"` //
}
type EarlyBirdActivityAddSkusV2ResponseDataResult struct {
	Success   bool       `json:"success"`    // [Required]
	ErrorCode *ErrorCode `json:"error_code"` // [Required]
	Repeated  string     `json:"repeated"`   // [Required]
	Retry     string     `json:"retry"`      // [Required]
}
type EarlyBirdActivityDeactivateSkusV2Response struct {
	BaseResponse                                               // Common response fields
	Result       *EarlyBirdActivityAddSkusV2ResponseDataResult `json:"result,omitempty"` //
}
type EarlyBirdActivityIsWhitelistSellerResponse struct {
	BaseResponse                                              // Common response fields
	Result       *CreateEarlyBirdActivityV2ResponseDataResult `json:"result,omitempty"` //
}
type ErrorCode struct {
	DisplayMessage string `json:"display_message"` // [Required]
	LogMessage     string `json:"log_message"`     // [Required]
	Key            string `json:"key"`             // [Required]
}
