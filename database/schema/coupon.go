// ent/schema/car.go

package schema

import (
	"time"

	"github.com/jianbo-zh/jydata/database/mixin"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Coupon 优惠卷
type Coupon struct {
	ent.Schema
}

// Fields of the Coupon.
func (Coupon) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Int("scenic_area_id").Optional().Nillable().Comment("景区ID"),
		field.String("name").Comment("名称"),
		field.String("coupon_no").Comment("优惠卷编号"),
		field.Int("user_id").Comment("用户编号"),
		field.Int("limit_amount").Comment("限制金额（单位：分）"),
		field.Int("coupon_amount").Comment("优惠金额（单位：分）"),
		field.Int("bind_order_id").Optional().Nillable().Comment("绑定订单ID"),
		field.Int("state").Default(1).Comment("使用状态（1-未使用、2-已使用、3-已过期）"),
		field.Time("valid_start_time").Optional().Comment("有效期开始时间"),
		field.Time("valid_end_time").Optional().Comment("有效期结束时间"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the Coupon.
func (Coupon) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Mixin of the Pet.
func (Coupon) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.SoftDeleteMixin{},
	}
}
