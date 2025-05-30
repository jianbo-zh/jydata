// ent/schema/order.go

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OrderBilling holds the schema definition for the OrderBilling entity.
type OrderBilling struct {
	ent.Schema
}

// Fields of the OrderBilling.
func (OrderBilling) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Comment("ID"),
		field.Int("type").Default(1).Comment("订单类型(1-按租车计费 2-按班次计费)"),
		field.Int("order_id").Unique().Comment("订单ID"),

		field.Int("start_time_price").Default(0).Comment("起步价(按时)收费价格(单位：分)"),
		field.Int("start_time_unit").Default(1).Comment("起步价(按时)计量单位（单位：秒）"),
		field.Int("normal_time_price").Default(0).Comment("时间收费价格(单位：分)"),
		field.Int("normal_time_unit").Default(1).Comment("时间计量单位（单位：秒）"),

		field.Float("cumulative_second").Default(0).Comment("累积时长(秒)"),
		field.Float("cumulative_meter").Default(0).Comment("累积里程(米)"),
		field.Int("cumulative_stop").Default(0).Comment("累积站点(站)"),

		field.Int("start_stop_price").Default(0).Comment("起步价(按站)收费价格(单位：分)"),
		field.Int("start_stop_unit").Default(1).Comment("起步价(按站)计量单位（单位：站）"),
		field.Int("normal_stop_price").Default(0).Comment("按站收费价格(单位：分)"),
		field.Int("normal_stop_unit").Default(1).Comment("按站计量单位（单位：站）"),

		field.Int("coupon_id").Default(0).Comment("优惠卷ID"),
		field.Int("coupon_limit_amount").Default(0).Comment("优惠限制金额（单位：分）"),
		field.Int("coupon_deduction_amount").Default(0).Comment("优惠卷抵扣金额（单位：分）"),

		field.Int("capped_amount").Default(0).Comment("封顶价格（单位：分）"),

		field.Int("state").Default(0).Comment("计时状态(0-未开始 1-计时中 2-暂停中 4-已结束)"),

		field.Time("start_time").Optional().Comment("开始时间"),
		field.Time("finish_time").Optional().Comment("结束时间"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

func (OrderBilling) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).Ref("billing").Unique().Field("order_id").Required(),
	}
}
