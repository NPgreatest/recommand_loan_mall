package mall

import (
	"github.com/gin-gonic/gin"
	v1 "main.go/api/v1"
)

type MallQueryRouter struct {
}

func (m *MallQueryRouter) InitMallQueryRouter(Router *gin.RouterGroup) {
	mallGoodsRouter := Router.Group("v1")
	var mallQueryApi = v1.ApiGroupApp.MallApiGroup.MallQueryApi
	{
		mallGoodsRouter.POST("query", mallQueryApi.TextToItem) // 获取分类数据
	}
}
