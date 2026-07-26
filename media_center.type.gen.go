package golazada

type CompleteCreateVideoResponse struct {
	BaseResponse         // Common response fields
	ResultCode    string `json:"result_code,omitempty"`    //
	ResultMessage string `json:"result_message,omitempty"` //
	VideoId       string `json:"video_id,omitempty"`       //
}
type GetVideoQuotaResponse struct {
	BaseResponse         // Common response fields
	CapacitySize  string `json:"capacity_size,omitempty"`  //
	ResultCode    string `json:"result_code,omitempty"`    //
	ResultMessage string `json:"result_message,omitempty"` //
	UsedSize      string `json:"used_size,omitempty"`      //
}
type GetVideoResponse struct {
	BaseResponse         // Common response fields
	CoverUrl      string `json:"cover_url,omitempty"`      //
	ResultCode    string `json:"result_code,omitempty"`    //
	ResultMessage string `json:"result_message,omitempty"` //
	State         string `json:"state,omitempty"`          //
	Title         string `json:"title,omitempty"`          //
	VideoUrl      string `json:"video_url,omitempty"`      //
}
type InitCreateVideoResponse struct {
	BaseResponse         // Common response fields
	ResultCode    string `json:"result_code,omitempty"`    //
	ResultMessage string `json:"result_message,omitempty"` //
	UploadId      string `json:"upload_id,omitempty"`      //
}
type RemoveVideoResponse struct {
	BaseResponse         // Common response fields
	ResultCode    string `json:"result_code,omitempty"`    //
	ResultMessage string `json:"result_message,omitempty"` //
}
type UploadVideoBlockResponse struct {
	BaseResponse         // Common response fields
	ETag          string `json:"e_tag,omitempty"`          //
	ResultCode    string `json:"result_code,omitempty"`    //
	ResultMessage string `json:"result_message,omitempty"` //
}
