package types

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
}

type RoutingPath struct {
	Segements []Segement
}

type StopStock struct {
	StopId int `json:"i,omitempty"` // 站点ID
	Sales  int `json:"s,omitempty"` // 售卖数量
	Ups    int `json:"u,omitempty"` // 上车数量
	Downs  int `json:"d,omitempty"` // 下车数量
}
