package mall

import (
	"errors"
	"time"

	"github.com/jinzhu/copier"
	"main.go/global"
	"main.go/model/common"
	"main.go/model/mall"
	mallReq "main.go/model/mall/request"
	mallRes "main.go/model/mall/response"
	"main.go/model/manage"
	"main.go/utils"
)

type MallShopCartService struct {
}

// GetMyShoppingCartItems 不分页
func (m *MallShopCartService) GetMyShoppingCartItems(userID int) (err error, cartItems []mallRes.CartItemResponse) {
	var shopCartItems []mall.MallShoppingCartItem
	var goodsInfos []manage.MallGoodsInfo
	global.GVA_DB.Where("user_id=? and is_deleted = 0", userID).Find(&shopCartItems)
	var goodsIds []int
	for _, shopcartItem := range shopCartItems {
		goodsIds = append(goodsIds, shopcartItem.GoodsId)
	}
	global.GVA_DB.Where("goods_id in ?", goodsIds).Find(&goodsInfos)
	goodsMap := make(map[int]manage.MallGoodsInfo)
	for _, goodsInfo := range goodsInfos {
		goodsMap[goodsInfo.GoodsId] = goodsInfo
	}
	for _, v := range shopCartItems {
		var cartItem mallRes.CartItemResponse
		copier.Copy(&cartItem, &v)
		if _, ok := goodsMap[v.GoodsId]; ok {
			goodsInfo := goodsMap[v.GoodsId]
			cartItem.GoodsName = goodsInfo.GoodsName
			cartItem.GoodsCoverImg = goodsInfo.GoodsCoverImg
			cartItem.SellingPrice = goodsInfo.SellingPrice
		}
		cartItems = append(cartItems, cartItem)
	}

	return
}

func (m *MallShopCartService) SaveMallCartItem(userID int, req mallReq.SaveCartItemParam) (err error) {
	if req.GoodsCount < 1 {
		return errors.New("商品数量不能小于 1 ！")

	}
	if req.GoodsCount > 5 {
		return errors.New("超出单个商品的最大购买数量！")
	}
	var shopCartItems []mall.MallShoppingCartItem
	// 是否已存在商品
	err = global.GVA_DB.Where("user_id = ? and goods_id = ? and is_deleted = 0", userID, req.GoodsId).Find(&shopCartItems).Error
	if err != nil {
		return errors.New("已存在！无需重复添加！")
	}
	err = global.GVA_DB.Where("goods_id = ? ", req.GoodsId).First(&manage.MallGoodsInfo{}).Error
	if err != nil {
		return errors.New(" 商品为空")
	}
	var total int64
	global.GVA_DB.Where("user_id =?  and is_deleted = 0", userID).Count(&total)
	if total > 20 {
		return errors.New("超出购物车最大容量！")
	}
	var shopCartItem mall.MallShoppingCartItem
	if err = copier.Copy(&shopCartItem, &req); err != nil {
		return err
	}
	shopCartItem.UserId = userID
	shopCartItem.CreateTime = common.JSONTime{Time: time.Now()}
	shopCartItem.UpdateTime = common.JSONTime{Time: time.Now()}
	err = global.GVA_DB.Save(&shopCartItem).Error
	return
}

func (m *MallShopCartService) UpdateMallCartItem(userID int, req mallReq.UpdateCartItemParam) (err error) {
	//超出单个商品的最大数量
	if req.GoodsCount > 5 {
		return errors.New("超出单个商品的最大购买数量！")
	}
	var shopCartItem mall.MallShoppingCartItem
	if err = global.GVA_DB.Where("cart_item_id=? and is_deleted = 0", req.CartItemId).First(&shopCartItem).Error; err != nil {
		return errors.New("未查询到记录！")
	}
	if shopCartItem.UserId != userID {
		return errors.New("禁止该操作！")
	}
	shopCartItem.GoodsCount = req.GoodsCount
	shopCartItem.UpdateTime = common.JSONTime{time.Now()}
	err = global.GVA_DB.Save(&shopCartItem).Error
	return
}

func (m *MallShopCartService) DeleteMallCartItem(userID int, id int) (err error) {
	var shopCartItem mall.MallShoppingCartItem
	err = global.GVA_DB.Where("cart_item_id = ? and is_deleted = 0", id).First(&shopCartItem).Error
	if err != nil {
		return
	}
	if userID != shopCartItem.UserId {
		return errors.New("禁止该操作！")
	}
	err = global.GVA_DB.Where("cart_item_id = ? and is_deleted = 0", id).UpdateColumns(&mall.MallShoppingCartItem{IsDeleted: 1}).Error
	return
}

func (m *MallShopCartService) GetCartItemsForSettle(userID int, cartItemIds []int) (err error, cartItemRes []mallRes.CartItemResponse) {
	var shopCartItems []mall.MallShoppingCartItem
	err = global.GVA_DB.Where("cart_item_id in (?) and user_id = ? and is_deleted = 0", cartItemIds, userID).Find(&shopCartItems).Error
	if err != nil {
		return
	}
	_, cartItemRes = getMallShoppingCartItemVOS(shopCartItems)
	//购物车算价
	var priceTotal float64
	for _, cartItem := range cartItemRes {
		priceTotal = priceTotal + float64(cartItem.GoodsCount)*cartItem.SellingPrice
	}
	return
}

// 购物车数据转换
func getMallShoppingCartItemVOS(cartItems []mall.MallShoppingCartItem) (err error, cartItemsRes []mallRes.CartItemResponse) {
	var goodsIds []int
	for _, cartItem := range cartItems {
		goodsIds = append(goodsIds, cartItem.GoodsId)
	}
	var newBeeMallGoods []manage.MallGoodsInfo
	err = global.GVA_DB.Where("goods_id in ?", goodsIds).Find(&newBeeMallGoods).Error
	if err != nil {
		return
	}

	newBeeMallGoodsMap := make(map[int]manage.MallGoodsInfo)
	for _, goodsInfo := range newBeeMallGoods {
		newBeeMallGoodsMap[goodsInfo.GoodsId] = goodsInfo
	}
	for _, cartItem := range cartItems {
		var cartItemRes mallRes.CartItemResponse
		copier.Copy(&cartItemRes, &cartItem)
		// 是否包含key
		if _, ok := newBeeMallGoodsMap[cartItemRes.GoodsId]; ok {
			newBeeMallGoodsTemp := newBeeMallGoodsMap[cartItemRes.GoodsId]
			cartItemRes.GoodsCoverImg = newBeeMallGoodsTemp.GoodsCoverImg
			goodsName := utils.SubStrLen(newBeeMallGoodsTemp.GoodsName, 28)
			cartItemRes.GoodsName = goodsName
			cartItemRes.SellingPrice = newBeeMallGoodsTemp.SellingPrice
			cartItemsRes = append(cartItemsRes, cartItemRes)
		}
	}
	return
}
