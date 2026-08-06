package golazada

import (
	"context"
	"encoding/json"
	"fmt"
)

type ReturnAndRefundService interface {
	// GetReverseOrderDetail Get the detailed information for a specific reverse order
	// Path: /order/reverse/return/detail/list
	GetReverseOrderDetail(ctx context.Context) (*GetReverseOrderDetailResponse, error)
	// GetReverseOrderHistoryList Get the communication history of the reverse order
	// Path: /order/reverse/return/history/list
	GetReverseOrderHistoryList(ctx context.Context) (*GetReverseOrderHistoryListResponse, error)
	// GetReverseOrderReasonList Get the list of reject reason. Need to be used in all refuse refund actions
	// Path: /order/reverse/reason/list
	GetReverseOrderReasonList(ctx context.Context) (*GetReverseOrderReasonListResponse, error)
	// GetReverseOrdersForSeller Use this API to get the list of items for a range of reverse orders.
	// Path: /reverse/getreverseordersforseller
	GetReverseOrdersForSeller(ctx context.Context, opt GetReverseOrdersForSellerRequest) (*GetReverseOrdersForSellerResponse, error)
	// InitReverseOrderCancel Seller initiates a cancelation
	// Path: /order/reverse/cancel/create
	InitReverseOrderCancel(ctx context.Context, req InitReverseOrderCancelRequest) (*InitReverseOrderCancelResponse, error)
	// InitReverseOrderCancelDecide Seller initiates a cancelation
	// Path: /order/reverse/cancel/seller/decide
	InitReverseOrderCancelDecide(ctx context.Context) (*InitReverseOrderCancelDecideResponse, error)
	// ReverseOrderOnlyRefundDecide Seller can use this API to operate only refund requests
	// Path: /order/reverse/onlyrefund/seller/decide
	ReverseOrderOnlyRefundDecide(ctx context.Context, req ReverseOrderOnlyRefundDecideRequest) (*ReverseOrderOnlyRefundDecideResponse, error)
	// ReverseOrderReturnUpdate Seller can use this API to action on return and refund related.
	// Path: /order/reverse/return/update
	ReverseOrderReturnUpdate(ctx context.Context, req ReverseOrderReturnUpdateRequest) (*ReverseOrderReturnUpdateResponse, error)
}

type ReturnAndRefundServiceOp[T any] struct {
	client *Client[T]
}

// GetReverseOrderDetail Get the detailed information for a specific reverse order
// Path: /order/reverse/return/detail/list
func (s *ReturnAndRefundServiceOp[T]) GetReverseOrderDetail(ctx context.Context) (*GetReverseOrderDetailResponse, error) {
	path := "/order/reverse/return/detail/list"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetReverseOrderDetailResponse)
	if string(wrapper.Data) != "null" && len(wrapper.Data) > 0 {
		if err := json.Unmarshal(wrapper.Data, &resp.Response); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}
	}
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetReverseOrderHistoryList Get the communication history of the reverse order
// Path: /order/reverse/return/history/list
func (s *ReturnAndRefundServiceOp[T]) GetReverseOrderHistoryList(ctx context.Context) (*GetReverseOrderHistoryListResponse, error) {
	path := "/order/reverse/return/history/list"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetReverseOrderHistoryListResponse)
	if string(wrapper.Data) != "null" && len(wrapper.Data) > 0 {
		if err := json.Unmarshal(wrapper.Data, &resp.Response); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}
	}
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetReverseOrderReasonList Get the list of reject reason. Need to be used in all refuse refund actions
// Path: /order/reverse/reason/list
func (s *ReturnAndRefundServiceOp[T]) GetReverseOrderReasonList(ctx context.Context) (*GetReverseOrderReasonListResponse, error) {
	path := "/order/reverse/reason/list"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetReverseOrderReasonListResponse)
	if string(wrapper.Data) != "null" && len(wrapper.Data) > 0 {
		if err := json.Unmarshal(wrapper.Data, &resp.Response); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}
	}
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetReverseOrdersForSeller Use this API to get the list of items for a range of reverse orders.
// Path: /reverse/getreverseordersforseller
func (s *ReturnAndRefundServiceOp[T]) GetReverseOrdersForSeller(ctx context.Context, opt GetReverseOrdersForSellerRequest) (*GetReverseOrdersForSellerResponse, error) {
	path := "/reverse/getreverseordersforseller"
	params := paramsFromStruct(opt)
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetReverseOrdersForSellerResponse)
	if string(wrapper.Data) != "null" && len(wrapper.Data) > 0 {
		if err := json.Unmarshal(wrapper.Data, &resp.Response); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}
	}
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// InitReverseOrderCancel Seller initiates a cancelation
// Path: /order/reverse/cancel/create
func (s *ReturnAndRefundServiceOp[T]) InitReverseOrderCancel(ctx context.Context, req InitReverseOrderCancelRequest) (*InitReverseOrderCancelResponse, error) {
	path := "/order/reverse/cancel/create"
	params := paramsFromStruct(req)
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(InitReverseOrderCancelResponse)
	if string(wrapper.Data) != "null" && len(wrapper.Data) > 0 {
		if err := json.Unmarshal(wrapper.Data, &resp.Response); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}
	}
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// InitReverseOrderCancelDecide Seller initiates a cancelation
// Path: /order/reverse/cancel/seller/decide
func (s *ReturnAndRefundServiceOp[T]) InitReverseOrderCancelDecide(ctx context.Context) (*InitReverseOrderCancelDecideResponse, error) {
	path := "/order/reverse/cancel/seller/decide"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(InitReverseOrderCancelDecideResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// ReverseOrderOnlyRefundDecide Seller can use this API to operate only refund requests
// Path: /order/reverse/onlyrefund/seller/decide
func (s *ReturnAndRefundServiceOp[T]) ReverseOrderOnlyRefundDecide(ctx context.Context, req ReverseOrderOnlyRefundDecideRequest) (*ReverseOrderOnlyRefundDecideResponse, error) {
	path := "/order/reverse/onlyrefund/seller/decide"
	params := paramsFromStruct(req)
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(ReverseOrderOnlyRefundDecideResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// ReverseOrderReturnUpdate Seller can use this API to action on return and refund related.
// Path: /order/reverse/return/update
func (s *ReturnAndRefundServiceOp[T]) ReverseOrderReturnUpdate(ctx context.Context, req ReverseOrderReturnUpdateRequest) (*ReverseOrderReturnUpdateResponse, error) {
	path := "/order/reverse/return/update"
	params := paramsFromStruct(req)
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(ReverseOrderReturnUpdateResponse)
	if string(wrapper.Data) != "null" && len(wrapper.Data) > 0 {
		if err := json.Unmarshal(wrapper.Data, &resp.Response); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}
	}
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}
