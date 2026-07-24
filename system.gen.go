package golazada

import (
	"context"
)

type SystemService interface {
	// StartExportByDataset Open the download operation
	// Path: /fbi/download/startExportByDataset
	StartExportByDataset(ctx context.Context) (*StartExportByDatasetResponse, error)
}

type SystemServiceOp[T any] struct {
	client *Client[T]
}

// StartExportByDataset Open the download operation
// Path: /fbi/download/startExportByDataset
func (s *SystemServiceOp[T]) StartExportByDataset(ctx context.Context) (*StartExportByDatasetResponse, error) {
	path := "/fbi/download/startExportByDataset"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(StartExportByDatasetResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}
