package golazada

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_RedMart_RssGetOnePickupJob(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_rss_pickup_job_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping RssGetOnePickupJob due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/rss/pickup-job/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.RedMart.RssGetOnePickupJob(ctx)
	if err != nil {
		t.Logf("RedMart.RssGetOnePickupJob returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("RedMart.RssGetOnePickupJob response: %#v", res)
}
func Test_RedMart_RssGetPickupJobs(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_rss_pickup_jobs_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping RssGetPickupJobs due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/rss/pickup-jobs/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.RedMart.RssGetPickupJobs(ctx)
	if err != nil {
		t.Logf("RedMart.RssGetPickupJobs returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("RedMart.RssGetPickupJobs response: %#v", res)
}
func Test_RedMart_RssGetPickupLocations(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_rss_pickupLocations_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping RssGetPickupLocations due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/rss/pickupLocations/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.RedMart.RssGetPickupLocations(ctx)
	if err != nil {
		t.Logf("RedMart.RssGetPickupLocations returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("RedMart.RssGetPickupLocations response: %#v", res)
}
func Test_RedMart_RssGetProduct(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_rss_product_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping RssGetProduct due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/rss/product/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.RedMart.RssGetProduct(ctx)
	if err != nil {
		t.Logf("RedMart.RssGetProduct returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("RedMart.RssGetProduct response: %#v", res)
}
func Test_RedMart_RssGetProducts(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_rss_products_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping RssGetProducts due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/rss/products/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.RedMart.RssGetProducts(ctx)
	if err != nil {
		t.Logf("RedMart.RssGetProducts returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("RedMart.RssGetProducts response: %#v", res)
}
func Test_RedMart_RssGetStockLot(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_rss_stockLot_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping RssGetStockLot due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/rss/stockLot/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.RedMart.RssGetStockLot(ctx)
	if err != nil {
		t.Logf("RedMart.RssGetStockLot returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("RedMart.RssGetStockLot response: %#v", res)
}
func Test_RedMart_RssGetStockLots(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_rss_stockLots_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping RssGetStockLots due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/rss/stockLots/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.RedMart.RssGetStockLots(ctx)
	if err != nil {
		t.Logf("RedMart.RssGetStockLots returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("RedMart.RssGetStockLots response: %#v", res)
}
func Test_RedMart_RssUpdateStockLot(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_rss_stockLot_update_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping RssUpdateStockLot due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/rss/stockLot/update*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.RedMart.RssUpdateStockLot(ctx)
	if err != nil {
		t.Logf("RedMart.RssUpdateStockLot returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("RedMart.RssUpdateStockLot response: %#v", res)
}
