package fieldstate

type OrderState int

const (
	// 订单状态(0-待付押金、1-待使用、2-进行中、3-待支付、4-待退费、5-已完成、6-已取消)
	OrderState_PendingDeposit = 0 // 待付押金
	OrderState_PendingUseCar  = 1 // 待使用
	OrderState_UsingCar       = 2 // 进行中
	OrderState_PendingPayment = 3 // 待支付
	OrderState_PendingRefund  = 4 // 待退费
	OrderState_Completed      = 5 // 已完成
	OrderState_Cancelled      = 6 // 已取消
)

const (
	DepositState_No      = 0 // 无押金
	DepositState_Pending = 1 // 待付押金
	DepositState_Paid    = 2 // 已付押金
)

const (
	EmergencyState_No     = 0 // 无紧急
	EmergencyState_Yes    = 1 // 紧急呼救
	EmergencyState_Cancel = 2 // 已取消紧急呼救
)
