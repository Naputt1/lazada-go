package golazada

import (
	"context"
)

type LazPayService interface {
	// CollectBenefit collect lazada marketplace benefit
	// Path: /insurance/promotion/collectBenefit
	CollectBenefit(ctx context.Context) (*CollectBenefitResponse, error)
	// ConsultPayment The interface is used for consult pay view. Will return pay view info including balance, coupon, credit card etc. If we have no available coupon, we will return pay method view with an empty list of coupon.
	// Path: /lazadapay/v1/debit/consult_payment
	ConsultPayment(ctx context.Context) (*ConsultPaymentResponse, error)
	// CreateSubscriptionToFusion Create User Subscription To Fusion
	// Path: /insurance/subscription/create
	CreateSubscriptionToFusion(ctx context.Context) (*CreateSubscriptionToFusionResponse, error)
	// DGUtiityPreCreateOrder This API provides an open interface for partner users to create DG orders
	// Path: /digital/service/createorder
	DGUtiityPreCreateOrder(ctx context.Context) (*DGUtiityPreCreateOrderResponse, error)
	// DGUtilityPreGetPaymentStatus get payment status
	// Path: /digital/service/getPaymentStatus
	DGUtilityPreGetPaymentStatus(ctx context.Context) (*DGUtilityPreGetPaymentStatusResponse, error)
	// DGUtilityPreUpdateFulfillemtStatus update fulfillemt status
	// Path: /digital/service/updateFulfillemtStatus
	DGUtilityPreUpdateFulfillemtStatus(ctx context.Context) (*DGUtilityPreUpdateFulfillemtStatusResponse, error)
	// DigitalAlterOrderStatus Change Lazada Digital Order Status
	// Path: /digital/order/alterStatus
	DigitalAlterOrderStatus(ctx context.Context) (*DigitalAlterOrderStatusResponse, error)
	// DigitalCreateOrder Create Digital Virtual Order
	// Path: /digital/order/create
	DigitalCreateOrder(ctx context.Context) (*DigitalCreateOrderResponse, error)
	// DigitalQueryOrder Query Lazada Digital Order Status
	// Path: /digital/order/getStatus
	DigitalQueryOrder(ctx context.Context) (*DigitalQueryOrderResponse, error)
	// GetSubscriptionToFusion Get User Subscription To Fusion
	// Path: /insurance/subscription/getSubscription
	GetSubscriptionToFusion(ctx context.Context) (*GetSubscriptionToFusionResponse, error)
	// InsuranceAlterOrderStatus Change Lazada Insurance Order Status
	// Path: /insurance/order/alterStatus
	InsuranceAlterOrderStatus(ctx context.Context) (*InsuranceAlterOrderStatusResponse, error)
	// InsuranceCreateOrder Lazada Insurance Create Order
	// Path: /insurance/order/create
	InsuranceCreateOrder(ctx context.Context) (*InsuranceCreateOrderResponse, error)
	// InsuranceGetPromotions get lazada marketplace  ump promotions
	// Path: /insurance/promotion/getPromotions
	InsuranceGetPromotions(ctx context.Context) (*InsuranceGetPromotionsResponse, error)
	// InsuranceQueryOrder Query Lazada Insurance Order Status
	// Path: /insurance/order/getStatus
	InsuranceQueryOrder(ctx context.Context) (*InsuranceQueryOrderResponse, error)
	// InsuranceRealTimeCDP 用户完成操作后，实时更新CDP人群
	// Path: /insurance/syncCDP
	InsuranceRealTimeCDP(ctx context.Context) (*InsuranceRealTimeCDPResponse, error)
	// LazadaCFOInvoiceRpaCallback Call RPA and return the official invoice
	// Path: /rpa/id/tax/callback
	LazadaCFOInvoiceRpaCallback(ctx context.Context) (*LazadaCFOInvoiceRpaCallbackResponse, error)
	// OpenServiceBalanceQuery Open Service Account Balance Info Query
	// Path: /wallet/open/service/balance/query
	OpenServiceBalanceQuery(ctx context.Context) (*OpenServiceBalanceQueryResponse, error)
	// OpenServiceKycQuery Open Service User KYC Info Query
	// Path: /wallet/open/service/kyc/query
	OpenServiceKycQuery(ctx context.Context) (*OpenServiceKycQueryResponse, error)
	// OpenServiceWithdrawApply Open Service Withdraw Apply
	// Path: /wallet/open/service/withdraw
	OpenServiceWithdrawApply(ctx context.Context) (*OpenServiceWithdrawApplyResponse, error)
	// OpenServiceWithdrawQuery Open Service Withdraw Query
	// Path: /wallet/open/service/withdraw/query
	OpenServiceWithdrawQuery(ctx context.Context) (*OpenServiceWithdrawQueryResponse, error)
	// QueryAddonOrder list user  addon order detail
	// Path: /insurance/addon/orders/query
	QueryAddonOrder(ctx context.Context) (*QueryAddonOrderResponse, error)
	// QueryBenefit get lazada marketplace benefit
	// Path: /insurance/promotion/queryBenefit
	QueryBenefit(ctx context.Context) (*QueryBenefitResponse, error)
	// Reconciliation Reconciliation
	// Path: /wallet/open/service/reconciliation
	Reconciliation(ctx context.Context) (*ReconciliationResponse, error)
	// RedeemMpVoucher 商城险域外voucher核销
	// Path: /insurance/voucher/redeemVoucher
	RedeemMpVoucher(ctx context.Context) (*RedeemMpVoucherResponse, error)
}

type LazPayServiceOp[T any] struct {
	client *Client[T]
}

// CollectBenefit collect lazada marketplace benefit
// Path: /insurance/promotion/collectBenefit
func (s *LazPayServiceOp[T]) CollectBenefit(ctx context.Context) (*CollectBenefitResponse, error) {
	path := "/insurance/promotion/collectBenefit"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(CollectBenefitResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// ConsultPayment The interface is used for consult pay view. Will return pay view info including balance, coupon, credit card etc. If we have no available coupon, we will return pay method view with an empty list of coupon.
// Path: /lazadapay/v1/debit/consult_payment
func (s *LazPayServiceOp[T]) ConsultPayment(ctx context.Context) (*ConsultPaymentResponse, error) {
	path := "/lazadapay/v1/debit/consult_payment"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(ConsultPaymentResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// CreateSubscriptionToFusion Create User Subscription To Fusion
// Path: /insurance/subscription/create
func (s *LazPayServiceOp[T]) CreateSubscriptionToFusion(ctx context.Context) (*CreateSubscriptionToFusionResponse, error) {
	path := "/insurance/subscription/create"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(CreateSubscriptionToFusionResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// DGUtiityPreCreateOrder This API provides an open interface for partner users to create DG orders
// Path: /digital/service/createorder
func (s *LazPayServiceOp[T]) DGUtiityPreCreateOrder(ctx context.Context) (*DGUtiityPreCreateOrderResponse, error) {
	path := "/digital/service/createorder"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(DGUtiityPreCreateOrderResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// DGUtilityPreGetPaymentStatus get payment status
// Path: /digital/service/getPaymentStatus
func (s *LazPayServiceOp[T]) DGUtilityPreGetPaymentStatus(ctx context.Context) (*DGUtilityPreGetPaymentStatusResponse, error) {
	path := "/digital/service/getPaymentStatus"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(DGUtilityPreGetPaymentStatusResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// DGUtilityPreUpdateFulfillemtStatus update fulfillemt status
// Path: /digital/service/updateFulfillemtStatus
func (s *LazPayServiceOp[T]) DGUtilityPreUpdateFulfillemtStatus(ctx context.Context) (*DGUtilityPreUpdateFulfillemtStatusResponse, error) {
	path := "/digital/service/updateFulfillemtStatus"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(DGUtilityPreUpdateFulfillemtStatusResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// DigitalAlterOrderStatus Change Lazada Digital Order Status
// Path: /digital/order/alterStatus
func (s *LazPayServiceOp[T]) DigitalAlterOrderStatus(ctx context.Context) (*DigitalAlterOrderStatusResponse, error) {
	path := "/digital/order/alterStatus"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(DigitalAlterOrderStatusResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// DigitalCreateOrder Create Digital Virtual Order
// Path: /digital/order/create
func (s *LazPayServiceOp[T]) DigitalCreateOrder(ctx context.Context) (*DigitalCreateOrderResponse, error) {
	path := "/digital/order/create"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(DigitalCreateOrderResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// DigitalQueryOrder Query Lazada Digital Order Status
// Path: /digital/order/getStatus
func (s *LazPayServiceOp[T]) DigitalQueryOrder(ctx context.Context) (*DigitalQueryOrderResponse, error) {
	path := "/digital/order/getStatus"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(DigitalQueryOrderResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetSubscriptionToFusion Get User Subscription To Fusion
// Path: /insurance/subscription/getSubscription
func (s *LazPayServiceOp[T]) GetSubscriptionToFusion(ctx context.Context) (*GetSubscriptionToFusionResponse, error) {
	path := "/insurance/subscription/getSubscription"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetSubscriptionToFusionResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// InsuranceAlterOrderStatus Change Lazada Insurance Order Status
// Path: /insurance/order/alterStatus
func (s *LazPayServiceOp[T]) InsuranceAlterOrderStatus(ctx context.Context) (*InsuranceAlterOrderStatusResponse, error) {
	path := "/insurance/order/alterStatus"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(InsuranceAlterOrderStatusResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// InsuranceCreateOrder Lazada Insurance Create Order
// Path: /insurance/order/create
func (s *LazPayServiceOp[T]) InsuranceCreateOrder(ctx context.Context) (*InsuranceCreateOrderResponse, error) {
	path := "/insurance/order/create"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(InsuranceCreateOrderResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// InsuranceGetPromotions get lazada marketplace  ump promotions
// Path: /insurance/promotion/getPromotions
func (s *LazPayServiceOp[T]) InsuranceGetPromotions(ctx context.Context) (*InsuranceGetPromotionsResponse, error) {
	path := "/insurance/promotion/getPromotions"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(InsuranceGetPromotionsResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// InsuranceQueryOrder Query Lazada Insurance Order Status
// Path: /insurance/order/getStatus
func (s *LazPayServiceOp[T]) InsuranceQueryOrder(ctx context.Context) (*InsuranceQueryOrderResponse, error) {
	path := "/insurance/order/getStatus"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(InsuranceQueryOrderResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// InsuranceRealTimeCDP 用户完成操作后，实时更新CDP人群
// Path: /insurance/syncCDP
func (s *LazPayServiceOp[T]) InsuranceRealTimeCDP(ctx context.Context) (*InsuranceRealTimeCDPResponse, error) {
	path := "/insurance/syncCDP"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(InsuranceRealTimeCDPResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// LazadaCFOInvoiceRpaCallback Call RPA and return the official invoice
// Path: /rpa/id/tax/callback
func (s *LazPayServiceOp[T]) LazadaCFOInvoiceRpaCallback(ctx context.Context) (*LazadaCFOInvoiceRpaCallbackResponse, error) {
	path := "/rpa/id/tax/callback"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(LazadaCFOInvoiceRpaCallbackResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// OpenServiceBalanceQuery Open Service Account Balance Info Query
// Path: /wallet/open/service/balance/query
func (s *LazPayServiceOp[T]) OpenServiceBalanceQuery(ctx context.Context) (*OpenServiceBalanceQueryResponse, error) {
	path := "/wallet/open/service/balance/query"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(OpenServiceBalanceQueryResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// OpenServiceKycQuery Open Service User KYC Info Query
// Path: /wallet/open/service/kyc/query
func (s *LazPayServiceOp[T]) OpenServiceKycQuery(ctx context.Context) (*OpenServiceKycQueryResponse, error) {
	path := "/wallet/open/service/kyc/query"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(OpenServiceKycQueryResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// OpenServiceWithdrawApply Open Service Withdraw Apply
// Path: /wallet/open/service/withdraw
func (s *LazPayServiceOp[T]) OpenServiceWithdrawApply(ctx context.Context) (*OpenServiceWithdrawApplyResponse, error) {
	path := "/wallet/open/service/withdraw"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(OpenServiceWithdrawApplyResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// OpenServiceWithdrawQuery Open Service Withdraw Query
// Path: /wallet/open/service/withdraw/query
func (s *LazPayServiceOp[T]) OpenServiceWithdrawQuery(ctx context.Context) (*OpenServiceWithdrawQueryResponse, error) {
	path := "/wallet/open/service/withdraw/query"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(OpenServiceWithdrawQueryResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// QueryAddonOrder list user  addon order detail
// Path: /insurance/addon/orders/query
func (s *LazPayServiceOp[T]) QueryAddonOrder(ctx context.Context) (*QueryAddonOrderResponse, error) {
	path := "/insurance/addon/orders/query"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(QueryAddonOrderResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// QueryBenefit get lazada marketplace benefit
// Path: /insurance/promotion/queryBenefit
func (s *LazPayServiceOp[T]) QueryBenefit(ctx context.Context) (*QueryBenefitResponse, error) {
	path := "/insurance/promotion/queryBenefit"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(QueryBenefitResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// Reconciliation Reconciliation
// Path: /wallet/open/service/reconciliation
func (s *LazPayServiceOp[T]) Reconciliation(ctx context.Context) (*ReconciliationResponse, error) {
	path := "/wallet/open/service/reconciliation"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(ReconciliationResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// RedeemMpVoucher 商城险域外voucher核销
// Path: /insurance/voucher/redeemVoucher
func (s *LazPayServiceOp[T]) RedeemMpVoucher(ctx context.Context) (*RedeemMpVoucherResponse, error) {
	path := "/insurance/voucher/redeemVoucher"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(RedeemMpVoucherResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}
