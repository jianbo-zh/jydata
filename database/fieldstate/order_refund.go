package fieldstate

type OrderRefundState int

const (
	OrderRefundState_Pending  = 1 // 待退款
	OrderRefundState_Refunded = 2 // 已退款
)
