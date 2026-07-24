package golazada

import (
	"context"
)

type LogisticsStationService interface {
	// CageValidation Validate if a cage is valid
	// Path: /logistics/station/cages/validate
	CageValidation(ctx context.Context) (*CageValidationResponse, error)
	// ConfirmInbound Confirm inbound. Call this API to inbound the scanned parcel and finish the inbound process
	// Path: /logistics/station/v1/confirm-inbound
	ConfirmInbound(ctx context.Context) (*ConfirmInboundResponse, error)
	// ConfirmParcelCollection Confirm customer collects or rejects parcel. This API is used after ValidateOTP success.
	// Path: /logistics/station/v1/cp/confirm-parcel-collection
	ConfirmParcelCollection(ctx context.Context) (*ConfirmParcelCollectionResponse, error)
	// CreateScannedParcel Create a scanned parcel. Call this API when scanning the tracking number on the parcel.
	// Path: /logistics/station/v1/scanned-parcels/create
	CreateScannedParcel(ctx context.Context) (*CreateScannedParcelResponse, error)
	// DeleteScannedParcel Delete scanned parcels by tracking number. This API is required when user deletes the scanned parcels
	// Path: /logistics/station/v1/scanned-parcels/delete
	DeleteScannedParcel(ctx context.Context) (*DeleteScannedParcelResponse, error)
	// DopConfirmInbound DOP confirm inbound
	// Path: /logistics/station/dop/confirm-inbound
	DopConfirmInbound(ctx context.Context) (*DopConfirmInboundResponse, error)
	// DopCreateScannedParcel DOP create scanned parcel
	// Path: /logistics/station/dop/scanned-parcels
	DopCreateScannedParcel(ctx context.Context) (*DopCreateScannedParcelResponse, error)
	// DopDeleteScannedParcel DOP delete scanned parcel
	// Path: /logistics/station/dop/scanned-parcels/delete
	DopDeleteScannedParcel(ctx context.Context) (*DopDeleteScannedParcelResponse, error)
	// DopGetInboundedParcel DOP get list scanned parcel
	// Path: /logistics/station/dop/inbounded-parcels/list
	DopGetInboundedParcel(ctx context.Context) (*DopGetInboundedParcelResponse, error)
	// DopGetScannedParcel DOP get list scanned parcel
	// Path: /logistics/station/dop/scanned-parcels/list
	DopGetScannedParcel(ctx context.Context) (*DopGetScannedParcelResponse, error)
	// GetCpScheduledPuParcel Get a list of parcels that are scheduled to be picked up for return to seller. These parcels are expired (no collection from customer), SLA breached or customer rejected. This API is used to help the agent prepare parcels before seller comes.
	// Path: /logistics/station/v1/cp/scheduled-pu-parcels/list
	GetCpScheduledPuParcel(ctx context.Context) (*GetCpScheduledPuParcelResponse, error)
	// GetInboundedParcel Get a list of inbounded parcels by a list of tracking numbers. This API is used for checking the status of inbounded parcels such as parcels picked up by LEX, picked up by 3PL, or collected by a customer.
	// Path: /logistics/station/v1/inbounded-parcels/list
	GetInboundedParcel(ctx context.Context) (*GetInboundedParcelResponse, error)
	// GetListAccessStation Get list access station by APP
	// Path: /logistics/station/list
	GetListAccessStation(ctx context.Context) (*GetListAccessStationResponse, error)
	// GetMetaData Get metadata such as reject reasons, etc
	// Path: /logistics/station/v1/metadata
	GetMetaData(ctx context.Context) (*GetMetaDataResponse, error)
	// GetScannedParcel Get a list of scanned parcels. This API is often used for synchronization purposes such as: user refreshes the page, partner system can call this API to get the list of scanned parcels again. This API is not required to call during operations.
	// Path: /logistics/station/v1/scanned-parcels/list
	GetScannedParcel(ctx context.Context) (*GetScannedParcelResponse, error)
	// SearchCustomerReturnParcel Search customer return parcel by at least 4 letters text. This API is to improve user experience, user can search for the tracking number instead of typing the full tracking number.
	// Path: /logistics/station/v1/dop/cr-parcels/search
	SearchCustomerReturnParcel(ctx context.Context) (*SearchCustomerReturnParcelResponse, error)
	// ValidateCage Validate if a cage is valid. This API is often called before starting inbound but it's not required.
	// Path: /logistics/station/v1/cages/validate
	ValidateCage(ctx context.Context) (*ValidateCageResponse, error)
	// ValidateOTP Validate if OTP of parcel is valid or not. This API is used for checking OTP before confirming collection.
	// Path: /logistics/station/v1/cp/validate-otp
	ValidateOTP(ctx context.Context) (*ValidateOTPResponse, error)
}

type LogisticsStationServiceOp[T any] struct {
	client *Client[T]
}

// CageValidation Validate if a cage is valid
// Path: /logistics/station/cages/validate
func (s *LogisticsStationServiceOp[T]) CageValidation(ctx context.Context) (*CageValidationResponse, error) {
	path := "/logistics/station/cages/validate"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(CageValidationResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// ConfirmInbound Confirm inbound. Call this API to inbound the scanned parcel and finish the inbound process
// Path: /logistics/station/v1/confirm-inbound
func (s *LogisticsStationServiceOp[T]) ConfirmInbound(ctx context.Context) (*ConfirmInboundResponse, error) {
	path := "/logistics/station/v1/confirm-inbound"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(ConfirmInboundResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// ConfirmParcelCollection Confirm customer collects or rejects parcel. This API is used after ValidateOTP success.
// Path: /logistics/station/v1/cp/confirm-parcel-collection
func (s *LogisticsStationServiceOp[T]) ConfirmParcelCollection(ctx context.Context) (*ConfirmParcelCollectionResponse, error) {
	path := "/logistics/station/v1/cp/confirm-parcel-collection"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(ConfirmParcelCollectionResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// CreateScannedParcel Create a scanned parcel. Call this API when scanning the tracking number on the parcel.
// Path: /logistics/station/v1/scanned-parcels/create
func (s *LogisticsStationServiceOp[T]) CreateScannedParcel(ctx context.Context) (*CreateScannedParcelResponse, error) {
	path := "/logistics/station/v1/scanned-parcels/create"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(CreateScannedParcelResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// DeleteScannedParcel Delete scanned parcels by tracking number. This API is required when user deletes the scanned parcels
// Path: /logistics/station/v1/scanned-parcels/delete
func (s *LogisticsStationServiceOp[T]) DeleteScannedParcel(ctx context.Context) (*DeleteScannedParcelResponse, error) {
	path := "/logistics/station/v1/scanned-parcels/delete"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(DeleteScannedParcelResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// DopConfirmInbound DOP confirm inbound
// Path: /logistics/station/dop/confirm-inbound
func (s *LogisticsStationServiceOp[T]) DopConfirmInbound(ctx context.Context) (*DopConfirmInboundResponse, error) {
	path := "/logistics/station/dop/confirm-inbound"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(DopConfirmInboundResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// DopCreateScannedParcel DOP create scanned parcel
// Path: /logistics/station/dop/scanned-parcels
func (s *LogisticsStationServiceOp[T]) DopCreateScannedParcel(ctx context.Context) (*DopCreateScannedParcelResponse, error) {
	path := "/logistics/station/dop/scanned-parcels"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(DopCreateScannedParcelResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// DopDeleteScannedParcel DOP delete scanned parcel
// Path: /logistics/station/dop/scanned-parcels/delete
func (s *LogisticsStationServiceOp[T]) DopDeleteScannedParcel(ctx context.Context) (*DopDeleteScannedParcelResponse, error) {
	path := "/logistics/station/dop/scanned-parcels/delete"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(DopDeleteScannedParcelResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// DopGetInboundedParcel DOP get list scanned parcel
// Path: /logistics/station/dop/inbounded-parcels/list
func (s *LogisticsStationServiceOp[T]) DopGetInboundedParcel(ctx context.Context) (*DopGetInboundedParcelResponse, error) {
	path := "/logistics/station/dop/inbounded-parcels/list"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(DopGetInboundedParcelResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// DopGetScannedParcel DOP get list scanned parcel
// Path: /logistics/station/dop/scanned-parcels/list
func (s *LogisticsStationServiceOp[T]) DopGetScannedParcel(ctx context.Context) (*DopGetScannedParcelResponse, error) {
	path := "/logistics/station/dop/scanned-parcels/list"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(DopGetScannedParcelResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetCpScheduledPuParcel Get a list of parcels that are scheduled to be picked up for return to seller. These parcels are expired (no collection from customer), SLA breached or customer rejected. This API is used to help the agent prepare parcels before seller comes.
// Path: /logistics/station/v1/cp/scheduled-pu-parcels/list
func (s *LogisticsStationServiceOp[T]) GetCpScheduledPuParcel(ctx context.Context) (*GetCpScheduledPuParcelResponse, error) {
	path := "/logistics/station/v1/cp/scheduled-pu-parcels/list"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetCpScheduledPuParcelResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetInboundedParcel Get a list of inbounded parcels by a list of tracking numbers. This API is used for checking the status of inbounded parcels such as parcels picked up by LEX, picked up by 3PL, or collected by a customer.
// Path: /logistics/station/v1/inbounded-parcels/list
func (s *LogisticsStationServiceOp[T]) GetInboundedParcel(ctx context.Context) (*GetInboundedParcelResponse, error) {
	path := "/logistics/station/v1/inbounded-parcels/list"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetInboundedParcelResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetListAccessStation Get list access station by APP
// Path: /logistics/station/list
func (s *LogisticsStationServiceOp[T]) GetListAccessStation(ctx context.Context) (*GetListAccessStationResponse, error) {
	path := "/logistics/station/list"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetListAccessStationResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetMetaData Get metadata such as reject reasons, etc
// Path: /logistics/station/v1/metadata
func (s *LogisticsStationServiceOp[T]) GetMetaData(ctx context.Context) (*GetMetaDataResponse, error) {
	path := "/logistics/station/v1/metadata"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetMetaDataResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetScannedParcel Get a list of scanned parcels. This API is often used for synchronization purposes such as: user refreshes the page, partner system can call this API to get the list of scanned parcels again. This API is not required to call during operations.
// Path: /logistics/station/v1/scanned-parcels/list
func (s *LogisticsStationServiceOp[T]) GetScannedParcel(ctx context.Context) (*GetScannedParcelResponse, error) {
	path := "/logistics/station/v1/scanned-parcels/list"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetScannedParcelResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// SearchCustomerReturnParcel Search customer return parcel by at least 4 letters text. This API is to improve user experience, user can search for the tracking number instead of typing the full tracking number.
// Path: /logistics/station/v1/dop/cr-parcels/search
func (s *LogisticsStationServiceOp[T]) SearchCustomerReturnParcel(ctx context.Context) (*SearchCustomerReturnParcelResponse, error) {
	path := "/logistics/station/v1/dop/cr-parcels/search"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(SearchCustomerReturnParcelResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// ValidateCage Validate if a cage is valid. This API is often called before starting inbound but it's not required.
// Path: /logistics/station/v1/cages/validate
func (s *LogisticsStationServiceOp[T]) ValidateCage(ctx context.Context) (*ValidateCageResponse, error) {
	path := "/logistics/station/v1/cages/validate"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(ValidateCageResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// ValidateOTP Validate if OTP of parcel is valid or not. This API is used for checking OTP before confirming collection.
// Path: /logistics/station/v1/cp/validate-otp
func (s *LogisticsStationServiceOp[T]) ValidateOTP(ctx context.Context) (*ValidateOTPResponse, error) {
	path := "/logistics/station/v1/cp/validate-otp"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(ValidateOTPResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}
