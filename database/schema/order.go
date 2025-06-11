// ent/schema/order.go

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Comment("ID"),
		field.Int("type").Default(1).Comment("订单类型(1-租车订单 2-班次订单)"),
		field.Int("period").Default(0).Comment("时间段(0-未知、1-工作日、2-周末、3-节假日)"),

		field.String("order_no").Unique().Comment("订单号"),
		field.String("wx_tx_id").Default("").Comment("订单支付ID"),

		field.Int("scenic_area_id").Comment("景区ID"),
		field.String("mch_id").Default("").Comment("商户号"),
		field.String("scenic_area_name").Comment("景区名称"),

		field.Int("user_id").Comment("用户id"),
		field.String("open_id").Comment("用户微信OpenID"),
		field.String("nickname").Comment("用户昵称"),
		field.String("phone").Comment("用户手机号"),

		field.Int("car_id").Comment("车辆ID"),
		field.String("device_id").Comment("设备ID"),
		field.String("car_name").Comment("汽车名称"),
		field.String("car_license_plate").Comment("车辆号牌"),
		field.Int("model_id").Comment("型号ID"),
		field.String("model_name").Comment("型号名称"),

		field.Int("coupon_id").Default(0).Comment("优惠卷ID"),
		field.String("coupon_name").Default("").Comment("优惠卷名称"),

		field.Int("use_mileage_meter").Default(0).Comment("行驶里程"),
		field.Int("use_time_second").Default(0).Comment("使用时长"),

		field.Int("deposit_amount").Default(0).Comment("押金金额"),
		field.Int("original_amount").Default(0).Comment("订单原价(折扣前价格)"),
		field.Int("order_amount").Default(0).Comment("实付金额"),
		field.Int("refunded_amount").Default(0).Comment("已退款金额"),
		field.Int("coupon_amount").Default(0).Comment("优惠金额"),

		field.Int("order_state").Default(0).Comment("订单状态(0-待支付、1-待使用、2-进行中、3-待支付、4-待退费、5-已完成、6-已取消)"),
		field.Int("deposit_state").Default(0).Comment("押金状态(0-无支付、1-待付支付、2-已支付)"),
		field.Int("emergency_state").Default(0).Comment("紧急状态(0-无紧急、1-紧急呼救、2-已取消紧急呼救)"),

		field.Bool("is_test_order").Default(false).Comment("是否测试订单"),
		field.Bool("is_cancel").Default(false).Comment("是否取消"),
		field.Bool("is_profit_sharing").Default(false).Comment("是否分润"),
		field.String("remark").Default("").Comment("订单备注"),

		field.Int("user_score").Default(5).Comment("用户评分"),
		field.String("user_comment").Default("").Comment("用户评价"),

		field.Time("deposit_time").Optional().Comment("支付押金时间"),
		field.Time("finish_time").Optional().Comment("完成时间"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("car", Car.Type).Ref("orders").Unique().Field("car_id").Required(),
		edge.From("background_scenic_area", ScenicArea.Type).Ref("orders").Unique().Field("scenic_area_id").Required(),
		edge.To("billing", OrderBilling.Type).Unique(),
		edge.To("refund", OrderRefund.Type),
		edge.To("sharing", OrderSharing.Type).Unique(),
	}
}
