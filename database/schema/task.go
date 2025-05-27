package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the User.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("type").Comment("类型"),
		field.Int("rel_id").Default(0).Comment("关联ID"),
		field.Any("rel_data").Optional().Comment("关联数据"),
		field.Int("state").Default(0).Comment("状态(0-待处理、1-待重试、2-执行成功、3-执行失败)"),
		field.Int("try_times").Default(0).Comment("重试次数"),
		field.String("error_msg").Default("").Comment("错误描述"),
		field.Time("start_time").Comment("开始时间"),
		field.Time("next_time").Comment("下次时间"),
		field.Time("end_time").Optional().Comment("结束时间"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return nil
}
