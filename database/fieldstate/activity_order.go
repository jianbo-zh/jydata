package fieldstate

type ActivityOrderState int

const (
	// 订单状态(0-待支付、1-已支付)
	ActivityOrderState_Pending = 0 // 待付押金
	ActivityOrderState_Paid    = 1 // 待使用
)
