package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// PermissionRule holds the schema definition for the PermissionRule entity.
type AppPush struct {
	ent.Schema
}

// Fields of the PermissionRule.
func (AppPush) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("user").Comment("用户ID"),
		field.String("title").Default("").Comment("消息标题"),
		field.String("content").Default("").Comment("消息内容"),
		field.Int("state").Default(1).Comment("状态：1-待处理，2-已成功，3-已失败"),
		field.String("remark").Default("").Comment("备注内容"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the PermissionRule.
func (AppPush) Edges() []ent.Edge {
	return []ent.Edge{}
}
