package database

import (
	"context"
	"time"

	"github.com/jianbo-zh/jydata/database/ent"
	"github.com/jianbo-zh/jydata/database/ent/car"
	"github.com/jianbo-zh/jydata/database/mixin"
)

type CarState int
type CarUseState int
type DrivingState int

const (
	CarEmergencyState_No     int = 0 // 无紧急
	CarEmergencyState_Yes    int = 1 // 紧急呼救
	CarEmergencyState_Cancel int = 2 // 已取消紧急呼救
)

// 车辆状态：1-未激活、2-未运营、3-运营中、4-维修中、5-中止运营
const (
	CarState_Inactivated   int = 1
	CarState_NoOperation   int = 2
	CarState_OnOperating   int = 3
	CarState_Maintenance   int = 4
	CarState_StopOperation int = 5
)

const (
	// 使用状态：1-空闲中、2-预约中、3-租用中
	CarUseState_Idle        = 1 // 空闲中
	CarUseState_Booking     = 2 // 预约中
	CarUseState_Using       = 3 // 租用中
	CarUseState_Dispatching = 6 // 调度中
)

const (
	CarIsOnline  = 1  // 在线
	CarIsOffline = -1 // 离线
)

const (
	// 2024.06.04 之后的
	DrivingState_Operation        DrivingState = 0b0000001 // 运营 - 1
	DrivingState_Maintain         DrivingState = 0b0000010 // 运维 - 2
	DrivingState_Lock             DrivingState = 0b0000100 // 锁车 - 4
	DrivingState_Manual           DrivingState = 0b0001000 // 手动 - 8
	DrivingState_Auto             DrivingState = 0b0010000 // 自动 - 16
	DrivingState_Remote           DrivingState = 0b0100000 // 远程 - 32
	DrivingState_Fault            DrivingState = 0b1000000 // 故障 - 64
	DrivingState_Bootup           DrivingState = 0         // 启动 - 0
	DrivingState_Operation_Lock   DrivingState = 0b0000101 // 运营锁车 - 5
	DrivingState_Operation_Manual DrivingState = 0b0001001 // 运营手动 - 9
	DrivingState_Operation_Auto   DrivingState = 0b0010001 // 运营自动 - 17
	DrivingState_Operation_Fault  DrivingState = 0b1000001 // 运营故障 - 65
	DrivingState_Maintain_Lock    DrivingState = 0b0000110 // 维护锁车 - 6
	DrivingState_Maintain_Manual  DrivingState = 0b0001010 // 维护手动 - 10
	DrivingState_Maintain_Auto    DrivingState = 0b0010010 // 维护自动 - 18
	DrivingState_Maintain_Remote  DrivingState = 0b0100010 // 维护远程 - 34
	DrivingState_Maintain_Fault   DrivingState = 0b1000010 // 维护故障 - 66
)

// 正常运营，无故障
func IsDrivingStateOperationNormal(drivingState int) bool {
	return int(DrivingState_Operation)&drivingState == int(DrivingState_Operation) &&
		int(DrivingState_Fault)&drivingState != int(DrivingState_Fault)
}

func IsDrivingStateOperationManual(drivingState int) bool {
	return int(DrivingState_Operation_Manual)&drivingState == int(DrivingState_Operation_Manual)
}

func IsDrivingStateOperationAuto(drivingState int) bool {
	return int(DrivingState_Operation_Auto)&drivingState == int(DrivingState_Operation_Auto)
}

// 正常运维，无故障
func IsDrivingStateMaintainNormal(drivingState int) bool {
	return int(DrivingState_Maintain)&drivingState == int(DrivingState_Maintain) &&
		int(DrivingState_Fault)&drivingState != int(DrivingState_Fault)
}

func IsDrivingStateOperation(drivingState int) bool {
	return int(DrivingState_Operation)&drivingState == int(DrivingState_Operation)
}

func IsDrivingStateMaintain(drivingState int) bool {
	return int(DrivingState_Maintain)&drivingState == int(DrivingState_Maintain)
}

func IsDrivingStateFault(drivingState int) bool {
	return int(DrivingState_Fault)&drivingState == int(DrivingState_Fault)
}

func IsDrivingStateManual(drivingState int) bool {
	return int(DrivingState_Manual)&drivingState == int(DrivingState_Manual)
}

func IsDrivingStateAuto(drivingState int) bool {
	return int(DrivingState_Auto)&drivingState == int(DrivingState_Auto)
}

func IsDrivingStateRemote(drivingState int) bool {
	return int(DrivingState_Remote)&drivingState == int(DrivingState_Remote)
}

func IsDrivingStateLock(drivingState int) bool {
	return int(DrivingState_Lock)&drivingState == int(DrivingState_Lock)
}

func IsDrivingStateBootup(drivingState int) bool {
	return drivingState == int(DrivingState_Bootup)
}

func (db *Database) GetCar(ctx context.Context, id int) (*ent.Car, error) {
	return db.MainDB().Car.Get(mixin.SkipSoftDelete(ctx), id)
}

func (db *Database) GetCarByDeviceID(ctx context.Context, deviceID string) (*ent.Car, error) {
	return db.MainDB().Car.Query().Where(car.DeviceIDEQ(deviceID)).First(ctx)
}

func (db *Database) GetCarsByIDs(ctx context.Context, ids []int) ([]*ent.Car, error) {
	return db.MainDB().Car.Query().Where(car.IDIn(ids...)).All(ctx)
}

func (db *Database) UpdateCarBooking(ctx context.Context, carID int, orderID int) (*ent.Car, error) {
	return db.MainDB().Car.UpdateOneID(carID).
		Where(car.StateEQ(CarState_OnOperating)).
		Where(car.UseStateEQ(CarUseState_Idle)).
		SetUseState(CarUseState_Booking).
		SetUseOrderID(orderID).
		Save(ctx)
}

func (db *Database) UpdateCarUsing(ctx context.Context, carID int, orderID int) (*ent.Car, error) {
	return db.MainDB().Car.UpdateOneID(carID).
		Where(car.Or(
			car.UseStateEQ(CarUseState_Idle),
			car.UseStateEQ(CarUseState_Booking))).
		SetUseState(CarUseState_Using).
		SetUseOrderID(orderID).
		Save(ctx)
}

func (db *Database) UpdateCarAliveTime(ctx context.Context, deviceID string, aliveTime time.Time) (int, error) {
	return db.MainDB().Car.Update().Where(car.DeviceIDEQ(deviceID)).
		SetAliveTime(time.Now()).
		Save(ctx)
}

func (db *Database) UpdateCarDispatching(ctx context.Context, carID int, useState int, dispatchTaskID int) (*ent.Car, error) {
	return db.MainDB().Car.UpdateOneID(carID).
		Where(car.UseStateEQ(useState)).
		Where(car.DispatchTaskIDEQ(0)).
		SetUseState(CarUseState_Dispatching).
		SetDispatchTaskID(dispatchTaskID).
		Save(ctx)
}

func (db *Database) UpdateCarDispatchEnd(ctx context.Context, carID int, useState int) (*ent.Car, error) {
	return db.MainDB().Car.UpdateOneID(carID).
		Where(car.UseStateEQ(CarUseState_Dispatching)).
		Where(car.DispatchTaskIDGT(0)).
		SetUseState(useState).
		SetDispatchTaskID(0).
		Save(ctx)
}
