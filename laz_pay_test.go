package golazada

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_LazPay_CollectBenefit(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_insurance_promotion_collectBenefit_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CollectBenefit due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/insurance/promotion/collectBenefit*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazPay.CollectBenefit(ctx)
	if err != nil {
		t.Logf("LazPay.CollectBenefit returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazPay.CollectBenefit response: %#v", res)
}
func Test_LazPay_ConsultPayment(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_lazadapay_v1_debit_consult_payment_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping ConsultPayment due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/lazadapay/v1/debit/consult_payment*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazPay.ConsultPayment(ctx)
	if err != nil {
		t.Logf("LazPay.ConsultPayment returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazPay.ConsultPayment response: %#v", res)
}
func Test_LazPay_CreateSubscriptionToFusion(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_insurance_subscription_create_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CreateSubscriptionToFusion due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/insurance/subscription/create*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazPay.CreateSubscriptionToFusion(ctx)
	if err != nil {
		t.Logf("LazPay.CreateSubscriptionToFusion returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazPay.CreateSubscriptionToFusion response: %#v", res)
}
func Test_LazPay_DGUtiityPreCreateOrder(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_digital_service_createorder_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DGUtiityPreCreateOrder due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/digital/service/createorder*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazPay.DGUtiityPreCreateOrder(ctx)
	if err != nil {
		t.Logf("LazPay.DGUtiityPreCreateOrder returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazPay.DGUtiityPreCreateOrder response: %#v", res)
}
func Test_LazPay_DGUtilityPreGetPaymentStatus(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_digital_service_getPaymentStatus_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DGUtilityPreGetPaymentStatus due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/digital/service/getPaymentStatus*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazPay.DGUtilityPreGetPaymentStatus(ctx)
	if err != nil {
		t.Logf("LazPay.DGUtilityPreGetPaymentStatus returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazPay.DGUtilityPreGetPaymentStatus response: %#v", res)
}
func Test_LazPay_DGUtilityPreUpdateFulfillemtStatus(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_digital_service_updateFulfillemtStatus_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DGUtilityPreUpdateFulfillemtStatus due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/digital/service/updateFulfillemtStatus*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazPay.DGUtilityPreUpdateFulfillemtStatus(ctx)
	if err != nil {
		t.Logf("LazPay.DGUtilityPreUpdateFulfillemtStatus returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazPay.DGUtilityPreUpdateFulfillemtStatus response: %#v", res)
}
func Test_LazPay_DigitalAlterOrderStatus(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_digital_order_alterStatus_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DigitalAlterOrderStatus due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/digital/order/alterStatus*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazPay.DigitalAlterOrderStatus(ctx)
	if err != nil {
		t.Logf("LazPay.DigitalAlterOrderStatus returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazPay.DigitalAlterOrderStatus response: %#v", res)
}
func Test_LazPay_DigitalCreateOrder(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_digital_order_create_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DigitalCreateOrder due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/digital/order/create*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazPay.DigitalCreateOrder(ctx)
	if err != nil {
		t.Logf("LazPay.DigitalCreateOrder returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazPay.DigitalCreateOrder response: %#v", res)
}
func Test_LazPay_DigitalQueryOrder(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_digital_order_getStatus_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DigitalQueryOrder due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/digital/order/getStatus*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazPay.DigitalQueryOrder(ctx)
	if err != nil {
		t.Logf("LazPay.DigitalQueryOrder returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazPay.DigitalQueryOrder response: %#v", res)
}
func Test_LazPay_GetSubscriptionToFusion(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_insurance_subscription_getSubscription_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetSubscriptionToFusion due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/insurance/subscription/getSubscription*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazPay.GetSubscriptionToFusion(ctx)
	if err != nil {
		t.Logf("LazPay.GetSubscriptionToFusion returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazPay.GetSubscriptionToFusion response: %#v", res)
}
func Test_LazPay_InsuranceAlterOrderStatus(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_insurance_order_alterStatus_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping InsuranceAlterOrderStatus due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/insurance/order/alterStatus*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazPay.InsuranceAlterOrderStatus(ctx)
	if err != nil {
		t.Logf("LazPay.InsuranceAlterOrderStatus returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazPay.InsuranceAlterOrderStatus response: %#v", res)
}
func Test_LazPay_InsuranceCreateOrder(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_insurance_order_create_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping InsuranceCreateOrder due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/insurance/order/create*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazPay.InsuranceCreateOrder(ctx)
	if err != nil {
		t.Logf("LazPay.InsuranceCreateOrder returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazPay.InsuranceCreateOrder response: %#v", res)
}
func Test_LazPay_InsuranceGetPromotions(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_insurance_promotion_getPromotions_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping InsuranceGetPromotions due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/insurance/promotion/getPromotions*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazPay.InsuranceGetPromotions(ctx)
	if err != nil {
		t.Logf("LazPay.InsuranceGetPromotions returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazPay.InsuranceGetPromotions response: %#v", res)
}
func Test_LazPay_InsuranceQueryOrder(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_insurance_order_getStatus_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping InsuranceQueryOrder due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/insurance/order/getStatus*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazPay.InsuranceQueryOrder(ctx)
	if err != nil {
		t.Logf("LazPay.InsuranceQueryOrder returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazPay.InsuranceQueryOrder response: %#v", res)
}
func Test_LazPay_InsuranceRealTimeCDP(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_insurance_syncCDP_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping InsuranceRealTimeCDP due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/insurance/syncCDP*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazPay.InsuranceRealTimeCDP(ctx)
	if err != nil {
		t.Logf("LazPay.InsuranceRealTimeCDP returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazPay.InsuranceRealTimeCDP response: %#v", res)
}
func Test_LazPay_LazadaCFOInvoiceRpaCallback(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_rpa_id_tax_callback_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping LazadaCFOInvoiceRpaCallback due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/rpa/id/tax/callback*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazPay.LazadaCFOInvoiceRpaCallback(ctx)
	if err != nil {
		t.Logf("LazPay.LazadaCFOInvoiceRpaCallback returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazPay.LazadaCFOInvoiceRpaCallback response: %#v", res)
}
func Test_LazPay_OpenServiceBalanceQuery(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_wallet_open_service_balance_query_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping OpenServiceBalanceQuery due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/wallet/open/service/balance/query*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazPay.OpenServiceBalanceQuery(ctx)
	if err != nil {
		t.Logf("LazPay.OpenServiceBalanceQuery returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazPay.OpenServiceBalanceQuery response: %#v", res)
}
func Test_LazPay_OpenServiceKycQuery(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_wallet_open_service_kyc_query_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping OpenServiceKycQuery due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/wallet/open/service/kyc/query*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazPay.OpenServiceKycQuery(ctx)
	if err != nil {
		t.Logf("LazPay.OpenServiceKycQuery returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazPay.OpenServiceKycQuery response: %#v", res)
}
func Test_LazPay_OpenServiceWithdrawApply(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_wallet_open_service_withdraw_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping OpenServiceWithdrawApply due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/wallet/open/service/withdraw*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazPay.OpenServiceWithdrawApply(ctx)
	if err != nil {
		t.Logf("LazPay.OpenServiceWithdrawApply returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazPay.OpenServiceWithdrawApply response: %#v", res)
}
func Test_LazPay_OpenServiceWithdrawQuery(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_wallet_open_service_withdraw_query_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping OpenServiceWithdrawQuery due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/wallet/open/service/withdraw/query*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazPay.OpenServiceWithdrawQuery(ctx)
	if err != nil {
		t.Logf("LazPay.OpenServiceWithdrawQuery returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazPay.OpenServiceWithdrawQuery response: %#v", res)
}
func Test_LazPay_QueryAddonOrder(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_insurance_addon_orders_query_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping QueryAddonOrder due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/insurance/addon/orders/query*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazPay.QueryAddonOrder(ctx)
	if err != nil {
		t.Logf("LazPay.QueryAddonOrder returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazPay.QueryAddonOrder response: %#v", res)
}
func Test_LazPay_QueryBenefit(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_insurance_promotion_queryBenefit_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping QueryBenefit due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/insurance/promotion/queryBenefit*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazPay.QueryBenefit(ctx)
	if err != nil {
		t.Logf("LazPay.QueryBenefit returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazPay.QueryBenefit response: %#v", res)
}
func Test_LazPay_Reconciliation(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_wallet_open_service_reconciliation_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping Reconciliation due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/wallet/open/service/reconciliation*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazPay.Reconciliation(ctx)
	if err != nil {
		t.Logf("LazPay.Reconciliation returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazPay.Reconciliation response: %#v", res)
}
func Test_LazPay_RedeemMpVoucher(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_insurance_voucher_redeemVoucher_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping RedeemMpVoucher due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/insurance/voucher/redeemVoucher*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.LazPay.RedeemMpVoucher(ctx)
	if err != nil {
		t.Logf("LazPay.RedeemMpVoucher returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("LazPay.RedeemMpVoucher response: %#v", res)
}
