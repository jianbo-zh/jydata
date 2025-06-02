package database

import (
	"context"
	"time"

	"github.com/jianbo-zh/jydata/database/ent"
	"github.com/jianbo-zh/jydata/database/ent/order"
	"github.com/jianbo-zh/jydata/database/fieldstate"
)

func (db *Database) CreateOrder(ctx context.Context, req *ent.Order) (*ent.Order, error) {
	return db.MainDB().Order.Create().
		SetType(req.Type).
		SetOrderNo(req.OrderNo).
		SetMchID(req.MchID).
		SetScenicAreaID(req.ScenicAreaID).
		SetScenicAreaName(req.ScenicAreaName).
		SetUserID(req.UserID).
		SetOpenID(req.OpenID).
		SetNickname(req.Nickname).
		SetPhone(req.Phone).
		SetCarID(req.CarID).
		SetDeviceID(req.DeviceID).
		SetCarName(req.CarName).
		SetCarLicensePlate(req.CarLicensePlate).
		SetModelID(req.ModelID).
		SetModelName(req.ModelName).
		SetDepositAmount(req.DepositAmount).
		SetOrderAmount(req.OrderAmount).
		SetOrderState(req.OrderState).
		SetDepositState(req.DepositState).
		SetCouponID(req.CouponID).
		SetCouponName(req.CouponName).
		SetCouponAmount(req.CouponAmount).
		SetEmergencyState(req.EmergencyState).
		SetIsProfitSharing(req.IsProfitSharing).
		SetIsTestOrder(req.IsTestOrder).
		Save(ctx)
}

func (db *Database) GetUserOngoingOrder(ctx context.Context, userID int) (*ent.Order, error) {
	return db.MainDB().Order.Query().
		Where(order.UserIDEQ(userID)).
		Where(order.OrderStateNotIn(fieldstate.OrderState_Completed, fieldstate.OrderState_Cancelled)).
		First(ctx)
}

func (db *Database) GetOrder(ctx context.Context, id int) (*ent.Order, error) {
	return db.MainDB().Order.Get(ctx, id)
}

func (db *Database) GetOrderByOrderNo(ctx context.Context, orderNo string) (*ent.Order, error) {
	return db.MainDB().Order.Query().Where(order.OrderNoEQ(orderNo)).Only(ctx)
}

func (db *Database) UpdateOrderDepositPaid(ctx context.Context, orderID int, txID string, depositTime time.Time) (*ent.Order, error) {
	return db.MainDB().Order.UpdateOneID(orderID).
		Where(order.OrderStateEQ(fieldstate.OrderState_PendingDeposit)).
		SetWxTxID(txID).
		SetOrderState(fieldstate.OrderState_UsingCar).
		SetDepositState(fieldstate.DepositState_Paid).
		SetDepositTime(depositTime).
		Save(ctx)
}
