package golazada

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
)

type LazLikeService interface {
	// McnContentCancelSchedulePublish McnContentCancelSchedulePublish
	// Path: /content/mcn/content/cancelScheduled
	McnContentCancelSchedulePublish(ctx context.Context) (*McnContentCancelSchedulePublishResponse, error)
	// McnContentCompleteCreateVideo After uploading all blocks of the video file, call McnContentCompleteCreateVideo to complete the video uploading process.
	//
	// Path: /content/mcn/video/block/commit
	McnContentCompleteCreateVideo(ctx context.Context) (*McnContentCompleteCreateVideoResponse, error)
	// McnContentCreate create content
	// Path: /content/mcn/content/create
	McnContentCreate(ctx context.Context) (*McnContentCreateResponse, error)
	// McnContentInitCreateVideo Initial an upload video process, this API will return the corresponding UploadID
	// Path: /content/mcn/video/block/create
	McnContentInitCreateVideo(ctx context.Context) (*McnContentInitCreateVideoResponse, error)
	// McnContentListCategory list mcn content categories
	// Path: /content/mcn/category/list
	McnContentListCategory(ctx context.Context) (*McnContentListCategoryResponse, error)
	// McnContentPropertyTagList list mcn content property tags
	// Path: /content/mcn/property/list
	McnContentPropertyTagList(ctx context.Context) (*McnContentPropertyTagListResponse, error)
	// McnContentReplySchedulePublish McnContentReplySchedulePublish
	// Path: /content/mcn/content/replySchedulePublish
	McnContentReplySchedulePublish(ctx context.Context) (*McnContentReplySchedulePublishResponse, error)
	// McnContentUploadImage upload image
	// Path: /content/mcn/image/upload
	McnContentUploadImage(ctx context.Context, filename string, reader io.Reader) (*McnContentUploadImageResponse, error)
	// McnContentUploadVideoBlock upload one block of video file
	// Path: /content/mcn/video/block/upload
	McnContentUploadVideoBlock(ctx context.Context, filename string, reader io.Reader) (*McnContentUploadVideoBlockResponse, error)
	// McnProductValidator Identify high risk products
	// Path: /content/mcn/product/validate
	McnProductValidator(ctx context.Context) (*McnProductValidatorResponse, error)
	// MCNQueryTagInfoByName MCNQueryTagInfoByName
	// Path: /content/mcn/content/queryTagInfosByName
	MCNQueryTagInfoByName(ctx context.Context) (*MCNQueryTagInfoByNameResponse, error)
	// McnSimilarProductSearch 相似商品搜索接口
	// Path: /content/mcn/similar/product/search
	McnSimilarProductSearch(ctx context.Context) (*McnSimilarProductSearchResponse, error)
	// QueryContentReviewRecords Query content audit records. Currently, querying records with audit results of low (block) is supported.The number of query contents is limited to 500 (adjustable).
	// Path: /content/mcn/content/queryReviewRecords
	QueryContentReviewRecords(ctx context.Context) (*QueryContentReviewRecordsResponse, error)
}

type LazLikeServiceOp[T any] struct {
	client *Client[T]
}

// McnContentCancelSchedulePublish McnContentCancelSchedulePublish
// Path: /content/mcn/content/cancelScheduled
func (s *LazLikeServiceOp[T]) McnContentCancelSchedulePublish(ctx context.Context) (*McnContentCancelSchedulePublishResponse, error) {
	path := "/content/mcn/content/cancelScheduled"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(McnContentCancelSchedulePublishResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// McnContentCompleteCreateVideo After uploading all blocks of the video file, call McnContentCompleteCreateVideo to complete the video uploading process.
//
// Path: /content/mcn/video/block/commit
func (s *LazLikeServiceOp[T]) McnContentCompleteCreateVideo(ctx context.Context) (*McnContentCompleteCreateVideoResponse, error) {
	path := "/content/mcn/video/block/commit"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(McnContentCompleteCreateVideoResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// McnContentCreate create content
// Path: /content/mcn/content/create
func (s *LazLikeServiceOp[T]) McnContentCreate(ctx context.Context) (*McnContentCreateResponse, error) {
	path := "/content/mcn/content/create"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(McnContentCreateResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// McnContentInitCreateVideo Initial an upload video process, this API will return the corresponding UploadID
// Path: /content/mcn/video/block/create
func (s *LazLikeServiceOp[T]) McnContentInitCreateVideo(ctx context.Context) (*McnContentInitCreateVideoResponse, error) {
	path := "/content/mcn/video/block/create"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(McnContentInitCreateVideoResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// McnContentListCategory list mcn content categories
// Path: /content/mcn/category/list
func (s *LazLikeServiceOp[T]) McnContentListCategory(ctx context.Context) (*McnContentListCategoryResponse, error) {
	path := "/content/mcn/category/list"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(McnContentListCategoryResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// McnContentPropertyTagList list mcn content property tags
// Path: /content/mcn/property/list
func (s *LazLikeServiceOp[T]) McnContentPropertyTagList(ctx context.Context) (*McnContentPropertyTagListResponse, error) {
	path := "/content/mcn/property/list"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(McnContentPropertyTagListResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// McnContentReplySchedulePublish McnContentReplySchedulePublish
// Path: /content/mcn/content/replySchedulePublish
func (s *LazLikeServiceOp[T]) McnContentReplySchedulePublish(ctx context.Context) (*McnContentReplySchedulePublishResponse, error) {
	path := "/content/mcn/content/replySchedulePublish"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(McnContentReplySchedulePublishResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// McnContentUploadImage upload image
// Path: /content/mcn/image/upload
func (s *LazLikeServiceOp[T]) McnContentUploadImage(ctx context.Context, filename string, reader io.Reader) (*McnContentUploadImageResponse, error) {
	path := "/content/mcn/image/upload"
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", filename)
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(part, reader); err != nil {
		return nil, err
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}

	wrapper, err := s.client.execute(ctx, "POST", path, nil, map[string][]byte{"image": {}})
	if err != nil {
		return nil, err
	}
	resp := new(McnContentUploadImageResponse)
	if err := json.Unmarshal(wrapper.Data, resp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return resp, nil
}

// McnContentUploadVideoBlock upload one block of video file
// Path: /content/mcn/video/block/upload
func (s *LazLikeServiceOp[T]) McnContentUploadVideoBlock(ctx context.Context, filename string, reader io.Reader) (*McnContentUploadVideoBlockResponse, error) {
	path := "/content/mcn/video/block/upload"
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", filename)
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(part, reader); err != nil {
		return nil, err
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}

	wrapper, err := s.client.execute(ctx, "POST", path, nil, map[string][]byte{"image": {}})
	if err != nil {
		return nil, err
	}
	resp := new(McnContentUploadVideoBlockResponse)
	if err := json.Unmarshal(wrapper.Data, resp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return resp, nil
}

// McnProductValidator Identify high risk products
// Path: /content/mcn/product/validate
func (s *LazLikeServiceOp[T]) McnProductValidator(ctx context.Context) (*McnProductValidatorResponse, error) {
	path := "/content/mcn/product/validate"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(McnProductValidatorResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// MCNQueryTagInfoByName MCNQueryTagInfoByName
// Path: /content/mcn/content/queryTagInfosByName
func (s *LazLikeServiceOp[T]) MCNQueryTagInfoByName(ctx context.Context) (*MCNQueryTagInfoByNameResponse, error) {
	path := "/content/mcn/content/queryTagInfosByName"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(MCNQueryTagInfoByNameResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// McnSimilarProductSearch 相似商品搜索接口
// Path: /content/mcn/similar/product/search
func (s *LazLikeServiceOp[T]) McnSimilarProductSearch(ctx context.Context) (*McnSimilarProductSearchResponse, error) {
	path := "/content/mcn/similar/product/search"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(McnSimilarProductSearchResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// QueryContentReviewRecords Query content audit records. Currently, querying records with audit results of low (block) is supported.The number of query contents is limited to 500 (adjustable).
// Path: /content/mcn/content/queryReviewRecords
func (s *LazLikeServiceOp[T]) QueryContentReviewRecords(ctx context.Context) (*QueryContentReviewRecordsResponse, error) {
	path := "/content/mcn/content/queryReviewRecords"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(QueryContentReviewRecordsResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}
