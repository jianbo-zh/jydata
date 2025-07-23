package database

import (
	"context"
	"time"

	"github.com/jianbo-zh/jydata/database/ent"
	"github.com/jianbo-zh/jydata/database/ent/car"
	"github.com/jianbo-zh/jydata/database/fieldstate"
	"github.com/jianbo-zh/jydata/database/mixin"
)

// 正常运营，无故障
func IsDrivingStateOperationNormal(drivingState int) bool {
	return int(fieldstate.DrivingState_Operation)&drivingState == int(fieldstate.DrivingState_Operation) &&
		int(fieldstate.DrivingState_Fault)&drivingState != int(fieldstate.DrivingState_Fault)
}

func IsDrivingStateOperationManual(drivingState int) bool {
	return int(fieldstate.DrivingState_Operation_Manual)&drivingState == int(fieldstate.DrivingState_Operation_Manual)
}

func IsDrivingStateOperationAuto(drivingState int) bool {
	return int(fieldstate.DrivingState_Operation_Auto)&drivingState == int(fieldstate.DrivingState_Operation_Auto)
}

// 正常运维，无故障
func IsDrivingStateMaintainNormal(drivingState int) bool {
	return int(fieldstate.DrivingState_Maintain)&drivingState == int(fieldstate.DrivingState_Maintain) &&
		int(fieldstate.DrivingState_Fault)&drivingState != int(fieldstate.DrivingState_Fault)
}

func IsDrivingStateOperation(drivingState int) bool {
	return int(fieldstate.DrivingState_Operation)&drivingState == int(fieldstate.DrivingState_Operation)
}

func IsDrivingStateMaintain(drivingState int) bool {
	return int(fieldstate.DrivingState_Maintain)&drivingState == int(fieldstate.DrivingState_Maintain)
}

func IsDrivingStateFault(drivingState int) bool {
	return int(fieldstate.DrivingState_Fault)&drivingState == int(fieldstate.DrivingState_Fault)
}

func IsDrivingStateManual(drivingState int) bool {
	return int(fieldstate.DrivingState_Manual)&drivingState == int(fieldstate.DrivingState_Manual)
}

func IsDrivingStateAuto(drivingState int) bool {
	return int(fieldstate.DrivingState_Auto)&drivingState == int(fieldstate.DrivingState_Auto)
}

func IsDrivingStateRemote(drivingState int) bool {
	return int(fieldstate.DrivingState_Remote)&drivingState == int(fieldstate.DrivingState_Remote)
}

func IsDrivingStateLock(drivingState int) bool {
	return int(fieldstate.DrivingState_Lock)&drivingState == int(fieldstate.DrivingState_Lock)
}

func IsDrivingStateBootup(drivingState int) bool {
	return drivingState == int(fieldstate.DrivingState_Bootup)
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
		Where(car.StateEQ(fieldstate.CarState_OnOperating)).
		Where(car.UseStateEQ(fieldstate.CarUseState_Idle)).
		SetUseState(fieldstate.CarUseState_Booking).
		SetUseOrderID(orderID).
		Save(ctx)
}

func (db *Database) UpdateCarUsing(ctx context.Context, carID int, orderID int) (*ent.Car, error) {
	return db.MainDB().Car.UpdateOneID(carID).
		Where(car.Or(
			car.UseStateEQ(fieldstate.CarUseState_Idle),
			car.UseStateEQ(fieldstate.CarUseState_Booking))).
		SetUseState(fieldstate.CarUseState_Using).
		SetUseOrderID(orderID).
		Save(ctx)
}

func (db *Database) UpdateCarAliveTime(ctx context.Context, deviceID string, aliveTime time.Time) (int, error) {
	return db.MainDB().Car.Update().Where(car.DeviceIDEQ(deviceID)).
		SetAliveTime(time.Now()).
		Save(ctx)
}

func (db *Database) UpdateCarDispatching(ctx context.Context, carID int, dispatchTaskID int, scheMode int) (*ent.Car, error) {
	return db.MainDB().Car.UpdateOneID(carID).
		Where(car.UseStateEQ(fieldstate.CarUseState_Idle)).
		Where(car.DispatchTaskIDEQ(0)).
		SetUseState(fieldstate.CarUseState_Dispatching).
		SetDispatchTaskID(dispatchTaskID).
		SetDispatchScheMode(scheMode).
		Save(ctx)
}

func (db *Database) UpdateCarDispatchEnd(ctx context.Context, carID int, useState int) (*ent.Car, error) {
	return db.MainDB().Car.UpdateOneID(carID).
		Where(car.UseStateEQ(fieldstate.CarUseState_Dispatching)).
		Where(car.DispatchTaskIDGT(0)).
		SetUseState(useState).
		SetDispatchTaskID(0).
		SetDispatchScheMode(0).
		Save(ctx)
}
