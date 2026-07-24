package golazada

import (
	"context"
)

type LazadaDGService interface {
	// DigitalServiceCdkCodeReceived 接受码商发码请求，给用户发送码。
	// Path: /digital/service/cdkCodeReceived
	DigitalServiceCdkCodeReceived(ctx context.Context) (*DigitalServiceCdkCodeReceivedResponse, error)
	// InstallServiceCallBack Install the service callback interface
	// Path: /digital/install/servicecallback
	InstallServiceCallBack(ctx context.Context) (*InstallServiceCallBackResponse, error)
	// InstallServiceCallBack1 Install the service callback interface
	// Path: /digital/test/install/servicecallback
	InstallServiceCallBack1(ctx context.Context) (*InstallServiceCallBack1Response, error)
	// InstallServiceCallBackForTest Install the service callback interface
	// Path: /digital/install/test/servicecallback
	InstallServiceCallBackForTest(ctx context.Context) (*InstallServiceCallBackForTestResponse, error)
	// InuranceNotication Third party insurance company callback interface
	//
	// Path: /digital/insurance/notification
	InuranceNotication(ctx context.Context) (*InuranceNoticationResponse, error)
	// InuranceNotication1 Third party insurance company callback interface
	//
	// Path: /digital/insurance/test/notificationcopy
	InuranceNotication1(ctx context.Context) (*InuranceNotication1Response, error)
	// InuranceNotifyLapse Insurance company push the callback notification to partners once the policy has been cancelled successfully
	// Path: /digital/insurance/notificationlapse
	InuranceNotifyLapse(ctx context.Context) (*InuranceNotifyLapseResponse, error)
}

type LazadaDGServiceOp[T any] struct {
	client *Client[T]
}

// DigitalServiceCdkCodeReceived 接受码商发码请求，给用户发送码。
// Path: /digital/service/cdkCodeReceived
func (s *LazadaDGServiceOp[T]) DigitalServiceCdkCodeReceived(ctx context.Context) (*DigitalServiceCdkCodeReceivedResponse, error) {
	path := "/digital/service/cdkCodeReceived"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(DigitalServiceCdkCodeReceivedResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// InstallServiceCallBack Install the service callback interface
// Path: /digital/install/servicecallback
func (s *LazadaDGServiceOp[T]) InstallServiceCallBack(ctx context.Context) (*InstallServiceCallBackResponse, error) {
	path := "/digital/install/servicecallback"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(InstallServiceCallBackResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// InstallServiceCallBack1 Install the service callback interface
// Path: /digital/test/install/servicecallback
func (s *LazadaDGServiceOp[T]) InstallServiceCallBack1(ctx context.Context) (*InstallServiceCallBack1Response, error) {
	path := "/digital/test/install/servicecallback"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(InstallServiceCallBack1Response)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// InstallServiceCallBackForTest Install the service callback interface
// Path: /digital/install/test/servicecallback
func (s *LazadaDGServiceOp[T]) InstallServiceCallBackForTest(ctx context.Context) (*InstallServiceCallBackForTestResponse, error) {
	path := "/digital/install/test/servicecallback"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(InstallServiceCallBackForTestResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// InuranceNotication Third party insurance company callback interface
//
// Path: /digital/insurance/notification
func (s *LazadaDGServiceOp[T]) InuranceNotication(ctx context.Context) (*InuranceNoticationResponse, error) {
	path := "/digital/insurance/notification"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(InuranceNoticationResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// InuranceNotication1 Third party insurance company callback interface
//
// Path: /digital/insurance/test/notificationcopy
func (s *LazadaDGServiceOp[T]) InuranceNotication1(ctx context.Context) (*InuranceNotication1Response, error) {
	path := "/digital/insurance/test/notificationcopy"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(InuranceNotication1Response)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// InuranceNotifyLapse Insurance company push the callback notification to partners once the policy has been cancelled successfully
// Path: /digital/insurance/notificationlapse
func (s *LazadaDGServiceOp[T]) InuranceNotifyLapse(ctx context.Context) (*InuranceNotifyLapseResponse, error) {
	path := "/digital/insurance/notificationlapse"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(InuranceNotifyLapseResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}
