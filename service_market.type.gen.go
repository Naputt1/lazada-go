package golazada

type ArticleBizOrders struct {
	OrderCycleStart string `json:"orderCycleStart"` // [Required]
	RefundFee       string `json:"refundFee"`       // [Required]
	ArticleItemName string `json:"articleItemName"` // [Required]
	BizType         string `json:"bizType"`         // [Required]
	ArticleName     string `json:"articleName"`     // [Required]
	TotalPayFee     string `json:"totalPayFee"`     // [Required]
	OrderId         string `json:"orderId"`         // [Required]
	OrderCycleEnd   string `json:"orderCycleEnd"`   // [Required]
	ItemCode        string `json:"itemCode"`        // [Required]
	Fee             string `json:"fee"`             // [Required]
	UserId          string `json:"userId"`          // [Required]
	Nick            string `json:"nick"`            // [Required]
	ActivityCode    string `json:"activityCode"`    // [Required]
	ItemName        string `json:"itemName"`        // [Required]
	OrderCycle      string `json:"orderCycle"`      // [Required]
	BizOrderId      string `json:"bizOrderId"`      // [Required]
	PromFee         string `json:"promFee"`         // [Required]
	Create          string `json:"create"`          // [Required]
	ArticleCode     string `json:"articleCode"`     // [Required]
}
type ServiceMarketAppKeyOrderQueryResponse struct {
	BaseResponse                                                  // Common response fields
	Result       *ServiceMarketAppKeyOrderQueryResponseDataResult `json:"result,omitempty"` //
}
type ServiceMarketAppKeyOrderQueryResponseDataResult struct {
	Data       *ServiceMarketAppKeyOrderQueryResponseDataResultData `json:"data"`       // [Required]
	Success    bool                                                 `json:"success"`    // [Required]
	ResultCode string                                               `json:"resultCode"` // [Required]
	Remark     string                                               `json:"remark"`     // [Required]
}
type ServiceMarketAppKeyOrderQueryResponseDataResultData struct {
	TotalItem        string             `json:"totalItem"`        // [Required]
	ArticleBizOrders []ArticleBizOrders `json:"articleBizOrders"` // [Required]
}
type ServiceMarketAppKeySubQueryResponse struct {
	BaseResponse                                                // Common response fields
	Result       *ServiceMarketAppKeySubQueryResponseDataResult `json:"result,omitempty"` //
}
type ServiceMarketAppKeySubQueryResponseDataResult struct {
	Data    []ServiceMarketAppKeySubQueryResponseDataResultData `json:"data"`    // [Required]
	Success bool                                                `json:"success"` // [Required]
}
type ServiceMarketAppKeySubQueryResponseDataResultData struct {
	Nick         string `json:"nick"`          // [Required]
	ItemCode     string `json:"item_code"`     // [Required]
	ExpireNotice string `json:"expire_notice"` // [Required]
	EndTime      string `json:"end_time"`      // [Required]
	ArticleName  string `json:"article_name"`  // [Required]
	ItemName     string `json:"item_name"`     // [Required]
	Autosub      string `json:"autosub"`       // [Required]
	ArticleCode  string `json:"article_code"`  // [Required]
	Status       string `json:"status"`        // [Required]
}
