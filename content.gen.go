package golazada

import (
	"context"
)

type ContentService interface {
	// CancelTask cancel tasks
	// Path: /content/ai/cancelTask
	CancelTask(ctx context.Context) (*CancelTaskResponse, error)
	// ChangeFace change face using lazada AI algorithm
	// Path: /content/ai/changeFace
	ChangeFace(ctx context.Context) (*ChangeFaceResponse, error)
	// ChangeProductBackground change product background using lazada AI algorithm
	// Path: /content/ai/changeProductBackground
	ChangeProductBackground(ctx context.Context) (*ChangeProductBackgroundResponse, error)
	// FixHand fixHand using lazada AI algorithm
	// Path: /content/ai/fixHand
	FixHand(ctx context.Context) (*FixHandResponse, error)
	// GetTaskStatus get task status
	// Path: /content/ai/getTaskStatus
	GetTaskStatus(ctx context.Context) (*GetTaskStatusResponse, error)
	// ProductImageMatch match product image
	// Path: /content/ai/productImageMatch
	ProductImageMatch(ctx context.Context) (*ProductImageMatchResponse, error)
	// TryOnCloth try on cloth using lazada AI algorithm
	// Path: /content/ai/tryOnCloth
	TryOnCloth(ctx context.Context) (*TryOnClothResponse, error)
}

type ContentServiceOp[T any] struct {
	client *Client[T]
}

// CancelTask cancel tasks
// Path: /content/ai/cancelTask
func (s *ContentServiceOp[T]) CancelTask(ctx context.Context) (*CancelTaskResponse, error) {
	path := "/content/ai/cancelTask"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(CancelTaskResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// ChangeFace change face using lazada AI algorithm
// Path: /content/ai/changeFace
func (s *ContentServiceOp[T]) ChangeFace(ctx context.Context) (*ChangeFaceResponse, error) {
	path := "/content/ai/changeFace"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(ChangeFaceResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// ChangeProductBackground change product background using lazada AI algorithm
// Path: /content/ai/changeProductBackground
func (s *ContentServiceOp[T]) ChangeProductBackground(ctx context.Context) (*ChangeProductBackgroundResponse, error) {
	path := "/content/ai/changeProductBackground"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(ChangeProductBackgroundResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// FixHand fixHand using lazada AI algorithm
// Path: /content/ai/fixHand
func (s *ContentServiceOp[T]) FixHand(ctx context.Context) (*FixHandResponse, error) {
	path := "/content/ai/fixHand"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(FixHandResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetTaskStatus get task status
// Path: /content/ai/getTaskStatus
func (s *ContentServiceOp[T]) GetTaskStatus(ctx context.Context) (*GetTaskStatusResponse, error) {
	path := "/content/ai/getTaskStatus"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetTaskStatusResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// ProductImageMatch match product image
// Path: /content/ai/productImageMatch
func (s *ContentServiceOp[T]) ProductImageMatch(ctx context.Context) (*ProductImageMatchResponse, error) {
	path := "/content/ai/productImageMatch"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(ProductImageMatchResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// TryOnCloth try on cloth using lazada AI algorithm
// Path: /content/ai/tryOnCloth
func (s *ContentServiceOp[T]) TryOnCloth(ctx context.Context) (*TryOnClothResponse, error) {
	path := "/content/ai/tryOnCloth"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(TryOnClothResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}
