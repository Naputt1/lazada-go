package golazada

import (
	"context"
	"encoding/json"
	"fmt"
)

type OrderService interface {
	// GetDocument Use this API to retrieve order-related documents, including invoices and shipping labels.
	// Path: /order/document/get
	GetDocument(ctx context.Context) (*GetDocumentResponse, error)
	// GetMultipleOrderItems Use this API to get the item information of one or more orders.（No more than 50 at a time）
	// Path: /orders/items/get
	GetMultipleOrderItems(ctx context.Context, opt GetMultipleOrderItemsRequest) (*GetMultipleOrderItemsResponse, error)
	// GetOrder Use this API to get the list of items for a single order.
	// Path: /order/get
	GetOrder(ctx context.Context) (*GetOrderResponse, error)
	// GetOrderItems Use this API to get the item information of an order.
	// Path: /order/items/get
	GetOrderItems(ctx context.Context, opt GetOrderItemsRequest) (*GetOrderItemsResponse, error)
	// GetOrders Use this API to get the list of items for a range of orders1..
	// Path: /orders/get
	GetOrders(ctx context.Context, opt GetOrdersRequest) (*GetOrdersResponse, error)
	// GetOVOOrders This interface is only applicable to the merchant side of the business and is used to set the maximum number of SKUs that certain merchants can sell per day
	// Path: /orders/ovo/get
	GetOVOOrders(ctx context.Context) (*GetOVOOrdersResponse, error)
	// OrderCancelValidate Seller can check whether the order can be canceled through this API and get corresponding reasons if not.
	// Path: /order/reverse/cancel/validate
	OrderCancelValidate(ctx context.Context, opt OrderCancelValidateRequest) (*OrderCancelValidateResponse, error)
	// SetInvoiceNumber Use this API to set the invoice number for the specified order.
	// Path: /order/invoice_number/set
	SetInvoiceNumber(ctx context.Context) (*SetInvoiceNumberResponse, error)
}

type OrderServiceOp[T any] struct {
	client *Client[T]
}

// GetDocument Use this API to retrieve order-related documents, including invoices and shipping labels.
// Path: /order/document/get
func (s *OrderServiceOp[T]) GetDocument(ctx context.Context) (*GetDocumentResponse, error) {
	path := "/order/document/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetDocumentResponse)
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

// GetMultipleOrderItems Use this API to get the item information of one or more orders.（No more than 50 at a time）
// Path: /orders/items/get
func (s *OrderServiceOp[T]) GetMultipleOrderItems(ctx context.Context, opt GetMultipleOrderItemsRequest) (*GetMultipleOrderItemsResponse, error) {
	path := "/orders/items/get"
	params := paramsFromStruct(opt)
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetMultipleOrderItemsResponse)
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

// GetOrder Use this API to get the list of items for a single order.
// Path: /order/get
func (s *OrderServiceOp[T]) GetOrder(ctx context.Context) (*GetOrderResponse, error) {
	path := "/order/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetOrderResponse)
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

// GetOrderItems Use this API to get the item information of an order.
// Path: /order/items/get
func (s *OrderServiceOp[T]) GetOrderItems(ctx context.Context, opt GetOrderItemsRequest) (*GetOrderItemsResponse, error) {
	path := "/order/items/get"
	params := paramsFromStruct(opt)
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetOrderItemsResponse)
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

// GetOrders Use this API to get the list of items for a range of orders1..
// Path: /orders/get
func (s *OrderServiceOp[T]) GetOrders(ctx context.Context, opt GetOrdersRequest) (*GetOrdersResponse, error) {
	path := "/orders/get"
	params := paramsFromStruct(opt)
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetOrdersResponse)
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

// GetOVOOrders This interface is only applicable to the merchant side of the business and is used to set the maximum number of SKUs that certain merchants can sell per day
// Path: /orders/ovo/get
func (s *OrderServiceOp[T]) GetOVOOrders(ctx context.Context) (*GetOVOOrdersResponse, error) {
	path := "/orders/ovo/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetOVOOrdersResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// OrderCancelValidate Seller can check whether the order can be canceled through this API and get corresponding reasons if not.
// Path: /order/reverse/cancel/validate
func (s *OrderServiceOp[T]) OrderCancelValidate(ctx context.Context, opt OrderCancelValidateRequest) (*OrderCancelValidateResponse, error) {
	path := "/order/reverse/cancel/validate"
	params := paramsFromStruct(opt)
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(OrderCancelValidateResponse)
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

// SetInvoiceNumber Use this API to set the invoice number for the specified order.
// Path: /order/invoice_number/set
func (s *OrderServiceOp[T]) SetInvoiceNumber(ctx context.Context) (*SetInvoiceNumberResponse, error) {
	path := "/order/invoice_number/set"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(SetInvoiceNumberResponse)
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
