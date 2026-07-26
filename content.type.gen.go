package golazada

type CancelTaskResponse struct {
	BaseResponse                               // Common response fields
	Result       *CancelTaskResponseDataResult `json:"result,omitempty"` //
}
type CancelTaskResponseDataResult struct {
	ResultMessage     string `json:"result_message"`      // [Required]
	Success           bool   `json:"success"`             // [Required]
	CanceledTaskCount string `json:"canceled_task_count"` // [Required]
	ResultCode        string `json:"result_code"`         // [Required]
}
type ChangeFaceResponse struct {
	BaseResponse                               // Common response fields
	Result       *ChangeFaceResponseDataResult `json:"result,omitempty"` //
}
type ChangeFaceResponseDataResult struct {
	ResultMessage string `json:"result_message"` // [Required]
	Success       bool   `json:"success"`        // [Required]
	ResultCode    string `json:"result_code"`    // [Required]
	TaskId        string `json:"task_id"`        // [Required]
}
type ChangeProductBackgroundResponse struct {
	BaseResponse                               // Common response fields
	Result       *ChangeFaceResponseDataResult `json:"result,omitempty"` //
}
type FixHandResponse struct {
	BaseResponse                               // Common response fields
	Result       *ChangeFaceResponseDataResult `json:"result,omitempty"` //
}
type GetTaskStatusResponse struct {
	BaseResponse                                  // Common response fields
	Result       *GetTaskStatusResponseDataResult `json:"result,omitempty"` //
}
type GetTaskStatusResponseDataResult struct {
	Data          interface{} `json:"data"`           // [Required]
	ResultMessage string      `json:"result_message"` // [Required]
	Success       bool        `json:"success"`        // [Required]
	ResultCode    string      `json:"result_code"`    // [Required]
	FailMessage   string      `json:"fail_message"`   // [Required]
	Status        string      `json:"status"`         // [Required]
}
type ProductImageMatchResponse struct {
	BaseResponse                                      // Common response fields
	Result       *ProductImageMatchResponseDataResult `json:"result,omitempty"` //
}
type ProductImageMatchResponseDataResult struct {
	ResultMessage  string        `json:"result_message"`   // [Required]
	Success        bool          `json:"success"`          // [Required]
	MatchImageUrls []interface{} `json:"match_image_urls"` // [Required]
	ResultCode     string        `json:"result_code"`      // [Required]
}
type TryOnClothResponse struct {
	BaseResponse                               // Common response fields
	Result       *ChangeFaceResponseDataResult `json:"result,omitempty"` //
}
