package fieldtype

type BillingStrategyMode int

const (
	BillingStrategyMode_Time    = 1 // 按时间
	BillingStrategyMode_Mileage = 2 // 按里程
	BillingStrategyMode_Stop    = 3 // 按站点
)
