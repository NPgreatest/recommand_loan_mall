package mall

import (
	"github.com/gin-gonic/gin"
	v1 "main.go/api/v1"
	"main.go/middleware"
)

type MallUserRouter struct {
}

func (m *MallUserRouter) InitMallUserRouter(Router *gin.RouterGroup) {
	mallUserRouter := Router.Group("v1").Use(middleware.UserJWTAuth())
	userRouter := Router.Group("v1")
	var mallUserApi = v1.ApiGroupApp.MallApiGroup.MallUserApi
	{
		mallUserRouter.PUT("/user/info", mallUserApi.UserInfoUpdate) //修改用户信息
		mallUserRouter.GET("/user/info", mallUserApi.GetUserInfo)    //获取用户信息
		mallUserRouter.POST("/user/finance", mallUserApi.SetUserFinance)
		mallUserRouter.GET("/user/finance", mallUserApi.GetUserFinance)
		mallUserRouter.POST("/user/try_loan", mallUserApi.GetUserLoan)
		mallUserRouter.POST("/user/loan", mallUserApi.DoLoan)
		mallUserRouter.POST("/user/payloan", mallUserApi.PayLoan)
	}
	{
		userRouter.POST("/user/register", mallUserApi.UserRegister) //用户注册
		userRouter.POST("/user/login", mallUserApi.UserLogin)       //登陆

	}

}
