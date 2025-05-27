package database

import (
	"context"
	"fmt"

	"github.com/jianbo-zh/jydata/database/ent/carcumulative"
	"github.com/jianbo-zh/jydata/database/fieldstate"
)

func (r *Database) IncrCarCumulativeTime(ctx context.Context, carID int, deviceID string, drivingState fieldstate.DrivingState, sec float32) error {
	update := r.MainDB().CarCumulative.Update().Where(carcumulative.DeviceIDEQ(deviceID))

	switch drivingState {
	case fieldstate.DrivingState_Operation_Lock:
		if rows, err := update.AddOperationLockDuration(sec).
			Save(ctx); err != nil {
			return fmt.Errorf("db.Update carcumulative operation lock time error: %w", err)

		} else if rows == 0 {
			if _, err = r.MainDB().CarCumulative.Create().
				SetCarID(carID).
				SetDeviceID(deviceID).
				SetOperationLockDuration(sec).
				Save(ctx); err != nil {
				return fmt.Errorf("db.Create carcumulative operation lock time error: %w", err)
			}
		}
	case fieldstate.DrivingState_Operation_Manual:
		if rows, err := update.AddOperationManualDuration(sec).
			Save(ctx); err != nil {
			return fmt.Errorf("db.Update carcumulative operation manual time error: %w", err)

		} else if rows == 0 {
			if _, err = r.MainDB().CarCumulative.Create().
				SetCarID(carID).
				SetDeviceID(deviceID).
				SetOperationManualDuration(sec).
				Save(ctx); err != nil {
				return fmt.Errorf("db.Create carcumulative operation manual time error: %w", err)
			}
		}
	case fieldstate.DrivingState_Operation_Auto:
		if rows, err := update.AddOperationAutoDuration(sec).
			Save(ctx); err != nil {
			return fmt.Errorf("db.Update carcumulative operation auto time error: %w", err)

		} else if rows == 0 {
			if _, err = r.MainDB().CarCumulative.Create().
				SetCarID(carID).
				SetDeviceID(deviceID).
				SetOperationAutoDuration(sec).
				Save(ctx); err != nil {
				return fmt.Errorf("db.Create carcumulative operation auto time error: %w", err)
			}
		}
	case fieldstate.DrivingState_Operation_Fault:
		if rows, err := update.AddOperationFaultDuration(sec).
			Save(ctx); err != nil {
			return fmt.Errorf("db.Update carcumulative operation fault time error: %w", err)

		} else if rows == 0 {
			if _, err = r.MainDB().CarCumulative.Create().
				SetCarID(carID).
				SetDeviceID(deviceID).
				SetOperationFaultDuration(sec).
				Save(ctx); err != nil {
				return fmt.Errorf("db.Create carcumulative operation fault time error: %w", err)
			}
		}
	case fieldstate.DrivingState_Maintain_Lock:
		if rows, err := update.AddMaintainLockDuration(sec).
			Save(ctx); err != nil {
			return fmt.Errorf("db.Update carcumulative maintain lock time error: %w", err)

		} else if rows == 0 {
			if _, err = r.MainDB().CarCumulative.Create().
				SetCarID(carID).
				SetDeviceID(deviceID).
				SetMaintainLockDuration(sec).
				Save(ctx); err != nil {
				return fmt.Errorf("db.Create carcumulative maintain lock time error: %w", err)
			}
		}
	case fieldstate.DrivingState_Maintain_Manual:
		if rows, err := update.AddMaintainManualDuration(sec).
			Save(ctx); err != nil {
			return fmt.Errorf("db.Update carcumulative maintain manual time error: %w", err)

		} else if rows == 0 {
			if _, err = r.MainDB().CarCumulative.Create().
				SetCarID(carID).
				SetDeviceID(deviceID).
				SetMaintainManualDuration(sec).
				Save(ctx); err != nil {
				return fmt.Errorf("db.Create carcumulative maintain manual time error: %w", err)
			}
		}
	case fieldstate.DrivingState_Maintain_Auto:
		if rows, err := update.AddMaintainAutoDuration(sec).
			Save(ctx); err != nil {
			return fmt.Errorf("db.Update carcumulative maintain auto time error: %w", err)

		} else if rows == 0 {
			if _, err = r.MainDB().CarCumulative.Create().
				SetCarID(carID).
				SetDeviceID(deviceID).
				SetMaintainAutoDuration(sec).
				Save(ctx); err != nil {
				return fmt.Errorf("db.Create carcumulative maintain auto time error: %w", err)
			}
		}
	case fieldstate.DrivingState_Maintain_Remote:
		if rows, err := update.AddMaintainRemoteDuration(sec).
			Save(ctx); err != nil {
			return fmt.Errorf("db.Update carcumulative maintain remote time error: %w", err)

		} else if rows == 0 {
			if _, err = r.MainDB().CarCumulative.Create().
				SetCarID(carID).
				SetDeviceID(deviceID).
				SetMaintainRemoteDuration(sec).
				Save(ctx); err != nil {
				return fmt.Errorf("db.Create carcumulative maintain remote time error: %w", err)
			}
		}
	case fieldstate.DrivingState_Maintain_Fault:
		if rows, err := update.AddMaintainFaultDuration(sec).
			Save(ctx); err != nil {
			return fmt.Errorf("db.Update carcumulative maintain fault time error: %w", err)

		} else if rows == 0 {
			if _, err = r.MainDB().CarCumulative.Create().
				SetCarID(carID).
				SetDeviceID(deviceID).
				SetMaintainFaultDuration(sec).
				Save(ctx); err != nil {
				return fmt.Errorf("db.Create carcumulative maintain fault time error: %w", err)
			}
		}
	}

	return nil
}
