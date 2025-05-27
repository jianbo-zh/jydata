package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type MapVersion struct {
	ent.Schema
}

// Fields of the User.
func (MapVersion) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Int("scenic_area_id").Comment("景区ID"),
		field.Int("file_id").Comment("文件ID"),
		field.String("file_path").Comment("文件路径"),
		field.String("version").Unique().Comment("版本号"),
		field.String("name").Default("").Comment("版本名称"),
		field.String("remark").Default("").Comment("版本备注"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the User.
func (MapVersion) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("background_scenic_area", ScenicArea.Type).
			Ref("map_versions").
			Unique().
			Field("scenic_area_id").
			Required(),
	}
}
