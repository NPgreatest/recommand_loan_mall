package response

type GoodsSearchResponse struct {
	GoodsId       int     `json:"goodsId"`
	GoodsName     string  `json:"goodsName"`
	GoodsIntro    string  `json:"goodsIntro"`
	GoodsCoverImg string  `json:"goodsCoverImg"`
	SellingPrice  float64 `json:"sellingPrice"`
}

type GoodsInfoDetailResponse struct {
	GoodsId            int      `json:"goodsId"`
	GoodsName          string   `json:"goodsName"`
	GoodsIntro         string   `json:"goodsIntro"`
	GoodsCoverImg      string   `json:"goodsCoverImg"`
	SellingPrice       float64  `json:"sellingPrice"`
	GoodsDetailContent string   `json:"goodsDetailContent"  `
	OriginalPrice      int      `json:"originalPrice" `
	Tag                string   `json:"tag" form:"tag" `
	GoodsCarouselList  []string `json:"goodsCarouselList" `
}
