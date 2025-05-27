package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Identity holds the schema definition for the Identity entity.
type Role struct {
	ent.Schema
}

// Fields of the Identity.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("name").NotEmpty().Comment("名称"),
		field.Ints("access_ids").Default([]int{}).Comment("权限id"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the Identity.
func (Role) Edges() []ent.Edge {
	return []ent.Edge{}
}
