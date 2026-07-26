package golazada

type AddAdgroupBatchResponse struct {
	BaseResponse          // Common response fields
	AnalyseTraceId string `json:"analyseTraceId,omitempty"` //
	ErrorMsg       string `json:"errorMsg,omitempty"`       //
	Result         string `json:"result,omitempty"`         //
}
type AddSolutionResponse struct {
	BaseResponse               // Common response fields
	AnalyseTraceId string      `json:"analyseTraceId,omitempty"` //
	ErrorMsg       string      `json:"errorMsg,omitempty"`       //
	Result         interface{} `json:"result,omitempty"`         //
}
type AudienceViewDTO struct {
	AdCrowdTag string `json:"adCrowdTag"` // [Required]
	Discount   string `json:"discount"`   // [Required]
}
type ClickserverResponse struct {
	BaseResponse                                       // Common response fields
	Result       *GetPickUpStoreListResponseDataResult `json:"result,omitempty"` //
}
type DeleteAdgroupBatchResponse struct {
	BaseResponse          // Common response fields
	AnalyseTraceId string `json:"analyseTraceId,omitempty"` //
	ErrorMsg       string `json:"errorMsg,omitempty"`       //
	Result         string `json:"result,omitempty"`         //
}
type DeleteCampaignResponse struct {
	BaseResponse          // Common response fields
	AnalyseTraceId string `json:"analyseTraceId,omitempty"` //
	ErrorMsg       string `json:"errorMsg,omitempty"`       //
	Result         string `json:"result,omitempty"`         //
}
type GetAccountSignInfoResponse struct {
	BaseResponse               // Common response fields
	AnalyseTraceId string      `json:"analyseTraceId,omitempty"` //
	ErrorMsg       string      `json:"errorMsg,omitempty"`       //
	Result         interface{} `json:"result,omitempty"`         //
}
type GetAutoTopUpOptionOneConfigResponse struct {
	BaseResponse                                                  // Common response fields
	AnalyseTraceId string                                         `json:"analyseTraceId,omitempty"` //
	ErrorMsg       string                                         `json:"errorMsg,omitempty"`       //
	Result         *GetAutoTopUpOptionOneConfigResponseDataResult `json:"result,omitempty"`         //
}
type GetAutoTopUpOptionOneConfigResponseDataResult struct {
	LimitAmount string `json:"limitAmount"` // [Required]
	TopUpAmount string `json:"topUpAmount"` // [Required]
	Status      string `json:"status"`      // [Required]
}
type GetCampaignCountResponse struct {
	BaseResponse          // Common response fields
	AnalyseTraceId string `json:"analyseTraceId,omitempty"` //
	ErrorMsg       string `json:"errorMsg,omitempty"`       //
	Result         string `json:"result,omitempty"`         //
}
type GetCampaignResponse struct {
	BaseResponse                                  // Common response fields
	AnalyseTraceId string                         `json:"analyseTraceId,omitempty"` //
	ErrorMsg       string                         `json:"errorMsg,omitempty"`       //
	Result         *GetCampaignResponseDataResult `json:"result,omitempty"`         //
}
type GetCampaignResponseDataResult struct {
	CampaignObjective string        `json:"campaignObjective"` // [Required]
	CampaignType      string        `json:"campaignType"`      // [Required]
	EndDate           string        `json:"endDate"`           // [Required]
	CampaignId        string        `json:"campaignId"`        // [Required]
	OnlineStatus      string        `json:"onlineStatus"`      // [Required]
	SwitchStatus      string        `json:"switchStatus"`      // [Required]
	Platform          []interface{} `json:"platform"`          // [Required]
	BudgetUsedAmount  string        `json:"budgetUsedAmount"`  // [Required]
	AutoItemSelect    string        `json:"autoItemSelect"`    // [Required]
	CampaignModel     string        `json:"campaignModel"`     // [Required]
	MaxBid            string        `json:"maxBid"`            // [Required]
	HaveAdCount       string        `json:"haveAdCount"`       // [Required]
	SceneId           string        `json:"sceneId"`           // [Required]
	AutoCreative      string        `json:"autoCreative"`      // [Required]
	CampaignName      string        `json:"campaignName"`      // [Required]
	StartDate         string        `json:"startDate"`         // [Required]
	DayBudget         string        `json:"dayBudget"`         // [Required]
}
type GetDiscoveryReportAdgroupResponse struct {
	BaseResponse                                              // Common response fields
	Result       *GetDiscoveryReportAdgroupResponseDataResult `json:"result,omitempty"` //
}
type GetDiscoveryReportAdgroupResponseDataResult struct {
	Result         []ResultResult `json:"result"`         // [Required]
	ErrorKey       string         `json:"errorKey"`       // [Required]
	ErrorDTOList   []interface{}  `json:"errorDTOList"`   // [Required]
	Success        bool           `json:"success"`        // [Required]
	AnalyseTraceId string         `json:"analyseTraceId"` // [Required]
	ErrorCode      string         `json:"errorCode"`      // [Required]
	TotalCount     string         `json:"totalCount"`     // [Required]
	ErrorMsg       string         `json:"errorMsg"`       // [Required]
}
type GetDiscoveryReportAudienceResponse struct {
	BaseResponse                                               // Common response fields
	Result       *GetDiscoveryReportAudienceResponseDataResult `json:"result,omitempty"` //
}
type GetDiscoveryReportAudienceResponseDataResult struct {
	Result         []ResponseDataResultResult `json:"result"`         // [Required]
	ErrorKey       string                     `json:"errorKey"`       // [Required]
	ErrorDTOList   []interface{}              `json:"errorDTOList"`   // [Required]
	Success        bool                       `json:"success"`        // [Required]
	AnalyseTraceId string                     `json:"analyseTraceId"` // [Required]
	ErrorCode      string                     `json:"errorCode"`      // [Required]
	TotalCount     string                     `json:"totalCount"`     // [Required]
	ErrorMsg       string                     `json:"errorMsg"`       // [Required]
}
type GetDiscoveryReportCampaignResponse struct {
	BaseResponse                                               // Common response fields
	Result       *GetDiscoveryReportCampaignResponseDataResult `json:"result,omitempty"` //
}
type GetDiscoveryReportCampaignResponseDataResult struct {
	Result         []GetDiscoveryReportCampaignResponseDataResultResult `json:"result"`         // [Required]
	ErrorKey       string                                               `json:"errorKey"`       // [Required]
	Success        bool                                                 `json:"success"`        // [Required]
	AnalyseTraceId string                                               `json:"analyseTraceId"` // [Required]
	ErrorCode      string                                               `json:"errorCode"`      // [Required]
	TotalCount     string                                               `json:"totalCount"`     // [Required]
	ErrorMsg       string                                               `json:"errorMsg"`       // [Required]
}
type GetDiscoveryReportCampaignResponseDataResultResult struct {
	Ctr               string `json:"ctr"`               // [Required]
	CampaignObjective string `json:"campaignObjective"` // [Required]
	CampaignType      string `json:"campaignType"`      // [Required]
	CampaignId        string `json:"campaignId"`        // [Required]
	StoreRevenue      string `json:"storeRevenue"`      // [Required]
	StoreCvr          string `json:"storeCvr"`          // [Required]
	StoreA2c          string `json:"storeA2c"`          // [Required]
	StoreOrders       string `json:"storeOrders"`       // [Required]
	ProductUnitSold   string `json:"productUnitSold"`   // [Required]
	Impressions       string `json:"impressions"`       // [Required]
	ProductCvr        string `json:"productCvr"`        // [Required]
	ProductOrders     string `json:"productOrders"`     // [Required]
	StoreRoi          string `json:"storeRoi"`          // [Required]
	Cpc               string `json:"cpc"`               // [Required]
	Spend             string `json:"spend"`             // [Required]
	Clicks            string `json:"clicks"`            // [Required]
	ProductRevenue    string `json:"productRevenue"`    // [Required]
	StoreUnitSold     string `json:"storeUnitSold"`     // [Required]
	CampaignName      string `json:"campaignName"`      // [Required]
	ProductType       string `json:"productType"`       // [Required]
	DayBudget         string `json:"dayBudget"`         // [Required]
	ProductA2c        string `json:"productA2c"`        // [Required]
}
type GetDiscoveryReportKeywordResponse struct {
	BaseResponse                                              // Common response fields
	Result       *GetDiscoveryReportKeywordResponseDataResult `json:"result,omitempty"` //
}
type GetDiscoveryReportKeywordResponseDataResult struct {
	Result         []GetDiscoveryReportKeywordResponseDataResultResult `json:"result"`         // [Required]
	ErrorKey       string                                              `json:"errorKey"`       // [Required]
	ErrorDTOList   []interface{}                                       `json:"errorDTOList"`   // [Required]
	Success        bool                                                `json:"success"`        // [Required]
	AnalyseTraceId string                                              `json:"analyseTraceId"` // [Required]
	ErrorCode      string                                              `json:"errorCode"`      // [Required]
	TotalCount     string                                              `json:"totalCount"`     // [Required]
	ErrorMsg       string                                              `json:"errorMsg"`       // [Required]
}
type GetDiscoveryReportKeywordResponseDataResultResult struct {
	ProductImageUrl string `json:"productImageUrl"` // [Required]
	Ctr             string `json:"ctr"`             // [Required]
	KeywordId       string `json:"keywordId"`       // [Required]
	CampaignId      string `json:"campaignId"`      // [Required]
	StoreRevenue    string `json:"storeRevenue"`    // [Required]
	StoreCvr        string `json:"storeCvr"`        // [Required]
	StoreA2c        string `json:"storeA2c"`        // [Required]
	StoreOrders     string `json:"storeOrders"`     // [Required]
	ProductUnitSold string `json:"productUnitSold"` // [Required]
	Impressions     string `json:"impressions"`     // [Required]
	ProductCvr      string `json:"productCvr"`      // [Required]
	ProductOrders   string `json:"productOrders"`   // [Required]
	StoreRoi        string `json:"storeRoi"`        // [Required]
	AdgroupId       string `json:"adgroupId"`       // [Required]
	AdgroupName     string `json:"adgroupName"`     // [Required]
	Cpc             string `json:"cpc"`             // [Required]
	Spend           string `json:"spend"`           // [Required]
	MaxBid          string `json:"maxBid"`          // [Required]
	StoreUnitSold   string `json:"storeUnitSold"`   // [Required]
	Clicks          string `json:"clicks"`          // [Required]
	ProductRevenue  string `json:"productRevenue"`  // [Required]
	Keyword         string `json:"keyword"`         // [Required]
	CampaignName    string `json:"campaignName"`    // [Required]
	ProductA2c      string `json:"productA2c"`      // [Required]
}
type GetLatestSignInfoResponse struct {
	BaseResponse               // Common response fields
	AnalyseTraceId string      `json:"analyseTraceId,omitempty"` //
	ErrorMsg       string      `json:"errorMsg,omitempty"`       //
	Result         interface{} `json:"result,omitempty"`         //
}
type GetPickUpStoreListResponseDataResult struct {
	BizExtMap      interface{} `json:"biz_ext_map"`      // [Required]
	Headers        interface{} `json:"headers"`          // [Required]
	MsgCode        string      `json:"msg_code"`         // [Required]
	HttpStatusCode string      `json:"http_status_code"` // [Required]
	Success        bool        `json:"success"`          // [Required]
	MsgInfo        string      `json:"msg_info"`         // [Required]
	Model          interface{} `json:"model"`            // [Required]
	MappingCode    string      `json:"mapping_code"`     // [Required]
}
type GetReportCampaignOnFIrstSlotResponse struct {
	BaseResponse                                                 // Common response fields
	Result       *GetReportCampaignOnFIrstSlotResponseDataResult `json:"result,omitempty"` //
}
type GetReportCampaignOnFIrstSlotResponseDataResult struct {
	Result         []GetReportCampaignOnFIrstSlotResponseDataResultResult `json:"result"`         // [Required]
	ErrorKey       string                                                 `json:"errorKey"`       // [Required]
	ErrorDTOList   []interface{}                                          `json:"errorDTOList"`   // [Required]
	Success        bool                                                   `json:"success"`        // [Required]
	AnalyseTraceId string                                                 `json:"analyseTraceId"` // [Required]
	ErrorCode      string                                                 `json:"errorCode"`      // [Required]
	TotalCount     string                                                 `json:"totalCount"`     // [Required]
	ErrorMsg       string                                                 `json:"errorMsg"`       // [Required]
}
type GetReportCampaignOnFIrstSlotResponseDataResultResult struct {
	Ctr               string `json:"ctr"`               // [Required]
	CampaignObjective string `json:"campaignObjective"` // [Required]
	CampaignType      string `json:"campaignType"`      // [Required]
	FirstImpShare     string `json:"firstImpShare"`     // [Required]
	CampaignId        string `json:"campaignId"`        // [Required]
	StoreRevenue      string `json:"storeRevenue"`      // [Required]
	StoreCvr          string `json:"storeCvr"`          // [Required]
	StoreA2c          string `json:"storeA2c"`          // [Required]
	StoreOrders       string `json:"storeOrders"`       // [Required]
	ProductUnitSold   string `json:"productUnitSold"`   // [Required]
	Impressions       string `json:"impressions"`       // [Required]
	ProductCvr        string `json:"productCvr"`        // [Required]
	ProductOrders     string `json:"productOrders"`     // [Required]
	StoreRoi          string `json:"storeRoi"`          // [Required]
	Cpc               string `json:"cpc"`               // [Required]
	Spend             string `json:"spend"`             // [Required]
	Clicks            string `json:"clicks"`            // [Required]
	ProductRevenue    string `json:"productRevenue"`    // [Required]
	StoreUnitSold     string `json:"storeUnitSold"`     // [Required]
	CampaignName      string `json:"campaignName"`      // [Required]
	ProductType       string `json:"productType"`       // [Required]
	DayBudget         string `json:"dayBudget"`         // [Required]
	ProductA2c        string `json:"productA2c"`        // [Required]
}
type GetReportOverviewMetricResponse struct {
	BaseResponse                                              // Common response fields
	AnalyseTraceId string                                     `json:"analyseTraceId,omitempty"` //
	ErrorMsg       string                                     `json:"errorMsg,omitempty"`       //
	Result         *GetReportOverviewMetricResponseDataResult `json:"result,omitempty"`         //
}
type GetReportOverviewMetricResponseDataResult struct {
	MetricList []string      `json:"metricList"` // [Required]
	DateList   []string      `json:"dateList"`   // [Required]
	HourList   []interface{} `json:"hourList"`   // [Required]
}
type GetReportOverviewResponse struct {
	BaseResponse                                        // Common response fields
	AnalyseTraceId string                               `json:"analyseTraceId,omitempty"` //
	ErrorMsg       string                               `json:"errorMsg,omitempty"`       //
	Result         *GetReportOverviewResponseDataResult `json:"result,omitempty"`         //
}
type GetReportOverviewResponseDataResult struct {
	LastReportOverviewDetailDTO *LastReportOverviewDetailDTO `json:"lastReportOverviewDetailDTO"` // [Required]
	ReportOverviewDetailDTO     *LastReportOverviewDetailDTO `json:"reportOverviewDetailDTO"`     // [Required]
}
type LastReportOverviewDetailDTO struct {
	Ctr         string `json:"ctr"`         // [Required]
	Revenue     string `json:"revenue"`     // [Required]
	Spend       string `json:"spend"`       // [Required]
	UnitsSold   string `json:"unitsSold"`   // [Required]
	Cpc         string `json:"cpc"`         // [Required]
	Clicks      string `json:"clicks"`      // [Required]
	Impressions string `json:"impressions"` // [Required]
	Roi         string `json:"roi"`         // [Required]
}
type ListCategoryResponse struct {
	BaseResponse                                    // Common response fields
	AnalyseTraceId string                           `json:"analyseTraceId,omitempty"` //
	ErrorMsg       string                           `json:"errorMsg,omitempty"`       //
	Result         []ListCategoryResponseDataResult `json:"result,omitempty"`         //
}
type ListCategoryResponseDataResult struct {
	Selectable string `json:"selectable"` // [Required]
	Label      string `json:"label"`      // [Required]
	Value      string `json:"value"`      // [Required]
	IsLeaf     string `json:"isLeaf"`     // [Required]
}
type ListKeywordByAdgroupResponse struct {
	BaseResponse                                            // Common response fields
	AnalyseTraceId string                                   `json:"analyseTraceId,omitempty"` //
	ErrorMsg       string                                   `json:"errorMsg,omitempty"`       //
	Result         []ListKeywordByAdgroupResponseDataResult `json:"result,omitempty"`         //
	TotalCount     string                                   `json:"totalCount,omitempty"`     //
}
type ListKeywordByAdgroupResponseDataResult struct {
	SuggestedPrice     string `json:"suggestedPrice"`     // [Required]
	ReservePrice       string `json:"reservePrice"`       // [Required]
	Currency           string `json:"currency"`           // [Required]
	SoftLowerLimit     string `json:"softLowerLimit"`     // [Required]
	Keyword            string `json:"keyword"`            // [Required]
	SoftUpperLimit     string `json:"softUpperLimit"`     // [Required]
	Relevance          string `json:"relevance"`          // [Required]
	SoftUpperLimitType string `json:"softUpperLimitType"` // [Required]
	HistoricalPV       string `json:"historicalPV"`       // [Required]
}
type ListKeywordByItemResponse struct {
	BaseResponse                                            // Common response fields
	AnalyseTraceId string                                   `json:"analyseTraceId,omitempty"` //
	ErrorMsg       string                                   `json:"errorMsg,omitempty"`       //
	Result         []ListKeywordByAdgroupResponseDataResult `json:"result,omitempty"`         //
}
type ModifyAutoTopUpOptionOneConfigResponse struct {
	BaseResponse          // Common response fields
	AnalyseTraceId string `json:"analyseTraceId,omitempty"` //
	ErrorMsg       string `json:"errorMsg,omitempty"`       //
	Result         string `json:"result,omitempty"`         //
}
type ResponseDataResultResult struct {
	ProductImageUrl string `json:"productImageUrl"` // [Required]
	Ctr             string `json:"ctr"`             // [Required]
	CampaignId      string `json:"campaignId"`      // [Required]
	StoreRevenue    string `json:"storeRevenue"`    // [Required]
	StoreCvr        string `json:"storeCvr"`        // [Required]
	StoreA2c        string `json:"storeA2c"`        // [Required]
	StoreOrders     string `json:"storeOrders"`     // [Required]
	ProductUnitSold string `json:"productUnitSold"` // [Required]
	Impressions     string `json:"impressions"`     // [Required]
	ProductCvr      string `json:"productCvr"`      // [Required]
	ProductOrders   string `json:"productOrders"`   // [Required]
	AudienceFakeId  string `json:"audienceFakeId"`  // [Required]
	StoreRoi        string `json:"storeRoi"`        // [Required]
	AdgroupId       string `json:"adgroupId"`       // [Required]
	AudienceGroup   string `json:"audienceGroup"`   // [Required]
	AdgroupName     string `json:"adgroupName"`     // [Required]
	Cpc             string `json:"cpc"`             // [Required]
	Spend           string `json:"spend"`           // [Required]
	Clicks          string `json:"clicks"`          // [Required]
	ProductRevenue  string `json:"productRevenue"`  // [Required]
	StoreUnitSold   string `json:"storeUnitSold"`   // [Required]
	CampaignName    string `json:"campaignName"`    // [Required]
	ProductA2c      string `json:"productA2c"`      // [Required]
}
type ResultResult struct {
	DateRange       string `json:"dateRange"`       // [Required]
	ProductUnitSold string `json:"productUnitSold"` // [Required]
	ProductCvr      string `json:"productCvr"`      // [Required]
	ProductOrders   string `json:"productOrders"`   // [Required]
	AdgroupId       string `json:"adgroupId"`       // [Required]
	AdgroupName     string `json:"adgroupName"`     // [Required]
	Cpc             string `json:"cpc"`             // [Required]
	Spend           string `json:"spend"`           // [Required]
	StoreUnitSold   string `json:"storeUnitSold"`   // [Required]
	ProductA2c      string `json:"productA2c"`      // [Required]
	ProductImageUrl string `json:"productImageUrl"` // [Required]
	Ctr             string `json:"ctr"`             // [Required]
	CampaignId      string `json:"campaignId"`      // [Required]
	StoreRevenue    string `json:"storeRevenue"`    // [Required]
	StoreCvr        string `json:"storeCvr"`        // [Required]
	StoreA2c        string `json:"storeA2c"`        // [Required]
	StoreOrders     string `json:"storeOrders"`     // [Required]
	Impressions     string `json:"impressions"`     // [Required]
	BidPrice        string `json:"bidPrice"`        // [Required]
	ItemId          string `json:"itemId"`          // [Required]
	StoreRoi        string `json:"storeRoi"`        // [Required]
	MaxBid          string `json:"maxBid"`          // [Required]
	Clicks          string `json:"clicks"`          // [Required]
	ProductRevenue  string `json:"productRevenue"`  // [Required]
	CampaignName    string `json:"campaignName"`    // [Required]
}
type SearchAdgroupListResponse struct {
	BaseResponse                                         // Common response fields
	AnalyseTraceId string                                `json:"analyseTraceId,omitempty"` //
	ErrorMsg       string                                `json:"errorMsg,omitempty"`       //
	Result         []SearchAdgroupListResponseDataResult `json:"result,omitempty"`         //
	TotalCount     string                                `json:"totalCount,omitempty"`     //
}
type SearchAdgroupListResponseDataResult struct {
	UnitsSold                 string            `json:"unitsSold"`                 // [Required]
	ProductOrders             string            `json:"productOrders"`             // [Required]
	CampaignSwitchStatus      string            `json:"campaignSwitchStatus"`      // [Required]
	AdAccountBalanceStatus    string            `json:"adAccountBalanceStatus"`    // [Required]
	Revenue                   string            `json:"revenue"`                   // [Required]
	AdgroupId                 string            `json:"adgroupId"`                 // [Required]
	AdgroupName               string            `json:"adgroupName"`               // [Required]
	ImageUrl                  string            `json:"imageUrl"`                  // [Required]
	Spend                     string            `json:"spend"`                     // [Required]
	Cpc                       string            `json:"cpc"`                       // [Required]
	CampaignScheduleStatus    string            `json:"campaignScheduleStatus"`    // [Required]
	AdSwitchStatus            string            `json:"adSwitchStatus"`            // [Required]
	AutoCreative              string            `json:"autoCreative"`              // [Required]
	Ctr                       string            `json:"ctr"`                       // [Required]
	CampaignDailyBudgetStatus string            `json:"campaignDailyBudgetStatus"` // [Required]
	ProductEligibleStatus     string            `json:"productEligibleStatus"`     // [Required]
	SellerEligibleStatus      string            `json:"sellerEligibleStatus"`      // [Required]
	StoreRevenue              string            `json:"storeRevenue"`              // [Required]
	StoreOrders               string            `json:"storeOrders"`               // [Required]
	Impressions               string            `json:"impressions"`               // [Required]
	StoreUnitsSold            string            `json:"storeUnitsSold"`            // [Required]
	BidPrice                  string            `json:"bidPrice"`                  // [Required]
	AudienceViewDTOList       []AudienceViewDTO `json:"audienceViewDTOList"`       // [Required]
	ItemId                    string            `json:"itemId"`                    // [Required]
	StoreRoi                  string            `json:"storeRoi"`                  // [Required]
	ProductStockStatus        string            `json:"productStockStatus"`        // [Required]
	AdApproveStatus           string            `json:"adApproveStatus"`           // [Required]
	Clicks                    string            `json:"clicks"`                    // [Required]
	Status                    string            `json:"status"`                    // [Required]
}
type SearchCampaignListResponse struct {
	BaseResponse                                          // Common response fields
	AnalyseTraceId string                                 `json:"analyseTraceId,omitempty"` //
	ErrorMsg       string                                 `json:"errorMsg,omitempty"`       //
	Result         []SearchCampaignListResponseDataResult `json:"result,omitempty"`         //
	TotalCount     string                                 `json:"totalCount,omitempty"`     //
}
type SearchCampaignListResponseDataResult struct {
	Ctr                       string `json:"ctr"`                       // [Required]
	CampaignDailyBudgetStatus string `json:"campaignDailyBudgetStatus"` // [Required]
	EndDate                   string `json:"endDate"`                   // [Required]
	StoreRevenue              string `json:"storeRevenue"`              // [Required]
	CampaignId                string `json:"campaignId"`                // [Required]
	StoreOrders               string `json:"storeOrders"`               // [Required]
	Impressions               string `json:"impressions"`               // [Required]
	StoreUnitsSold            string `json:"storeUnitsSold"`            // [Required]
	CampaignSwitchStatus      string `json:"campaignSwitchStatus"`      // [Required]
	AdAccountBalanceStatus    string `json:"adAccountBalanceStatus"`    // [Required]
	StoreRoi                  string `json:"storeRoi"`                  // [Required]
	DailyBudget               string `json:"dailyBudget"`               // [Required]
	Cpc                       string `json:"cpc"`                       // [Required]
	Spend                     string `json:"spend"`                     // [Required]
	CampaignScheduleStatus    string `json:"campaignScheduleStatus"`    // [Required]
	Clicks                    string `json:"clicks"`                    // [Required]
	CampaignName              string `json:"campaignName"`              // [Required]
	HaveActiveAdStatus        string `json:"haveActiveAdStatus"`        // [Required]
	StartDate                 string `json:"startDate"`                 // [Required]
	Status                    string `json:"status"`                    // [Required]
}
type SearchKeywordResponse struct {
	BaseResponse                                            // Common response fields
	AnalyseTraceId string                                   `json:"analyseTraceId,omitempty"` //
	ErrorMsg       string                                   `json:"errorMsg,omitempty"`       //
	Result         []ListKeywordByAdgroupResponseDataResult `json:"result,omitempty"`         //
	TotalCount     string                                   `json:"totalCount,omitempty"`     //
}
type SearchProductWithPageResponse struct {
	BaseResponse                                             // Common response fields
	AnalyseTraceId string                                    `json:"analyseTraceId,omitempty"` //
	ErrorMsg       string                                    `json:"errorMsg,omitempty"`       //
	Result         []SearchProductWithPageResponseDataResult `json:"result,omitempty"`         //
	TotalCount     string                                    `json:"totalCount,omitempty"`     //
}
type SearchProductWithPageResponseDataResult struct {
	AvgSalesVolume     string   `json:"avgSalesVolume"`     // [Required]
	IsDigitalUtilities string   `json:"isDigitalUtilities"` // [Required]
	Inventory          string   `json:"inventory"`          // [Required]
	ProductName        string   `json:"productName"`        // [Required]
	BidPrice           string   `json:"bidPrice"`           // [Required]
	Ipv                string   `json:"ipv"`                // [Required]
	Tags               []string `json:"tags"`               // [Required]
	ItemId             string   `json:"itemId"`             // [Required]
	CompetitionIndex   string   `json:"competitionIndex"`   // [Required]
	ImageUrl           string   `json:"imageUrl"`           // [Required]
	IsBan              string   `json:"isBan"`              // [Required]
	PdpLink            string   `json:"pdpLink"`            // [Required]
	ContentScore       string   `json:"contentScore"`       // [Required]
	RetailPrice        string   `json:"retailPrice"`        // [Required]
	CategoryId         string   `json:"categoryId"`         // [Required]
	Cvr                string   `json:"cvr"`                // [Required]
}
type SignResponse struct {
	BaseResponse               // Common response fields
	AnalyseTraceId string      `json:"analyseTraceId,omitempty"` //
	ErrorMsg       string      `json:"errorMsg,omitempty"`       //
	Result         interface{} `json:"result,omitempty"`         //
}
type UpdateAdgroupBatchResponse struct {
	BaseResponse          // Common response fields
	AnalyseTraceId string `json:"analyseTraceId,omitempty"` //
	ErrorMsg       string `json:"errorMsg,omitempty"`       //
	Result         string `json:"result,omitempty"`         //
}
type UpdateCampaignResponse struct {
	BaseResponse               // Common response fields
	AnalyseTraceId string      `json:"analyseTraceId,omitempty"` //
	ErrorMsg       string      `json:"errorMsg,omitempty"`       //
	Result         interface{} `json:"result,omitempty"`         //
}
