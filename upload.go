package golazada

import (
	"context"
	"encoding/json"
	"fmt"
)

// UploadImageBytes uploads a single image from raw bytes to the Lazada site.
// Unlike the generated UploadImage method (which streams an io.Reader), this
// takes an in-memory []byte so callers can upload data they already hold.
// Declared here as a package function (not on the ProductService interface) so
// it survives regeneration of the generated files.
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
