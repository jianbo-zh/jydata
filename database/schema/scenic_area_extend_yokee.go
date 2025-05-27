package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ScenicAreaExtendYokee 景区扩展表=>九识站点
type ScenicAreaExtendYokee struct {
	ent.Schema
}

// Fields of the User.
func (ScenicAreaExtendYokee) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Int("scenic_area_id").Comment("景点ID"),
		field.String("yokee_app_id").Comment("九识分配的AppID"),
		field.String("yokee_app_key").Comment("九识分配的AppKey"),
		field.Int("yokee_org_id").Comment("九识分配的公司ID"),
		field.String("yokee_org_name").Comment("九识分配的公司名称"),
		field.Int("yokee_station_id").Comment("九识分配的站点ID"),
		field.String("yokee_station_name").Comment("九识分配的站点名称"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the User.
func (ScenicAreaExtendYokee) Edges() []ent.Edge {
	return []ent.Edge{}
}
