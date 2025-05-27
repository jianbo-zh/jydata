package fieldstate

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
