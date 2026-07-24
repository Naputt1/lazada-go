package golazada

import (
	"context"
)

type ETicketsService interface {
	// GetOrderItemsFromBarCode E-Ticcket certificate query Open API
	// Path: /eticket/code/query
	GetOrderItemsFromBarCode(ctx context.Context) (*GetOrderItemsFromBarCodeResponse, error)
	// GlobalEticketMerchantMaAvailable the callback interface before consume  code
	// Path: /eticket/ma/available
	GlobalEticketMerchantMaAvailable(ctx context.Context) (*GlobalEticketMerchantMaAvailableResponse, error)
	// GlobalEticketMerchantMaConsume consume ma
	// Path: /eticket/ma/consume
	GlobalEticketMerchantMaConsume(ctx context.Context) (*GlobalEticketMerchantMaConsumeResponse, error)
	// GlobalEticketMerchantMaFailsend the callback interface when send code failed
	// Path: /eticket/ma/failsend
	GlobalEticketMerchantMaFailsend(ctx context.Context) (*GlobalEticketMerchantMaFailsendResponse, error)
	// GlobalEticketMerchantMaQuery the callback interface that query lazada platform ma
	// Path: /eticket/ma/query
	GlobalEticketMerchantMaQuery(ctx context.Context) (*GlobalEticketMerchantMaQueryResponse, error)
	// GlobalEticketMerchantMaQueryTbMa the callback interface that query tb ma
	// Path: /eticket/ma/queryTbMa
	GlobalEticketMerchantMaQueryTbMa(ctx context.Context) (*GlobalEticketMerchantMaQueryTbMaResponse, error)
	// GlobalEticketMerchantMaSend the callback interface when merchant send code successful
	// Path: /eticket/ma/send
	GlobalEticketMerchantMaSend(ctx context.Context) (*GlobalEticketMerchantMaSendResponse, error)
	// RedeemOrderItems Certificate Consume Open API
	// Path: /eticket/code/consume
	RedeemOrderItems(ctx context.Context) (*RedeemOrderItemsResponse, error)
}

type ETicketsServiceOp[T any] struct {
	client *Client[T]
}

// GetOrderItemsFromBarCode E-Ticcket certificate query Open API
// Path: /eticket/code/query
func (s *ETicketsServiceOp[T]) GetOrderItemsFromBarCode(ctx context.Context) (*GetOrderItemsFromBarCodeResponse, error) {
	path := "/eticket/code/query"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetOrderItemsFromBarCodeResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GlobalEticketMerchantMaAvailable the callback interface before consume  code
// Path: /eticket/ma/available
func (s *ETicketsServiceOp[T]) GlobalEticketMerchantMaAvailable(ctx context.Context) (*GlobalEticketMerchantMaAvailableResponse, error) {
	path := "/eticket/ma/available"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(GlobalEticketMerchantMaAvailableResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GlobalEticketMerchantMaConsume consume ma
// Path: /eticket/ma/consume
func (s *ETicketsServiceOp[T]) GlobalEticketMerchantMaConsume(ctx context.Context) (*GlobalEticketMerchantMaConsumeResponse, error) {
	path := "/eticket/ma/consume"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(GlobalEticketMerchantMaConsumeResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GlobalEticketMerchantMaFailsend the callback interface when send code failed
// Path: /eticket/ma/failsend
func (s *ETicketsServiceOp[T]) GlobalEticketMerchantMaFailsend(ctx context.Context) (*GlobalEticketMerchantMaFailsendResponse, error) {
	path := "/eticket/ma/failsend"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(GlobalEticketMerchantMaFailsendResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GlobalEticketMerchantMaQuery the callback interface that query lazada platform ma
// Path: /eticket/ma/query
func (s *ETicketsServiceOp[T]) GlobalEticketMerchantMaQuery(ctx context.Context) (*GlobalEticketMerchantMaQueryResponse, error) {
	path := "/eticket/ma/query"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(GlobalEticketMerchantMaQueryResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GlobalEticketMerchantMaQueryTbMa the callback interface that query tb ma
// Path: /eticket/ma/queryTbMa
func (s *ETicketsServiceOp[T]) GlobalEticketMerchantMaQueryTbMa(ctx context.Context) (*GlobalEticketMerchantMaQueryTbMaResponse, error) {
	path := "/eticket/ma/queryTbMa"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(GlobalEticketMerchantMaQueryTbMaResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GlobalEticketMerchantMaSend the callback interface when merchant send code successful
// Path: /eticket/ma/send
func (s *ETicketsServiceOp[T]) GlobalEticketMerchantMaSend(ctx context.Context) (*GlobalEticketMerchantMaSendResponse, error) {
	path := "/eticket/ma/send"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(GlobalEticketMerchantMaSendResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// RedeemOrderItems Certificate Consume Open API
// Path: /eticket/code/consume
func (s *ETicketsServiceOp[T]) RedeemOrderItems(ctx context.Context) (*RedeemOrderItemsResponse, error) {
	path := "/eticket/code/consume"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(RedeemOrderItemsResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}
