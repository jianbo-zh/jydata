package database

import (
	"context"
	"time"

	"github.com/jianbo-zh/jydata/database/ent"
	"github.com/jianbo-zh/jydata/database/ent/schetask"
	"github.com/jianbo-zh/jydata/database/fieldstate"
	"github.com/jianbo-zh/jydata/database/schema/types"
	"github.com/jianbo-zh/jylib/util"

	"entgo.io/ent/dialect/sql"
)

func (m *Database) AddScheTask(ctx context.Context, scheTask *ent.ScheTask) (*ent.ScheTask, error) {
	return m.MainDB().ScheTask.Create().
		SetUserType(scheTask.UserType).
		SetUserID(scheTask.UserID).
		SetScenicAreaID(scheTask.ScenicAreaID).
		SetCarID(scheTask.CarID).
		SetDeviceID(scheTask.DeviceID).
		SetDestID(scheTask.DestID).
		SetDestLon(scheTask.DestLon).
		SetDestLat(scheTask.DestLat).
		SetType(scheTask.Type).
		SetLoadLimit(scheTask.LoadLimit).
		SetState(scheTask.State).
		SetAbnormalState(scheTask.AbnormalState).
		SetRemark(scheTask.Remark).
		SetRoutingPath(types.RoutingPath{}).
		SetNillableEndTime(nil).
		Save(ctx)
}

func (m *Database) GetScheTask(ctx context.Context, taskID int) (*ent.ScheTask, error) {
	return m.MainDB().ScheTask.Get(ctx, taskID)
}

func (m *Database) GetLastScheTask(ctx context.Context, carID int) (*ent.ScheTask, error) {
	return m.MainDB().ScheTask.Query().
		Where(schetask.CarIDEQ(carID)).
		Order(schetask.ByID(sql.OrderDesc())).
		First(ctx)
}

func (m *Database) GetDoingScheTasks(ctx context.Context) ([]*ent.ScheTask, error) {
	return m.MainDB().ScheTask.Query().
		Where(schetask.StateNotIn(
			int(fieldstate.ScheTaskState_Finished),
			int(fieldstate.ScheTaskState_Canceled),
			int(fieldstate.ScheTaskState_Stopped),
		)).
		All(ctx)
}

func (m *Database) UpdateScheTaskStart(ctx context.Context, taskID int, scheState fieldstate.ScheTaskState, scheAbnormalState fieldstate.ScheTaskAbnormalState, rpath types.RoutingPath) (*ent.ScheTask, error) {
	return m.MainDB().ScheTask.UpdateOneID(taskID).
		SetState(int(scheState)).
		SetAbnormalState(int(scheAbnormalState)).
		SetRoutingPath(rpath).
		Save(ctx)
}

func (m *Database) UpdateScheTaskState(ctx context.Context, taskID int, scheState fieldstate.ScheTaskState, scheAbnormalState fieldstate.ScheTaskAbnormalState, remark string) (*ent.ScheTask, error) {
	update := m.MainDB().ScheTask.UpdateOneID(taskID).
		SetState(int(scheState)).
		SetAbnormalState(int(scheAbnormalState)).
		SetRemark(remark)

	if util.InSlice(int(scheState), []int{
		int(fieldstate.ScheTaskState_Finished),
		int(fieldstate.ScheTaskState_Canceled),
		int(fieldstate.ScheTaskState_Stopped),
	}) {
		update.SetEndTime(time.Now())
	}

	return update.Save(ctx)
}
