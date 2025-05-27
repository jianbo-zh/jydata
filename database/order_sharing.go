package database

import (
	"context"

	"github.com/jianbo-zh/jydata/database/ent"
)

type OrderSharingState int

const (
	OrderSharingState_Pending   = 1 // 待处理
	OrderSharingState_Shared    = 2 // 已完成
	OrderSharingState_Canceled  = 3 // 已取消（实际分润金额为0时）
	OrderSharingState_Initiated = 4 // 已发起退款

)

func (db *Database) CreateOrderSharing(ctx context.Context, sharing *ent.OrderSharing) (*ent.OrderSharing, error) {
	return db.MainDB().OrderSharing.Create().
		SetOrderID(sharing.OrderID).
		SetSharingNo(sharing.SharingNo).
		SetWxSharingID("").
		SetSharingAmount(sharing.SharingAmount).
		SetReceivers(sharing.Receivers).
		SetState(sharing.State).
		SetRemark(sharing.Remark).
		Save(ctx)
}

func (db *Database) GetOrderSharing(ctx context.Context, id int) (*ent.OrderSharing, error) {
	return db.MainDB().OrderSharing.Get(ctx, id)
}
