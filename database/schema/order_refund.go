// ent/schema/order.go

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// 订单退款
type OrderRefund struct {
	ent.Schema
}

func (OrderRefund) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Comment("ID"),
		field.Int("type").Default(0).Comment("退款类型（0-结算退款 1-运管退款 2-申诉退款）"),
		field.Int("initiator_id").Default(0).Comment("退款发起人ID"),
		field.Int("scenic_area_id").Default(0).Comment("景区ID"),
		field.Int("order_id").Comment("订单ID"),
		field.Int("order_appeal_id").Optional().Nillable().Comment("订单申诉ID"),
		field.String("order_no").Default("").Comment("订单编号"),
		field.String("refund_no").Unique().Comment("退款单号"),
		field.String("wx_refund_id").Default("").Comment("微信退款ID"),
		field.Int("refund_amount").Default(0).Comment("退款金额"),
		field.Int("state").Default(1).Comment("退款状态(1-待退款、2-已退款)"),
		field.String("remark").Default("").Comment("退款备注"),
		field.String("errmsg").Default("").Comment("失败原因"),
		field.Time("finish_time").Nillable().Optional().Comment("完成时间"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

func (OrderRefund) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).Ref("refund").Unique().Field("order_id").Required(),
	}
}
