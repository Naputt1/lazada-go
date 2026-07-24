package golazada

import (
	"context"
)

type FreeShippingService interface {
	// FreeShippingActivate activate free shipping promotion
	// Path: /promotion/freeshipping/activate
	FreeShippingActivate(ctx context.Context) (*FreeShippingActivateResponse, error)
	// FreeShippingAddSelectedProductSKU add sku for free shipping promotion
	// Path: /promotion/freeshipping/product/sku/add
	FreeShippingAddSelectedProductSKU(ctx context.Context) (*FreeShippingAddSelectedProductSKUResponse, error)
	// FreeShippingCreate create a new free shipping promotion
	// Path: /promotion/freeshipping/create
	FreeShippingCreate(ctx context.Context) (*FreeShippingCreateResponse, error)
	// FreeShippingDeactivate deactivate free shipping promotion
	// Path: /promotion/freeshipping/deactivate
	FreeShippingDeactivate(ctx context.Context) (*FreeShippingDeactivateResponse, error)
	// FreeShippingDeleteSelectedProductSKU delete sku for free shipping promotion
	// Path: /promotion/freeshipping/product/sku/remove
	FreeShippingDeleteSelectedProductSKU(ctx context.Context) (*FreeShippingDeleteSelectedProductSKUResponse, error)
	// FreeShippingDeliveryOptionsQuery query free shipping promotion delivery options
	// Path: /promotion/freeshipping/deliveryoptions/get
	FreeShippingDeliveryOptionsQuery(ctx context.Context) (*FreeShippingDeliveryOptionsQueryResponse, error)
	// FreeShippingGet get free shipping promotion
	// Path: /promotion/freeshipping/get
	FreeShippingGet(ctx context.Context) (*FreeShippingGetResponse, error)
	// FreeShippingList query free shipping promotion list
	// Path: /promotion/freeshippings/get
	FreeShippingList(ctx context.Context) (*FreeShippingListResponse, error)
	// FreeShippingRegionsQuery query free shipping promotion regions
	// Path: /promotion/freeshipping/regions/get
	FreeShippingRegionsQuery(ctx context.Context) (*FreeShippingRegionsQueryResponse, error)
	// FreeShippingSelectedProductList query free shipping promotion selected product list
	// Path: /promotion/freeshipping/products/get
	FreeShippingSelectedProductList(ctx context.Context) (*FreeShippingSelectedProductListResponse, error)
	// FreeShippingUpdate update free shipping promotion
	// Path: /promotion/freeshipping/update
	FreeShippingUpdate(ctx context.Context) (*FreeShippingUpdateResponse, error)
}

type FreeShippingServiceOp[T any] struct {
	client *Client[T]
}

// FreeShippingActivate activate free shipping promotion
// Path: /promotion/freeshipping/activate
func (s *FreeShippingServiceOp[T]) FreeShippingActivate(ctx context.Context) (*FreeShippingActivateResponse, error) {
	path := "/promotion/freeshipping/activate"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(FreeShippingActivateResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// FreeShippingAddSelectedProductSKU add sku for free shipping promotion
// Path: /promotion/freeshipping/product/sku/add
func (s *FreeShippingServiceOp[T]) FreeShippingAddSelectedProductSKU(ctx context.Context) (*FreeShippingAddSelectedProductSKUResponse, error) {
	path := "/promotion/freeshipping/product/sku/add"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(FreeShippingAddSelectedProductSKUResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// FreeShippingCreate create a new free shipping promotion
// Path: /promotion/freeshipping/create
func (s *FreeShippingServiceOp[T]) FreeShippingCreate(ctx context.Context) (*FreeShippingCreateResponse, error) {
	path := "/promotion/freeshipping/create"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(FreeShippingCreateResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// FreeShippingDeactivate deactivate free shipping promotion
// Path: /promotion/freeshipping/deactivate
func (s *FreeShippingServiceOp[T]) FreeShippingDeactivate(ctx context.Context) (*FreeShippingDeactivateResponse, error) {
	path := "/promotion/freeshipping/deactivate"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(FreeShippingDeactivateResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// FreeShippingDeleteSelectedProductSKU delete sku for free shipping promotion
// Path: /promotion/freeshipping/product/sku/remove
func (s *FreeShippingServiceOp[T]) FreeShippingDeleteSelectedProductSKU(ctx context.Context) (*FreeShippingDeleteSelectedProductSKUResponse, error) {
	path := "/promotion/freeshipping/product/sku/remove"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(FreeShippingDeleteSelectedProductSKUResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// FreeShippingDeliveryOptionsQuery query free shipping promotion delivery options
// Path: /promotion/freeshipping/deliveryoptions/get
func (s *FreeShippingServiceOp[T]) FreeShippingDeliveryOptionsQuery(ctx context.Context) (*FreeShippingDeliveryOptionsQueryResponse, error) {
	path := "/promotion/freeshipping/deliveryoptions/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(FreeShippingDeliveryOptionsQueryResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// FreeShippingGet get free shipping promotion
// Path: /promotion/freeshipping/get
func (s *FreeShippingServiceOp[T]) FreeShippingGet(ctx context.Context) (*FreeShippingGetResponse, error) {
	path := "/promotion/freeshipping/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(FreeShippingGetResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// FreeShippingList query free shipping promotion list
// Path: /promotion/freeshippings/get
func (s *FreeShippingServiceOp[T]) FreeShippingList(ctx context.Context) (*FreeShippingListResponse, error) {
	path := "/promotion/freeshippings/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(FreeShippingListResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// FreeShippingRegionsQuery query free shipping promotion regions
// Path: /promotion/freeshipping/regions/get
func (s *FreeShippingServiceOp[T]) FreeShippingRegionsQuery(ctx context.Context) (*FreeShippingRegionsQueryResponse, error) {
	path := "/promotion/freeshipping/regions/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(FreeShippingRegionsQueryResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// FreeShippingSelectedProductList query free shipping promotion selected product list
// Path: /promotion/freeshipping/products/get
func (s *FreeShippingServiceOp[T]) FreeShippingSelectedProductList(ctx context.Context) (*FreeShippingSelectedProductListResponse, error) {
	path := "/promotion/freeshipping/products/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(FreeShippingSelectedProductListResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// FreeShippingUpdate update free shipping promotion
// Path: /promotion/freeshipping/update
func (s *FreeShippingServiceOp[T]) FreeShippingUpdate(ctx context.Context) (*FreeShippingUpdateResponse, error) {
	path := "/promotion/freeshipping/update"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(FreeShippingUpdateResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}
