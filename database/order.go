package database

import (
	"context"
	"time"

	"github.com/jianbo-zh/jydata/database/ent"
	"github.com/jianbo-zh/jydata/database/ent/order"
)

type OrderState int

const (
	// 订单状态(0-待付押金、1-待使用、2-进行中、3-待支付、4-待退费、5-已完成、6-已取消)
	OrderState_PendingDeposit = 0 // 待付押金
	OrderState_PendingUseCar  = 1 // 待使用
	OrderState_UsingCar       = 2 // 进行中
	OrderState_PendingPayment = 3 // 待支付
	OrderState_PendingRefund  = 4 // 待退费
	OrderState_Completed      = 5 // 已完成
	OrderState_Cancelled      = 6 // 已取消
)

const (
	DepositState_No      = 0 // 无押金
	DepositState_Pending = 1 // 待付押金
	DepositState_Paid    = 2 // 已付押金
)

const (
	EmergencyState_No     = 0 // 无紧急
	EmergencyState_Yes    = 1 // 紧急呼救
	EmergencyState_Cancel = 2 // 已取消紧急呼救
)

func (db *Database) CreateOrder(ctx context.Context, req *ent.Order) (*ent.Order, error) {
	return db.MainDB().Order.Create().
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
		SetEmergencyState(req.EmergencyState).
		SetIsProfitSharing(req.IsProfitSharing).
		SetIsTestOrder(req.IsTestOrder).
		Save(ctx)
}

func (db *Database) GetUserOngoingOrder(ctx context.Context, userID int) (*ent.Order, error) {
	return db.MainDB().Order.Query().
		Where(order.UserIDEQ(userID)).
		Where(order.OrderStateNotIn(OrderState_Completed, OrderState_Cancelled)).
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
		Where(order.OrderStateEQ(OrderState_PendingDeposit)).
		SetWxTxID(txID).
		SetOrderState(OrderState_UsingCar).
		SetDepositState(DepositState_Paid).
		SetDepositTime(depositTime).
		Save(ctx)
}
