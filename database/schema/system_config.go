package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// SystemConfig holds the schema definition for the SystemConfig entity.
type SystemConfig struct {
	ent.Schema
}

// Fields of the SystemConfig.
func (SystemConfig) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.Enum("group").Values("base", "pay", "payment").Comment("分组"),
		field.String("name").Comment("名称"),
		field.String("key").NotEmpty().Comment("键"),
		field.String("value").Comment("值"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the SystemConfig.
func (SystemConfig) Edges() []ent.Edge {
	return []ent.Edge{}
}
