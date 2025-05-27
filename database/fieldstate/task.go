package fieldstate

// 任务状态
type TaskState int

// 状态(0-待处理、1-待重试、2-执行成功、3-执行失败)
const (
	TaskState_Pending = 0 // 待处理
	TaskState_Retry   = 1 // 待重试
	TaskState_Success = 2 // 执行成功
	TaskState_Failed  = 3 // 执行失败
)
