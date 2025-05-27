// ent/schema/file.go

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Feedback struct {
	ent.Schema
}

func (Feedback) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Comment("ID"),
		field.Int("user_id").Comment("用户ID"),
		field.Int("scenic_area_id").Comment("景区ID"),
		field.String("content").MaxLen(1000).Comment("反馈内容"),
		field.Int("state").Default(1).Comment("状态(1-待处理 2-已处理)"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

func (Feedback) Edges() []ent.Edge {
	return nil
}
