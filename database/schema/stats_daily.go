// ent/schema/poi.go

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// StatsDaily holds the schema definition for the Poi entity.
type StatsDaily struct {
	ent.Schema
}

// Fields of the StatsDaily.
func (StatsDaily) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Comment("ID"),
		field.Int("scenic_area_id").Comment("景区ID"),
		field.String("y").Comment("年"),
		field.String("ym").Comment("年月"),
		field.String("ymd").Comment("年月日"),
		field.Int("order_amount").Comment("订单额"),
		field.Int("order_count").Comment("订单数"),
		field.Int("order_duration").Comment("订单时长"),
		field.Int("order_distance").Comment("订单距离"),
		field.Int("register_user_count").Comment("注册用户数"),
		field.Int("order_user_count").Comment("下单用户数"),
		field.Int("operation_car_count").Comment("运营车辆数"),
		field.Int("operation_car_duration").Default(0).Comment("运营时长"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
	}
}

// Edges of the StatsDaily.
func (StatsDaily) Edges() []ent.Edge {
	return []ent.Edge{}
}
