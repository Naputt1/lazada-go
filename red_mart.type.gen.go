package golazada

type PickupLocations struct {
	Id int64 `json:"id"` // [Required]
}
type ResultDataItems struct {
	ShipmentsInfo     []ShipmentsInfo `json:"shipmentsInfo"`     // [Required]
	QtyFulfilled      int64           `json:"qtyFulfilled"`      // [Required]
	Size              string          `json:"size"`              // [Required]
	Rpc               int64           `json:"rpc"`               // [Required]
	Qty               int64           `json:"qty"`               // [Required]
	ImageUrl          string          `json:"imageUrl"`          // [Required]
	Name              string          `json:"name"`              // [Required]
	Vpc               string          `json:"vpc"`               // [Required]
	MinimumExpiryDate int64           `json:"minimumExpiryDate"` // [Required]
	Sku               string          `json:"sku"`               // [Required]
}
type RssGetOnePickupJobResponse struct {
	BaseResponse                                       // Common response fields
	Result       *RssGetOnePickupJobResponseDataResult `json:"result,omitempty"` //
}
type RssGetOnePickupJobResponseDataResult struct {
	Data         *RssGetOnePickupJobResponseDataResultData `json:"data"`         // [Required]
	Success      bool                                      `json:"success"`      // [Required]
	ErrorMessage string                                    `json:"errorMessage"` // [Required]
}
type RssGetOnePickupJobResponseDataResultData struct {
	PreferredPickupTime    string            `json:"preferredPickupTime"`    // [Required]
	AmendabilityCutOffDate int64             `json:"amendabilityCutOffDate"` // [Required]
	PickedAt               int64             `json:"pickedAt"`               // [Required]
	QtyFulfilledCount      int64             `json:"qtyFulfilledCount"`      // [Required]
	Id                     int64             `json:"id"`                     // [Required]
	Category               string            `json:"category"`               // [Required]
	Items                  []ResultDataItems `json:"items"`                  // [Required]
	ScheduledAt            int64             `json:"scheduledAt"`            // [Required]
	Status                 string            `json:"status"`                 // [Required]
	QtyCount               int64             `json:"qtyCount"`               // [Required]
}
type RssGetPickupJobsResponse struct {
	BaseResponse                                     // Common response fields
	Result       *RssGetPickupJobsResponseDataResult `json:"result,omitempty"` //
}
type RssGetPickupJobsResponseDataResult struct {
	Data         []RssGetOnePickupJobResponseDataResultData `json:"data"`         // [Required]
	Success      bool                                       `json:"success"`      // [Required]
	ErrorMessage string                                     `json:"errorMessage"` // [Required]
}
type RssGetPickupLocationsResponse struct {
	BaseResponse                                          // Common response fields
	Result       *RssGetPickupLocationsResponseDataResult `json:"result,omitempty"` //
}
type RssGetPickupLocationsResponseDataResult struct {
	Total        int64                                         `json:"total"`        // [Required]
	Data         []RssGetPickupLocationsResponseDataResultData `json:"data"`         // [Required]
	Success      bool                                          `json:"success"`      // [Required]
	ErrorMessage string                                        `json:"errorMessage"` // [Required]
	PageSize     string                                        `json:"pageSize"`     // [Required]
	Page         string                                        `json:"page"`         // [Required]
}
type RssGetPickupLocationsResponseDataResultData struct {
	Country      string `json:"country"`      // [Required]
	City         string `json:"city"`         // [Required]
	PostalCode   string `json:"postalCode"`   // [Required]
	Name         string `json:"name"`         // [Required]
	AddressLine1 string `json:"addressLine1"` // [Required]
	AddressLine2 string `json:"addressLine2"` // [Required]
	Id           int64  `json:"id"`           // [Required]
}
type RssGetProductResponse struct {
	BaseResponse                                  // Common response fields
	Result       *RssGetProductResponseDataResult `json:"result,omitempty"` //
}
type RssGetProductResponseDataResult struct {
	Data         *RssGetProductResponseDataResultData `json:"data"`         // [Required]
	Success      bool                                 `json:"success"`      // [Required]
	ErrorMessage string                               `json:"errorMessage"` // [Required]
}
type RssGetProductResponseDataResultData struct {
	ProductCode     string            `json:"productCode"`     // [Required]
	Rpc             int64             `json:"rpc"`             // [Required]
	Title           string            `json:"title"`           // [Required]
	Barcodes        []string          `json:"barcodes"`        // [Required]
	PickupLocations []PickupLocations `json:"pickupLocations"` // [Required]
	Status          string            `json:"status"`          // [Required]
}
type RssGetProductsResponse struct {
	BaseResponse                                   // Common response fields
	Result       *RssGetProductsResponseDataResult `json:"result,omitempty"` //
}
type RssGetProductsResponseDataResult struct {
	Total        int64                                 `json:"total"`        // [Required]
	Data         []RssGetProductResponseDataResultData `json:"data"`         // [Required]
	Success      bool                                  `json:"success"`      // [Required]
	ErrorMessage string                                `json:"errorMessage"` // [Required]
	PageSize     string                                `json:"pageSize"`     // [Required]
	Page         string                                `json:"page"`         // [Required]
}
type RssGetStockLotResponse struct {
	BaseResponse                                   // Common response fields
	Result       *RssGetStockLotResponseDataResult `json:"result,omitempty"` //
}
type RssGetStockLotResponseDataResult struct {
	Data         *RssGetStockLotResponseDataResultData `json:"data"`         // [Required]
	Success      bool                                  `json:"success"`      // [Required]
	ErrorMessage string                                `json:"errorMessage"` // [Required]
}
type RssGetStockLotResponseDataResultData struct {
	QuantityAvailableForSale   int64 `json:"quantityAvailableForSale"`   // [Required]
	QuantityScheduledForPickup int64 `json:"quantityScheduledForPickup"` // [Required]
	Id                         int64 `json:"id"`                         // [Required]
	QuantityAtPickupLocation   int64 `json:"quantityAtPickupLocation"`   // [Required]
}
type RssGetStockLotsResponse struct {
	BaseResponse                                    // Common response fields
	Result       *RssGetStockLotsResponseDataResult `json:"result,omitempty"` //
}
type RssGetStockLotsResponseDataResult struct {
	Data         []RssGetStockLotResponseDataResultData `json:"data"`         // [Required]
	Success      bool                                   `json:"success"`      // [Required]
	ErrorMessage string                                 `json:"errorMessage"` // [Required]
}
type RssUpdateStockLotResponse struct {
	BaseResponse                                   // Common response fields
	Result       *RssGetStockLotResponseDataResult `json:"result,omitempty"` //
}
type ShipmentsInfo struct {
	OrderId string `json:"orderId"` // [Required]
	Qty     int64  `json:"qty"`     // [Required]
}
