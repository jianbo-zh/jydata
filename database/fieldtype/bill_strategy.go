package fieldtype

type BillingStrategyMode int

const (
	BillingStrategyMode_Time    BillingStrategyMode = 1 // 按时间
	BillingStrategyMode_Mileage BillingStrategyMode = 2 // 按里程
	BillingStrategyMode_Stop    BillingStrategyMode = 3 // 按站点
)
