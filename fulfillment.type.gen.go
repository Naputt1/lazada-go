package golazada

type ConfirmCollectForDBSResponse struct {
	BaseResponse                                         // Common response fields
	Result       *ConfirmCollectForDBSResponseDataResult `json:"result,omitempty"` //
}
type ConfirmCollectForDBSResponseDataResult struct {
	ErrorMsg  string                                      `json:"error_msg"`  // [Required]
	Data      *ConfirmCollectForDBSResponseDataResultData `json:"data"`       // [Required]
	Success   bool                                        `json:"success"`    // [Required]
	ErrorCode string                                      `json:"error_code"` // [Required]
}
type ConfirmCollectForDBSResponseDataResultData struct {
	Packages []Packages `json:"packages"` // [Required]
}
type ConfirmDeliveryForDBSResponse struct {
	BaseResponse                                         // Common response fields
	Result       *ConfirmCollectForDBSResponseDataResult `json:"result,omitempty"` //
}
type DeliverDigitalResponse struct {
	BaseResponse                                   // Common response fields
	Result       *DeliverDigitalResponseDataResult `json:"result,omitempty"` //
}
type DeliverDigitalResponseDataResult struct {
	Data      *DeliverDigitalResponseDataResultData `json:"data"`      // [Required]
	Success   bool                                  `json:"success"`   // [Required]
	ErrorCode string                                `json:"errorCode"` // [Required]
	ErrorMsg  string                                `json:"errorMsg"`  // [Required]
}
type DeliverDigitalResponseDataResultData struct {
	Orders []Orders `json:"orders"` // [Required]
}
type FailedDeliveryForDBSResponse struct {
	BaseResponse                                         // Common response fields
	Result       *ConfirmCollectForDBSResponseDataResult `json:"result,omitempty"` //
}
type GetDocumentReq struct {
	DocType       string                   `json:"doc_type"`                  // [Required]
	Packages      []GetDocumentReqPackages `json:"packages"`                  // [Required]
	PrintItemList *bool                    `json:"print_item_list,omitempty"` // [Optional]
}
type GetDocumentReqPackages struct {
	PackageId string `json:"package_id"` // [Required]
}
type GetShipmentProviderResponse struct {
	BaseResponse                                        // Common response fields
	Result       *GetShipmentProviderResponseDataResult `json:"result,omitempty"` //
}
type GetShipmentProviderResponseDataResult struct {
	ErrorMsg  string                                     `json:"error_msg"`  // [Required]
	Data      *GetShipmentProviderResponseDataResultData `json:"data"`       // [Required]
	Success   bool                                       `json:"success"`    // [Required]
	ErrorCode string                                     `json:"error_code"` // [Required]
}
type GetShipmentProviderResponseDataResultData struct {
	PlatformDefault      string              `json:"platform_default"`       // [Required]
	ShipmentProviders    []ShipmentProviders `json:"shipment_providers"`     // [Required]
	ShippingAllocateType string              `json:"shipping_allocate_type"` // [Required]
}
type Module struct {
	Result string `json:"result"` // [Required]
}
type OrderItem struct {
	Msg         string `json:"msg"`           // [Required]
	OrderItemId int64  `json:"order_item_id"` // [Required]
	ItemErrCode string `json:"item_err_code"` // [Required]
	Retry       string `json:"retry"`         // [Required]
}
type Orders struct {
	OrderItemList []OrderItem `json:"order_item_list"` // [Required]
	OrderId       int64       `json:"order_id"`        // [Required]
}
type Packages struct {
	Msg         string `json:"msg"`           // [Required]
	ItemErrCode string `json:"item_err_code"` // [Required]
	PackageId   string `json:"package_id"`    // [Required]
	Retry       string `json:"retry"`         // [Required]
}
type PackageStatusUpdateForDBSResponse struct {
	BaseResponse                        // Common response fields
	ErrorCode    *ResponseDataErrorCode `json:"errorCode,omitempty"` //
	Module       *Module                `json:"module,omitempty"`    //
}
type PackOrder struct {
	OrderItemList []PackOrderOrderItem `json:"order_item_list"` // [Required]
	OrderId       int64                `json:"order_id"`        // [Required]
}
type PackOrderOrderItem struct {
	OrderItemId      int64  `json:"order_item_id"`     // [Required]
	Msg              string `json:"msg"`               // [Required]
	ItemErrCode      string `json:"item_err_code"`     // [Required]
	TrackingNumber   string `json:"tracking_number"`   // [Required]
	ShipmentProvider string `json:"shipment_provider"` // [Required]
	PackageId        string `json:"package_id"`        // [Required]
	Retry            string `json:"retry"`             // [Required]
}
type PackRequest struct {
	OrderItemIds string `json:"order_item_ids" url:"order_item_ids"` // [Required]
}
type PackResponse struct {
	BaseResponse                         // Common response fields
	Response     PackResponseData        `json:"data"` // Response data
}
type PackResponseData struct {
	PackOrderList []PackOrder `json:"pack_order_list"` // [Required]
}
type PackResponseDataResult struct {
	ErrorMsg  string                      `json:"error_msg"`  // [Required]
	Data      *PackResponseDataResultData `json:"data"`       // [Required]
	Success   bool                        `json:"success"`    // [Required]
	ErrorCode string                      `json:"error_code"` // [Required]
}
type PackResponseDataResultData struct {
	PackOrderList []PackOrder `json:"pack_order_list"` // [Required]
}
type PrintAWBRequest struct {
	GetDocumentReq *GetDocumentReq `json:"getDocumentReq"` // [Required]
}
type PrintAWBResponse struct {
	BaseResponse                      // Common response fields
	Response     PrintAWBResponseData `json:"result,omitempty"` //
}
type PrintAWBResponseData struct {
	Result *PrintAWBResponseDataResult `json:"result"` // Response data
}
type PrintAWBResponseDataResult struct {
	ErrorMsg  string                          `json:"error_msg"`  // [Required]
	Data      *PrintAWBResponseDataResultData `json:"data"`       // [Required]
	Success   bool                            `json:"success"`    // [Required]
	ErrorCode string                          `json:"error_code"` // [Required]
}
type PrintAWBResponseDataResultData struct {
	File    string `json:"file"`     // [Required]
	PdfUrl  string `json:"pdf_url"`  // [Required]
	DocType string `json:"doc_type"` // [Required]
}
type ReadyToShipRequest struct {
	OrderItemIds string `json:"order_item_ids" url:"order_item_ids"` // [Required]
}
type ReadyToShipResponse struct {
	BaseResponse                                         // Common response fields
	Response     ReadyToShipResponseData `json:"data"` // Response data
}
type ReadyToShipResponseData struct {
	TipContent string `json:"tip_content"` // [Required]
	TipType    string `json:"tip_type"`    // [Required]
}
type RecreatePackageResponse struct {
	BaseResponse                                         // Common response fields
	Result       *ConfirmCollectForDBSResponseDataResult `json:"result,omitempty"` //
}
type ShipmentProviders struct {
	Name         string `json:"name"`          // [Required]
	ProviderCode string `json:"provider_code"` // [Required]
}
