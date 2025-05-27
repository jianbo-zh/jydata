package fieldtype

type OrderRefundType int

const (
	OrderRefundType_Settlement OrderRefundType = 0 // 结算退款
	OrderRefundType_Manual     OrderRefundType = 1 // 运管退款
)
