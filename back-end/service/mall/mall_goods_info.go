package mall

import (
	"errors"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/mall"
	mallRes "main.go/model/mall/response"
	"main.go/model/manage"
	"main.go/utils"
)

type MallGoodsInfoService struct {
}

// MallGoodsListBySearch 商品搜索分页
func (m *MallGoodsInfoService) MallGoodsListBySearch(pageNumber int, goodsCategoryId int, keyword string, orderBy string) (err error, searchGoodsList []mallRes.GoodsSearchResponse, total int64) {
	// 根据搜索条件查询
	var goodsList []manage.MallGoodsInfo
	db := global.GVA_DB.Model(&manage.MallGoodsInfo{})
	if keyword != "" {
		db.Where("goods_name like ? or goods_intro like ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	if goodsCategoryId >= 0 {
		db.Where("goods_category_id= ?", goodsCategoryId)
	}
	err = db.Count(&total).Error
	if err != nil {
		global.GVA_LOG.Error("count total page err", zap.Error(err))
		return err, nil, 0
	}
	switch orderBy {
	case "new":
		db.Order("goods_id desc")
	case "price":
		db.Order("selling_price asc")
	default:
		db.Order("stock_num desc")
	}
	limit := 10
	offset := 10 * (pageNumber - 1)
	err = db.Limit(limit).Offset(offset).Find(&goodsList).Error
	// 返回查询结果
	for _, goods := range goodsList {
		searchGoods := mallRes.GoodsSearchResponse{
			GoodsId:       goods.GoodsId,
			GoodsName:     utils.SubStrLen(goods.GoodsName, 28),
			GoodsIntro:    utils.SubStrLen(goods.GoodsIntro, 28),
			GoodsCoverImg: goods.GoodsCoverImg,
			SellingPrice:  goods.SellingPrice,
		}
		searchGoodsList = append(searchGoodsList, searchGoods)
	}
	return
}

// GetMallGoodsInfo 获取商品信息
func (m *MallGoodsInfoService) GetMallGoodsInfo(id int) (err error, res mallRes.GoodsInfoDetailResponse) {
	var mallGoodsInfo manage.MallGoodsInfo
	err = global.GVA_DB.Where("goods_id = ?", id).First(&mallGoodsInfo).Error
	if mallGoodsInfo.GoodsSellStatus != 0 {
		return errors.New("商品已下架"), mallRes.GoodsInfoDetailResponse{}
	}
	err = copier.Copy(&res, &mallGoodsInfo)
	if err != nil {
		return err, mallRes.GoodsInfoDetailResponse{}
	}
	var list []string
	list = append(list, mallGoodsInfo.GoodsCarousel)
	res.GoodsCarouselList = list

	return
}

// GetMallGoodsInfo 获取评论信息
func (m *MallGoodsInfoService) GetMallReviewInfo(id int, pageNumber int) (err error, mallReviewInfo []*mallRes.MallGoodsReviewResponse) {
	var total int64
	db := global.GVA_DB.Model(&mall.MallGoodsReview{})
	db.Where("goods_id = ?", id)
	err = db.Count(&total).Error
	if err != nil {
		global.GVA_LOG.Error("count total page err", zap.Error(err))
		return err, nil
	}
	limit := 10
	offset := 10 * (pageNumber - 1)
	var reviews []mall.MallGoodsReview
	err = db.Limit(limit).Offset(offset).Find(&reviews).Error
	if err != nil {
		global.GVA_LOG.Error("find review err", zap.Error(err))
		return err, nil
	}
	copier.Copy(&mallReviewInfo, &reviews)
	ids := make([]int64, 0)
	for _, items := range reviews {
		ids = append(ids, items.UserId)
	}
	var users []mall.MallUser
	err = global.GVA_DB.Where("user_id IN ?", ids).Find(&users).Error
	if err != nil {
		global.GVA_LOG.Error("find review corresponding user err", zap.Error(err))
		return err, nil
	}
	infoMap := make(map[int]mall.MallUser)
	for _, items := range users {
		infoMap[items.UserId] = items
	}
	for _, items := range mallReviewInfo {
		items.Avatar = infoMap[items.UserId].Avatar
		items.NickName = infoMap[items.UserId].NickName
	}
	return
}
