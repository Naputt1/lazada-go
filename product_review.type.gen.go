package golazada

type GetHistoryReviewIdListResponse struct {
	BaseResponse                                    // Common response fields
	Response     GetHistoryReviewIdListResponseData `json:"data"` // Response data
}
type GetHistoryReviewIdListResponseData struct {
	Current  string   `json:"current"`   // [Required]
	Total    int64    `json:"total"`     // [Required]
	IdList   []string `json:"id_list"`   // [Required]
	PageSize int64    `json:"page_size"` // [Required]
}
type GetReviewListByIdListResponse struct {
	BaseResponse                                   // Common response fields
	Response     GetReviewListByIdListResponseData `json:"data"` // Response data
}
type GetReviewListByIdListResponseData struct {
	OutdatedReviews []string `json:"outdated_reviews"` // [Required]
	ReviewList      []Review `json:"review_list"`      // [Required]
}
type Ratings struct {
	SellerRating    string `json:"seller_rating"`    // [Required]
	OverallRating   string `json:"overall_rating"`   // [Required]
	LogisticsRating string `json:"logistics_rating"` // [Required]
	ProductRating   string `json:"product_rating"`   // [Required]
}
type Review struct {
	ReviewImages  []interface{}  `json:"review_images"`  // [Required]
	CanReply      string         `json:"can_reply"`      // [Required]
	CreateTime    string         `json:"create_time"`    // [Required]
	SubmitTime    string         `json:"submit_time"`    // [Required]
	ReviewContent string         `json:"review_content"` // [Required]
	Ratings       *Ratings       `json:"ratings"`        // [Required]
	ProductId     int64          `json:"product_id"`     // [Required]
	ReviewVideos  []ReviewVideos `json:"review_videos"`  // [Required]
	Id            int64          `json:"id"`             // [Required]
	SellerReply   string         `json:"seller_reply"`   // [Required]
	OrderId       int64          `json:"order_id"`       // [Required]
	ReviewType    string         `json:"review_type"`    // [Required]
}
type ReviewVideos struct {
	VideoUrl      string `json:"video_url"`       // [Required]
	VideoCoverUrl string `json:"video_cover_url"` // [Required]
}
type SubmitSellerReplyResponse struct {
	BaseResponse        // Common response fields
	Response     string `json:"data"` // Response data
}
