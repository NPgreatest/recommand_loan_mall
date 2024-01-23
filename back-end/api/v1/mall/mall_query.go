package mall

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/common/response"
	mallReq "main.go/model/mall/request"
)

type MallQueryApi struct {
}

func (m *MallQueryApi) TextToItem(c *gin.Context) {
	var queryParam mallReq.QueryParam
	_ = c.ShouldBindJSON(&queryParam)
	if err, recommendItems := mallQueryService.TextToItem(queryParam.QueryString); err != nil {
		global.GVA_LOG.Error("获取推荐失败", zap.Error(err))
		response.FailWithMessage("获取推荐失败:"+err.Error(), c)
	} else {
		response.OkWithData(recommendItems, c)
	}
}
