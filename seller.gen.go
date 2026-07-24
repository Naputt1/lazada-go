package golazada

import (
	"context"
)

type SellerService interface {
	// BatchQueryFollowStatus Query whether these customers follow this seller.
	// Path: /shop/follow/status/batch/query
	BatchQueryFollowStatus(ctx context.Context) (*BatchQueryFollowStatusResponse, error)
	// GetCountryInfo getCountryInfo
	// Path: /seller/cb/country/get
	GetCountryInfo(ctx context.Context) (*GetCountryInfoResponse, error)
	// GetPickUpStoreList return the list of pick up store infomation for requested Seller
	// Path: /rc/store/list/get
	GetPickUpStoreList(ctx context.Context) (*GetPickUpStoreListResponse, error)
	// GetSeller Get seller information by current seller ID.
	// Path: /seller/get
	GetSeller(ctx context.Context) (*GetSellerResponse, error)
	// GetSellerMetricsById Provide seller metrics data of the specific seller, like positive seller rating, ship on time rate and etc.
	// Path: /seller/metrics/get
	GetSellerMetricsById(ctx context.Context) (*GetSellerMetricsByIdResponse, error)
	// GetSellerPerformance Provide the performance metrics of the current seller, such as positive seller rating, ship on time, etc.
	// Path: /seller/performance/get
	GetSellerPerformance(ctx context.Context) (*GetSellerPerformanceResponse, error)
	// GetSellerRegisterInfo getSellerRegisterInfo
	// Path: /seller/cb/register/info
	GetSellerRegisterInfo(ctx context.Context) (*GetSellerRegisterInfoResponse, error)
	// GetSubAddress get location info
	// Path: /seller/cb/country/location/get
	GetSubAddress(ctx context.Context) (*GetSubAddressResponse, error)
	// GetWarehouseBySellerId get warehouse by seller id
	// Path: /rc/warehouse/get
	GetWarehouseBySellerId(ctx context.Context) (*GetWarehouseBySellerIdResponse, error)
	// PaymentBinding paymentBinding
	// Path: /seller/cb/payment/config
	PaymentBinding(ctx context.Context) (*PaymentBindingResponse, error)
	// QueryBuyboxHuntingInfo SPU竞价接口
	// Path: /hunting/buybox/get
	QueryBuyboxHuntingInfo(ctx context.Context) (*QueryBuyboxHuntingInfoResponse, error)
	// QueryWarehouseDetailInfoBySellerId query warehouse detail info by seller id
	// Path: /rc/warehouse/detail/get
	QueryWarehouseDetailInfoBySellerId(ctx context.Context) (*QueryWarehouseDetailInfoBySellerIdResponse, error)
	// SaveSellerWarehouseInfo Api to create or edit the seller warehouse info except the "default"
	// dropshipping warehouse and the return warehouse.
	// Path: /rc/sellerWarehouse/saveWarehouseInfo
	SaveSellerWarehouseInfo(ctx context.Context) (*SaveSellerWarehouseInfoResponse, error)
	// SellerCenterMsgList seller center msg box
	// Path: /sellercenter/msg/list
	SellerCenterMsgList(ctx context.Context) (*SellerCenterMsgListResponse, error)
	// SellerFieldVerify verify seller info field
	// Path: /seller/cb/register/fieldcheck
	SellerFieldVerify(ctx context.Context) (*SellerFieldVerifyResponse, error)
	// SellerPolicyFetch Fetch seller policy information
	// Path: /seller/policy/fetch
	SellerPolicyFetch(ctx context.Context) (*SellerPolicyFetchResponse, error)
	// SynchronizeSellerItemArConfig synchronize seller item ar config
	// Path: /seller/ar/config/syn
	SynchronizeSellerItemArConfig(ctx context.Context) (*SynchronizeSellerItemArConfigResponse, error)
}

type SellerServiceOp[T any] struct {
	client *Client[T]
}

// BatchQueryFollowStatus Query whether these customers follow this seller.
// Path: /shop/follow/status/batch/query
func (s *SellerServiceOp[T]) BatchQueryFollowStatus(ctx context.Context) (*BatchQueryFollowStatusResponse, error) {
	path := "/shop/follow/status/batch/query"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(BatchQueryFollowStatusResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetCountryInfo getCountryInfo
// Path: /seller/cb/country/get
func (s *SellerServiceOp[T]) GetCountryInfo(ctx context.Context) (*GetCountryInfoResponse, error) {
	path := "/seller/cb/country/get"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(GetCountryInfoResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetPickUpStoreList return the list of pick up store infomation for requested Seller
// Path: /rc/store/list/get
func (s *SellerServiceOp[T]) GetPickUpStoreList(ctx context.Context) (*GetPickUpStoreListResponse, error) {
	path := "/rc/store/list/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetPickUpStoreListResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetSeller Get seller information by current seller ID.
// Path: /seller/get
func (s *SellerServiceOp[T]) GetSeller(ctx context.Context) (*GetSellerResponse, error) {
	path := "/seller/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetSellerResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetSellerMetricsById Provide seller metrics data of the specific seller, like positive seller rating, ship on time rate and etc.
// Path: /seller/metrics/get
func (s *SellerServiceOp[T]) GetSellerMetricsById(ctx context.Context) (*GetSellerMetricsByIdResponse, error) {
	path := "/seller/metrics/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetSellerMetricsByIdResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetSellerPerformance Provide the performance metrics of the current seller, such as positive seller rating, ship on time, etc.
// Path: /seller/performance/get
func (s *SellerServiceOp[T]) GetSellerPerformance(ctx context.Context) (*GetSellerPerformanceResponse, error) {
	path := "/seller/performance/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetSellerPerformanceResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetSellerRegisterInfo getSellerRegisterInfo
// Path: /seller/cb/register/info
func (s *SellerServiceOp[T]) GetSellerRegisterInfo(ctx context.Context) (*GetSellerRegisterInfoResponse, error) {
	path := "/seller/cb/register/info"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(GetSellerRegisterInfoResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetSubAddress get location info
// Path: /seller/cb/country/location/get
func (s *SellerServiceOp[T]) GetSubAddress(ctx context.Context) (*GetSubAddressResponse, error) {
	path := "/seller/cb/country/location/get"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(GetSubAddressResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetWarehouseBySellerId get warehouse by seller id
// Path: /rc/warehouse/get
func (s *SellerServiceOp[T]) GetWarehouseBySellerId(ctx context.Context) (*GetWarehouseBySellerIdResponse, error) {
	path := "/rc/warehouse/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetWarehouseBySellerIdResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// PaymentBinding paymentBinding
// Path: /seller/cb/payment/config
func (s *SellerServiceOp[T]) PaymentBinding(ctx context.Context) (*PaymentBindingResponse, error) {
	path := "/seller/cb/payment/config"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(PaymentBindingResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// QueryBuyboxHuntingInfo SPU竞价接口
// Path: /hunting/buybox/get
func (s *SellerServiceOp[T]) QueryBuyboxHuntingInfo(ctx context.Context) (*QueryBuyboxHuntingInfoResponse, error) {
	path := "/hunting/buybox/get"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(QueryBuyboxHuntingInfoResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// QueryWarehouseDetailInfoBySellerId query warehouse detail info by seller id
// Path: /rc/warehouse/detail/get
func (s *SellerServiceOp[T]) QueryWarehouseDetailInfoBySellerId(ctx context.Context) (*QueryWarehouseDetailInfoBySellerIdResponse, error) {
	path := "/rc/warehouse/detail/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(QueryWarehouseDetailInfoBySellerIdResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// SaveSellerWarehouseInfo Api to create or edit the seller warehouse info except the "default"
// dropshipping warehouse and the return warehouse.
// Path: /rc/sellerWarehouse/saveWarehouseInfo
func (s *SellerServiceOp[T]) SaveSellerWarehouseInfo(ctx context.Context) (*SaveSellerWarehouseInfoResponse, error) {
	path := "/rc/sellerWarehouse/saveWarehouseInfo"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(SaveSellerWarehouseInfoResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// SellerCenterMsgList seller center msg box
// Path: /sellercenter/msg/list
func (s *SellerServiceOp[T]) SellerCenterMsgList(ctx context.Context) (*SellerCenterMsgListResponse, error) {
	path := "/sellercenter/msg/list"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(SellerCenterMsgListResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// SellerFieldVerify verify seller info field
// Path: /seller/cb/register/fieldcheck
func (s *SellerServiceOp[T]) SellerFieldVerify(ctx context.Context) (*SellerFieldVerifyResponse, error) {
	path := "/seller/cb/register/fieldcheck"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(SellerFieldVerifyResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// SellerPolicyFetch Fetch seller policy information
// Path: /seller/policy/fetch
func (s *SellerServiceOp[T]) SellerPolicyFetch(ctx context.Context) (*SellerPolicyFetchResponse, error) {
	path := "/seller/policy/fetch"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(SellerPolicyFetchResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// SynchronizeSellerItemArConfig synchronize seller item ar config
// Path: /seller/ar/config/syn
func (s *SellerServiceOp[T]) SynchronizeSellerItemArConfig(ctx context.Context) (*SynchronizeSellerItemArConfigResponse, error) {
	path := "/seller/ar/config/syn"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(SynchronizeSellerItemArConfigResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}
