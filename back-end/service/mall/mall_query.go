package mall

import (
	"fmt"
	"main.go/global"
	mallRes "main.go/model/mall/response"
	"main.go/model/manage"
	"main.go/utils"
)

type MallQueryService struct {
}

func (m *MallQueryService) TextToItem(token string) (err error, cartItems []mallRes.RecommendResponse) {
	embedding, err := utils.GetEmbedding(token)
	if err != nil {
		return err, nil
	}
	fmt.Println(len(embedding))
	ids, err := utils.PgvectorGetId(embedding)
	fmt.Println(ids)
	var resItems []manage.MallGoodsInfo
	err = global.GVA_DB.Where("goods_id IN (?)", ids).Find(&resItems).Error
	if err != nil {
		return err, nil
	}
	for _, items := range resItems {
		cartItems = append(cartItems, mallRes.RecommendResponse{
			GoodsId:       items.GoodsId,
			GoodsName:     items.GoodsName,
			GoodsCoverImg: items.GoodsCoverImg,
			SellingPrice:  items.SellingPrice,
		})
	}
	return
}
