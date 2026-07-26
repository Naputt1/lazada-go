package golazada

type GetMessagesResponse struct {
	BaseResponse                         // Common response fields
	Response     GetMessagesResponseData `json:"data"`                  // Response data
	ErrCode      string                  `json:"err_code,omitempty"`    //
	ErrMessage   string                  `json:"err_message,omitempty"` //
}
type GetMessagesResponseData struct {
	LastMessageId string    `json:"last_message_id"` // [Required]
	MessageList   []Message `json:"message_list"`    // [Required]
	NextStartTime string    `json:"next_start_time"` // [Required]
	HasMore       string    `json:"has_more"`        // [Required]
}
type GetSessionDetailResponse struct {
	BaseResponse                              // Common response fields
	Response     GetSessionDetailResponseData `json:"data"`                  // Response data
	ErrCode      string                       `json:"err_code,omitempty"`    //
	ErrMessage   string                       `json:"err_message,omitempty"` //
}
type GetSessionDetailResponseData struct {
	Summary         string   `json:"summary"`           // [Required]
	UnreadCount     string   `json:"unread_count"`      // [Required]
	LastMessageId   string   `json:"last_message_id"`   // [Required]
	HeadUrl         string   `json:"head_url"`          // [Required]
	SelfPosition    string   `json:"self_position"`     // [Required]
	LastMessageTime string   `json:"last_message_time"` // [Required]
	SiteId          string   `json:"site_id"`           // [Required]
	SessionId       string   `json:"session_id"`        // [Required]
	Title           string   `json:"title"`             // [Required]
	BuyerId         string   `json:"buyer_id"`          // [Required]
	ToPosition      string   `json:"to_position"`       // [Required]
	Tags            []string `json:"tags"`              // [Required]
}
type GetSessionListResponse struct {
	BaseResponse                            // Common response fields
	Response     GetSessionListResponseData `json:"data"`                  // Response data
	ErrCode      string                     `json:"err_code,omitempty"`    //
	ErrMessage   string                     `json:"err_message,omitempty"` //
}
type GetSessionListResponseData struct {
	SessionList   []Session `json:"session_list"`    // [Required]
	NextStartTime string    `json:"next_start_time"` // [Required]
	HasMore       string    `json:"has_more"`        // [Required]
	LastSessionId string    `json:"last_session_id"` // [Required]
}
type Message struct {
	FromAccountType string `json:"from_account_type"` // [Required]
	ProcessMsg      string `json:"process_msg"`       // [Required]
	SessionId       string `json:"session_id"`        // [Required]
	MessageId       string `json:"message_id"`        // [Required]
	Type            string `json:"type"`              // [Required]
	Content         string `json:"content"`           // [Required]
	ToAccountId     string `json:"to_account_id"`     // [Required]
	SendTime        string `json:"send_time"`         // [Required]
	AutoReply       string `json:"auto_reply"`        // [Required]
	ToAccountType   string `json:"to_account_type"`   // [Required]
	SiteId          string `json:"site_id"`           // [Required]
	TemplateId      string `json:"template_id"`       // [Required]
	FromAccountId   string `json:"from_account_id"`   // [Required]
	Status          string `json:"status"`            // [Required]
}
type MessageRecallResponse struct {
	BaseResponse        // Common response fields
	ErrCode      string `json:"err_code,omitempty"`    //
	ErrMessage   string `json:"err_message,omitempty"` //
}
type OpenSessionResponse struct {
	BaseResponse        // Common response fields
	SessionId    string `json:"session_id,omitempty"` //
}
type ReadSessionResponse struct {
	BaseResponse        // Common response fields
	ErrCode      string `json:"err_code,omitempty"`    //
	ErrMessage   string `json:"err_message,omitempty"` //
}
type SendMessageResponse struct {
	BaseResponse                         // Common response fields
	Response     SendMessageResponseData `json:"data"`                  // Response data
	ErrCode      string                  `json:"err_code,omitempty"`    //
	ErrMessage   string                  `json:"err_message,omitempty"` //
}
type SendMessageResponseData struct {
	MessageId   string `json:"message_id"`   // [Required]
	TemplateId  string `json:"template_id"`  // [Required]
	CurrentTime string `json:"current_time"` // [Required]
}
type Session struct {
	Summary         string   `json:"summary"`           // [Required]
	UnreadCount     string   `json:"unread_count"`      // [Required]
	LastMessageId   string   `json:"last_message_id"`   // [Required]
	HeadUrl         string   `json:"head_url"`          // [Required]
	SelfPosition    string   `json:"self_position"`     // [Required]
	SiteId          string   `json:"site_id"`           // [Required]
	LastMessageTime string   `json:"last_message_time"` // [Required]
	SessionId       string   `json:"session_id"`        // [Required]
	BuyerId         string   `json:"buyer_id"`          // [Required]
	Title           string   `json:"title"`             // [Required]
	ToPosition      string   `json:"to_position"`       // [Required]
	Tags            []string `json:"tags"`              // [Required]
}
