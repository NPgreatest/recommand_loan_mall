package mall

// MallUserFinance 用户财务信息结构体
type MallUserFinance struct {
	UserId          int64   `json:"userId" form:"userId" gorm:"primarykey"`
	MonthlyIncome   float64 `json:"monthlyIncome" form:"monthlyIncome" gorm:"column:monthly_income;comment:每月收入;type:decimal(10,2);"`
	MonthlyExpenses float64 `json:"monthlyExpenses" form:"monthlyExpenses" gorm:"column:monthly_expenses;comment:每月支出;type:decimal(10,2);"`
	CreditScore     int     `json:"creditScore" form:"creditScore" gorm:"column:credit_score;comment:信用评分;type:int;"`
	DebtStatus      float64 `json:"debtStatus" form:"debtStatus" gorm:"column:debt_status;comment:债务状况;type:decimal(10,2);"`
}

// TableName 设置MallUserFinance的表名为mall_user_finance
func (MallUserFinance) TableName() string {
	return "mall_user_finance"
}
