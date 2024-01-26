package mall

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/common/response"
	mallReq "main.go/model/mall/request"
	"main.go/utils"
	"strconv"
)

type MallUserApi struct {
}

func (m *MallUserApi) UserRegister(c *gin.Context) {
	var req mallReq.RegisterUserParam
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.MallUserRegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := mallUserService.RegisterUser(req); err != nil {
		global.GVA_LOG.Error("创建失败", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
	}
	response.OkWithMessage("创建成功", c)
}

func (m *MallUserApi) UserInfoUpdate(c *gin.Context) {
	var req mallReq.UpdateUserInfoParam
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.GVA_LOG.Error("参数错误", zap.Error(err))
		response.FailWithMessage("参数错误"+err.Error(), c)
		return
	}
	userID, _ := utils.VerifyToken(c.GetHeader("Authorization"))
	if err := mallUserService.UpdateUserInfo(userID, req); err != nil {
		global.GVA_LOG.Error("更新用户信息失败", zap.Error(err))
		response.FailWithMessage("更新用户信息失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (m *MallUserApi) GetUserInfo(c *gin.Context) {
	userID, _ := utils.VerifyToken(c.GetHeader("Authorization"))
	if err, userDetail := mallUserService.GetUserDetail(userID); err != nil {
		global.GVA_LOG.Error("未查询到记录", zap.Error(err))
		response.FailWithMessage("未查询到记录", c)
	} else {
		response.OkWithData(userDetail, c)
	}
}

func (m *MallUserApi) SetUserFinance(c *gin.Context) {
	userID, _ := utils.VerifyToken(c.GetHeader("Authorization"))
	var req mallReq.UserSetFinance
	_ = c.ShouldBindJSON(&req)
	iuserID, _ := strconv.Atoi(userID)
	if err := mallUserService.SetUserFinance(iuserID, req); err != nil {
		global.GVA_LOG.Error("设置预算失败", zap.Error(err))
		response.FailWithMessage("设置预算失败", c)
	} else {
		response.OkWithMessage("设置成功", c)
	}
}

func (m *MallUserApi) GetUserFinance(c *gin.Context) {
	userID, _ := utils.VerifyToken(c.GetHeader("Authorization"))
	iuserID, _ := strconv.Atoi(userID)
	if err, financeDetail := mallUserService.GetUserFinance(iuserID); err != nil {
		global.GVA_LOG.Error("未查询到记录", zap.Error(err))
		response.FailWithMessage("未查询到记录", c)
	} else {
		response.OkWithData(financeDetail, c)
	}
}

func (m *MallUserApi) UserLogin(c *gin.Context) {
	var req mallReq.UserLoginParam
	_ = c.ShouldBindJSON(&req)
	if err, token := mallUserService.UserLogin(req); err != nil {
		response.FailWithMessage("登陆失败", c)
	} else {
		response.OkWithData(token.Token, c)
	}
}
