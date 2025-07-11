package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// CarsFlightExtendYokee 车辆发车班次
type CarsFlightExtendYokee struct {
	ent.Schema
}

// Fields of the CarsFlightExtendYokee.
func (CarsFlightExtendYokee) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Comment("ID"),
		field.Int("flight_id").Unique().Comment("班次ID"),
		field.Int("yokee_dispatch_id").Comment("九识分配的任务ID"),
		// field.Float32("yokee_speed_limit").Comment("速度限制，单位m/s"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

func (CarsFlightExtendYokee) Edges() []ent.Edge {
	return []ent.Edge{}
}
