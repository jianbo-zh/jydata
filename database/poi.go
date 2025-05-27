package database

import (
	"context"

	"github.com/jianbo-zh/jydata/database/ent"
	"github.com/jianbo-zh/jydata/database/ent/poi"
)

type PoiType int

const (
	PoiType_Unknown         PoiType = 0
	PoiType_ParkingSpot     PoiType = 1 // 停车点
	PoiType_TouristSpot     PoiType = 2 // 景点
	PoiType_Toilet          PoiType = 3 // 厕所
	PoiType_Shop            PoiType = 4 // 商铺
	PoiType_EntranceAndExit PoiType = 5 // 出入口
)

func (m *Database) GetScenicAreaPois(ctx context.Context, scenicAreaID int, levels []int) ([]*ent.Poi, error) {
	query := m.MainDB().Poi.Query().Where(poi.ScenicAreaIDEQ(scenicAreaID))
	if len(levels) > 0 {
		query.Where(poi.LevelIn(levels...))
	}

	return query.All(ctx)
}

// CreatePOI 添加poi点
func (m *Database) CreatePOI(ctx context.Context, epoi *ent.Poi) (*ent.Poi, error) {
	return m.MainDB().Poi.Create().
		SetName(epoi.Name).
		SetScenicAreaID(epoi.ScenicAreaID).
		SetType(epoi.Type).
		SetWgsLon(epoi.WgsLon).
		SetWgsLat(epoi.WgsLat).
		SetGcjLon(epoi.GcjLon).
		SetGcjLat(epoi.GcjLat).
		SetBdLon(epoi.BdLon).
		SetBdLat(epoi.BdLat).
		SetIntroText(epoi.IntroText).
		SetImageIds(epoi.ImageIds).
		SetAudioID(epoi.AudioID).
		SetVideoID(epoi.VideoID).
		SetBroadcastRadius(epoi.BroadcastRadius).
		SetParkingRadius(epoi.ParkingRadius).
		SetLevel(epoi.Level).
		SetStopHeading(epoi.StopHeading).
		Save(ctx)
}
