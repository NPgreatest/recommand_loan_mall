package mall

type MallUserFinance struct {
	UserID            int64   `json:"user_id" gorm:"column:user_id;"`
	Gender            bool    `json:"gender" gorm:"column:gender;"`
	Dependents        int16   `json:"dependents" gorm:"column:dependents;"`
	Married           bool    `json:"married" gorm:"column:married;"`
	Education         bool    `json:"education" gorm:"column:education;"`
	SelfEmployed      bool    `json:"self_employed" gorm:"column:self_employed;"`
	ApplicantIncome   float64 `json:"applicant_income" gorm:"column:applicant_income;"`
	CoapplicantIncome float64 `json:"coapplicant_income" gorm:"column:coapplicant_income;"`
	City              bool    `json:"city" gorm:"column:city;"`
	Amount            int     `json:"amount" gorm:"column:amount;"`
	Term              int     `json:"term" gorm:"column:term;"`
	Debt              int     `json:"debt" gorm:"column:debt;"`
}

// TableName 设置MallUserFinance的表名为mall_user_finance
func (MallUserFinance) TableName() string {
	return "mall_user_finance"
}
