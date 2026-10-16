package golazada

import (
	"context"
	"encoding/json"
	"fmt"
)

type StoreFlashSaleService interface {
	// ActivateStoreFlashSale activate Store Flash Sale
	// Path: /promotion/storeflashsale/activate
	ActivateStoreFlashSale(ctx context.Context) (*ActivateStoreFlashSaleResponse, error)
	// CreateStoreFlashSale create a new Store Flash Sale
	// Path: /promotion/storeflashsale/create
	CreateStoreFlashSale(ctx context.Context) (*CreateStoreFlashSaleResponse, error)
	// DeactivateStoreFlashSale deactivate Store Flash Sale
	// Path: /promotion/storeflashsale/deactivate
	DeactivateStoreFlashSale(ctx context.Context) (*DeactivateStoreFlashSaleResponse, error)
	// GetStoreFlashSale getStoreFlashSale
	// Path: /promotion/storeflashsale/get
	GetStoreFlashSale(ctx context.Context) (*GetStoreFlashSaleResponse, error)
	// ListStoreFlashSale list Store Flash Sale
	// Path: /promotion/storeflashsale/list
	ListStoreFlashSale(ctx context.Context) (*ListStoreFlashSaleResponse, error)
	// UpdateStoreFlashSale update storeflashsale
	// Path: /promotion/storeflashsale/update
	UpdateStoreFlashSale(ctx context.Context) (*UpdateStoreFlashSaleResponse, error)
}

type StoreFlashSaleServiceOp[T any] struct {
	client *Client[T]
}

// ActivateStoreFlashSale activate Store Flash Sale
// Path: /promotion/storeflashsale/activate
func (s *StoreFlashSaleServiceOp[T]) ActivateStoreFlashSale(ctx context.Context) (*ActivateStoreFlashSaleResponse, error) {
	path := "/promotion/storeflashsale/activate"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(ActivateStoreFlashSaleResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// CreateStoreFlashSale create a new Store Flash Sale
// Path: /promotion/storeflashsale/create
func (s *StoreFlashSaleServiceOp[T]) CreateStoreFlashSale(ctx context.Context) (*CreateStoreFlashSaleResponse, error) {
	path := "/promotion/storeflashsale/create"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(CreateStoreFlashSaleResponse)
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

// DeactivateStoreFlashSale deactivate Store Flash Sale
// Path: /promotion/storeflashsale/deactivate
func (s *StoreFlashSaleServiceOp[T]) DeactivateStoreFlashSale(ctx context.Context) (*DeactivateStoreFlashSaleResponse, error) {
	path := "/promotion/storeflashsale/deactivate"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(DeactivateStoreFlashSaleResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetStoreFlashSale getStoreFlashSale
// Path: /promotion/storeflashsale/get
func (s *StoreFlashSaleServiceOp[T]) GetStoreFlashSale(ctx context.Context) (*GetStoreFlashSaleResponse, error) {
	path := "/promotion/storeflashsale/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetStoreFlashSaleResponse)
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

// ListStoreFlashSale list Store Flash Sale
// Path: /promotion/storeflashsale/list
func (s *StoreFlashSaleServiceOp[T]) ListStoreFlashSale(ctx context.Context) (*ListStoreFlashSaleResponse, error) {
	path := "/promotion/storeflashsale/list"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(ListStoreFlashSaleResponse)
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

// UpdateStoreFlashSale update storeflashsale
// Path: /promotion/storeflashsale/update
func (s *StoreFlashSaleServiceOp[T]) UpdateStoreFlashSale(ctx context.Context) (*UpdateStoreFlashSaleResponse, error) {
	path := "/promotion/storeflashsale/update"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(UpdateStoreFlashSaleResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}
