package golazada

import (
	"context"
)

type FinanceService interface {
	// GetPayoutStatus Get your transaction statements  created after the provided date
	// Path: /finance/payout/status/get
	GetPayoutStatus(ctx context.Context) (*GetPayoutStatusResponse, error)
	// QueryAccountTransactions Query Account Transactions
	// Path: /finance/transaction/accountTransactions/query
	QueryAccountTransactions(ctx context.Context) (*QueryAccountTransactionsResponse, error)
	// QueryLogisticsFeeDetail Api is provided for finance and seller to query logistics fee details from slb.
	// Path: /lbs/slb/queryLogisticsFeeDetail
	QueryLogisticsFeeDetail(ctx context.Context) (*QueryLogisticsFeeDetailResponse, error)
	// QueryTransactionDetails API to query seller transaction details within specific date range.
	// Path: /finance/transaction/details/get
	QueryTransactionDetails(ctx context.Context) (*QueryTransactionDetailsResponse, error)
}

type FinanceServiceOp[T any] struct {
	client *Client[T]
}

// GetPayoutStatus Get your transaction statements  created after the provided date
// Path: /finance/payout/status/get
func (s *FinanceServiceOp[T]) GetPayoutStatus(ctx context.Context) (*GetPayoutStatusResponse, error) {
	path := "/finance/payout/status/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetPayoutStatusResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// QueryAccountTransactions Query Account Transactions
// Path: /finance/transaction/accountTransactions/query
func (s *FinanceServiceOp[T]) QueryAccountTransactions(ctx context.Context) (*QueryAccountTransactionsResponse, error) {
	path := "/finance/transaction/accountTransactions/query"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(QueryAccountTransactionsResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// QueryLogisticsFeeDetail Api is provided for finance and seller to query logistics fee details from slb.
// Path: /lbs/slb/queryLogisticsFeeDetail
func (s *FinanceServiceOp[T]) QueryLogisticsFeeDetail(ctx context.Context) (*QueryLogisticsFeeDetailResponse, error) {
	path := "/lbs/slb/queryLogisticsFeeDetail"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(QueryLogisticsFeeDetailResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// QueryTransactionDetails API to query seller transaction details within specific date range.
// Path: /finance/transaction/details/get
func (s *FinanceServiceOp[T]) QueryTransactionDetails(ctx context.Context) (*QueryTransactionDetailsResponse, error) {
	path := "/finance/transaction/details/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(QueryTransactionDetailsResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}
