// ent/schema/order.go

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// OrderExtendFlight holds the schema definition for the OrderExtendFlight entity.
type OrderExtendFlight struct {
	ent.Schema
}

// Fields of the OrderExtendFlight.
func (OrderExtendFlight) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Comment("ID"),
		field.Int("order_id").Comment("订单编号"),
		field.Int("flight_id").Default(0).Comment("班次ID"),
		field.String("flight_no").Default("").Comment("班次编号"),
		field.Int("route_id").Default(0).Comment("路线ID"),
		field.String("route_name").Default("").Comment("路线名称"),
		field.Int("start_stop_id").Comment("上车点"),
		field.Int("start_stop_index").Default(0).Comment("上车点序号"),
		field.Int("end_stop_id").Comment("下车点"),
		field.Int("end_stop_index").Default(0).Comment("下车点序号"),
		field.Int("ticket_count").Default(0).Comment("购票数量"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

func (OrderExtendFlight) Edges() []ent.Edge {
	return []ent.Edge{}
}
