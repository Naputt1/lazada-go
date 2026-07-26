package golazada

type GetStoreCustomPageResponse struct {
	BaseResponse                                // Common response fields
	Response     GetStoreCustomPageResponseData `json:"data"` // Response data
}
type GetStoreCustomPageResponseData struct {
	Result       *GetStoreCustomPageResponseDataResult `json:"result"`        // [Required]
	ErrorMessage string                                `json:"error_message"` // [Required]
	Success      bool                                  `json:"success"`       // [Required]
	Error        string                                `json:"error"`         // [Required]
}
type GetStoreCustomPageResponseDataResult struct {
	PageList []ResultPage    `json:"page_list"` // [Required]
	PageInfo *ResultPageInfo `json:"page_info"` // [Required]
}
type ResultPage struct {
	DecoratePageUrl        string `json:"decorate_page_url"`         // [Required]
	WirelessPagePreviewUrl string `json:"wireless_page_preview_url"` // [Required]
	WirelessEndTime        string `json:"wireless_end_time"`         // [Required]
	TimedPublishTime       string `json:"timed_publish_time"`        // [Required]
	RelatePageId           string `json:"relate_page_id"`            // [Required]
	ClientType             string `json:"client_type"`               // [Required]
	PcEndTime              string `json:"pc_end_time"`               // [Required]
	PcPagePreviewUrl       string `json:"pc_page_preview_url"`       // [Required]
	PageId                 string `json:"page_id"`                   // [Required]
	Path                   string `json:"path"`                      // [Required]
	WirelessPageViewUrl    string `json:"wireless_page_view_url"`    // [Required]
	PageViewUrl            string `json:"page_view_url"`             // [Required]
	LastEditTime           string `json:"last_edit_time"`            // [Required]
	PublishTime            string `json:"publish_time"`              // [Required]
	QrUrl                  string `json:"qr_url"`                    // [Required]
	PageName               string `json:"page_name"`                 // [Required]
	StatusKey              string `json:"status_key"`                // [Required]
}
type ResultPageInfo struct {
	TotalCount  int64  `json:"total_count"`  // [Required]
	CurrentPage string `json:"current_page"` // [Required]
}
