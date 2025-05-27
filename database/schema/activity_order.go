// ent/schema/order.go

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Order holds the schema definition for the Order entity.
type ActivityOrder struct {
	ent.Schema
}

// Fields of the Order.
func (ActivityOrder) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Comment("ID"),
		field.String("order_no").Unique().Comment("订单号"),
		field.String("wx_tx_id").Default("").Comment("订单支付ID"),
		field.Int("scenic_area_id").Comment("景区ID"),
		field.String("mch_id").Default("").Comment("商户号"),
		field.String("scenic_area_name").Comment("景区名称"),
		field.Int("user_id").Comment("用户id"),
		field.String("open_id").Comment("用户微信OpenID"),
		field.Int("order_amount").Default(0).Comment("订单金额"),
		field.Int("order_state").Default(0).Comment("订单状态(0-待支付、1-已支付)"),
		field.String("remark").Default("").Comment("订单备注"),
		field.Time("paid_time").Optional().Comment("支付时间"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

func (ActivityOrder) Edges() []ent.Edge {
	return []ent.Edge{}
}
