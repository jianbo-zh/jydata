// ent/schema/poi.go

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// StatsDailyScenicArea holds the schema definition for the Poi entity.
type StatsDailyScenicArea struct {
	ent.Schema
}

// Fields of the StatsDailyScenicArea.
func (StatsDailyScenicArea) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Comment("ID"),
		field.Int("scenic_area_id").Comment("景区ID"),
		field.Int("y").Comment("年"),
		field.Int("m").Comment("月"),
		field.Int("d").Comment("日"),
		field.Int("week").Comment("第几周"),
		field.Int("weekday").Comment("星期几(1,2,3,4,5,6,7)"),
		field.Int("period").Comment("时间段(1-工作日、2-周末、3-节假日)"),

		field.Int("user_count_login").Default(0).Comment("登录用户数"),
		field.Int("user_count_register").Default(0).Comment("注册用户数"),
		field.Int("user_count_order").Default(0).Comment("下单用户数"),
		field.Int("user_count_loginorder").Default(0).Comment("登录且下单用户数"),

		field.Int("order_count_create").Default(0).Comment("新增订单数"),
		field.Int("order_count_finish").Default(0).Comment("完成订单数"),
		field.Int("order_count_cancel").Default(0).Comment("取消订单数"),

		field.Int("order_amount").Default(0).Comment("订单额（分）"),
		field.Int("order_duration").Default(0).Comment("订单时长（秒）"),
		field.Int("order_mileage").Default(0).Comment("订单距离（米）"),

		field.Int("car_count_deploy").Default(0).Comment("投放车辆数量"),
		field.Int("car_count_operation").Default(0).Comment("运营车辆数量"),

		field.Float32("car_operation_manual_duration").Default(0).Comment("运营手动时长（秒）"),
		field.Float32("car_operation_auto_duration").Default(0).Comment("运营自动时长（秒）"),
		field.Float32("car_operation_normal_duration").Default(0).Comment("车辆正常态运营时长（秒）"),
		field.Float32("car_operation_fault_duration").Default(0).Comment("车辆故障态运营时长（秒）"),
		field.Float32("car_operation_mileage").Default(0).Comment("车辆运营里程（米）"),
		field.Uint32("car_operation_fault_times").Default(0).Comment("车辆运营故障次数"),

		field.Float32("car_maintain_manual_duration").Default(0).Comment("运维手动时长（秒）"),
		field.Float32("car_maintain_auto_duration").Default(0).Comment("运维自动时长（秒）"),
		field.Float32("car_maintain_normal_duration").Default(0).Comment("车辆正常态运维时长（秒）"),
		field.Float32("car_maintain_fault_duration").Default(0).Comment("车辆故障态运维时长（秒）"),
		field.Float32("car_maintain_mileage").Default(0).Comment("车辆运维里程（米）"),
		field.Uint32("car_maintain_fault_times").Default(0).Comment("车辆运维故障次数"),

		field.Int("car_alarm_times").Default(0).Comment("车辆告警次数"),

		field.Time("stats_time").Comment("统计的时间"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
	}
}

// Edges of the StatsDailyScenicArea.
func (StatsDailyScenicArea) Edges() []ent.Edge {
	return []ent.Edge{}
}
