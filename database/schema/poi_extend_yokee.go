package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// PoiExtendYokee POI扩展表=>九识停车点
type PoiExtendYokee struct {
	ent.Schema
}

// Fields of the User.
func (PoiExtendYokee) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Int("poi_id").Comment("POI点"),
		field.Int("yokee_station_id").Comment("九识站点ID"),
		field.String("yokee_station_name").Comment("九识站点ID"),
		field.Int("yokee_stop_id").Comment("九识停车点ID"),
		field.String("yokee_stop_name").Comment("九识停车点名称"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the User.
func (PoiExtendYokee) Edges() []ent.Edge {
	return []ent.Edge{}
}
