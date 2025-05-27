package database

import (
	"context"
	"time"

	"github.com/jianbo-zh/jydata/database/ent"
	"github.com/jianbo-zh/jydata/database/ent/orderbilling"
)

type OrderBillingState int

const (
	// 0-未开始 1-计时中 2-暂停中 4-已结束
	OrderBillingState_Init  OrderBillingState = 0 // 未开始
	OrderBillingState_Start OrderBillingState = 1 // 计时中
	OrderBillingState_Stop  OrderBillingState = 2 // 暂停中
	OrderBillingState_End   OrderBillingState = 3 // 已结束
)

func (db *Database) CreateOrderBilling(ctx context.Context, billing *ent.OrderBilling) (*ent.OrderBilling, error) {
	return db.MainDB().OrderBilling.Create().
		SetOrderID(billing.OrderID).
		SetState(int(OrderBillingState_Init)).
		SetStartTimePrice(billing.StartTimePrice).
		SetStartTimeUnit(billing.StartTimeUnit).
		SetNormalTimePrice(billing.NormalTimePrice).
		SetNormalTimeUnit(billing.NormalTimeUnit).
		SetCappedAmount(billing.CappedAmount).
		SetNillableStartTime(nil).
		SetNillableFinishTime(nil).
		Save(ctx)
}

func (db *Database) GetOrderBillingByOrderID(ctx context.Context, orderID int) (*ent.OrderBilling, error) {
	return db.MainDB().OrderBilling.Query().
		Where(orderbilling.OrderIDEQ(orderID)).
		First(ctx)
}

func (db *Database) UpdateOrderBillingStart(ctx context.Context, id int) (*ent.OrderBilling, error) {
	return db.MainDB().OrderBilling.UpdateOneID(id).
		SetStartTime(time.Now()).
		SetState(int(OrderBillingState_Start)).
		Save(ctx)
}

func (db *Database) UpdateOrderBillingEnd(ctx context.Context, id int, totalSec int) (*ent.OrderBilling, error) {
	return db.MainDB().OrderBilling.UpdateOneID(id).
		SetCumulativeSecond(float64(totalSec)).
		ClearStartTime().
		SetFinishTime(time.Now()).
		SetState(int(OrderBillingState_End)).
		Save(ctx)
}

func (db *Database) UpdateOrderBillingStop(ctx context.Context, id int, incrSec float64) (*ent.OrderBilling, error) {
	return db.MainDB().OrderBilling.UpdateOneID(id).
		AddCumulativeSecond(incrSec).
		ClearStartTime().
		SetState(int(OrderBillingState_Stop)).
		Save(ctx)
}

func (db *Database) IncrOrderBillingMeter(ctx context.Context, id int, distance float64) (*ent.OrderBilling, error) {
	return db.MainDB().OrderBilling.UpdateOneID(id).
		AddCumulativeMeter(distance).
		Save(ctx)
}
