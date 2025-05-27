package fieldtype

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
