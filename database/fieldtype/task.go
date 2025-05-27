package fieldtype

// 任务类型
type TaskType string

const (
	TaskType_OrderTimeoutAbortCheck   = "order_timeout_abort_check"   // 超时中止（支付押金超时）
	TaskType_OrderExpiredCloseCheck   = "order_expired_close_check"   // 订单最长时间检查
	TaskType_OrderRefund              = "order_refund"                // 订单退款
	TaskType_OrderProfitSharingResult = "order_profit_sharing_result" // 分润结果查询
)
