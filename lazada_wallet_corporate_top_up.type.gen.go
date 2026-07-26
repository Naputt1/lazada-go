package golazada

type DirectTransferQueryResponse struct {
	BaseResponse             // Common response fields
	AccountNumber     string `json:"account_number,omitempty"`      //
	Amount            string `json:"amount,omitempty"`              //
	Deposit           string `json:"deposit,omitempty"`             //
	TransferOrderId   string `json:"transfer_order_id,omitempty"`   //
	TransferRequestId string `json:"transfer_request_id,omitempty"` //
}
type DirectTransferRequestResponse struct {
	BaseResponse             // Common response fields
	AccountNumber     string `json:"account_number,omitempty"`      //
	Amount            string `json:"amount,omitempty"`              //
	Deposit           string `json:"deposit,omitempty"`             //
	TransferOrderId   string `json:"transfer_order_id,omitempty"`   //
	TransferRequestId string `json:"transfer_request_id,omitempty"` //
	Withdrawable      string `json:"withdrawable,omitempty"`        //
}
type GiftCodeQueryResponse struct {
	BaseResponse                  // Common response fields
	CreateStatus    string        `json:"create_status,omitempty"`     //
	CurrentPage     string        `json:"current_page,omitempty"`      //
	Deposit         string        `json:"deposit,omitempty"`           //
	PageSize        int64         `json:"page_size,omitempty"`         //
	Records         []interface{} `json:"records,omitempty"`           //
	TotalNumber     string        `json:"total_number,omitempty"`      //
	TotalPage       int64         `json:"total_page,omitempty"`        //
	TransferOrderId string        `json:"transfer_order_id,omitempty"` //
}
type GiftCodeRequestResponse struct {
	BaseResponse           // Common response fields
	CreateStatus    string `json:"create_status,omitempty"`     //
	Deposit         string `json:"deposit,omitempty"`           //
	TotalNumber     string `json:"total_number,omitempty"`      //
	TransferOrderId string `json:"transfer_order_id,omitempty"` //
}
type Reconciliation1Response struct {
	BaseResponse        // Common response fields
	Res          string `json:"res,omitempty"` //
}
