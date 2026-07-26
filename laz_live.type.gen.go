package golazada

type HighlightProductResponse struct {
	BaseResponse                              // Common response fields
	Response     HighlightProductResponseData `json:"data"` // Response data
}
type HighlightProductResponseData struct {
	Success bool `json:"success"` // [Required]
}
