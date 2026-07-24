package golazada

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
)

type MediaCenterService interface {
	// CompleteCreateVideo After uploading all blocks of the video file,  call CompleteCreateVideo to complete the video uploading process.
	// Path: /media/video/block/commit
	CompleteCreateVideo(ctx context.Context) (*CompleteCreateVideoResponse, error)
	// GetVideo You call this action to get video info after uploading.
	// Path: /media/video/get
	GetVideo(ctx context.Context) (*GetVideoResponse, error)
	// GetVideoQuota You call this api to get the capacity quota of seller.
	// Path: /media/video/quota/get
	GetVideoQuota(ctx context.Context) (*GetVideoQuotaResponse, error)
	// InitCreateVideo A seller starts to upload a video file
	// Path: /media/video/block/create
	InitCreateVideo(ctx context.Context) (*InitCreateVideoResponse, error)
	// RemoveVideo You can this api to delete a video file permanently.
	// Path: /media/video/remove
	RemoveVideo(ctx context.Context) (*RemoveVideoResponse, error)
	// UploadVideoBlock The API is used to upload one block of origin video file. The video file can split into multiple files. For example, a 8MB video file can be split into three blocks. 3MB, 3MB and 2MB. These three blocks can be uploaded by calling UploadVideoBlock three times.
	// Path: /media/video/block/upload
	UploadVideoBlock(ctx context.Context, filename string, reader io.Reader) (*UploadVideoBlockResponse, error)
}

type MediaCenterServiceOp[T any] struct {
	client *Client[T]
}

// CompleteCreateVideo After uploading all blocks of the video file,  call CompleteCreateVideo to complete the video uploading process.
// Path: /media/video/block/commit
func (s *MediaCenterServiceOp[T]) CompleteCreateVideo(ctx context.Context) (*CompleteCreateVideoResponse, error) {
	path := "/media/video/block/commit"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(CompleteCreateVideoResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetVideo You call this action to get video info after uploading.
// Path: /media/video/get
func (s *MediaCenterServiceOp[T]) GetVideo(ctx context.Context) (*GetVideoResponse, error) {
	path := "/media/video/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetVideoResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetVideoQuota You call this api to get the capacity quota of seller.
// Path: /media/video/quota/get
func (s *MediaCenterServiceOp[T]) GetVideoQuota(ctx context.Context) (*GetVideoQuotaResponse, error) {
	path := "/media/video/quota/get"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetVideoQuotaResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// InitCreateVideo A seller starts to upload a video file
// Path: /media/video/block/create
func (s *MediaCenterServiceOp[T]) InitCreateVideo(ctx context.Context) (*InitCreateVideoResponse, error) {
	path := "/media/video/block/create"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(InitCreateVideoResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// RemoveVideo You can this api to delete a video file permanently.
// Path: /media/video/remove
func (s *MediaCenterServiceOp[T]) RemoveVideo(ctx context.Context) (*RemoveVideoResponse, error) {
	path := "/media/video/remove"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(RemoveVideoResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// UploadVideoBlock The API is used to upload one block of origin video file. The video file can split into multiple files. For example, a 8MB video file can be split into three blocks. 3MB, 3MB and 2MB. These three blocks can be uploaded by calling UploadVideoBlock three times.
// Path: /media/video/block/upload
func (s *MediaCenterServiceOp[T]) UploadVideoBlock(ctx context.Context, filename string, reader io.Reader) (*UploadVideoBlockResponse, error) {
	path := "/media/video/block/upload"
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
	resp := new(UploadVideoBlockResponse)
	if err := json.Unmarshal(wrapper.Data, resp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return resp, nil
}
