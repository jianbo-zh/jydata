// ent/schema/order.go

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// 订单申诉
type OrderAppeal struct {
	ent.Schema
}

func (OrderAppeal) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Comment("ID"),
		field.Int("scenic_area_id").Comment("景区ID"),
		field.Int("user_id").Comment("用户ID"),
		field.Int("order_id").Comment("订单ID"),
		field.String("order_no").Comment("订单编号"),
		field.Int("type").Default(1).Comment("1-多付款 2-其他"),
		field.Int("end_stop_id").Default(0).Comment("停车点ID"),
		field.Int("end_stop_image_id").Default(0).Comment("停车点图片"),
		field.Int("state").Default(1).Comment("申述状态(1-待审核、2-申诉成功、3-申诉失败、4-已取消)"),
		field.Int("refund_amount").Default(0).Comment("退款金额"),
		field.String("user_comment").Default("").Comment("申诉内容"),
		field.String("review_comment").Default("").Comment("审核备注"),
		field.Time("refund_time").Optional().Nillable().Comment("退款时间"),
		field.Time("review_time").Optional().Nillable().Comment("审核时间"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

func (OrderAppeal) Edges() []ent.Edge {
	return []ent.Edge{}
}
