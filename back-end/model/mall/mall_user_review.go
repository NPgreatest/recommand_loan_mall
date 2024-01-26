package mall

// MallGoodsReview 代表mall_goods_review表的结构体
type MallGoodsReview struct {
	GoodsID       int64  `json:"goodsId" form:"goodsId" gorm:"column:goods_id;comment:商品ID;"`
	UserId        int64  `json:"userId" form:"userId" gorm:"column:user_id;comment:评论用户ID;"`
	ReviewTime    int    `json:"reviewTime" form:"reviewTime" gorm:"column:review_time;comment:评论时间;"`
	ReviewStar    int    `json:"reviewStar" form:"reviewStar" gorm:"column:review_star;comment:评论星级;"`
	ReviewTitle   string `json:"reviewTitle" form:"reviewTitle" gorm:"column:review_title;comment:评论标题;type:text;"`
	ReviewContent string `json:"reviewContent" form:"reviewContent" gorm:"column:review_content;comment:评论内容;type:text;"`
}

// TableName 设置MallGoodsReview的表名为mall_goods_review
func (MallGoodsReview) TableName() string {
	return "mall_goods_review"
}
