package database

import (
	"context"
	"time"

	"github.com/jianbo-zh/jydata/database/ent"
	"github.com/jianbo-zh/jydata/database/fieldtype"
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
