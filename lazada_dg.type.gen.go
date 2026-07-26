package golazada

type DigitalServiceCdkCodeReceivedResponse struct {
	BaseResponse        // Common response fields
	ResultCode   string `json:"result_code,omitempty"` //
	ResultMsg    string `json:"result_msg,omitempty"`  //
}
type InstallServiceCallBack1Response struct {
	BaseResponse         // Common response fields
	ExtendInfo    string `json:"extendInfo,omitempty"`    //
	ResultCode    string `json:"resultCode,omitempty"`    //
	ResultMsg     string `json:"resultMsg,omitempty"`     //
	TransactionId string `json:"transactionId,omitempty"` //
}
type InstallServiceCallBackForTestResponse struct {
	BaseResponse         // Common response fields
	ExtendInfo    string `json:"extendInfo,omitempty"`    //
	ResultCode    string `json:"resultCode,omitempty"`    //
	ResultMsg     string `json:"resultMsg,omitempty"`     //
	TransactionId string `json:"transactionId,omitempty"` //
}
type InstallServiceCallBackResponse struct {
	BaseResponse         // Common response fields
	ExtendInfo    string `json:"extendInfo,omitempty"`    //
	ResultCode    string `json:"resultCode,omitempty"`    //
	ResultMsg     string `json:"resultMsg,omitempty"`     //
	TransactionId string `json:"transactionId,omitempty"` //
}
type InuranceNotication1Response struct {
	BaseResponse         // Common response fields
	ErrorCode     string `json:"errorCode,omitempty"`     //
	ErrorMsg      string `json:"errorMsg,omitempty"`      //
	ExtendInfo    string `json:"extendInfo,omitempty"`    //
	TransactionId string `json:"transactionId,omitempty"` //
}
type InuranceNoticationResponse struct {
	BaseResponse         // Common response fields
	ErrorCode     string `json:"errorCode,omitempty"`     //
	ErrorMsg      string `json:"errorMsg,omitempty"`      //
	ExtendInfo    string `json:"extendInfo,omitempty"`    //
	TransactionId string `json:"transactionId,omitempty"` //
}
type InuranceNotifyLapseResponse struct {
	BaseResponse         // Common response fields
	ErrorCode     string `json:"errorCode,omitempty"`     //
	ErrorMsg      string `json:"errorMsg,omitempty"`      //
	ExtendInfo    string `json:"extendInfo,omitempty"`    //
	TransactionId string `json:"transactionId,omitempty"` //
}
