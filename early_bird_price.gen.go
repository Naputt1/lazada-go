package golazada

import (
	"context"
)

type EarlyBirdPriceService interface {
	// CreateEarlyBirdActivityV2 early bird price activity create
	// Path: /activity/early/bird/create/v2
	CreateEarlyBirdActivityV2(ctx context.Context) (*CreateEarlyBirdActivityV2Response, error)
	// EarlyBirdActivityAddSkusV2 add skus for early bird activity
	// Path: /activity/early/bird/addSkus/v2
	EarlyBirdActivityAddSkusV2(ctx context.Context) (*EarlyBirdActivityAddSkusV2Response, error)
	// EarlyBirdActivityDeactivateSkusV2 deactivate Skus for early bird acivity
	// Path: /activity/early/bird/deactivateSkus/v2
	EarlyBirdActivityDeactivateSkusV2(ctx context.Context) (*EarlyBirdActivityDeactivateSkusV2Response, error)
	// EarlyBirdActivityIsWhitelistSeller is whitelist seller for early bird acivity
	// Path: /activity/early/bird/isWhitelistSeller
	EarlyBirdActivityIsWhitelistSeller(ctx context.Context) (*EarlyBirdActivityIsWhitelistSellerResponse, error)
}

type EarlyBirdPriceServiceOp[T any] struct {
	client *Client[T]
}

// CreateEarlyBirdActivityV2 early bird price activity create
// Path: /activity/early/bird/create/v2
func (s *EarlyBirdPriceServiceOp[T]) CreateEarlyBirdActivityV2(ctx context.Context) (*CreateEarlyBirdActivityV2Response, error) {
	path := "/activity/early/bird/create/v2"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(CreateEarlyBirdActivityV2Response)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// EarlyBirdActivityAddSkusV2 add skus for early bird activity
// Path: /activity/early/bird/addSkus/v2
func (s *EarlyBirdPriceServiceOp[T]) EarlyBirdActivityAddSkusV2(ctx context.Context) (*EarlyBirdActivityAddSkusV2Response, error) {
	path := "/activity/early/bird/addSkus/v2"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(EarlyBirdActivityAddSkusV2Response)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// EarlyBirdActivityDeactivateSkusV2 deactivate Skus for early bird acivity
// Path: /activity/early/bird/deactivateSkus/v2
func (s *EarlyBirdPriceServiceOp[T]) EarlyBirdActivityDeactivateSkusV2(ctx context.Context) (*EarlyBirdActivityDeactivateSkusV2Response, error) {
	path := "/activity/early/bird/deactivateSkus/v2"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(EarlyBirdActivityDeactivateSkusV2Response)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// EarlyBirdActivityIsWhitelistSeller is whitelist seller for early bird acivity
// Path: /activity/early/bird/isWhitelistSeller
func (s *EarlyBirdPriceServiceOp[T]) EarlyBirdActivityIsWhitelistSeller(ctx context.Context) (*EarlyBirdActivityIsWhitelistSellerResponse, error) {
	path := "/activity/early/bird/isWhitelistSeller"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(EarlyBirdActivityIsWhitelistSellerResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}
