package golazada

import (
	"context"
	"encoding/json"
	"fmt"
)

type ProductReviewService interface {
	// GetHistoryReviewIdList Get history review id list for one seller(reviews within 3 months can be get)
	// Path: /review/seller/history/list
	GetHistoryReviewIdList(ctx context.Context) (*GetHistoryReviewIdListResponse, error)
	// GetReviewListByIdList get review list by id list, need get id list first
	// Path: /review/seller/list/v2
	GetReviewListByIdList(ctx context.Context) (*GetReviewListByIdListResponse, error)
	// SubmitSellerReply submit seller reply for customers review
	// Path: /review/seller/reply/add
	SubmitSellerReply(ctx context.Context) (*SubmitSellerReplyResponse, error)
}

type ProductReviewServiceOp[T any] struct {
	client *Client[T]
}

// GetHistoryReviewIdList Get history review id list for one seller(reviews within 3 months can be get)
// Path: /review/seller/history/list
func (s *ProductReviewServiceOp[T]) GetHistoryReviewIdList(ctx context.Context) (*GetHistoryReviewIdListResponse, error) {
	path := "/review/seller/history/list"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetHistoryReviewIdListResponse)
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

// GetReviewListByIdList get review list by id list, need get id list first
// Path: /review/seller/list/v2
func (s *ProductReviewServiceOp[T]) GetReviewListByIdList(ctx context.Context) (*GetReviewListByIdListResponse, error) {
	path := "/review/seller/list/v2"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetReviewListByIdListResponse)
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

// SubmitSellerReply submit seller reply for customers review
// Path: /review/seller/reply/add
func (s *ProductReviewServiceOp[T]) SubmitSellerReply(ctx context.Context) (*SubmitSellerReplyResponse, error) {
	path := "/review/seller/reply/add"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(SubmitSellerReplyResponse)
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
