// ent/schema/poi.go

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Poi holds the schema definition for the Poi entity.
type Poi struct {
	ent.Schema
}

// Fields of the POI.
func (Poi) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Comment("ID"),
		field.String("name").Comment("景点名称"),
		field.String("alias").Default("").Comment("别名"),
		field.Int("scenic_area_id").Comment("景区ID"),
		field.Int("type").Comment("景点类型(0-其他，1-停车点 2-景点 3-厕所 4-商铺 5-出入口 6-公交站点)"),
		field.Float("wgs_lon").Default(0).Comment("经度(地球坐标系)"),
		field.Float("wgs_lat").Default(0).Comment("纬度(地球坐标系)"),
		field.Float("gcj_lon").Default(0).Comment("经度(火星坐标系)"),
		field.Float("gcj_lat").Default(0).Comment("纬度(火星坐标系)"),
		field.Float("bd_lon").Default(0).Comment("经度(百度坐标系)"),
		field.Float("bd_lat").Default(0).Comment("纬度(百度坐标系)"),
		field.Float("stop_heading").Default(0).Comment("停车点车辆朝向"),
		field.String("intro_text").SchemaType(map[string]string{
			dialect.MySQL: "varchar(6000)", // Override MySQL.
		}).Default("").Comment("简介文本"),
		field.Ints("image_ids").Default(nil).Comment("图片集"),
		field.Int("audio_id").Default(0).Comment("语音播报"),
		field.Int("video_id").Default(0).Comment("介绍视频"),
		field.Int("broadcast_radius").Default(0).Comment("播报半径"),
		field.Int("parking_radius").Default(0).Comment("停车半径"),
		field.Int("level").Default(1).Comment("POI层级"),
		field.Int("extend_yokee_id").Optional().Nillable().Comment("Yokee扩展ID"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the POI.
func (Poi) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("background_scenic_area", ScenicArea.Type).
			Ref("pois").
			Unique().
			Field("scenic_area_id").
			Required(),
	}

}
