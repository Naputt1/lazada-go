package golazada

type GetChannelcodeByFirstMileNoResponse struct {
	BaseResponse                                                // Common response fields
	Result       *GetChannelcodeByFirstMileNoResponseDataResult `json:"result,omitempty"` //
}
type GetChannelcodeByFirstMileNoResponseDataResult struct {
	Success   bool          `json:"success"`   // [Required]
	Module    []interface{} `json:"module"`    // [Required]
	ErrorCode string        `json:"errorCode"` // [Required]
	ErrorMsg  string        `json:"errorMsg"`  // [Required]
}
type GetLazadaBigbagPDFLableResponse struct {
	BaseResponse                                            // Common response fields
	Result       *GetLazadaBigbagPDFLableResponseDataResult `json:"result,omitempty"` //
}
type GetLazadaBigbagPDFLableResponseDataResult struct {
	Data      interface{} `json:"data"`      // [Required]
	Success   bool        `json:"success"`   // [Required]
	ErrorCode string      `json:"errorCode"` // [Required]
	ErrorMsg  string      `json:"errorMsg"`  // [Required]
}
type LazadaBigbagCancelResponse struct {
	BaseResponse                                       // Common response fields
	Result       *LazadaBigbagCancelResponseDataResult `json:"result,omitempty"` //
}
type LazadaBigbagCancelResponseDataResult struct {
	ErrorMsg  string      `json:"error_msg"`  // [Required]
	Data      interface{} `json:"data"`       // [Required]
	Success   bool        `json:"success"`    // [Required]
	ErrorCode string      `json:"error_code"` // [Required]
}
type LazadaBigbagCollectionPointsResponse struct {
	BaseResponse                                                 // Common response fields
	Result       *LazadaBigbagCollectionPointsResponseDataResult `json:"result,omitempty"` //
}
type LazadaBigbagCollectionPointsResponseDataResult struct {
	ErroMsg   string                                              `json:"erroMsg"`   // [Required]
	Data      *LazadaBigbagCollectionPointsResponseDataResultData `json:"data"`      // [Required]
	Success   bool                                                `json:"success"`   // [Required]
	ErrorCode string                                              `json:"errorCode"` // [Required]
}
type LazadaBigbagCollectionPointsResponseDataResultData struct {
	PageSize         string        `json:"pageSize"`         // [Required]
	ItemList         []interface{} `json:"itemList"`         // [Required]
	TotalCount       string        `json:"totalCount"`       // [Required]
	CurrentPageIndex string        `json:"currentPageIndex"` // [Required]
	PageTotalNum     string        `json:"pageTotalNum"`     // [Required]
}
type LazadaBigbagCommitResponse struct {
	BaseResponse                                       // Common response fields
	Result       *LazadaBigbagCommitResponseDataResult `json:"result,omitempty"` //
}
type LazadaBigbagCommitResponseDataResult struct {
	Data      *LazadaBigbagCommitResponseDataResultData `json:"data"`      // [Required]
	Success   bool                                      `json:"success"`   // [Required]
	ErrorCode string                                    `json:"errorCode"` // [Required]
	ErrorMsg  string                                    `json:"errorMsg"`  // [Required]
}
type LazadaBigbagCommitResponseDataResultData struct {
	HandoverContentId   string `json:"handoverContentId"`   // [Required]
	HandoverContentCode string `json:"handoverContentCode"` // [Required]
	HandoverOrderId     string `json:"handoverOrderId"`     // [Required]
}
type LazadaBigbagUpdateResponse struct {
	BaseResponse                                       // Common response fields
	Result       *LazadaBigbagUpdateResponseDataResult `json:"result,omitempty"` //
}
type LazadaBigbagUpdateResponseDataResult struct {
	ErroMsg   string      `json:"erroMsg"`   // [Required]
	Data      interface{} `json:"data"`      // [Required]
	Success   bool        `json:"success"`   // [Required]
	ErrorCode string      `json:"errorCode"` // [Required]
}
type LazadaSellerAccountBindResponse struct {
	BaseResponse                                            // Common response fields
	Result       *GetLazadaBigbagPDFLableResponseDataResult `json:"result,omitempty"` //
}
type QueryAddressInformaitonResponse struct {
	BaseResponse                                            // Common response fields
	Result       *QueryAddressInformaitonResponseDataResult `json:"result,omitempty"` //
}
type QueryAddressInformaitonResponseDataResult struct {
	Data      *QueryAddressInformaitonResponseDataResultData `json:"data"`      // [Required]
	Success   bool                                           `json:"success"`   // [Required]
	ErrorCode string                                         `json:"errorCode"` // [Required]
	ErrorMsg  string                                         `json:"errorMsg"`  // [Required]
}
type QueryAddressInformaitonResponseDataResultData struct {
	MatchDetailAddress string `json:"matchDetailAddress"` // [Required]
	AddressId          string `json:"addressId"`          // [Required]
}
type QueryLazadaBigbagInfoResponse struct {
	BaseResponse                                            // Common response fields
	Result       *GetLazadaBigbagPDFLableResponseDataResult `json:"result,omitempty"` //
}
