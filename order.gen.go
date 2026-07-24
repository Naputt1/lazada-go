package golazada

import (
	"context"
)

type OrderService interface {
	// GetDocument Use this API to retrieve order-related documents, including invoices and shipping labels.
	// Path: /order/document/get
	GetDocument(ctx context.Context) (*GetDocumentResponse, error)
	// GetMultipleOrderItems Use this API to get the item information of one or more orders.（No more than 50 at a time）
	// Path: /orders/items/get
	GetMultipleOrderItems(ctx context.Context) (*GetMultipleOrderItemsResponse, error)
	// GetOrder Use this API to get the list of items for a single order.
	// Path: /order/get
	GetOrder(ctx context.Context) (*GetOrderResponse, error)
	// GetOrderItems Use this API to get the item information of an order.
	// Path: /order/items/get
	GetOrderItems(ctx context.Context) (*GetOrderItemsResponse, error)
	// GetOrders Use this API to get the list of items for a range of orders1..
	// Path: /orders/get
	GetOrders(ctx context.Context) (*GetOrdersResponse, error)
	// GetOVOOrders This interface is only applicable to the merchant side of the business and is used to set the maximum number of SKUs that certain merchants can sell per day
	// Path: /orders/ovo/get
	GetOVOOrders(ctx context.Context) (*GetOVOOrdersResponse, error)
	// OrderCancelValidate Seller can check whether the order can be canceled through this API and get corresponding reasons if not.
	// Path: /order/reverse/cancel/validate
	OrderCancelValidate(ctx context.Context) (*OrderCancelValidateResponse, error)
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
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetMultipleOrderItems Use this API to get the item information of one or more orders.（No more than 50 at a time）
// Path: /orders/items/get
func (s *OrderServiceOp[T]) GetMultipleOrderItems(ctx context.Context) (*GetMultipleOrderItemsResponse, error) {
	path := "/orders/items/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetMultipleOrderItemsResponse)
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
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetOrderItems Use this API to get the item information of an order.
// Path: /order/items/get
func (s *OrderServiceOp[T]) GetOrderItems(ctx context.Context) (*GetOrderItemsResponse, error) {
	path := "/order/items/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetOrderItemsResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetOrders Use this API to get the list of items for a range of orders1..
// Path: /orders/get
func (s *OrderServiceOp[T]) GetOrders(ctx context.Context) (*GetOrdersResponse, error) {
	path := "/orders/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetOrdersResponse)
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
func (s *OrderServiceOp[T]) OrderCancelValidate(ctx context.Context) (*OrderCancelValidateResponse, error) {
	path := "/order/reverse/cancel/validate"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(OrderCancelValidateResponse)
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
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}
