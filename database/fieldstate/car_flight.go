package fieldstate

// 车辆班次状态
type CarFlightState int

// 1-待发车 2-行驶中 2-已结束 3-已取消
const (
	CarFlightState_Pending = 1
	CarFlightState_Going   = 2
	CarFlightState_Finish  = 3
	CarFlightState_Cancel  = 4
)
