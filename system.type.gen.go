package golazada

type StartExportByDatasetResponse struct {
	BaseResponse                                         // Common response fields
	Result       *StartExportByDatasetResponseDataResult `json:"result,omitempty"` //
}
type StartExportByDatasetResponseDataResult struct {
	ReturnCode            string      `json:"returnCode"`            // [Required]
	ReturnValue           interface{} `json:"returnValue"`           // [Required]
	ReturnErrorStackTrace string      `json:"returnErrorStackTrace"` // [Required]
	ReturnMessage         string      `json:"returnMessage"`         // [Required]
}
