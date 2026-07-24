package golazada

import (
	"context"
)

type ServiceMarketService interface {
	// ServiceMarketAppKeyOrderQuery Query user order list for specific App on Service Market
	// Path: /service/market/order/query
	ServiceMarketAppKeyOrderQuery(ctx context.Context) (*ServiceMarketAppKeyOrderQueryResponse, error)
	// ServiceMarketAppKeySubQuery Query user subscription info for specific App on Service Market
	// Path: /service/market/subs/query
	ServiceMarketAppKeySubQuery(ctx context.Context) (*ServiceMarketAppKeySubQueryResponse, error)
}

type ServiceMarketServiceOp[T any] struct {
	client *Client[T]
}

// ServiceMarketAppKeyOrderQuery Query user order list for specific App on Service Market
// Path: /service/market/order/query
func (s *ServiceMarketServiceOp[T]) ServiceMarketAppKeyOrderQuery(ctx context.Context) (*ServiceMarketAppKeyOrderQueryResponse, error) {
	path := "/service/market/order/query"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(ServiceMarketAppKeyOrderQueryResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// ServiceMarketAppKeySubQuery Query user subscription info for specific App on Service Market
// Path: /service/market/subs/query
func (s *ServiceMarketServiceOp[T]) ServiceMarketAppKeySubQuery(ctx context.Context) (*ServiceMarketAppKeySubQueryResponse, error) {
	path := "/service/market/subs/query"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(ServiceMarketAppKeySubQueryResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}
