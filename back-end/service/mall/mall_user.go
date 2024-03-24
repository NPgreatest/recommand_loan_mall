package mall

import (
	"errors"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"log"
	"main.go/global"
	"main.go/model/common"
	"main.go/model/mall"
	mallReq "main.go/model/mall/request"
	mallRes "main.go/model/mall/response"
	"main.go/utils"
	"os/exec"
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

func (m *MallUserService) SetUserFinance(userID int, finance *mallReq.UserSetFinance) error {
	var existingFinance mall.MallUserFinance
	result := global.GVA_DB.Where("user_id = ?", userID).First(&existingFinance)
	if result.Error == nil {
		updateMap := map[string]interface{}{
			"user_id":            int64(userID), // 通常不需要更新主键
			"gender":             finance.Gender,
			"dependents":         finance.Dependents,
			"married":            finance.Married,
			"education":          finance.Education,
			"self_employed":      finance.SelfEmployed,
			"applicant_income":   finance.ApplicantIncome,
			"coapplicant_income": finance.CoapplicantIncome,
			"city":               finance.City,
		}
		return global.GVA_DB.Model(&existingFinance).Where("user_id =?", userID).Updates(updateMap).Error
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		existingFinance = mall.MallUserFinance{
			UserID:            int64(userID),
			Gender:            finance.Gender,
			Dependents:        finance.Dependents,
			Married:           finance.Married,
			Education:         finance.Education,
			SelfEmployed:      finance.SelfEmployed,
			ApplicantIncome:   finance.ApplicantIncome,
			CoapplicantIncome: finance.CoapplicantIncome,
			City:              finance.City,
		}
		return global.GVA_DB.Create(&existingFinance).Error
	} else {
		return result.Error
	}
}

func (m *MallUserService) GetUserFinance(userID int) (err error, finance mall.MallUserFinance) {
	err = global.GVA_DB.Where("user_id =?", userID).First(&finance).Error
	return
}
func (m *MallUserService) GetUserLoan(userID int, req *mallReq.UserGetLoanReq) (err error, ok bool) {
	var financeInfo *mall.MallUserFinance
	err = global.GVA_DB.Where("user_id = ?", userID).First(&financeInfo).Error
	if err != nil {
		return err, false
	}
	args := preparePythonArguments(financeInfo, req)
	cmd := exec.Command("python", append([]string{"./python-script/main.py"}, args...)...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("执行Python脚本时出错: %v", err)
		return err, false
	}
	log.Printf("Python脚本输出: %s", output)
	res := string(output)
	if res[0] == '0' {
		return nil, false
	} else {
		return nil, true
	}
}

func (m *MallUserService) DoLoan(userID int, req *mallReq.UserGetLoanReq) (err error) {
	var financeInfo *mall.MallUserFinance
	err = global.GVA_DB.Where("user_id = ?", userID).First(&financeInfo).Error
	if err != nil {
		return
	}
	if financeInfo.Debt > 0 {
		return errors.New("目前贷款未还清")
	}
	financeInfo.Debt = req.Amount
	financeInfo.Amount += req.Amount
	financeInfo.Term = req.Term
	return global.GVA_DB.Where("user_id =?", userID).UpdateColumns(&financeInfo).Error
}

func (m *MallUserService) PayLoan(userID int, amount int) (err error) {
	var financeInfo *mall.MallUserFinance
	err = global.GVA_DB.Where("user_id = ?", userID).First(&financeInfo).Error
	if err != nil {
		return
	}
	if financeInfo.Debt > 0 {
		if financeInfo.Debt <= amount {
			financeInfo.Amount += amount - financeInfo.Debt
			financeInfo.Debt = 0
			financeInfo.Term = 0
		} else {
			financeInfo.Debt -= amount
		}
	} else {
		financeInfo.Amount += amount
		financeInfo.Term = 0
	}
	updateMap := map[string]interface{}{
		"user_id":            int64(userID), // 通常不需要更新主键
		"gender":             financeInfo.Gender,
		"dependents":         financeInfo.Dependents,
		"married":            financeInfo.Married,
		"education":          financeInfo.Education,
		"self_employed":      financeInfo.SelfEmployed,
		"applicant_income":   financeInfo.ApplicantIncome,
		"coapplicant_income": financeInfo.CoapplicantIncome,
		"city":               financeInfo.City,
		"amount":             financeInfo.Amount,
		"term":               financeInfo.Term,
		"debt":               financeInfo.Debt,
	}
	return global.GVA_DB.Model(&financeInfo).Where("user_id =?", userID).Updates(updateMap).Error
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

func preparePythonArguments(financeInfo *mall.MallUserFinance, req *mallReq.UserGetLoanReq) []string {
	// 将Golang的数据结构转换为Python脚本需要的字符串参数形式
	// 注意：这里的转换逻辑需要根据Python脚本的期望输入格式进行调整
	args := []string{
		utils.GenderToString(financeInfo.Gender),
		utils.MarriedToString(financeInfo.Married),
		strconv.Itoa(int(financeInfo.Dependents)),
		utils.EducationToString(financeInfo.Education),
		utils.SelfEmployedToString(financeInfo.SelfEmployed),
		strconv.FormatFloat(financeInfo.ApplicantIncome, 'f', -1, 64),
		strconv.FormatFloat(financeInfo.CoapplicantIncome, 'f', -1, 64),
		strconv.Itoa(req.Amount),
		strconv.Itoa(req.Term),
		"1", //TODO 根据用户购买记录预测
		utils.CityToString(financeInfo.City),
	}
	return args
}
