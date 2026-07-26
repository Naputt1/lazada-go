package golazada

import (
	"context"
	"encoding/json"
	"fmt"
)

type InstantMessagingService interface {
	// GetMessages Get message list
	// Path: /im/message/list
	GetMessages(ctx context.Context) (*GetMessagesResponse, error)
	// GetSessionDetail get session detail by sessionid
	// Path: /im/session/get
	GetSessionDetail(ctx context.Context) (*GetSessionDetailResponse, error)
	// GetSessionList query seller session list
	// Path: /im/session/list
	GetSessionList(ctx context.Context) (*GetSessionListResponse, error)
	// MessageRecall message recall
	// Path: /im/message/recall
	MessageRecall(ctx context.Context) (*MessageRecallResponse, error)
	// OpenSession open a new conversation
	// Path: /im/session/open
	OpenSession(ctx context.Context) (*OpenSessionResponse, error)
	// ReadSession session read
	// Path: /im/session/read
	ReadSession(ctx context.Context) (*ReadSessionResponse, error)
	// SendMessage send message
	// Path: /im/message/send
	SendMessage(ctx context.Context) (*SendMessageResponse, error)
}

type InstantMessagingServiceOp[T any] struct {
	client *Client[T]
}

// GetMessages Get message list
// Path: /im/message/list
func (s *InstantMessagingServiceOp[T]) GetMessages(ctx context.Context) (*GetMessagesResponse, error) {
	path := "/im/message/list"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetMessagesResponse)
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

// GetSessionDetail get session detail by sessionid
// Path: /im/session/get
func (s *InstantMessagingServiceOp[T]) GetSessionDetail(ctx context.Context) (*GetSessionDetailResponse, error) {
	path := "/im/session/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetSessionDetailResponse)
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

// GetSessionList query seller session list
// Path: /im/session/list
func (s *InstantMessagingServiceOp[T]) GetSessionList(ctx context.Context) (*GetSessionListResponse, error) {
	path := "/im/session/list"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetSessionListResponse)
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

// MessageRecall message recall
// Path: /im/message/recall
func (s *InstantMessagingServiceOp[T]) MessageRecall(ctx context.Context) (*MessageRecallResponse, error) {
	path := "/im/message/recall"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(MessageRecallResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// OpenSession open a new conversation
// Path: /im/session/open
func (s *InstantMessagingServiceOp[T]) OpenSession(ctx context.Context) (*OpenSessionResponse, error) {
	path := "/im/session/open"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(OpenSessionResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// ReadSession session read
// Path: /im/session/read
func (s *InstantMessagingServiceOp[T]) ReadSession(ctx context.Context) (*ReadSessionResponse, error) {
	path := "/im/session/read"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(ReadSessionResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// SendMessage send message
// Path: /im/message/send
func (s *InstantMessagingServiceOp[T]) SendMessage(ctx context.Context) (*SendMessageResponse, error) {
	path := "/im/message/send"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(SendMessageResponse)
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
