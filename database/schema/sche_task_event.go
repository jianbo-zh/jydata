// ent/schema/car.go

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Car holds the schema definition for the Car entity.
type ScheTaskEvent struct {
	ent.Schema
}

// Fields of the Car.
func (ScheTaskEvent) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Int("sche_task_id").Comment("调度ID"),
		field.Int("state").Comment("调度状态（2-调度中 3-已暂停 5-停滞不前 6-已完成 7-已取消 8-系统终止 9-异常状态）"),
		field.Int("abnormal_state").Default(0).Comment("异常子状态（1-无异常 2-离线 3-道路外 4-有载人或物 5-驾驶状态异常）"),
		field.String("remark").Comment("日志记录"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
	}
}

// Edges of the Car.
func (ScheTaskEvent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sche_task", ScheTask.Type).
			Ref("events").
			Unique().
			Field("sche_task_id").
			Required(),
	}
}
