package golazada

import (
	"context"
	"encoding/json"
	"fmt"
)

type SellerVoucherService interface {
	// SellerVoucheDeleteSelectedProductSKU delete seller voucher promotion product sku
	// Path: /promotion/voucher/product/sku/remove
	SellerVoucheDeleteSelectedProductSKU(ctx context.Context) (*SellerVoucheDeleteSelectedProductSKUResponse, error)
	// SellerVoucherActivate activate seller voucher promotion
	// Path: /promotion/voucher/activate
	SellerVoucherActivate(ctx context.Context) (*SellerVoucherActivateResponse, error)
	// SellerVoucherAddSelectedProductSKU add seller voucher promotion product sku
	// Path: /promotion/voucher/product/sku/add
	SellerVoucherAddSelectedProductSKU(ctx context.Context) (*SellerVoucherAddSelectedProductSKUResponse, error)
	// SellerVoucherCreate create a new seller voucher promotion
	// Path: /promotion/voucher/create
	SellerVoucherCreate(ctx context.Context) (*SellerVoucherCreateResponse, error)
	// SellerVoucherDeactivate deactivate seller voucher promotion
	// Path: /promotion/voucher/deactivate
	SellerVoucherDeactivate(ctx context.Context) (*SellerVoucherDeactivateResponse, error)
	// SellerVoucherDetailQuery get a seller voucher promotion detail
	// Path: /promotion/voucher/get
	SellerVoucherDetailQuery(ctx context.Context) (*SellerVoucherDetailQueryResponse, error)
	// SellerVoucherList query seller voucher promotion list
	// Path: /promotion/vouchers/get
	SellerVoucherList(ctx context.Context) (*SellerVoucherListResponse, error)
	// SellerVoucherSelectedProductList query seller voucher selected products list
	// Path: /promotion/voucher/products/get
	SellerVoucherSelectedProductList(ctx context.Context) (*SellerVoucherSelectedProductListResponse, error)
	// SellerVoucherUpdate update a existing seller voucher promotion
	// Path: /promotion/voucher/update
	SellerVoucherUpdate(ctx context.Context) (*SellerVoucherUpdateResponse, error)
}

type SellerVoucherServiceOp[T any] struct {
	client *Client[T]
}

// SellerVoucheDeleteSelectedProductSKU delete seller voucher promotion product sku
// Path: /promotion/voucher/product/sku/remove
func (s *SellerVoucherServiceOp[T]) SellerVoucheDeleteSelectedProductSKU(ctx context.Context) (*SellerVoucheDeleteSelectedProductSKUResponse, error) {
	path := "/promotion/voucher/product/sku/remove"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(SellerVoucheDeleteSelectedProductSKUResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// SellerVoucherActivate activate seller voucher promotion
// Path: /promotion/voucher/activate
func (s *SellerVoucherServiceOp[T]) SellerVoucherActivate(ctx context.Context) (*SellerVoucherActivateResponse, error) {
	path := "/promotion/voucher/activate"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(SellerVoucherActivateResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// SellerVoucherAddSelectedProductSKU add seller voucher promotion product sku
// Path: /promotion/voucher/product/sku/add
func (s *SellerVoucherServiceOp[T]) SellerVoucherAddSelectedProductSKU(ctx context.Context) (*SellerVoucherAddSelectedProductSKUResponse, error) {
	path := "/promotion/voucher/product/sku/add"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(SellerVoucherAddSelectedProductSKUResponse)
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

// SellerVoucherCreate create a new seller voucher promotion
// Path: /promotion/voucher/create
func (s *SellerVoucherServiceOp[T]) SellerVoucherCreate(ctx context.Context) (*SellerVoucherCreateResponse, error) {
	path := "/promotion/voucher/create"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(SellerVoucherCreateResponse)
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

// SellerVoucherDeactivate deactivate seller voucher promotion
// Path: /promotion/voucher/deactivate
func (s *SellerVoucherServiceOp[T]) SellerVoucherDeactivate(ctx context.Context) (*SellerVoucherDeactivateResponse, error) {
	path := "/promotion/voucher/deactivate"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(SellerVoucherDeactivateResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// SellerVoucherDetailQuery get a seller voucher promotion detail
// Path: /promotion/voucher/get
func (s *SellerVoucherServiceOp[T]) SellerVoucherDetailQuery(ctx context.Context) (*SellerVoucherDetailQueryResponse, error) {
	path := "/promotion/voucher/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(SellerVoucherDetailQueryResponse)
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

// SellerVoucherList query seller voucher promotion list
// Path: /promotion/vouchers/get
func (s *SellerVoucherServiceOp[T]) SellerVoucherList(ctx context.Context) (*SellerVoucherListResponse, error) {
	path := "/promotion/vouchers/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(SellerVoucherListResponse)
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

// SellerVoucherSelectedProductList query seller voucher selected products list
// Path: /promotion/voucher/products/get
func (s *SellerVoucherServiceOp[T]) SellerVoucherSelectedProductList(ctx context.Context) (*SellerVoucherSelectedProductListResponse, error) {
	path := "/promotion/voucher/products/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(SellerVoucherSelectedProductListResponse)
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

// SellerVoucherUpdate update a existing seller voucher promotion
// Path: /promotion/voucher/update
func (s *SellerVoucherServiceOp[T]) SellerVoucherUpdate(ctx context.Context) (*SellerVoucherUpdateResponse, error) {
	path := "/promotion/voucher/update"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(SellerVoucherUpdateResponse)
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
