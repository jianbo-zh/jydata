package fieldtype

type PoiType int

const (
	PoiType_Unknown         PoiType = 0
	PoiType_ParkingSpot     PoiType = 1 // 停车点
	PoiType_TouristSpot     PoiType = 2 // 景点
	PoiType_Toilet          PoiType = 3 // 厕所
	PoiType_Shop            PoiType = 4 // 商铺
	PoiType_EntranceAndExit PoiType = 5 // 出入口
)
