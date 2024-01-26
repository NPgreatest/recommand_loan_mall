package response

type MallGoodsReviewResponse struct {
	GoodsID       int64  `json:"goodsId"`
	UserId        int    `json:"userId" `
	NickName      string `json:"nickName"`
	Avatar        string `json:"avatar"`
	ReviewTime    int    `json:"reviewTime"`
	ReviewStar    int    `json:"reviewStar"`
	ReviewTitle   string `json:"reviewTitle"`
	ReviewContent string `json:"reviewContent"`
}
