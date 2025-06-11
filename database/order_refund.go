package database

import (
	"context"

	"github.com/jianbo-zh/jydata/database/ent"
	"github.com/jianbo-zh/jydata/database/ent/orderrefund"

	"entgo.io/ent/dialect/sql"
)

func (db *Database) CreateOrderRefund(ctx context.Context, refund *ent.OrderRefund) (*ent.OrderRefund, error) {
	return db.MainDB().OrderRefund.Create().
		SetType(refund.Type).
		SetInitiatorID(refund.InitiatorID).
		SetScenicAreaID(refund.ScenicAreaID).
		SetOrderID(refund.OrderID).
		SetOrderNo(refund.OrderNo).
		SetRefundNo(refund.RefundNo).
		SetWxRefundID("").
		SetRefundAmount(refund.RefundAmount).
		SetState(refund.State).
		SetRemark(refund.Remark).
		SetNillableOrderAppealID(refund.OrderAppealID).
		Save(ctx)
}

func (db *Database) GetOrderRefund(ctx context.Context, id int) (*ent.OrderRefund, error) {
	return db.MainDB().OrderRefund.Get(ctx, id)
}

func (db *Database) GetOrderRefundByOrderID(ctx context.Context, orderID int) (*ent.OrderRefund, error) {
	return db.MainDB().OrderRefund.Query().
		Where(orderrefund.OrderIDEQ(orderID)).
		Order(orderrefund.ByID(sql.OrderDesc())).
		First(ctx)
}

func (db *Database) GetOrderRefundByRefundNo(ctx context.Context, refundNo string) (*ent.OrderRefund, error) {
	return db.MainDB().OrderRefund.Query().Where(orderrefund.RefundNoEQ(refundNo)).First(ctx)
}
