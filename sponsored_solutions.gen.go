package golazada

import (
	"context"
)

type SponsoredSolutionsService interface {
	// AddAdgroupBatch Do add adgroup for one campaign.
	// Path: /sponsor/solutions/adgroup/addAdgroupBatch
	AddAdgroupBatch(ctx context.Context) (*AddAdgroupBatchResponse, error)
	// AddSolution Add sponsor solution
	// Path: /sponsor/solutions/addSolution
	AddSolution(ctx context.Context) (*AddSolutionResponse, error)
	// Clickserver aidc click server interface
	// Path: /gproject/ads/aidc/click
	Clickserver(ctx context.Context) (*ClickserverResponse, error)
	// DeleteAdgroupBatch Delete adgroup batch.
	// Path: /sponsor/solutions/adgroup/deleteAdgroupBatch
	DeleteAdgroupBatch(ctx context.Context) (*DeleteAdgroupBatchResponse, error)
	// DeleteCampaign Delete campaign.
	// Path: /sponsor/solutions/campaign/deleteCampaign
	DeleteCampaign(ctx context.Context) (*DeleteCampaignResponse, error)
	// GetAccountSignInfo Get seller account sign status.
	// Path: /sponsor/solutions/account/getAccountSignInfo
	GetAccountSignInfo(ctx context.Context) (*GetAccountSignInfoResponse, error)
	// GetAutoTopUpOptionOneConfig Get auto top up option one config.
	// Path: /sponsor/solutions/wallet/getAutoTopUpOptionOneConfig
	GetAutoTopUpOptionOneConfig(ctx context.Context) (*GetAutoTopUpOptionOneConfigResponse, error)
	// GetCampaign Get campaign list with bizCode by seller.
	// Path: /sponsor/solutions/campaign/getCampaign
	GetCampaign(ctx context.Context) (*GetCampaignResponse, error)
	// GetCampaignCount Get campaign count with bizCode for each solution type.
	// Path: /sponsor/solutions/campaign/getCampaignCount
	GetCampaignCount(ctx context.Context) (*GetCampaignCountResponse, error)
	// GetDiscoveryReportAdgroup Get sponsored discovery report adgroup level
	// Path: /sponsor/solutions/report/getDiscoveryReportAdgroup
	GetDiscoveryReportAdgroup(ctx context.Context) (*GetDiscoveryReportAdgroupResponse, error)
	// GetDiscoveryReportAudience Get sponsored discovery report audience level
	// Path: /sponsor/solutions/report/getDiscoveryReportAudience
	GetDiscoveryReportAudience(ctx context.Context) (*GetDiscoveryReportAudienceResponse, error)
	// GetDiscoveryReportCampaign Get sponsored discovery report campaign level
	// Path: /sponsor/solutions/report/getDiscoveryReportCampaign
	GetDiscoveryReportCampaign(ctx context.Context) (*GetDiscoveryReportCampaignResponse, error)
	// GetDiscoveryReportKeyword Get sponsored discovery report keyword level
	// Path: /sponsor/solutions/report/getDiscoveryReportKeyword
	GetDiscoveryReportKeyword(ctx context.Context) (*GetDiscoveryReportKeywordResponse, error)
	// GetLatestSignInfo Get the latest url of sign(T&C).
	// Path: /sponsor/solutions/account/getLatestSignInfo
	GetLatestSignInfo(ctx context.Context) (*GetLatestSignInfoResponse, error)
	// GetReportCampaignOnFIrstSlot Get sponsored discovery report campaign first slot
	// Path: /sponsor/solutions/report/getReportCampaignOnPrePlacement
	GetReportCampaignOnFIrstSlot(ctx context.Context) (*GetReportCampaignOnFIrstSlotResponse, error)
	// GetReportOverview Get report overview.
	// Path: /sponsor/solutions/report/getReportOverview
	GetReportOverview(ctx context.Context) (*GetReportOverviewResponse, error)
	// GetReportOverviewMetric get report overview metric
	// Path: /sponsor/solutions/report/getReportOverviewMetric
	GetReportOverviewMetric(ctx context.Context) (*GetReportOverviewMetricResponse, error)
	// ListCategory list category
	// Path: /sponsor/solutions/category/listCategory
	ListCategory(ctx context.Context) (*ListCategoryResponse, error)
	// ListKeywordByAdgroup List keyword by adgroup.
	// Path: /sponsor/solutions/keyword/listKeywordByAdgroup
	ListKeywordByAdgroup(ctx context.Context) (*ListKeywordByAdgroupResponse, error)
	// ListKeywordByItem List keyword by item.
	// Path: /sponsor/solutions/keyword/listKeywordByItem
	ListKeywordByItem(ctx context.Context) (*ListKeywordByItemResponse, error)
	// ModifyAutoTopUpOptionOneConfig Modify auto top up option one config.1. each country has differect tax rate
	// 2. we have minimum and maximam top-up amount limitation.For SG, min=5, max = 8,333,333,330;for PH, min=100,Max=17,895,600;for TH, min=100,max=8,333,333,300;for VN, min=50,000,max=833,333,300,000;for MY,min=10,max=8,333,333,330;for ID,min=25,000,max=8,333,333,000.the api timeout is 3s, max qps is 100, make sure do not over these num, especially qps, otherwise you may be blacklisted or limited request count for a while.
	// Path: /sponsor/solutions/wallet/modifyAutoTopUpOptionOneConfig
	ModifyAutoTopUpOptionOneConfig(ctx context.Context) (*ModifyAutoTopUpOptionOneConfigResponse, error)
	// SearchAdgroupList Search adgroup with bizCode by seller.
	// Path: /sponsor/solutions/adgroup/searchAdgroupList
	SearchAdgroupList(ctx context.Context) (*SearchAdgroupListResponse, error)
	// SearchCampaignList Search campaign list with bizCode for sellers.
	// Path: /sponsor/solutions/campaign/searchCampaignList
	SearchCampaignList(ctx context.Context) (*SearchCampaignListResponse, error)
	// SearchKeyword Search keyword with specific word.
	// Path: /sponsor/solutions/keyword/searchKeyword
	SearchKeyword(ctx context.Context) (*SearchKeywordResponse, error)
	// SearchProductWithPage Search product.
	// Path: /sponsor/solutions/product/searchProductWithPage
	SearchProductWithPage(ctx context.Context) (*SearchProductWithPageResponse, error)
	// Sign Description: Do sign for seller. Seller or agencies can use this api to sign up the t&c.
	// Timeout Period： the api timeout is 10s, max qps is 300, make sure do not over these num, especially qps, otherwise you may be blacklisted or limited request count for a while.
	// Path: /sponsor/solutions/account/sign
	Sign(ctx context.Context) (*SignResponse, error)
	// UpdateAdgroupBatch Update adgroup batch.
	// Path: /sponsor/solutions/adgroup/updateAdgroupBatch
	UpdateAdgroupBatch(ctx context.Context) (*UpdateAdgroupBatchResponse, error)
	// UpdateCampaign Update campaign with status field.
	// Path: /sponsor/solutions/campaign/updateCampaign
	UpdateCampaign(ctx context.Context) (*UpdateCampaignResponse, error)
}

type SponsoredSolutionsServiceOp[T any] struct {
	client *Client[T]
}

// AddAdgroupBatch Do add adgroup for one campaign.
// Path: /sponsor/solutions/adgroup/addAdgroupBatch
func (s *SponsoredSolutionsServiceOp[T]) AddAdgroupBatch(ctx context.Context) (*AddAdgroupBatchResponse, error) {
	path := "/sponsor/solutions/adgroup/addAdgroupBatch"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(AddAdgroupBatchResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// AddSolution Add sponsor solution
// Path: /sponsor/solutions/addSolution
func (s *SponsoredSolutionsServiceOp[T]) AddSolution(ctx context.Context) (*AddSolutionResponse, error) {
	path := "/sponsor/solutions/addSolution"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(AddSolutionResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// Clickserver aidc click server interface
// Path: /gproject/ads/aidc/click
func (s *SponsoredSolutionsServiceOp[T]) Clickserver(ctx context.Context) (*ClickserverResponse, error) {
	path := "/gproject/ads/aidc/click"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(ClickserverResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// DeleteAdgroupBatch Delete adgroup batch.
// Path: /sponsor/solutions/adgroup/deleteAdgroupBatch
func (s *SponsoredSolutionsServiceOp[T]) DeleteAdgroupBatch(ctx context.Context) (*DeleteAdgroupBatchResponse, error) {
	path := "/sponsor/solutions/adgroup/deleteAdgroupBatch"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(DeleteAdgroupBatchResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// DeleteCampaign Delete campaign.
// Path: /sponsor/solutions/campaign/deleteCampaign
func (s *SponsoredSolutionsServiceOp[T]) DeleteCampaign(ctx context.Context) (*DeleteCampaignResponse, error) {
	path := "/sponsor/solutions/campaign/deleteCampaign"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(DeleteCampaignResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetAccountSignInfo Get seller account sign status.
// Path: /sponsor/solutions/account/getAccountSignInfo
func (s *SponsoredSolutionsServiceOp[T]) GetAccountSignInfo(ctx context.Context) (*GetAccountSignInfoResponse, error) {
	path := "/sponsor/solutions/account/getAccountSignInfo"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetAccountSignInfoResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetAutoTopUpOptionOneConfig Get auto top up option one config.
// Path: /sponsor/solutions/wallet/getAutoTopUpOptionOneConfig
func (s *SponsoredSolutionsServiceOp[T]) GetAutoTopUpOptionOneConfig(ctx context.Context) (*GetAutoTopUpOptionOneConfigResponse, error) {
	path := "/sponsor/solutions/wallet/getAutoTopUpOptionOneConfig"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(GetAutoTopUpOptionOneConfigResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetCampaign Get campaign list with bizCode by seller.
// Path: /sponsor/solutions/campaign/getCampaign
func (s *SponsoredSolutionsServiceOp[T]) GetCampaign(ctx context.Context) (*GetCampaignResponse, error) {
	path := "/sponsor/solutions/campaign/getCampaign"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(GetCampaignResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetCampaignCount Get campaign count with bizCode for each solution type.
// Path: /sponsor/solutions/campaign/getCampaignCount
func (s *SponsoredSolutionsServiceOp[T]) GetCampaignCount(ctx context.Context) (*GetCampaignCountResponse, error) {
	path := "/sponsor/solutions/campaign/getCampaignCount"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(GetCampaignCountResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetDiscoveryReportAdgroup Get sponsored discovery report adgroup level
// Path: /sponsor/solutions/report/getDiscoveryReportAdgroup
func (s *SponsoredSolutionsServiceOp[T]) GetDiscoveryReportAdgroup(ctx context.Context) (*GetDiscoveryReportAdgroupResponse, error) {
	path := "/sponsor/solutions/report/getDiscoveryReportAdgroup"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(GetDiscoveryReportAdgroupResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetDiscoveryReportAudience Get sponsored discovery report audience level
// Path: /sponsor/solutions/report/getDiscoveryReportAudience
func (s *SponsoredSolutionsServiceOp[T]) GetDiscoveryReportAudience(ctx context.Context) (*GetDiscoveryReportAudienceResponse, error) {
	path := "/sponsor/solutions/report/getDiscoveryReportAudience"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(GetDiscoveryReportAudienceResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetDiscoveryReportCampaign Get sponsored discovery report campaign level
// Path: /sponsor/solutions/report/getDiscoveryReportCampaign
func (s *SponsoredSolutionsServiceOp[T]) GetDiscoveryReportCampaign(ctx context.Context) (*GetDiscoveryReportCampaignResponse, error) {
	path := "/sponsor/solutions/report/getDiscoveryReportCampaign"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(GetDiscoveryReportCampaignResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetDiscoveryReportKeyword Get sponsored discovery report keyword level
// Path: /sponsor/solutions/report/getDiscoveryReportKeyword
func (s *SponsoredSolutionsServiceOp[T]) GetDiscoveryReportKeyword(ctx context.Context) (*GetDiscoveryReportKeywordResponse, error) {
	path := "/sponsor/solutions/report/getDiscoveryReportKeyword"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(GetDiscoveryReportKeywordResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetLatestSignInfo Get the latest url of sign(T&C).
// Path: /sponsor/solutions/account/getLatestSignInfo
func (s *SponsoredSolutionsServiceOp[T]) GetLatestSignInfo(ctx context.Context) (*GetLatestSignInfoResponse, error) {
	path := "/sponsor/solutions/account/getLatestSignInfo"
	var params map[string]string
	wrapper, err := s.client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}
	resp := new(GetLatestSignInfoResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetReportCampaignOnFIrstSlot Get sponsored discovery report campaign first slot
// Path: /sponsor/solutions/report/getReportCampaignOnPrePlacement
func (s *SponsoredSolutionsServiceOp[T]) GetReportCampaignOnFIrstSlot(ctx context.Context) (*GetReportCampaignOnFIrstSlotResponse, error) {
	path := "/sponsor/solutions/report/getReportCampaignOnPrePlacement"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(GetReportCampaignOnFIrstSlotResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetReportOverview Get report overview.
// Path: /sponsor/solutions/report/getReportOverview
func (s *SponsoredSolutionsServiceOp[T]) GetReportOverview(ctx context.Context) (*GetReportOverviewResponse, error) {
	path := "/sponsor/solutions/report/getReportOverview"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(GetReportOverviewResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// GetReportOverviewMetric get report overview metric
// Path: /sponsor/solutions/report/getReportOverviewMetric
func (s *SponsoredSolutionsServiceOp[T]) GetReportOverviewMetric(ctx context.Context) (*GetReportOverviewMetricResponse, error) {
	path := "/sponsor/solutions/report/getReportOverviewMetric"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(GetReportOverviewMetricResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// ListCategory list category
// Path: /sponsor/solutions/category/listCategory
func (s *SponsoredSolutionsServiceOp[T]) ListCategory(ctx context.Context) (*ListCategoryResponse, error) {
	path := "/sponsor/solutions/category/listCategory"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(ListCategoryResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// ListKeywordByAdgroup List keyword by adgroup.
// Path: /sponsor/solutions/keyword/listKeywordByAdgroup
func (s *SponsoredSolutionsServiceOp[T]) ListKeywordByAdgroup(ctx context.Context) (*ListKeywordByAdgroupResponse, error) {
	path := "/sponsor/solutions/keyword/listKeywordByAdgroup"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(ListKeywordByAdgroupResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// ListKeywordByItem List keyword by item.
// Path: /sponsor/solutions/keyword/listKeywordByItem
func (s *SponsoredSolutionsServiceOp[T]) ListKeywordByItem(ctx context.Context) (*ListKeywordByItemResponse, error) {
	path := "/sponsor/solutions/keyword/listKeywordByItem"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(ListKeywordByItemResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// ModifyAutoTopUpOptionOneConfig Modify auto top up option one config.1. each country has differect tax rate
// 2. we have minimum and maximam top-up amount limitation.For SG, min=5, max = 8,333,333,330;for PH, min=100,Max=17,895,600;for TH, min=100,max=8,333,333,300;for VN, min=50,000,max=833,333,300,000;for MY,min=10,max=8,333,333,330;for ID,min=25,000,max=8,333,333,000.the api timeout is 3s, max qps is 100, make sure do not over these num, especially qps, otherwise you may be blacklisted or limited request count for a while.
// Path: /sponsor/solutions/wallet/modifyAutoTopUpOptionOneConfig
func (s *SponsoredSolutionsServiceOp[T]) ModifyAutoTopUpOptionOneConfig(ctx context.Context) (*ModifyAutoTopUpOptionOneConfigResponse, error) {
	path := "/sponsor/solutions/wallet/modifyAutoTopUpOptionOneConfig"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(ModifyAutoTopUpOptionOneConfigResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// SearchAdgroupList Search adgroup with bizCode by seller.
// Path: /sponsor/solutions/adgroup/searchAdgroupList
func (s *SponsoredSolutionsServiceOp[T]) SearchAdgroupList(ctx context.Context) (*SearchAdgroupListResponse, error) {
	path := "/sponsor/solutions/adgroup/searchAdgroupList"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(SearchAdgroupListResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// SearchCampaignList Search campaign list with bizCode for sellers.
// Path: /sponsor/solutions/campaign/searchCampaignList
func (s *SponsoredSolutionsServiceOp[T]) SearchCampaignList(ctx context.Context) (*SearchCampaignListResponse, error) {
	path := "/sponsor/solutions/campaign/searchCampaignList"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(SearchCampaignListResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// SearchKeyword Search keyword with specific word.
// Path: /sponsor/solutions/keyword/searchKeyword
func (s *SponsoredSolutionsServiceOp[T]) SearchKeyword(ctx context.Context) (*SearchKeywordResponse, error) {
	path := "/sponsor/solutions/keyword/searchKeyword"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(SearchKeywordResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// SearchProductWithPage Search product.
// Path: /sponsor/solutions/product/searchProductWithPage
func (s *SponsoredSolutionsServiceOp[T]) SearchProductWithPage(ctx context.Context) (*SearchProductWithPageResponse, error) {
	path := "/sponsor/solutions/product/searchProductWithPage"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(SearchProductWithPageResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// Sign Description: Do sign for seller. Seller or agencies can use this api to sign up the t&c.
// Timeout Period： the api timeout is 10s, max qps is 300, make sure do not over these num, especially qps, otherwise you may be blacklisted or limited request count for a while.
// Path: /sponsor/solutions/account/sign
func (s *SponsoredSolutionsServiceOp[T]) Sign(ctx context.Context) (*SignResponse, error) {
	path := "/sponsor/solutions/account/sign"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(SignResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// UpdateAdgroupBatch Update adgroup batch.
// Path: /sponsor/solutions/adgroup/updateAdgroupBatch
func (s *SponsoredSolutionsServiceOp[T]) UpdateAdgroupBatch(ctx context.Context) (*UpdateAdgroupBatchResponse, error) {
	path := "/sponsor/solutions/adgroup/updateAdgroupBatch"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(UpdateAdgroupBatchResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}

// UpdateCampaign Update campaign with status field.
// Path: /sponsor/solutions/campaign/updateCampaign
func (s *SponsoredSolutionsServiceOp[T]) UpdateCampaign(ctx context.Context) (*UpdateCampaignResponse, error) {
	path := "/sponsor/solutions/campaign/updateCampaign"
	var params map[string]string
	wrapper, err := s.client.Post(ctx, path, params, nil)
	if err != nil {
		return nil, err
	}
	resp := new(UpdateCampaignResponse)
	resp.Code = wrapper.Code
	resp.Type = wrapper.Type
	resp.Message = wrapper.Message
	resp.RequestID = wrapper.RequestID
	return resp, nil
}
