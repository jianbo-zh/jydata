package database

import (
	"context"
	"time"

	"github.com/jianbo-zh/jydata/database/ent"
	"github.com/jianbo-zh/jydata/database/ent/orderbilling"
	"github.com/jianbo-zh/jydata/database/fieldstate"
)

func (db *Database) CreateOrderBilling(ctx context.Context, billing *ent.OrderBilling) (*ent.OrderBilling, error) {
	return db.MainDB().OrderBilling.Create().
		SetType(billing.Type).
		SetOrderID(billing.OrderID).
		SetState(int(fieldstate.OrderBillingState_Init)).
		SetStartTimePrice(billing.StartTimePrice).
		SetStartTimeUnit(billing.StartTimeUnit).
		SetNormalTimePrice(billing.NormalTimePrice).
		SetNormalTimeUnit(billing.NormalTimeUnit).
		SetStartStopPrice(billing.StartStopPrice).
		SetStartStopUnit(billing.StartStopUnit).
		SetNormalStopPrice(billing.NormalStopPrice).
		SetNormalStopUnit(billing.NormalStopUnit).
		SetCumulativeStop(billing.CumulativeStop).
		SetTicketCount(billing.TicketCount).
		SetCappedAmount(billing.CappedAmount).
		SetCouponID(billing.CouponID).
		SetCouponLimitAmount(billing.CouponLimitAmount).
		SetCouponDeductionAmount(billing.CouponDeductionAmount).
		SetState(billing.State).
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
		SetState(int(fieldstate.OrderBillingState_Start)).
		Save(ctx)
}

func (db *Database) UpdateOrderBillingEnd(ctx context.Context, id int, totalSec int) (*ent.OrderBilling, error) {
	return db.MainDB().OrderBilling.UpdateOneID(id).
		SetCumulativeSecond(float64(totalSec)).
		ClearStartTime().
		SetFinishTime(time.Now()).
		SetState(int(fieldstate.OrderBillingState_End)).
		Save(ctx)
}

func (db *Database) UpdateOrderBillingStop(ctx context.Context, id int, incrSec float64) (*ent.OrderBilling, error) {
	return db.MainDB().OrderBilling.UpdateOneID(id).
		AddCumulativeSecond(incrSec).
		ClearStartTime().
		SetState(int(fieldstate.OrderBillingState_Stop)).
		Save(ctx)
}

func (db *Database) IncrOrderBillingMeter(ctx context.Context, id int, distance float64) (*ent.OrderBilling, error) {
	return db.MainDB().OrderBilling.UpdateOneID(id).
		AddCumulativeMeter(distance).
		Save(ctx)
}
