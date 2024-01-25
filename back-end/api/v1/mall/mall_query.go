package mall

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/common/response"
	mallReq "main.go/model/mall/request"
	mallRes "main.go/model/mall/response"
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

func (m *MallQueryApi) AdvancedRecommend(c *gin.Context) {
	var queryParam mallReq.RecommendQueryParam
	_ = c.ShouldBindJSON(&queryParam)
	var res []mallRes.RecommendResponse
	var err error
	if err, res = mallQueryService.FineTuneGetList(queryParam.QueryString); err != nil {
		global.GVA_LOG.Error("从Fine-Tune模型获取推荐失败", zap.Error(err))
		response.FailWithMessage("从Fine-Tune模型获取推荐失败:"+err.Error(), c)
		return
	}
	response.OkWithData(res, c)
}
