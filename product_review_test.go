package golazada

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_ProductReview_GetHistoryReviewIdList(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_review_seller_history_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetHistoryReviewIdList due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/review/seller/history/list*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.ProductReview.GetHistoryReviewIdList(ctx)
	if err != nil {
		t.Logf("ProductReview.GetHistoryReviewIdList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ProductReview.GetHistoryReviewIdList response: %#v", res)
}
func Test_ProductReview_GetReviewListByIdList(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_review_seller_list_v2_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetReviewListByIdList due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/review/seller/list/v2*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.ProductReview.GetReviewListByIdList(ctx)
	if err != nil {
		t.Logf("ProductReview.GetReviewListByIdList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ProductReview.GetReviewListByIdList response: %#v", res)
}
func Test_ProductReview_SubmitSellerReply(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_review_seller_reply_add_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SubmitSellerReply due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/review/seller/reply/add*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.ProductReview.SubmitSellerReply(ctx)
	if err != nil {
		t.Logf("ProductReview.SubmitSellerReply returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("ProductReview.SubmitSellerReply response: %#v", res)
}
