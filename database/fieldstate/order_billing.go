package fieldstate

type OrderBillingState int

const (
	// 0-未开始 1-计时中 2-暂停中 4-已结束
	OrderBillingState_Init  OrderBillingState = 0 // 未开始
	OrderBillingState_Start OrderBillingState = 1 // 计时中
	OrderBillingState_Stop  OrderBillingState = 2 // 暂停中
	OrderBillingState_End   OrderBillingState = 3 // 已结束
)
