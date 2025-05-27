package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type ScenicAreaMap struct {
	ent.Schema
}

// Fields of the User.
func (ScenicAreaMap) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Int("scenic_area_id").Unique().Comment("景区ID"),
		field.String("name").Comment("地图名称"),
		field.Int("base_map_file_id").Comment("规划地图"),
		field.Int("routing_map_file_id").Comment("路由地图"),
		field.Int("sim_map_file_id").Comment("dv使用，降采样"),
		field.Int("fence_map_file_id").Comment("围栏地图"),
		field.Int("carui_map_file_id").Default(0).Comment("CarUI地图"),
		field.String("carui_ne_coord").Default("").Comment("CarUI坐标(东北)"),
		field.String("carui_sw_coord").Default("").Comment("CarUI坐标(西南)"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the User.
func (ScenicAreaMap) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("background_scenic_area", ScenicArea.Type).
			Ref("map").
			Unique().
			Field("scenic_area_id").
			Required(),
	}
}
