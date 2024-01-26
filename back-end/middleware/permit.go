package middleware

//func AuthPermit(permidID int) gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		userID, _ := utils.VerifyToken(ctx.GetHeader("Authorization"))
//		var c int64
//		db := global.GVA_DB.Model(&mall.Permission{})
//		var userInfo mall.User
//		err := global.GVA_DB.Where("login_name =?", userID).First(&userInfo).Error
//		if err != nil {
//			response.FailWithMessage("验证权限出错", ctx)
//			ctx.Abort()
//			return
//		}
//		err = db.Where("user_id = ? and permission_id = ?", userInfo.UserId, permidID).Count(&c).Error
//		if err != nil {
//			response.FailWithMessage("验证权限出错", ctx)
//			ctx.Abort()
//			return
//		}
//		if c == 0 {
//			response.FailWithDetailed(nil, "用户没有权限", ctx)
//			ctx.Abort()
//			return
//		}
//		ctx.Next()
//	}
//}
