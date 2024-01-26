package mall

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/common/response"
	mallReq "main.go/model/mall/request"
	"main.go/utils"
)

type MallShopCartApi struct {
}

func (m *MallShopCartApi) CartItemList(c *gin.Context) {
	userID, _ := utils.VerifyToken(c.GetHeader("Authorization"))
	iuserID, _ := strconv.Atoi(userID)
	if err, shopCartItem := mallShopCartService.GetMyShoppingCartItems(iuserID); err != nil {
		global.GVA_LOG.Error("获取购物车失败", zap.Error(err))
		response.FailWithMessage("获取购物车失败:"+err.Error(), c)
	} else {
		response.OkWithData(shopCartItem, c)
	}
}

func (m *MallShopCartApi) SaveMallShoppingCartItem(c *gin.Context) {
	userID, _ := utils.VerifyToken(c.GetHeader("Authorization"))
	iuserID, _ := strconv.Atoi(userID)
	var req mallReq.SaveCartItemParam
	_ = c.ShouldBindJSON(&req)
	if err := mallShopCartService.SaveMallCartItem(iuserID, req); err != nil {
		global.GVA_LOG.Error("添加购物车失败", zap.Error(err))
		response.FailWithMessage("添加购物车失败:"+err.Error(), c)
	}
	response.OkWithMessage("添加购物车成功", c)
}

func (m *MallShopCartApi) UpdateMallShoppingCartItem(c *gin.Context) {
	userID, _ := utils.VerifyToken(c.GetHeader("Authorization"))
	iuserID, _ := strconv.Atoi(userID)
	var req mallReq.UpdateCartItemParam
	_ = c.ShouldBindJSON(&req)
	if err := mallShopCartService.UpdateMallCartItem(iuserID, req); err != nil {
		global.GVA_LOG.Error("修改购物车失败", zap.Error(err))
		response.FailWithMessage("修改购物车失败:"+err.Error(), c)
	}
	response.OkWithMessage("修改购物车成功", c)
}

func (m *MallShopCartApi) DelMallShoppingCartItem(c *gin.Context) {
	userID, _ := utils.VerifyToken(c.GetHeader("Authorization"))
	iuserID, _ := strconv.Atoi(userID)
	id, _ := strconv.Atoi(c.Param("newBeeMallShoppingCartItemId"))
	if err := mallShopCartService.DeleteMallCartItem(iuserID, id); err != nil {
		global.GVA_LOG.Error("修改购物车失败", zap.Error(err))
		response.FailWithMessage("修改购物车失败:"+err.Error(), c)
	} else {
		response.OkWithMessage("修改购物车成功", c)
	}
}

func (m *MallShopCartApi) ToSettle(c *gin.Context) {
	cartItemIdsStr := c.Query("cartItemIds")
	userID, _ := utils.VerifyToken(c.GetHeader("Authorization"))
	iuserID, _ := strconv.Atoi(userID)
	cartItemIds := utils.StrToInt(cartItemIdsStr)
	if err, cartItemRes := mallShopCartService.GetCartItemsForSettle(iuserID, cartItemIds); err != nil {
		global.GVA_LOG.Error("获取购物明细异常：", zap.Error(err))
		response.FailWithMessage("获取购物明细异常:"+err.Error(), c)
	} else {
		response.OkWithData(cartItemRes, c)
	}

}
