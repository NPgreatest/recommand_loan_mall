package response

type RecommendResponse struct {
	GoodsId       int     `json:"goodsId"`
	GoodsName     string  `json:"goodsName"`
	GoodsCoverImg string  `json:"goodsCoverImg"`
	SellingPrice  float64 `json:"sellingPrice"`
}
