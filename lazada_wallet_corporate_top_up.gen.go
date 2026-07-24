package golazada

import (
	"context"
)

type LazadaWalletCorporateTopUpService interface {
	// DirectTransferQuery Direct Transfer - Query
	// Path: /wallet/transfer/query
	DirectTransferQuery(ctx context.Context) (*DirectTransferQueryResponse, error)
	// DirectTransferRequest Direct Transfer - Request to transfer
	// Path: /wallet/transfer/request
	DirectTransferRequest(ctx context.Context) (*DirectTransferRequestResponse, error)
	// GiftCodeQuery Gift Code - Query
	// Path: /wallet/giftcode/query
	GiftCodeQuery(ctx context.Context) (*GiftCodeQueryResponse, error)
	// GiftCodeRequest Gift Code - Request
	// Path: /wallet/giftcode/request
	GiftCodeRequest(ctx context.Context) (*GiftCodeRequestResponse, error)
	// Reconciliation1 Corporate TopUp - Reconciliation
	// Path: /wallet/open/reconciliation
	Reconciliation1(ctx context.Context) (*Reconciliation1Response, error)
}

type LazadaWalletCorporateTopUpServiceOp[T any] struct {
	client *Client[T]
}

// DirectTransferQuery Direct Transfer - Query
// Path: /wallet/transfer/query
func (s *LazadaWalletCorporateTopUpServiceOp[T]) DirectTransferQuery(ctx context.Context) (*DirectTransferQueryResponse, error) {
	path := "/wallet/transfer/query"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(DirectTransferQueryResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// DirectTransferRequest Direct Transfer - Request to transfer
// Path: /wallet/transfer/request
func (s *LazadaWalletCorporateTopUpServiceOp[T]) DirectTransferRequest(ctx context.Context) (*DirectTransferRequestResponse, error) {
	path := "/wallet/transfer/request"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(DirectTransferRequestResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GiftCodeQuery Gift Code - Query
// Path: /wallet/giftcode/query
func (s *LazadaWalletCorporateTopUpServiceOp[T]) GiftCodeQuery(ctx context.Context) (*GiftCodeQueryResponse, error) {
	path := "/wallet/giftcode/query"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(GiftCodeQueryResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GiftCodeRequest Gift Code - Request
// Path: /wallet/giftcode/request
func (s *LazadaWalletCorporateTopUpServiceOp[T]) GiftCodeRequest(ctx context.Context) (*GiftCodeRequestResponse, error) {
	path := "/wallet/giftcode/request"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(GiftCodeRequestResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// Reconciliation1 Corporate TopUp - Reconciliation
// Path: /wallet/open/reconciliation
func (s *LazadaWalletCorporateTopUpServiceOp[T]) Reconciliation1(ctx context.Context) (*Reconciliation1Response, error) {
	path := "/wallet/open/reconciliation"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(Reconciliation1Response)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}
