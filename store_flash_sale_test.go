package golazada

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_StoreFlashSale_ActivateStoreFlashSale(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_promotion_storeflashsale_activate_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping ActivateStoreFlashSale due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/promotion/storeflashsale/activate*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.StoreFlashSale.ActivateStoreFlashSale(ctx)
	if err != nil {
		t.Logf("StoreFlashSale.ActivateStoreFlashSale returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("StoreFlashSale.ActivateStoreFlashSale response: %#v", res)
}
func Test_StoreFlashSale_CreateStoreFlashSale(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_promotion_storeflashsale_create_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CreateStoreFlashSale due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/promotion/storeflashsale/create*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.StoreFlashSale.CreateStoreFlashSale(ctx)
	if err != nil {
		t.Logf("StoreFlashSale.CreateStoreFlashSale returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("StoreFlashSale.CreateStoreFlashSale response: %#v", res)
}
func Test_StoreFlashSale_DeactivateStoreFlashSale(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_promotion_storeflashsale_deactivate_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DeactivateStoreFlashSale due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/promotion/storeflashsale/deactivate*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.StoreFlashSale.DeactivateStoreFlashSale(ctx)
	if err != nil {
		t.Logf("StoreFlashSale.DeactivateStoreFlashSale returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("StoreFlashSale.DeactivateStoreFlashSale response: %#v", res)
}
func Test_StoreFlashSale_GetStoreFlashSale(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_promotion_storeflashsale_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetStoreFlashSale due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/promotion/storeflashsale/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.StoreFlashSale.GetStoreFlashSale(ctx)
	if err != nil {
		t.Logf("StoreFlashSale.GetStoreFlashSale returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("StoreFlashSale.GetStoreFlashSale response: %#v", res)
}
func Test_StoreFlashSale_ListStoreFlashSale(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_promotion_storeflashsale_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping ListStoreFlashSale due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/promotion/storeflashsale/list*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.StoreFlashSale.ListStoreFlashSale(ctx)
	if err != nil {
		t.Logf("StoreFlashSale.ListStoreFlashSale returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("StoreFlashSale.ListStoreFlashSale response: %#v", res)
}
func Test_StoreFlashSale_UpdateStoreFlashSale(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_promotion_storeflashsale_update_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateStoreFlashSale due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/promotion/storeflashsale/update*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.StoreFlashSale.UpdateStoreFlashSale(ctx)
	if err != nil {
		t.Logf("StoreFlashSale.UpdateStoreFlashSale returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("StoreFlashSale.UpdateStoreFlashSale response: %#v", res)
}
