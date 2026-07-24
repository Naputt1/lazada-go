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

func Test_LazLike_McnContentCancelSchedulePublish(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_content_mcn_content_cancelScheduled_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping McnContentCancelSchedulePublish due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/content/mcn/content/cancelScheduled*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazLike.McnContentCancelSchedulePublish(ctx)
	if err != nil {
		t.Logf("LazLike.McnContentCancelSchedulePublish returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazLike.McnContentCancelSchedulePublish response: %#v", res)
}
func Test_LazLike_McnContentCompleteCreateVideo(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_content_mcn_video_block_commit_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping McnContentCompleteCreateVideo due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/content/mcn/video/block/commit*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazLike.McnContentCompleteCreateVideo(ctx)
	if err != nil {
		t.Logf("LazLike.McnContentCompleteCreateVideo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazLike.McnContentCompleteCreateVideo response: %#v", res)
}
func Test_LazLike_McnContentCreate(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_content_mcn_content_create_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping McnContentCreate due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/content/mcn/content/create*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazLike.McnContentCreate(ctx)
	if err != nil {
		t.Logf("LazLike.McnContentCreate returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazLike.McnContentCreate response: %#v", res)
}
func Test_LazLike_McnContentInitCreateVideo(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_content_mcn_video_block_create_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping McnContentInitCreateVideo due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/content/mcn/video/block/create*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazLike.McnContentInitCreateVideo(ctx)
	if err != nil {
		t.Logf("LazLike.McnContentInitCreateVideo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazLike.McnContentInitCreateVideo response: %#v", res)
}
func Test_LazLike_McnContentListCategory(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_content_mcn_category_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping McnContentListCategory due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/content/mcn/category/list*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazLike.McnContentListCategory(ctx)
	if err != nil {
		t.Logf("LazLike.McnContentListCategory returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazLike.McnContentListCategory response: %#v", res)
}
func Test_LazLike_McnContentPropertyTagList(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_content_mcn_property_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping McnContentPropertyTagList due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/content/mcn/property/list*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazLike.McnContentPropertyTagList(ctx)
	if err != nil {
		t.Logf("LazLike.McnContentPropertyTagList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazLike.McnContentPropertyTagList response: %#v", res)
}
func Test_LazLike_McnContentReplySchedulePublish(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_content_mcn_content_replySchedulePublish_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping McnContentReplySchedulePublish due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/content/mcn/content/replySchedulePublish*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazLike.McnContentReplySchedulePublish(ctx)
	if err != nil {
		t.Logf("LazLike.McnContentReplySchedulePublish returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazLike.McnContentReplySchedulePublish response: %#v", res)
}
func Test_LazLike_McnContentUploadImage(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_content_mcn_image_upload_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping McnContentUploadImage due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/content/mcn/image/upload*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazLike.McnContentUploadImage(ctx, "test.jpg", strings.NewReader("test data"))
	if err != nil {
		t.Logf("LazLike.McnContentUploadImage returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazLike.McnContentUploadImage response: %#v", res)
}
func Test_LazLike_McnContentUploadVideoBlock(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_content_mcn_video_block_upload_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping McnContentUploadVideoBlock due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/content/mcn/video/block/upload*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazLike.McnContentUploadVideoBlock(ctx, "test.jpg", strings.NewReader("test data"))
	if err != nil {
		t.Logf("LazLike.McnContentUploadVideoBlock returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazLike.McnContentUploadVideoBlock response: %#v", res)
}
func Test_LazLike_McnProductValidator(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_content_mcn_product_validate_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping McnProductValidator due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/content/mcn/product/validate*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazLike.McnProductValidator(ctx)
	if err != nil {
		t.Logf("LazLike.McnProductValidator returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazLike.McnProductValidator response: %#v", res)
}
func Test_LazLike_MCNQueryTagInfoByName(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_content_mcn_content_queryTagInfosByName_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping MCNQueryTagInfoByName due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/content/mcn/content/queryTagInfosByName*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazLike.MCNQueryTagInfoByName(ctx)
	if err != nil {
		t.Logf("LazLike.MCNQueryTagInfoByName returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazLike.MCNQueryTagInfoByName response: %#v", res)
}
func Test_LazLike_McnSimilarProductSearch(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_content_mcn_similar_product_search_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping McnSimilarProductSearch due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/content/mcn/similar/product/search*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazLike.McnSimilarProductSearch(ctx)
	if err != nil {
		t.Logf("LazLike.McnSimilarProductSearch returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazLike.McnSimilarProductSearch response: %#v", res)
}
func Test_LazLike_QueryContentReviewRecords(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_content_mcn_content_queryReviewRecords_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping QueryContentReviewRecords due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/content/mcn/content/queryReviewRecords*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazLike.QueryContentReviewRecords(ctx)
	if err != nil {
		t.Logf("LazLike.QueryContentReviewRecords returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazLike.QueryContentReviewRecords response: %#v", res)
}
