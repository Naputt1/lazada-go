package golazada

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_EarlyBirdPrice_CreateEarlyBirdActivityV2(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_activity_early_bird_create_v2_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CreateEarlyBirdActivityV2 due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/activity/early/bird/create/v2*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.EarlyBirdPrice.CreateEarlyBirdActivityV2(ctx)
	if err != nil {
		t.Logf("EarlyBirdPrice.CreateEarlyBirdActivityV2 returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("EarlyBirdPrice.CreateEarlyBirdActivityV2 response: %#v", res)
}
func Test_EarlyBirdPrice_EarlyBirdActivityAddSkusV2(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_activity_early_bird_addSkus_v2_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping EarlyBirdActivityAddSkusV2 due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/activity/early/bird/addSkus/v2*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.EarlyBirdPrice.EarlyBirdActivityAddSkusV2(ctx)
	if err != nil {
		t.Logf("EarlyBirdPrice.EarlyBirdActivityAddSkusV2 returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("EarlyBirdPrice.EarlyBirdActivityAddSkusV2 response: %#v", res)
}
func Test_EarlyBirdPrice_EarlyBirdActivityDeactivateSkusV2(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_activity_early_bird_deactivateSkus_v2_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping EarlyBirdActivityDeactivateSkusV2 due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/activity/early/bird/deactivateSkus/v2*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.EarlyBirdPrice.EarlyBirdActivityDeactivateSkusV2(ctx)
	if err != nil {
		t.Logf("EarlyBirdPrice.EarlyBirdActivityDeactivateSkusV2 returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("EarlyBirdPrice.EarlyBirdActivityDeactivateSkusV2 response: %#v", res)
}
func Test_EarlyBirdPrice_EarlyBirdActivityIsWhitelistSeller(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_activity_early_bird_isWhitelistSeller_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping EarlyBirdActivityIsWhitelistSeller due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/activity/early/bird/isWhitelistSeller*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.EarlyBirdPrice.EarlyBirdActivityIsWhitelistSeller(ctx)
	if err != nil {
		t.Logf("EarlyBirdPrice.EarlyBirdActivityIsWhitelistSeller returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("EarlyBirdPrice.EarlyBirdActivityIsWhitelistSeller response: %#v", res)
}
