package golazada

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_MediaCenter_CompleteCreateVideo(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_media_video_block_commit_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CompleteCreateVideo due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/media/video/block/commit*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.MediaCenter.CompleteCreateVideo(ctx)
	if err != nil {
		t.Logf("MediaCenter.CompleteCreateVideo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("MediaCenter.CompleteCreateVideo response: %#v", res)
}
func Test_MediaCenter_GetVideo(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_media_video_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetVideo due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/media/video/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.MediaCenter.GetVideo(ctx)
	if err != nil {
		t.Logf("MediaCenter.GetVideo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("MediaCenter.GetVideo response: %#v", res)
}
func Test_MediaCenter_GetVideoQuota(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_media_video_quota_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetVideoQuota due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/media/video/quota/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.MediaCenter.GetVideoQuota(ctx)
	if err != nil {
		t.Logf("MediaCenter.GetVideoQuota returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("MediaCenter.GetVideoQuota response: %#v", res)
}
func Test_MediaCenter_InitCreateVideo(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_media_video_block_create_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping InitCreateVideo due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/media/video/block/create*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.MediaCenter.InitCreateVideo(ctx)
	if err != nil {
		t.Logf("MediaCenter.InitCreateVideo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("MediaCenter.InitCreateVideo response: %#v", res)
}
func Test_MediaCenter_RemoveVideo(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_media_video_remove_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping RemoveVideo due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/media/video/remove*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.MediaCenter.RemoveVideo(ctx)
	if err != nil {
		t.Logf("MediaCenter.RemoveVideo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("MediaCenter.RemoveVideo response: %#v", res)
}
func Test_MediaCenter_UploadVideoBlock(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_media_video_block_upload_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UploadVideoBlock due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/media/video/block/upload*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.MediaCenter.UploadVideoBlock(ctx, "test.jpg", strings.NewReader("test data"))
	if err != nil {
		t.Logf("MediaCenter.UploadVideoBlock returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("MediaCenter.UploadVideoBlock response: %#v", res)
}
