package database

import (
	"context"
	"time"

	"github.com/jianbo-zh/jydata/database/ent"
	"github.com/jianbo-zh/jydata/database/ent/activityorder"
)

type ActivityOrderState int

const (
	// 订单状态(0-待支付、1-已支付)
	ActivityOrderState_Pending = 0 // 待付押金
	ActivityOrderState_Paid    = 1 // 待使用
)

func (db *Database) CreateActivityOrder(ctx context.Context, req *ent.ActivityOrder) (*ent.ActivityOrder, error) {
	return db.MainDB().ActivityOrder.Create().
		SetOrderNo(req.OrderNo).
		SetScenicAreaID(req.ScenicAreaID).
		SetScenicAreaName(req.ScenicAreaName).
		SetUserID(req.UserID).
		SetOpenID(req.OpenID).
		SetOrderAmount(req.OrderAmount).
		SetOrderState(req.OrderState).
		SetRemark(req.Remark).
		Save(ctx)
}

func (db *Database) GetActivityOrder(ctx context.Context, id int) (*ent.ActivityOrder, error) {
	return db.MainDB().ActivityOrder.Get(ctx, id)
}

func (db *Database) GetActivityOrderByOrderNo(ctx context.Context, orderNo string) (*ent.ActivityOrder, error) {
	return db.MainDB().ActivityOrder.Query().Where(activityorder.OrderNoEQ(orderNo)).Only(ctx)
}

func (db *Database) UpdateActivityOrderPaid(ctx context.Context, orderID int, txID string, paidTime time.Time) (*ent.ActivityOrder, error) {
	return db.MainDB().ActivityOrder.UpdateOneID(orderID).
		Where(activityorder.OrderStateEQ(ActivityOrderState_Pending)).
		SetWxTxID(txID).
		SetOrderState(ActivityOrderState_Paid).
		SetPaidTime(paidTime).
		Save(ctx)
}
