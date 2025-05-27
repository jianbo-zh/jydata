package database

import (
	"context"

	"github.com/jianbo-zh/jydata/database/ent"
	"github.com/jianbo-zh/jydata/database/ent/orderrefund"

	"entgo.io/ent/dialect/sql"
)

type OrderRefundType int

const (
	OrderRefundType_Settlement OrderRefundType = 0 // 结算退款
	OrderRefundType_Manual     OrderRefundType = 1 // 运管退款
)

type OrderRefundState int

const (
	OrderRefundState_Pending  = 1 // 待退款
	OrderRefundState_Refunded = 2 // 已退款
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
