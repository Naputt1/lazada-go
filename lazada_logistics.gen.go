package golazada

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
)

type LazadaLogisticsService interface {
	// CreateCustomerAccountRelationshipByOTP Create customer account relationship for external by OTP
	// Path: /logistics/epis/customers/external_relationships_bundle
	CreateCustomerAccountRelationshipByOTP(ctx context.Context) (*CreateCustomerAccountRelationshipByOTPResponse, error)
	// CreateCustomerAccountRelationshipForExternal External partner calls LAZADA to create account relationship
	// Path: /logistics/epis/customers/external_relationships
	CreateCustomerAccountRelationshipForExternal(ctx context.Context) (*CreateCustomerAccountRelationshipForExternalResponse, error)
	// CreateOrUpdateCustomerWarehouse External partner calls LAZADA to create or update warehouses
	// Path: /logistics/epis/customers/warehouses
	CreateOrUpdateCustomerWarehouse(ctx context.Context) (*CreateOrUpdateCustomerWarehouseResponse, error)
	// EpisGetDeliveryOptions External partner call EPIS to get delivery options for package
	// Path: /logistics/epis/service/delivery_options
	EpisGetDeliveryOptions(ctx context.Context) (*EpisGetDeliveryOptionsResponse, error)
	// EpisPackageCancellation External partner call EPIS to cancel package
	// Path: /logistics/epis/packages/cancel
	EpisPackageCancellation(ctx context.Context) (*EpisPackageCancellationResponse, error)
	// EpisPackageCancellationV3 External partner call EPIS to cancel FFM + DEL package
	// Path: /logistics/epis/packages/cancel/v3
	EpisPackageCancellationV3(ctx context.Context) (*EpisPackageCancellationV3Response, error)
	// EpisPackageConsignment External partner call EPIS to consign package to get the tracking number and be able to print AWB after consign
	// Path: /logistics/epis/packages/consign
	EpisPackageConsignment(ctx context.Context) (*EpisPackageConsignmentResponse, error)
	// EpisPackageConsignmentV2 External partner call EPIS to consign FFM + DEL package to get the tracking number and be able to print AWB after consign
	// Path: /logistics/epis/packages/consign/v2
	EpisPackageConsignmentV2(ctx context.Context) (*EpisPackageConsignmentV2Response, error)
	// EpisPackageCreation External partner call EPIS to create package
	// Path: /logistics/epis/packages
	EpisPackageCreation(ctx context.Context) (*EpisPackageCreationResponse, error)
	// EpisPackageInfoUpdate External partner call EPIS to update package info after RTS
	// Path: /logistics/epis/packages/update
	EpisPackageInfoUpdate(ctx context.Context) (*EpisPackageInfoUpdateResponse, error)
	// EpisPackagePrintAwb External partner call LAZADA to print AWB
	// Path: /logistics/epis/packages/awb
	EpisPackagePrintAwb(ctx context.Context) (*EpisPackagePrintAwbResponse, error)
	// EpisPackageReadyToBeShipped External partner calls EPIS to mark a package as ready to be shipped
	// Path: /logistics/epis/packages/rts
	EpisPackageReadyToBeShipped(ctx context.Context) (*EpisPackageReadyToBeShippedResponse, error)
	// EpisPackageReAttempt Send re-attempt package request
	// Path: /logistics/epis/packages/reattempt
	EpisPackageReAttempt(ctx context.Context) (*EpisPackageReAttemptResponse, error)
	// EpisUploadAwbFulfillment External partner call EPIS to upload awb for fulfillment
	// Path: /logistics/epis/fulfillment/upload_awb
	EpisUploadAwbFulfillment(ctx context.Context, filename string, reader io.Reader) (*EpisUploadAwbFulfillmentResponse, error)
	// EpisXspaceCreate Create Xspace case
	// Path: /logistics/epis/xspace/create
	EpisXspaceCreate(ctx context.Context) (*EpisXspaceCreateResponse, error)
	// EpisXspaceGetDetail Get Xspace case detail
	// Path: /logistics/epis/xspace/detail
	EpisXspaceGetDetail(ctx context.Context) (*EpisXspaceGetDetailResponse, error)
	// EpisXspaceQuery Query Xspace case
	// Path: /logistics/epis/xspace/query
	EpisXspaceQuery(ctx context.Context) (*EpisXspaceQueryResponse, error)
	// EpisXspaceRateTicket Rate Xspace ticket
	// Path: /logistics/epis/xspace/rate
	EpisXspaceRateTicket(ctx context.Context) (*EpisXspaceRateTicketResponse, error)
	// EstimateShippingFee Estimate shipping fee
	// Path: /logistics/epis/estimate_shipping_fee
	EstimateShippingFee(ctx context.Context) (*EstimateShippingFeeResponse, error)
	// GetShippingFee Estimate package shipping fee (Estimated & Actual)
	// Path: /logistics/epis/get_shipping_fee
	GetShippingFee(ctx context.Context) (*GetShippingFeeResponse, error)
}

type LazadaLogisticsServiceOp[T any] struct {
	client *Client[T]
}

// CreateCustomerAccountRelationshipByOTP Create customer account relationship for external by OTP
// Path: /logistics/epis/customers/external_relationships_bundle
func (s *LazadaLogisticsServiceOp[T]) CreateCustomerAccountRelationshipByOTP(ctx context.Context) (*CreateCustomerAccountRelationshipByOTPResponse, error) {
	path := "/logistics/epis/customers/external_relationships_bundle"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(CreateCustomerAccountRelationshipByOTPResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// CreateCustomerAccountRelationshipForExternal External partner calls LAZADA to create account relationship
// Path: /logistics/epis/customers/external_relationships
func (s *LazadaLogisticsServiceOp[T]) CreateCustomerAccountRelationshipForExternal(ctx context.Context) (*CreateCustomerAccountRelationshipForExternalResponse, error) {
	path := "/logistics/epis/customers/external_relationships"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(CreateCustomerAccountRelationshipForExternalResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// CreateOrUpdateCustomerWarehouse External partner calls LAZADA to create or update warehouses
// Path: /logistics/epis/customers/warehouses
func (s *LazadaLogisticsServiceOp[T]) CreateOrUpdateCustomerWarehouse(ctx context.Context) (*CreateOrUpdateCustomerWarehouseResponse, error) {
	path := "/logistics/epis/customers/warehouses"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(CreateOrUpdateCustomerWarehouseResponse)
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

// EpisGetDeliveryOptions External partner call EPIS to get delivery options for package
// Path: /logistics/epis/service/delivery_options
func (s *LazadaLogisticsServiceOp[T]) EpisGetDeliveryOptions(ctx context.Context) (*EpisGetDeliveryOptionsResponse, error) {
	path := "/logistics/epis/service/delivery_options"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(EpisGetDeliveryOptionsResponse)
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

// EpisPackageCancellation External partner call EPIS to cancel package
// Path: /logistics/epis/packages/cancel
func (s *LazadaLogisticsServiceOp[T]) EpisPackageCancellation(ctx context.Context) (*EpisPackageCancellationResponse, error) {
	path := "/logistics/epis/packages/cancel"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(EpisPackageCancellationResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// EpisPackageCancellationV3 External partner call EPIS to cancel FFM + DEL package
// Path: /logistics/epis/packages/cancel/v3
func (s *LazadaLogisticsServiceOp[T]) EpisPackageCancellationV3(ctx context.Context) (*EpisPackageCancellationV3Response, error) {
	path := "/logistics/epis/packages/cancel/v3"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(EpisPackageCancellationV3Response)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// EpisPackageConsignment External partner call EPIS to consign package to get the tracking number and be able to print AWB after consign
// Path: /logistics/epis/packages/consign
func (s *LazadaLogisticsServiceOp[T]) EpisPackageConsignment(ctx context.Context) (*EpisPackageConsignmentResponse, error) {
	path := "/logistics/epis/packages/consign"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(EpisPackageConsignmentResponse)
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

// EpisPackageConsignmentV2 External partner call EPIS to consign FFM + DEL package to get the tracking number and be able to print AWB after consign
// Path: /logistics/epis/packages/consign/v2
func (s *LazadaLogisticsServiceOp[T]) EpisPackageConsignmentV2(ctx context.Context) (*EpisPackageConsignmentV2Response, error) {
	path := "/logistics/epis/packages/consign/v2"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(EpisPackageConsignmentV2Response)
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

// EpisPackageCreation External partner call EPIS to create package
// Path: /logistics/epis/packages
func (s *LazadaLogisticsServiceOp[T]) EpisPackageCreation(ctx context.Context) (*EpisPackageCreationResponse, error) {
	path := "/logistics/epis/packages"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(EpisPackageCreationResponse)
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

// EpisPackageInfoUpdate External partner call EPIS to update package info after RTS
// Path: /logistics/epis/packages/update
func (s *LazadaLogisticsServiceOp[T]) EpisPackageInfoUpdate(ctx context.Context) (*EpisPackageInfoUpdateResponse, error) {
	path := "/logistics/epis/packages/update"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(EpisPackageInfoUpdateResponse)
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

// EpisPackagePrintAwb External partner call LAZADA to print AWB
// Path: /logistics/epis/packages/awb
func (s *LazadaLogisticsServiceOp[T]) EpisPackagePrintAwb(ctx context.Context) (*EpisPackagePrintAwbResponse, error) {
	path := "/logistics/epis/packages/awb"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(EpisPackagePrintAwbResponse)
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

// EpisPackageReadyToBeShipped External partner calls EPIS to mark a package as ready to be shipped
// Path: /logistics/epis/packages/rts
func (s *LazadaLogisticsServiceOp[T]) EpisPackageReadyToBeShipped(ctx context.Context) (*EpisPackageReadyToBeShippedResponse, error) {
	path := "/logistics/epis/packages/rts"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(EpisPackageReadyToBeShippedResponse)
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

// EpisPackageReAttempt Send re-attempt package request
// Path: /logistics/epis/packages/reattempt
func (s *LazadaLogisticsServiceOp[T]) EpisPackageReAttempt(ctx context.Context) (*EpisPackageReAttemptResponse, error) {
	path := "/logistics/epis/packages/reattempt"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(EpisPackageReAttemptResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// EpisUploadAwbFulfillment External partner call EPIS to upload awb for fulfillment
// Path: /logistics/epis/fulfillment/upload_awb
func (s *LazadaLogisticsServiceOp[T]) EpisUploadAwbFulfillment(ctx context.Context, filename string, reader io.Reader) (*EpisUploadAwbFulfillmentResponse, error) {
	path := "/logistics/epis/fulfillment/upload_awb"
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", filename)
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(part, reader); err != nil {
		return nil, err
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}

	wrapper, err := s.client.execute(ctx, "POST", path, nil, map[string][]byte{"image": {}})
	if err != nil {
		return nil, err
	}
	resp := new(EpisUploadAwbFulfillmentResponse)
	if err := json.Unmarshal(wrapper.Data, resp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return resp, nil
}

// EpisXspaceCreate Create Xspace case
// Path: /logistics/epis/xspace/create
func (s *LazadaLogisticsServiceOp[T]) EpisXspaceCreate(ctx context.Context) (*EpisXspaceCreateResponse, error) {
	path := "/logistics/epis/xspace/create"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(EpisXspaceCreateResponse)
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

// EpisXspaceGetDetail Get Xspace case detail
// Path: /logistics/epis/xspace/detail
func (s *LazadaLogisticsServiceOp[T]) EpisXspaceGetDetail(ctx context.Context) (*EpisXspaceGetDetailResponse, error) {
	path := "/logistics/epis/xspace/detail"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(EpisXspaceGetDetailResponse)
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

// EpisXspaceQuery Query Xspace case
// Path: /logistics/epis/xspace/query
func (s *LazadaLogisticsServiceOp[T]) EpisXspaceQuery(ctx context.Context) (*EpisXspaceQueryResponse, error) {
	path := "/logistics/epis/xspace/query"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(EpisXspaceQueryResponse)
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

// EpisXspaceRateTicket Rate Xspace ticket
// Path: /logistics/epis/xspace/rate
func (s *LazadaLogisticsServiceOp[T]) EpisXspaceRateTicket(ctx context.Context) (*EpisXspaceRateTicketResponse, error) {
	path := "/logistics/epis/xspace/rate"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(EpisXspaceRateTicketResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// EstimateShippingFee Estimate shipping fee
// Path: /logistics/epis/estimate_shipping_fee
func (s *LazadaLogisticsServiceOp[T]) EstimateShippingFee(ctx context.Context) (*EstimateShippingFeeResponse, error) {
	path := "/logistics/epis/estimate_shipping_fee"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(EstimateShippingFeeResponse)
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

// GetShippingFee Estimate package shipping fee (Estimated & Actual)
// Path: /logistics/epis/get_shipping_fee
func (s *LazadaLogisticsServiceOp[T]) GetShippingFee(ctx context.Context) (*GetShippingFeeResponse, error) {
	path := "/logistics/epis/get_shipping_fee"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetShippingFeeResponse)
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
