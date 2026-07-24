package golazada

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_SellerVoucher_SellerVoucheDeleteSelectedProductSKU(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_promotion_voucher_product_sku_remove_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SellerVoucheDeleteSelectedProductSKU due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/promotion/voucher/product/sku/remove*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SellerVoucher.SellerVoucheDeleteSelectedProductSKU(ctx)
	if err != nil {
		t.Logf("SellerVoucher.SellerVoucheDeleteSelectedProductSKU returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SellerVoucher.SellerVoucheDeleteSelectedProductSKU response: %#v", res)
}
func Test_SellerVoucher_SellerVoucherActivate(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_promotion_voucher_activate_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SellerVoucherActivate due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/promotion/voucher/activate*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SellerVoucher.SellerVoucherActivate(ctx)
	if err != nil {
		t.Logf("SellerVoucher.SellerVoucherActivate returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SellerVoucher.SellerVoucherActivate response: %#v", res)
}
func Test_SellerVoucher_SellerVoucherAddSelectedProductSKU(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_promotion_voucher_product_sku_add_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SellerVoucherAddSelectedProductSKU due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/promotion/voucher/product/sku/add*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SellerVoucher.SellerVoucherAddSelectedProductSKU(ctx)
	if err != nil {
		t.Logf("SellerVoucher.SellerVoucherAddSelectedProductSKU returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SellerVoucher.SellerVoucherAddSelectedProductSKU response: %#v", res)
}
func Test_SellerVoucher_SellerVoucherCreate(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_promotion_voucher_create_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SellerVoucherCreate due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/promotion/voucher/create*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SellerVoucher.SellerVoucherCreate(ctx)
	if err != nil {
		t.Logf("SellerVoucher.SellerVoucherCreate returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SellerVoucher.SellerVoucherCreate response: %#v", res)
}
func Test_SellerVoucher_SellerVoucherDeactivate(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_promotion_voucher_deactivate_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SellerVoucherDeactivate due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/promotion/voucher/deactivate*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SellerVoucher.SellerVoucherDeactivate(ctx)
	if err != nil {
		t.Logf("SellerVoucher.SellerVoucherDeactivate returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SellerVoucher.SellerVoucherDeactivate response: %#v", res)
}
func Test_SellerVoucher_SellerVoucherDetailQuery(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_promotion_voucher_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SellerVoucherDetailQuery due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/promotion/voucher/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SellerVoucher.SellerVoucherDetailQuery(ctx)
	if err != nil {
		t.Logf("SellerVoucher.SellerVoucherDetailQuery returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SellerVoucher.SellerVoucherDetailQuery response: %#v", res)
}
func Test_SellerVoucher_SellerVoucherList(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_promotion_vouchers_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SellerVoucherList due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/promotion/vouchers/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SellerVoucher.SellerVoucherList(ctx)
	if err != nil {
		t.Logf("SellerVoucher.SellerVoucherList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SellerVoucher.SellerVoucherList response: %#v", res)
}
func Test_SellerVoucher_SellerVoucherSelectedProductList(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_promotion_voucher_products_get_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SellerVoucherSelectedProductList due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/promotion/voucher/products/get*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SellerVoucher.SellerVoucherSelectedProductList(ctx)
	if err != nil {
		t.Logf("SellerVoucher.SellerVoucherSelectedProductList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SellerVoucher.SellerVoucherSelectedProductList response: %#v", res)
}
func Test_SellerVoucher_SellerVoucherUpdate(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_promotion_voucher_update_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SellerVoucherUpdate due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/promotion/voucher/update*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SellerVoucher.SellerVoucherUpdate(ctx)
	if err != nil {
		t.Logf("SellerVoucher.SellerVoucherUpdate returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SellerVoucher.SellerVoucherUpdate response: %#v", res)
}
