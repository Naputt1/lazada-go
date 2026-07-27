package golazada

import (
	"context"
	"encoding/json"
	"fmt"
)

func (s *ProductServiceOp[T]) UploadImageBytes(ctx context.Context, filename string, data []byte) (*UploadImageResponse, error) {
	path := "/image/upload"
	wrapper, err := s.client.Post(ctx, path, nil, map[string][]byte{filename: data})
	if err != nil {
		return nil, err
	}
	resp := new(UploadImageResponse)
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
