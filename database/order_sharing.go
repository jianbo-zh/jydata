package database

import (
	"context"

	"github.com/jianbo-zh/jydata/database/ent"
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
