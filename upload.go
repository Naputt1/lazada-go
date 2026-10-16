package golazada

import (
	"context"
	"encoding/json"
	"fmt"
)

// UploadImageBytes uploads an image from raw bytes. Unlike the generated UploadImage
// (io.Reader), it takes an in-memory []byte. Declared here as a package function so it
// survives regeneration of the generated files.
func UploadImageBytes[T any](ctx context.Context, client *Client[T], filename string, data []byte) (*UploadImageResponse, error) {
	path := "/image/upload"
	wrapper, err := client.Post(ctx, path, nil, map[string][]byte{filename: data})
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
