package database

import (
	"context"
	"time"

	"github.com/jianbo-zh/jydata/database/ent"
	"github.com/jianbo-zh/jydata/database/fieldtype"
)

const (
	TaskTimeout_OrderAbortCheck          = 1 * time.Minute  // 1分钟
	TaskTimeout_OrderCancelCheck         = 5 * time.Minute  // 5分钟
	TaskTimeout_OrderExpirationCheck     = 24 * time.Hour   // 24小时
	TaskTimeout_OrderRefund              = 1 * time.Second  // 1秒
	TaskTimeout_OrderProfitSharingResult = 10 * time.Minute // 10分钟
)

// AddOrderRefundTask 添加退款任务
func (m *Database) AddOrderRefundTask(ctx context.Context, taskTime time.Time, refundID int) (*ent.Task, error) {
	return m.MainDB().Task.Create().
		SetType(fieldtype.TaskType_OrderRefund).
		SetRelID(refundID).
		SetRelData(nil).
		SetStartTime(taskTime).
		SetNextTime(taskTime).
		Save(ctx)
}

// AddOrderTimeoutAbortCheckTask 超时未付押金检查
func (m *Database) AddOrderTimeoutAbortCheckTask(ctx context.Context, taskTime time.Time, orderID int) (*ent.Task, error) {
	return m.MainDB().Task.Create().
		SetType(fieldtype.TaskType_OrderTimeoutAbortCheck).
		SetRelID(orderID).
		SetRelData(nil).
		SetStartTime(taskTime).
		SetNextTime(taskTime).
		Save(ctx)
}

// AddOrderExpiredCloseCheckTask 订单过期检查
func (m *Database) AddOrderExpiredCloseCheckTask(ctx context.Context, taskTime time.Time, orderID int) (*ent.Task, error) {
	return m.MainDB().Task.Create().
		SetType(fieldtype.TaskType_OrderExpiredCloseCheck).
		SetRelID(orderID).
		SetRelData(nil).
		SetStartTime(taskTime).
		SetNextTime(taskTime).
		Save(ctx)
}

// AddOrderProfitSharingCheckTask 订单分润结果任务
func (m *Database) AddOrderProfitSharingResultTask(ctx context.Context, taskTime time.Time, sharingID int) (*ent.Task, error) {
	return m.MainDB().Task.Create().
		SetType(fieldtype.TaskType_OrderProfitSharingResult).
		SetRelID(sharingID).
		SetStartTime(taskTime).
		SetNextTime(taskTime).
		Save(ctx)
}
