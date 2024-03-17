package request

// 用户注册
type RegisterUserParam struct {
	LoginName string `json:"loginName"`
	Password  string `json:"password"`
	Avatar    string `json:"avatar"`
}

// 更新用户信息
type UpdateUserInfoParam struct {
	NickName      string `json:"nickName"`
	PasswordMd5   string `json:"passwordMd5"`
	IntroduceSign string `json:"introduceSign"`
	Avatar        string `json:"avatar"`
}

type UserLoginParam struct {
	LoginName   string `json:"loginName"`
	PasswordMd5 string `json:"passwordMd5"`
}

type UserSetFinance struct {
	Gender            bool    `json:"gender"`
	Dependents        int16   `json:"dependents"`
	Married           bool    `json:"married"`
	Education         bool    `json:"education"`
	SelfEmployed      bool    `json:"self_employed"`
	ApplicantIncome   float64 `json:"applicant_income"`
	CoapplicantIncome float64 `json:"coapplicant_income"`
	City              bool    `json:"city"`
}
