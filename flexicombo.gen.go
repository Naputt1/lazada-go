package golazada

import (
	"context"
	"encoding/json"
	"fmt"
)

type FlexicomboService interface {
	// ActivateFlexiCombo activate flexi combo
	// Path: /promotion/flexicombo/activate
	ActivateFlexiCombo(ctx context.Context) (*ActivateFlexiComboResponse, error)
	// AddFlexiComboProducts add flexi combo products
	// Path: /promotion/flexicombo/products/add
	AddFlexiComboProducts(ctx context.Context) (*AddFlexiComboProductsResponse, error)
	// CreateFlexiCombo create a  new promotion flexi combo
	// Path: /promotion/flexicombo/create
	CreateFlexiCombo(ctx context.Context) (*CreateFlexiComboResponse, error)
	// DeactivateFlexiCombo deactivate flexi combo
	// Path: /promotion/flexicombo/deactivate
	DeactivateFlexiCombo(ctx context.Context) (*DeactivateFlexiComboResponse, error)
	// DeleteFlexiComboProducts delete flexi combo products
	// Path: /promotion/flexicombo/products/delete
	DeleteFlexiComboProducts(ctx context.Context) (*DeleteFlexiComboProductsResponse, error)
	// GetFlexiComboDetails get promotion flexi combo detail by id
	// Path: /promotion/flexicombo/details
	GetFlexiComboDetails(ctx context.Context) (*GetFlexiComboDetailsResponse, error)
	// ListFlexiCombo list flexi combo
	// Path: /promotion/flexicombo/list
	ListFlexiCombo(ctx context.Context) (*ListFlexiComboResponse, error)
	// ListFlexiComboProducts list flexi combo products
	// Path: /promotion/flexicombo/products/list
	ListFlexiComboProducts(ctx context.Context) (*ListFlexiComboProductsResponse, error)
	// UpdateFlexiCombo update flexi combo
	// Path: /promotion/flexicombo/update
	UpdateFlexiCombo(ctx context.Context) (*UpdateFlexiComboResponse, error)
}

type FlexicomboServiceOp[T any] struct {
	client *Client[T]
}

// ActivateFlexiCombo activate flexi combo
// Path: /promotion/flexicombo/activate
func (s *FlexicomboServiceOp[T]) ActivateFlexiCombo(ctx context.Context) (*ActivateFlexiComboResponse, error) {
	path := "/promotion/flexicombo/activate"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(ActivateFlexiComboResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// AddFlexiComboProducts add flexi combo products
// Path: /promotion/flexicombo/products/add
func (s *FlexicomboServiceOp[T]) AddFlexiComboProducts(ctx context.Context) (*AddFlexiComboProductsResponse, error) {
	path := "/promotion/flexicombo/products/add"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(AddFlexiComboProductsResponse)
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

// CreateFlexiCombo create a  new promotion flexi combo
// Path: /promotion/flexicombo/create
func (s *FlexicomboServiceOp[T]) CreateFlexiCombo(ctx context.Context) (*CreateFlexiComboResponse, error) {
	path := "/promotion/flexicombo/create"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(CreateFlexiComboResponse)
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

// DeactivateFlexiCombo deactivate flexi combo
// Path: /promotion/flexicombo/deactivate
func (s *FlexicomboServiceOp[T]) DeactivateFlexiCombo(ctx context.Context) (*DeactivateFlexiComboResponse, error) {
	path := "/promotion/flexicombo/deactivate"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(DeactivateFlexiComboResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// DeleteFlexiComboProducts delete flexi combo products
// Path: /promotion/flexicombo/products/delete
func (s *FlexicomboServiceOp[T]) DeleteFlexiComboProducts(ctx context.Context) (*DeleteFlexiComboProductsResponse, error) {
	path := "/promotion/flexicombo/products/delete"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(DeleteFlexiComboProductsResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetFlexiComboDetails get promotion flexi combo detail by id
// Path: /promotion/flexicombo/details
func (s *FlexicomboServiceOp[T]) GetFlexiComboDetails(ctx context.Context) (*GetFlexiComboDetailsResponse, error) {
	path := "/promotion/flexicombo/details"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetFlexiComboDetailsResponse)
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

// ListFlexiCombo list flexi combo
// Path: /promotion/flexicombo/list
func (s *FlexicomboServiceOp[T]) ListFlexiCombo(ctx context.Context) (*ListFlexiComboResponse, error) {
	path := "/promotion/flexicombo/list"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(ListFlexiComboResponse)
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

// ListFlexiComboProducts list flexi combo products
// Path: /promotion/flexicombo/products/list
func (s *FlexicomboServiceOp[T]) ListFlexiComboProducts(ctx context.Context) (*ListFlexiComboProductsResponse, error) {
	path := "/promotion/flexicombo/products/list"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(ListFlexiComboProductsResponse)
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

// UpdateFlexiCombo update flexi combo
// Path: /promotion/flexicombo/update
func (s *FlexicomboServiceOp[T]) UpdateFlexiCombo(ctx context.Context) (*UpdateFlexiComboResponse, error) {
	path := "/promotion/flexicombo/update"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(UpdateFlexiComboResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}
