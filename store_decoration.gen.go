package golazada

import (
	"context"
	"encoding/json"
	"fmt"
)

type StoreDecorationService interface {
	// GetStoreCustomPage GetStoreCustomPagevice
	//
	// Path: /store/custom/page/get
	GetStoreCustomPage(ctx context.Context) (*GetStoreCustomPageResponse, error)
}

type StoreDecorationServiceOp[T any] struct {
	client *Client[T]
}

// GetStoreCustomPage GetStoreCustomPagevice
//
// Path: /store/custom/page/get
func (s *StoreDecorationServiceOp[T]) GetStoreCustomPage(ctx context.Context) (*GetStoreCustomPageResponse, error) {
	path := "/store/custom/page/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetStoreCustomPageResponse)
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
