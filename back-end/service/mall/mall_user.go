package mall

import (
	"errors"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"main.go/global"
	"main.go/model/common"
	"main.go/model/mall"
	mallReq "main.go/model/mall/request"
	mallRes "main.go/model/mall/response"
	"main.go/utils"
	"strconv"
	"time"
)

type MallUserService struct {
}

// RegisterUser 注册用户
func (m *MallUserService) RegisterUser(req mallReq.RegisterUserParam) (err error) {
	if !errors.Is(global.GVA_DB.Where("login_name =?", req.LoginName).First(&mall.MallUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同用户名")
	}

	return global.GVA_DB.Create(&mall.MallUser{
		LoginName:     req.LoginName,
		PasswordMd5:   utils.MD5V([]byte(req.Password)),
		IntroduceSign: "2024毕业设计",
		NickName:      "请设置昵称",
		CreateTime:    common.JSONTime{Time: time.Now()},
		Avatar:        "51.jpg",
	}).Error

}

func (m *MallUserService) UpdateUserInfo(userID string, req mallReq.UpdateUserInfoParam) (err error) {
	var userInfo mall.MallUser
	err = global.GVA_DB.Where("user_id =?", userID).First(&userInfo).Error
	// 若密码为空字符，则表明用户不打算修改密码，使用原密码保存
	if !(req.PasswordMd5 == "") {
		userInfo.PasswordMd5 = utils.MD5V([]byte(req.PasswordMd5))
	}
	if req.Avatar != "" {
		address, err := utils.WriteImages(req.Avatar)
		if err != nil {
			return err
		}
		userInfo.Avatar = address
	}
	userInfo.NickName = req.NickName
	userInfo.IntroduceSign = req.IntroduceSign
	err = global.GVA_DB.Where("user_id =?", userID).UpdateColumns(&userInfo).Error
	return
}

func (m *MallUserService) GetUserDetail(userID string) (err error, userDetail mallRes.MallUserDetailResponse) {
	if err != nil {
		return errors.New("不存在的用户"), userDetail
	}
	var userInfo mall.MallUser
	err = global.GVA_DB.Where("user_id =?", userID).First(&userInfo).Error
	if err != nil {
		return errors.New("用户信息获取失败"), userDetail
	}
	err = copier.Copy(&userDetail, &userInfo)
	return
}

func (m *MallUserService) SetUserFinance(userID int, finance mallReq.UserSetFinance) (err error) {
	if err != nil {
		return errors.New("不存在的用户")
	}

	var existingFinance mall.MallUserFinance
	result := global.GVA_DB.Where("user_id = ?", userID).First(&existingFinance)
	if result.Error == nil {
		existingFinance.MonthlyIncome = finance.MonthlyIncome
		existingFinance.MonthlyExpenses = finance.MonthlyExpenses
		existingFinance.CreditScore = finance.CreditScore
		existingFinance.DebtStatus = finance.DebtStatus
		err = global.GVA_DB.Save(&existingFinance).Error
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		existingFinance = mall.MallUserFinance{
			UserId:          int64(userID),
			MonthlyIncome:   finance.MonthlyIncome,
			MonthlyExpenses: finance.MonthlyExpenses,
			CreditScore:     finance.CreditScore,
			DebtStatus:      finance.DebtStatus,
		}
		err = global.GVA_DB.Create(&existingFinance).Error
	} else {
		err = result.Error
	}
	return
}

func (m *MallUserService) GetUserFinance(userID int) (err error, finance mall.MallUserFinance) {
	err = global.GVA_DB.Where("user_id =?", userID).First(&finance).Error
	return
}

func (m *MallUserService) UserLogin(params mallReq.UserLoginParam) (err error, userToken mall.MallUserToken) {
	var user mall.MallUser
	err = global.GVA_DB.Where("login_name=? AND password_md5=?", params.LoginName, params.PasswordMd5).First(&user).Error
	if user != (mall.MallUser{}) && err == nil {
		struserToken, err := utils.CreateToken(strconv.Itoa(user.UserId), time.Hour*24)
		if err != nil {
			return err, userToken
		}
		userToken.UserId = user.UserId
		userToken.Token = struserToken
		userToken.UpdateTime = time.Now()
		userToken.ExpireTime = time.Now().Add(time.Hour * 24)
	}
	return err, userToken
}
