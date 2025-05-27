package database

import (
	"context"
	"jiuyouzx/pkg/util"
	"time"

	"github.com/jianbo-zh/jydata/database/ent"
	"github.com/jianbo-zh/jydata/database/ent/schetask"
	"github.com/jianbo-zh/jydata/database/schema/types"

	"entgo.io/ent/dialect/sql"
)

type ScheTaskState int

// 1-初始化 2-调度中 3-已暂停 5-停滞不前 6-已完成 7-已取消 8-系统终止 9-异常状态
const (
	// 初始化
	ScheTaskState_Init ScheTaskState = 1
	// 调度中
	ScheTaskState_Running ScheTaskState = 2
	// 已暂停
	ScheTaskState_Paused ScheTaskState = 3
	// 停滞中
	ScheTaskState_Stalled ScheTaskState = 5
	// 已完成
	ScheTaskState_Finished ScheTaskState = 6
	// 已取消
	ScheTaskState_Canceled ScheTaskState = 7
	// 系统终止
	ScheTaskState_Stopped ScheTaskState = 8
	// 异常状态
	ScheTaskState_Abnormal ScheTaskState = 9
)

type ScheTaskAbnormalState int

// 1-无异常 2-离线 3-道路外 4-有载人或物 5-驾驶状态异常 6-紧急制动中
const (
	// 没有异常
	ScheTaskAbnormalState_None ScheTaskAbnormalState = 1
	// 车辆离线
	ScheTaskAbnormalState_Offline ScheTaskAbnormalState = 2
	// 车辆在道路外
	ScheTaskAbnormalState_Outside ScheTaskAbnormalState = 3
	// 车辆有载人或物
	ScheTaskAbnormalState_HadLoad ScheTaskAbnormalState = 4
	// 车辆驾驶状态异常
	ScheTaskAbnormalState_DrivingState ScheTaskAbnormalState = 5
	// 车辆紧急制动中
	ScheTaskAbnormalState_Estop ScheTaskAbnormalState = 6
)

type ScheTaskType int

// 1-运营调度 2-运维调度
const (
	ScheTaskType_Operation ScheTaskType = 1 // 运营调度
	ScheTaskType_Maintain  ScheTaskType = 2 // 运维调度
)

type ScheTaskLoadLimit int

const (
	ScheTaskLoadLimit_None     ScheTaskLoadLimit = 1 // 没有限制
	ScheTaskLoadLimit_NoLoad   ScheTaskLoadLimit = 2 // 无负载调度
	ScheTaskLoadLimit_MustLoad ScheTaskLoadLimit = 3 // 有负载调度
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
			int(ScheTaskState_Finished),
			int(ScheTaskState_Canceled),
			int(ScheTaskState_Stopped),
		)).
		All(ctx)
}

func (m *Database) UpdateScheTaskStart(ctx context.Context, taskID int, scheState ScheTaskState, scheAbnormalState ScheTaskAbnormalState, rpath types.RoutingPath) (*ent.ScheTask, error) {
	return m.MainDB().ScheTask.UpdateOneID(taskID).
		SetState(int(scheState)).
		SetAbnormalState(int(scheAbnormalState)).
		SetRoutingPath(rpath).
		Save(ctx)
}

func (m *Database) UpdateScheTaskState(ctx context.Context, taskID int, scheState ScheTaskState, scheAbnormalState ScheTaskAbnormalState, remark string) (*ent.ScheTask, error) {
	update := m.MainDB().ScheTask.UpdateOneID(taskID).
		SetState(int(scheState)).
		SetAbnormalState(int(scheAbnormalState)).
		SetRemark(remark)

	if util.InSlice(int(scheState), []int{
		int(ScheTaskState_Finished),
		int(ScheTaskState_Canceled),
		int(ScheTaskState_Stopped),
	}) {
		update.SetEndTime(time.Now())
	}

	return update.Save(ctx)
}
