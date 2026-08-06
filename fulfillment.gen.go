package golazada

import (
	"context"
	"encoding/json"
	"fmt"
)

type FulfillmentService interface {
	// ConfirmCollectForDBS Use this API to mark an sof order item as being collected.
	// Path: /order/package/sof/collect
	ConfirmCollectForDBS(ctx context.Context) (*ConfirmCollectForDBSResponse, error)
	// ConfirmDeliveryForDBS Use this API to mark an sof order item as being delivered.
	// Path: /order/package/sof/delivered
	ConfirmDeliveryForDBS(ctx context.Context) (*ConfirmDeliveryForDBSResponse, error)
	// DeliverDigital Use this API to mark a digital order item as being delivered.
	// Path: /order/digital/delivered
	DeliverDigital(ctx context.Context) (*DeliverDigitalResponse, error)
	// FailedDeliveryForDBS Use this API to mark an sof order item as being delivered failed
	// Path: /order/package/sof/failed_delivery
	FailedDeliveryForDBS(ctx context.Context) (*FailedDeliveryForDBSResponse, error)
	// GetShipmentProvider Use this API to get the list of all active shipping providers, which is needed when working with the PackOrder API.
	// Path: /order/shipment/providers/get
	GetShipmentProvider(ctx context.Context) (*GetShipmentProviderResponse, error)
	// Pack Use this API to mark an order item as being packed.
	// Path: /order/fulfill/pack
	Pack(ctx context.Context, opt PackRequest) (*PackResponse, error)
	// PackageStatusUpdateForDBS DBS package status update.
	// This interface is only open to some stores
	// Path: /order/package/sof/status/update
	PackageStatusUpdateForDBS(ctx context.Context) (*PackageStatusUpdateForDBSResponse, error)
	// PrintAWB Use this API to retrieve order-related documents, only for shipping labels.
	// Path: /order/package/document/get
	PrintAWB(ctx context.Context, req PrintAWBRequest) (*PrintAWBResponse, error)
	// ReadyToShip Use this API to mark an order item as being ready to ship.
	// Path: /order/rts
	ReadyToShip(ctx context.Context, opt ReadyToShipRequest) (*ReadyToShipResponse, error)
	// RecreatePackage Use this API to mark a package item as being repack.
	// Path: /order/package/repack
	RecreatePackage(ctx context.Context) (*RecreatePackageResponse, error)
}

type FulfillmentServiceOp[T any] struct {
	client *Client[T]
}

// ConfirmCollectForDBS Use this API to mark an sof order item as being collected.
// Path: /order/package/sof/collect
func (s *FulfillmentServiceOp[T]) ConfirmCollectForDBS(ctx context.Context) (*ConfirmCollectForDBSResponse, error) {
	path := "/order/package/sof/collect"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(ConfirmCollectForDBSResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// ConfirmDeliveryForDBS Use this API to mark an sof order item as being delivered.
// Path: /order/package/sof/delivered
func (s *FulfillmentServiceOp[T]) ConfirmDeliveryForDBS(ctx context.Context) (*ConfirmDeliveryForDBSResponse, error) {
	path := "/order/package/sof/delivered"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(ConfirmDeliveryForDBSResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// DeliverDigital Use this API to mark a digital order item as being delivered.
// Path: /order/digital/delivered
func (s *FulfillmentServiceOp[T]) DeliverDigital(ctx context.Context) (*DeliverDigitalResponse, error) {
	path := "/order/digital/delivered"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(DeliverDigitalResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// FailedDeliveryForDBS Use this API to mark an sof order item as being delivered failed
// Path: /order/package/sof/failed_delivery
func (s *FulfillmentServiceOp[T]) FailedDeliveryForDBS(ctx context.Context) (*FailedDeliveryForDBSResponse, error) {
	path := "/order/package/sof/failed_delivery"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(FailedDeliveryForDBSResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetShipmentProvider Use this API to get the list of all active shipping providers, which is needed when working with the PackOrder API.
// Path: /order/shipment/providers/get
func (s *FulfillmentServiceOp[T]) GetShipmentProvider(ctx context.Context) (*GetShipmentProviderResponse, error) {
	path := "/order/shipment/providers/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetShipmentProviderResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// Pack Use this API to mark an order item as being packed.
// Path: /order/pack
func (s *FulfillmentServiceOp[T]) Pack(ctx context.Context, opt PackRequest) (*PackResponse, error) {
	path := "/order/pack"
	params := paramsFromStruct(opt)
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(PackResponse)
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

// PackageStatusUpdateForDBS DBS package status update.
// This interface is only open to some stores
// Path: /order/package/sof/status/update
func (s *FulfillmentServiceOp[T]) PackageStatusUpdateForDBS(ctx context.Context) (*PackageStatusUpdateForDBSResponse, error) {
	path := "/order/package/sof/status/update"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(PackageStatusUpdateForDBSResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// PrintAWB Use this API to retrieve order-related documents, only for shipping labels.
// Path: /order/package/document/get
func (s *FulfillmentServiceOp[T]) PrintAWB(ctx context.Context, req PrintAWBRequest) (*PrintAWBResponse, error) {
	path := "/order/package/document/get"
	params := paramsFromStruct(req)
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(PrintAWBResponse)
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

// ReadyToShip Use this API to mark an order item as being ready to ship.
// Path: /order/rts
func (s *FulfillmentServiceOp[T]) ReadyToShip(ctx context.Context, opt ReadyToShipRequest) (*ReadyToShipResponse, error) {
	path := "/order/rts"
	params := paramsFromStruct(opt)
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(ReadyToShipResponse)
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

// RecreatePackage Use this API to mark a package item as being repack.
// Path: /order/package/repack
func (s *FulfillmentServiceOp[T]) RecreatePackage(ctx context.Context) (*RecreatePackageResponse, error) {
	path := "/order/package/repack"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(RecreatePackageResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}
