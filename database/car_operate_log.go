package database

type CarOperateState int

// 车辆操作结果：1-失败、2-成功
const (
	CarOperateState_Succ = 2
	CarOperateState_Fail = 1
)
