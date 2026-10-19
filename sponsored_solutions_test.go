package golazada

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_SponsoredSolutions_AdAlgoSellerBaseInfo(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solution_smax_adAlgoSellerBaseInfo_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping AdAlgoSellerBaseInfo due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sponsor/solution/smax/adAlgoSellerBaseInfo*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.AdAlgoSellerBaseInfo(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.AdAlgoSellerBaseInfo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.AdAlgoSellerBaseInfo response: %#v", res)
}
func Test_SponsoredSolutions_AddAdgroupBatch(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solutions_adgroup_addAdgroupBatch_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping AddAdgroupBatch due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sponsor/solutions/adgroup/addAdgroupBatch*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.AddAdgroupBatch(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.AddAdgroupBatch returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.AddAdgroupBatch response: %#v", res)
}
func Test_SponsoredSolutions_AddSMaxAASolution(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solution_smax_addSMaxAASolution_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping AddSMaxAASolution due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sponsor/solution/smax/addSMaxAASolution*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.AddSMaxAASolution(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.AddSMaxAASolution returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.AddSMaxAASolution response: %#v", res)
}
func Test_SponsoredSolutions_AddSolution(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solutions_addSolution_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping AddSolution due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sponsor/solutions/addSolution*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.AddSolution(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.AddSolution returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.AddSolution response: %#v", res)
}
func Test_SponsoredSolutions_Clickserver(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_gproject_ads_aidc_click_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping Clickserver due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/gproject/ads/aidc/click*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.Clickserver(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.Clickserver returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.Clickserver response: %#v", res)
}
func Test_SponsoredSolutions_DeleteAdgroupBatch(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solutions_adgroup_deleteAdgroupBatch_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DeleteAdgroupBatch due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sponsor/solutions/adgroup/deleteAdgroupBatch*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.DeleteAdgroupBatch(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.DeleteAdgroupBatch returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.DeleteAdgroupBatch response: %#v", res)
}
func Test_SponsoredSolutions_DeleteCampaign(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solutions_campaign_deleteCampaign_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DeleteCampaign due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sponsor/solutions/campaign/deleteCampaign*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.DeleteCampaign(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.DeleteCampaign returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.DeleteCampaign response: %#v", res)
}
func Test_SponsoredSolutions_GetAccountSignInfo(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solutions_account_getAccountSignInfo_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetAccountSignInfo due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/sponsor/solutions/account/getAccountSignInfo*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.GetAccountSignInfo(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.GetAccountSignInfo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.GetAccountSignInfo response: %#v", res)
}
func Test_SponsoredSolutions_GetAutoTopUpOptionOneConfig(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solutions_wallet_getAutoTopUpOptionOneConfig_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetAutoTopUpOptionOneConfig due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sponsor/solutions/wallet/getAutoTopUpOptionOneConfig*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.GetAutoTopUpOptionOneConfig(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.GetAutoTopUpOptionOneConfig returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.GetAutoTopUpOptionOneConfig response: %#v", res)
}
func Test_SponsoredSolutions_GetBlackItemDetail(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solution_smax_getBlackItemDetail_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetBlackItemDetail due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sponsor/solution/smax/getBlackItemDetail*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.GetBlackItemDetail(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.GetBlackItemDetail returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.GetBlackItemDetail response: %#v", res)
}
func Test_SponsoredSolutions_GetCampaign(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solutions_campaign_getCampaign_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetCampaign due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sponsor/solutions/campaign/getCampaign*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.GetCampaign(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.GetCampaign returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.GetCampaign response: %#v", res)
}
func Test_SponsoredSolutions_GetCampaignCount(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solutions_campaign_getCampaignCount_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetCampaignCount due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sponsor/solutions/campaign/getCampaignCount*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.GetCampaignCount(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.GetCampaignCount returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.GetCampaignCount response: %#v", res)
}
func Test_SponsoredSolutions_GetDiscoveryReportAdgroup(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solutions_report_getDiscoveryReportAdgroup_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetDiscoveryReportAdgroup due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sponsor/solutions/report/getDiscoveryReportAdgroup*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.GetDiscoveryReportAdgroup(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.GetDiscoveryReportAdgroup returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.GetDiscoveryReportAdgroup response: %#v", res)
}
func Test_SponsoredSolutions_GetDiscoveryReportAudience(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solutions_report_getDiscoveryReportAudience_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetDiscoveryReportAudience due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sponsor/solutions/report/getDiscoveryReportAudience*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.GetDiscoveryReportAudience(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.GetDiscoveryReportAudience returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.GetDiscoveryReportAudience response: %#v", res)
}
func Test_SponsoredSolutions_GetDiscoveryReportCampaign(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solutions_report_getDiscoveryReportCampaign_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetDiscoveryReportCampaign due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sponsor/solutions/report/getDiscoveryReportCampaign*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.GetDiscoveryReportCampaign(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.GetDiscoveryReportCampaign returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.GetDiscoveryReportCampaign response: %#v", res)
}
func Test_SponsoredSolutions_GetDiscoveryReportKeyword(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solutions_report_getDiscoveryReportKeyword_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetDiscoveryReportKeyword due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sponsor/solutions/report/getDiscoveryReportKeyword*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.GetDiscoveryReportKeyword(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.GetDiscoveryReportKeyword returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.GetDiscoveryReportKeyword response: %#v", res)
}
func Test_SponsoredSolutions_GetLatestSignInfo(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solutions_account_getLatestSignInfo_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetLatestSignInfo due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("%s/sponsor/solutions/account/getLatestSignInfo*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.GetLatestSignInfo(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.GetLatestSignInfo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.GetLatestSignInfo response: %#v", res)
}
func Test_SponsoredSolutions_GetReportCampaignOnFIrstSlot(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solutions_report_getReportCampaignOnPrePlacement_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetReportCampaignOnFIrstSlot due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sponsor/solutions/report/getReportCampaignOnPrePlacement*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.GetReportCampaignOnFIrstSlot(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.GetReportCampaignOnFIrstSlot returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.GetReportCampaignOnFIrstSlot response: %#v", res)
}
func Test_SponsoredSolutions_GetReportOverview(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solution_smax_report_getReportOverview_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetReportOverview due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sponsor/solution/smax/report/getReportOverview*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.GetReportOverview(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.GetReportOverview returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.GetReportOverview response: %#v", res)
}
func Test_SponsoredSolutions_GetReportOverview1(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solutions_report_getReportOverview_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetReportOverview1 due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sponsor/solutions/report/getReportOverview*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.GetReportOverview1(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.GetReportOverview1 returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.GetReportOverview1 response: %#v", res)
}
func Test_SponsoredSolutions_GetReportOverviewMetric(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solution_smax_report_getReportOverviewMetric_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetReportOverviewMetric due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sponsor/solution/smax/report/getReportOverviewMetric*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.GetReportOverviewMetric(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.GetReportOverviewMetric returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.GetReportOverviewMetric response: %#v", res)
}
func Test_SponsoredSolutions_GetReportOverviewMetric1(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solutions_report_getReportOverviewMetric_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetReportOverviewMetric1 due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sponsor/solutions/report/getReportOverviewMetric*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.GetReportOverviewMetric1(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.GetReportOverviewMetric1 returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.GetReportOverviewMetric1 response: %#v", res)
}
func Test_SponsoredSolutions_ListCategory(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solutions_category_listCategory_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping ListCategory due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sponsor/solutions/category/listCategory*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.ListCategory(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.ListCategory returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.ListCategory response: %#v", res)
}
func Test_SponsoredSolutions_ListKeywordByAdgroup(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solutions_keyword_listKeywordByAdgroup_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping ListKeywordByAdgroup due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sponsor/solutions/keyword/listKeywordByAdgroup*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.ListKeywordByAdgroup(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.ListKeywordByAdgroup returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.ListKeywordByAdgroup response: %#v", res)
}
func Test_SponsoredSolutions_ListKeywordByItem(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solutions_keyword_listKeywordByItem_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping ListKeywordByItem due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sponsor/solutions/keyword/listKeywordByItem*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.ListKeywordByItem(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.ListKeywordByItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.ListKeywordByItem response: %#v", res)
}
func Test_SponsoredSolutions_ModifyAutoTopUpOptionOneConfig(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solutions_wallet_modifyAutoTopUpOptionOneConfig_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping ModifyAutoTopUpOptionOneConfig due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sponsor/solutions/wallet/modifyAutoTopUpOptionOneConfig*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.ModifyAutoTopUpOptionOneConfig(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.ModifyAutoTopUpOptionOneConfig returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.ModifyAutoTopUpOptionOneConfig response: %#v", res)
}
func Test_SponsoredSolutions_SearchAdgroupList(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solutions_adgroup_searchAdgroupList_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SearchAdgroupList due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sponsor/solutions/adgroup/searchAdgroupList*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.SearchAdgroupList(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.SearchAdgroupList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.SearchAdgroupList response: %#v", res)
}
func Test_SponsoredSolutions_SearchCampaignList(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solutions_campaign_searchCampaignList_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SearchCampaignList due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sponsor/solutions/campaign/searchCampaignList*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.SearchCampaignList(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.SearchCampaignList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.SearchCampaignList response: %#v", res)
}
func Test_SponsoredSolutions_SearchKeyword(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solutions_keyword_searchKeyword_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SearchKeyword due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sponsor/solutions/keyword/searchKeyword*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.SearchKeyword(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.SearchKeyword returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.SearchKeyword response: %#v", res)
}
func Test_SponsoredSolutions_SearchProductWithPage(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solutions_product_searchProductWithPage_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SearchProductWithPage due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sponsor/solutions/product/searchProductWithPage*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.SearchProductWithPage(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.SearchProductWithPage returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.SearchProductWithPage response: %#v", res)
}
func Test_SponsoredSolutions_SearchSmaxCampaign(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solution_smax_searchSmaxCampaign_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SearchSmaxCampaign due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sponsor/solution/smax/searchSmaxCampaign*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.SearchSmaxCampaign(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.SearchSmaxCampaign returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.SearchSmaxCampaign response: %#v", res)
}
func Test_SponsoredSolutions_Sign(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solutions_account_sign_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping Sign due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sponsor/solutions/account/sign*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.Sign(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.Sign returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.Sign response: %#v", res)
}
func Test_SponsoredSolutions_SmaxAaGmvPrediction(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solution_smax_smaxAaGmvPrediction_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SmaxAaGmvPrediction due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sponsor/solution/smax/smaxAaGmvPrediction*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.SmaxAaGmvPrediction(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.SmaxAaGmvPrediction returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.SmaxAaGmvPrediction response: %#v", res)
}
func Test_SponsoredSolutions_UpdateAdgroupBatch(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solutions_adgroup_updateAdgroupBatch_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateAdgroupBatch due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sponsor/solutions/adgroup/updateAdgroupBatch*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.UpdateAdgroupBatch(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.UpdateAdgroupBatch returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.UpdateAdgroupBatch response: %#v", res)
}
func Test_SponsoredSolutions_UpdateCampaign(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solutions_campaign_updateCampaign_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateCampaign due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sponsor/solutions/campaign/updateCampaign*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.UpdateCampaign(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.UpdateCampaign returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.UpdateCampaign response: %#v", res)
}
func Test_SponsoredSolutions_UpdateItemBlackList(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solution_smax_updateItemBlackList_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateItemBlackList due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sponsor/solution/smax/updateItemBlackList*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.UpdateItemBlackList(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.UpdateItemBlackList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.UpdateItemBlackList response: %#v", res)
}
func Test_SponsoredSolutions_UpdateSmaxCampaign(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solution_smax_updateSmaxCampaign_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateSmaxCampaign due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sponsor/solution/smax/updateSmaxCampaign*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.UpdateSmaxCampaign(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.UpdateSmaxCampaign returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.UpdateSmaxCampaign response: %#v", res)
}
func Test_SponsoredSolutions_UpdateSmaxStatusWithoutAd(t *testing.T) {
	setup()
	defer teardown()

	serverURL := client.getServerURL()
	fixture := "_sponsor_solution_smax_updateSmaxStatusWithoutAd_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateSmaxStatusWithoutAd due to missing fixture: %v", err)
	}

	mockResp := map[string]interface{}{
		"code": "0",
		"data": data,
	}
	mockData, _ := json.Marshal(mockResp)

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/sponsor/solution/smax/updateSmaxStatusWithoutAd*", serverURL),
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, string(mockData))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		},
	)

	ctx := context.Background()
	res, err := client.SponsoredSolutions.UpdateSmaxStatusWithoutAd(ctx)
	if err != nil {
		t.Logf("SponsoredSolutions.UpdateSmaxStatusWithoutAd returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("SponsoredSolutions.UpdateSmaxStatusWithoutAd response: %#v", res)
}
