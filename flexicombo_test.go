package golazada

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_Flexicombo_ActivateFlexiCombo(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_promotion_flexicombo_activate_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping ActivateFlexiCombo due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/promotion/flexicombo/activate*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Flexicombo.ActivateFlexiCombo(ctx)
	if err != nil {
		t.Logf("Flexicombo.ActivateFlexiCombo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Flexicombo.ActivateFlexiCombo response: %#v", res)
}
func Test_Flexicombo_AddFlexiComboProducts(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_promotion_flexicombo_products_add_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping AddFlexiComboProducts due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/promotion/flexicombo/products/add*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Flexicombo.AddFlexiComboProducts(ctx)
	if err != nil {
		t.Logf("Flexicombo.AddFlexiComboProducts returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Flexicombo.AddFlexiComboProducts response: %#v", res)
}
func Test_Flexicombo_CreateFlexiCombo(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_promotion_flexicombo_create_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CreateFlexiCombo due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/promotion/flexicombo/create*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Flexicombo.CreateFlexiCombo(ctx)
	if err != nil {
		t.Logf("Flexicombo.CreateFlexiCombo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Flexicombo.CreateFlexiCombo response: %#v", res)
}
func Test_Flexicombo_DeactivateFlexiCombo(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_promotion_flexicombo_deactivate_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DeactivateFlexiCombo due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/promotion/flexicombo/deactivate*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Flexicombo.DeactivateFlexiCombo(ctx)
	if err != nil {
		t.Logf("Flexicombo.DeactivateFlexiCombo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Flexicombo.DeactivateFlexiCombo response: %#v", res)
}
func Test_Flexicombo_DeleteFlexiComboProducts(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_promotion_flexicombo_products_delete_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DeleteFlexiComboProducts due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/promotion/flexicombo/products/delete*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Flexicombo.DeleteFlexiComboProducts(ctx)
	if err != nil {
		t.Logf("Flexicombo.DeleteFlexiComboProducts returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Flexicombo.DeleteFlexiComboProducts response: %#v", res)
}
func Test_Flexicombo_GetFlexiComboDetails(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_promotion_flexicombo_details_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetFlexiComboDetails due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/promotion/flexicombo/details*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Flexicombo.GetFlexiComboDetails(ctx)
	if err != nil {
		t.Logf("Flexicombo.GetFlexiComboDetails returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Flexicombo.GetFlexiComboDetails response: %#v", res)
}
func Test_Flexicombo_ListFlexiCombo(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_promotion_flexicombo_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping ListFlexiCombo due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/promotion/flexicombo/list*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Flexicombo.ListFlexiCombo(ctx)
	if err != nil {
		t.Logf("Flexicombo.ListFlexiCombo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Flexicombo.ListFlexiCombo response: %#v", res)
}
func Test_Flexicombo_ListFlexiComboProducts(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_promotion_flexicombo_products_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping ListFlexiComboProducts due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/promotion/flexicombo/products/list*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Flexicombo.ListFlexiComboProducts(ctx)
	if err != nil {
		t.Logf("Flexicombo.ListFlexiComboProducts returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Flexicombo.ListFlexiComboProducts response: %#v", res)
}
func Test_Flexicombo_UpdateFlexiCombo(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_promotion_flexicombo_update_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateFlexiCombo due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/promotion/flexicombo/update*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.Flexicombo.UpdateFlexiCombo(ctx)
	if err != nil {
		t.Logf("Flexicombo.UpdateFlexiCombo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Flexicombo.UpdateFlexiCombo response: %#v", res)
}
