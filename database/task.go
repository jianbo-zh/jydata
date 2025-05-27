package database

import (
	"context"
	"time"

	"github.com/jianbo-zh/jydata/database/ent"
)

// 任务类型
type TaskType string

// 任务状态
type TaskState int

const (
	TaskType_OrderTimeoutAbortCheck   = "order_timeout_abort_check"   // 超时中止（支付押金超时）
	TaskType_OrderExpiredCloseCheck   = "order_expired_close_check"   // 订单最长时间检查
	TaskType_OrderRefund              = "order_refund"                // 订单退款
	TaskType_OrderProfitSharingResult = "order_profit_sharing_result" // 分润结果查询
)

const (
	TaskTimeout_OrderAbortCheck          = 1 * time.Minute  // 1分钟
	TaskTimeout_OrderCancelCheck         = 5 * time.Minute  // 5分钟
	TaskTimeout_OrderExpirationCheck     = 24 * time.Hour   // 24小时
	TaskTimeout_OrderRefund              = 1 * time.Second  // 1秒
	TaskTimeout_OrderProfitSharingResult = 10 * time.Minute // 10分钟
)

// 状态(0-待处理、1-待重试、2-执行成功、3-执行失败)
const (
	TaskState_Pending = 0 // 待处理
	TaskState_Retry   = 1 // 待重试
	TaskState_Success = 2 // 执行成功
	TaskState_Failed  = 3 // 执行失败
)

// AddOrderRefundTask 添加退款任务
func (m *Database) AddOrderRefundTask(ctx context.Context, taskTime time.Time, refundID int) (*ent.Task, error) {
	return m.MainDB().Task.Create().
		SetType(TaskType_OrderRefund).
		SetRelID(refundID).
		SetRelData(nil).
		SetStartTime(taskTime).
		SetNextTime(taskTime).
		Save(ctx)
}

// AddOrderTimeoutAbortCheckTask 超时未付押金检查
func (m *Database) AddOrderTimeoutAbortCheckTask(ctx context.Context, taskTime time.Time, orderID int) (*ent.Task, error) {
	return m.MainDB().Task.Create().
		SetType(TaskType_OrderTimeoutAbortCheck).
		SetRelID(orderID).
		SetRelData(nil).
		SetStartTime(taskTime).
		SetNextTime(taskTime).
		Save(ctx)
}

// AddOrderExpiredCloseCheckTask 订单过期检查
func (m *Database) AddOrderExpiredCloseCheckTask(ctx context.Context, taskTime time.Time, orderID int) (*ent.Task, error) {
	return m.MainDB().Task.Create().
		SetType(TaskType_OrderExpiredCloseCheck).
		SetRelID(orderID).
		SetRelData(nil).
		SetStartTime(taskTime).
		SetNextTime(taskTime).
		Save(ctx)
}

// AddOrderProfitSharingCheckTask 订单分润结果任务
func (m *Database) AddOrderProfitSharingResultTask(ctx context.Context, taskTime time.Time, sharingID int) (*ent.Task, error) {
	return m.MainDB().Task.Create().
		SetType(TaskType_OrderProfitSharingResult).
		SetRelID(sharingID).
		SetStartTime(taskTime).
		SetNextTime(taskTime).
		Save(ctx)
}
