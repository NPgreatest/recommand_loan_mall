package mall

type MallUserFinance struct {
	UserID            int64   `json:"user_id"`
	Gender            bool    `json:"gender"`
	Dependents        int16   `json:"dependents"`
	Married           bool    `json:"married"`
	Education         bool    `json:"education"`
	SelfEmployed      bool    `json:"self_employed"`
	ApplicantIncome   float64 `json:"applicant_income"`
	CoapplicantIncome float64 `json:"coapplicant_income"`
	City              bool    `json:"city"`
}

// TableName 设置MallUserFinance的表名为mall_user_finance
func (MallUserFinance) TableName() string {
	return "mall_user_finance"
}
