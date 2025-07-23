package fieldtype

type ScheMode int

const (
	ScheMode_UserFreeScheduling     ScheMode = 1 // 【用户】自由调度
	ScheMode_OpsAttractionCustomers ScheMode = 2 // 【运营】揽客模式
	ScheMode_OpsDeploymentVehicles  ScheMode = 3 // 【运营】部署模式
)

type ScheArgsLoadLimit int

const (
	ScheArgsLoadLimit_None     ScheArgsLoadLimit = 1 // 没有限制
	ScheArgsLoadLimit_NoLoad   ScheArgsLoadLimit = 2 // 无负载调度
	ScheArgsLoadLimit_MustLoad ScheArgsLoadLimit = 3 // 有负载调度
)

type ScheArgsAroundObstacles int

const (
	ScheArgsAroundObstacles_No  ScheArgsAroundObstacles = 0 // 不绕障
	ScheArgsAroundObstacles_Yes ScheArgsAroundObstacles = 1 // 绕障
)

type ScheArgsEndLockState int

const (
	ScheArgsEndLockState_OperationLock ScheArgsEndLockState = 5 // 不锁车
	ScheArgsEndLockState_MaintainLock  ScheArgsEndLockState = 6 // 锁车
)
