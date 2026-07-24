package golazada

import (
	"context"
)

type FirstMileBigbagonlyForCNService interface {
	// GetChannelcodeByFirstMileNo get channelcode by first mile No
	// Path: /logistics/cngfc/fulfill/getchannelcode
	GetChannelcodeByFirstMileNo(ctx context.Context) (*GetChannelcodeByFirstMileNoResponse, error)
	// GetLazadaBigbagPDFLable Get Lazada Bigbag PDF Lable
	// Path: /logistics/cnpms/bigbag/lable/getPdf
	GetLazadaBigbagPDFLable(ctx context.Context) (*GetLazadaBigbagPDFLableResponse, error)
	// LazadaBigbagCancel Lazada Bigbag cancel
	// Path: /logistics/cnpms/bigbag/cancel
	LazadaBigbagCancel(ctx context.Context) (*LazadaBigbagCancelResponse, error)
	// LazadaBigbagCollectionPoints Lazada bigbag query collection points
	// Path: /logistics/cnpms/bigbag/querycollection
	LazadaBigbagCollectionPoints(ctx context.Context) (*LazadaBigbagCollectionPointsResponse, error)
	// LazadaBigbagCommit Lazada bigbag commit
	// Path: /logistics/cnpms/bigbag/commit
	LazadaBigbagCommit(ctx context.Context) (*LazadaBigbagCommitResponse, error)
	// LazadaBigbagUpdate Lazada bigbag update
	// Path: /logistics/cnpms/bigbag/update
	LazadaBigbagUpdate(ctx context.Context) (*LazadaBigbagUpdateResponse, error)
	// LazadaSellerAccountBind Lazada seller account bind for big bag pick up
	// Path: /logistics/cnpms/account/bind
	LazadaSellerAccountBind(ctx context.Context) (*LazadaSellerAccountBindResponse, error)
	// QueryAddressInformaiton Query Address Informaiton
	// Path: /logistics/cnpms/address/query
	QueryAddressInformaiton(ctx context.Context) (*QueryAddressInformaitonResponse, error)
	// QueryLazadaBigbagInfo Query Lazada Bigbag Info
	// Path: /logistics/cnpms/bigbag/query
	QueryLazadaBigbagInfo(ctx context.Context) (*QueryLazadaBigbagInfoResponse, error)
}

type FirstMileBigbagonlyForCNServiceOp[T any] struct {
	client *Client[T]
}

// GetChannelcodeByFirstMileNo get channelcode by first mile No
// Path: /logistics/cngfc/fulfill/getchannelcode
func (s *FirstMileBigbagonlyForCNServiceOp[T]) GetChannelcodeByFirstMileNo(ctx context.Context) (*GetChannelcodeByFirstMileNoResponse, error) {
	path := "/logistics/cngfc/fulfill/getchannelcode"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetChannelcodeByFirstMileNoResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetLazadaBigbagPDFLable Get Lazada Bigbag PDF Lable
// Path: /logistics/cnpms/bigbag/lable/getPdf
func (s *FirstMileBigbagonlyForCNServiceOp[T]) GetLazadaBigbagPDFLable(ctx context.Context) (*GetLazadaBigbagPDFLableResponse, error) {
	path := "/logistics/cnpms/bigbag/lable/getPdf"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetLazadaBigbagPDFLableResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// LazadaBigbagCancel Lazada Bigbag cancel
// Path: /logistics/cnpms/bigbag/cancel
func (s *FirstMileBigbagonlyForCNServiceOp[T]) LazadaBigbagCancel(ctx context.Context) (*LazadaBigbagCancelResponse, error) {
	path := "/logistics/cnpms/bigbag/cancel"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(LazadaBigbagCancelResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// LazadaBigbagCollectionPoints Lazada bigbag query collection points
// Path: /logistics/cnpms/bigbag/querycollection
func (s *FirstMileBigbagonlyForCNServiceOp[T]) LazadaBigbagCollectionPoints(ctx context.Context) (*LazadaBigbagCollectionPointsResponse, error) {
	path := "/logistics/cnpms/bigbag/querycollection"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(LazadaBigbagCollectionPointsResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// LazadaBigbagCommit Lazada bigbag commit
// Path: /logistics/cnpms/bigbag/commit
func (s *FirstMileBigbagonlyForCNServiceOp[T]) LazadaBigbagCommit(ctx context.Context) (*LazadaBigbagCommitResponse, error) {
	path := "/logistics/cnpms/bigbag/commit"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(LazadaBigbagCommitResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// LazadaBigbagUpdate Lazada bigbag update
// Path: /logistics/cnpms/bigbag/update
func (s *FirstMileBigbagonlyForCNServiceOp[T]) LazadaBigbagUpdate(ctx context.Context) (*LazadaBigbagUpdateResponse, error) {
	path := "/logistics/cnpms/bigbag/update"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(LazadaBigbagUpdateResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// LazadaSellerAccountBind Lazada seller account bind for big bag pick up
// Path: /logistics/cnpms/account/bind
func (s *FirstMileBigbagonlyForCNServiceOp[T]) LazadaSellerAccountBind(ctx context.Context) (*LazadaSellerAccountBindResponse, error) {
	path := "/logistics/cnpms/account/bind"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(LazadaSellerAccountBindResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// QueryAddressInformaiton Query Address Informaiton
// Path: /logistics/cnpms/address/query
func (s *FirstMileBigbagonlyForCNServiceOp[T]) QueryAddressInformaiton(ctx context.Context) (*QueryAddressInformaitonResponse, error) {
	path := "/logistics/cnpms/address/query"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(QueryAddressInformaitonResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// QueryLazadaBigbagInfo Query Lazada Bigbag Info
// Path: /logistics/cnpms/bigbag/query
func (s *FirstMileBigbagonlyForCNServiceOp[T]) QueryLazadaBigbagInfo(ctx context.Context) (*QueryLazadaBigbagInfoResponse, error) {
	path := "/logistics/cnpms/bigbag/query"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(QueryLazadaBigbagInfoResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}
