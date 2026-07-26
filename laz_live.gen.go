package golazada

import (
	"context"
	"encoding/json"
	"fmt"
)

type LazLiveService interface {
	// HighlightProduct highlight product
	// Path: /lazlive/product/highlight
	HighlightProduct(ctx context.Context) (*HighlightProductResponse, error)
}

type LazLiveServiceOp[T any] struct {
	client *Client[T]
}

// HighlightProduct highlight product
// Path: /lazlive/product/highlight
func (s *LazLiveServiceOp[T]) HighlightProduct(ctx context.Context) (*HighlightProductResponse, error) {
	path := "/lazlive/product/highlight"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(HighlightProductResponse)
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
