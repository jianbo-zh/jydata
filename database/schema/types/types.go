package types

import "github.com/jianbo-zh/jydata/database/fieldtype"

type OrderSharingReceiver struct {
	Type         int     `json:"type,omitempty"` // 1-商户 2-个人
	Account      string  `json:"account,omitempty"`
	Name         string  `json:"name,omitempty"`
	SharingRatio float64 `json:"ratio,omitempty"`
}

type Coord struct {
	Id  int
	Lat float64
	Lon float64
}

type Segement struct {
	Points []Coord
	Length float64
}

type RoutingPath struct {
	Segements []Segement
	Length    float64
}

type StopStock struct {
	StopId int `json:"i,omitempty"` // 站点ID
	Sales  int `json:"s,omitempty"` // 售卖数量
	Ups    int `json:"u,omitempty"` // 上车数量
	Downs  int `json:"d,omitempty"` // 下车数量
}

type OtaContent struct {
	ReqSource           int    `json:"rs,omitempty"`
	BtContent           string `json:"bc,omitempty"`
	BtContentSignature  string `json:"bcs,omitempty"`
	OtaContent          string `json:"oc,omitempty"`
	OtaContentSignature string `json:"ocs,omitempty"`
}

type OtaProgress struct {
	Progress  []*PkgProgress
	Completed bool
}

type PkgProgress struct {
	Name     string
	Stage    string
	Progress float64
	Result   int
}

type ScheArgs struct {
	LoadLimit       fieldtype.ScheArgsLoadLimit       `json:"load_limit,omitempty"`       // 负载限制(0-无限制 1-无负载调度 2-有负载调度)
	AroundObstacles fieldtype.ScheArgsAroundObstacles `json:"around_obstacles,omitempty"` // 是否绕过障碍物(0-不绕障 1-绕障)
	EndLockState    fieldtype.ScheArgsEndLockState    `json:"end_lock_state,omitempty"`   // 结束锁车态（见车辆状态表）
}
