package fieldstate

type OrderSharingState int

const (
	OrderSharingState_Pending   = 1 // 待处理
	OrderSharingState_Shared    = 2 // 已完成
	OrderSharingState_Canceled  = 3 // 已取消（实际分润金额为0时）
	OrderSharingState_Initiated = 4 // 已发起退款

)
