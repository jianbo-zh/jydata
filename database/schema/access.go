package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// PermissionRule holds the schema definition for the PermissionRule entity.
type Access struct {
	ent.Schema
}

// Fields of the PermissionRule.
func (Access) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.Int("pid").Comment("上级"),
		field.String("code").Comment("权限码"),
		field.String("name").Comment("权限名称"),
		field.String("url").Comment("url"),
		field.Int("sort").Comment("排序"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the PermissionRule.
func (Access) Edges() []ent.Edge {
	return []ent.Edge{}
}
