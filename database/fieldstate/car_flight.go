package fieldstate

// 车辆班次状态
type CarFlightState int

// 1-待发车 2-路线规划中 3-行驶中 4-到停靠点 5-到目的地 6-已取消
const (
	CarFlightState_Pending CarFlightState = 1
	CarFlightState_Routing CarFlightState = 2
	CarFlightState_Onway   CarFlightState = 3
	CarFlightState_Atstop  CarFlightState = 4
	CarFlightState_Finish  CarFlightState = 5
	CarFlightState_Cancel  CarFlightState = 6
)
