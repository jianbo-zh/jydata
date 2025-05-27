// ent/schema/poi.go

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// StatsDailyCar holds the schema definition for the Poi entity.
type StatsDailyCar struct {
	ent.Schema
}

// Fields of the StatsDailyCar.
func (StatsDailyCar) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Comment("ID"),
		field.Int("scenic_area_id").Comment("景区ID"),
		field.Int("car_id").Comment("车辆ID"),
		field.Int("y").Comment("年"),
		field.Int("m").Comment("月"),
		field.Int("d").Comment("日"),
		field.Int("week").Comment("第几周"),
		field.Int("weekday").Comment("星期几(1,2,3,4,5,6,7)"),
		field.Int("period").Comment("时间段(1-工作日、2-周末、3-节假日)"),

		field.Int("order_count_create").Default(0).Comment("新建订单数"),
		field.Int("order_count_finish").Default(0).Comment("完成订单数"),
		field.Int("order_count_cancel").Default(0).Comment("取消订单数"),
		field.Int("order_amount").Default(0).Comment("订单额"),
		field.Int("order_duration").Default(0).Comment("订单时长"),
		field.Int("order_mileage").Default(0).Comment("订单距离"),
		field.Int("order_score").Default(0).Comment("订单评分"),
		field.Float32("operation_lock_duration").Default(0).Comment("运营锁车时长（秒）"),
		field.Float32("operation_manual_duration").Default(0).Comment("运营手动时长（秒）"),
		field.Float32("operation_auto_duration").Default(0).Comment("运营自动时长（秒）"),
		field.Float32("operation_fault_duration").Default(0).Comment("运营故障时长（秒）"),
		field.Uint32("operation_lock_times").Default(0).Comment("运营锁车次数"),
		field.Uint32("operation_manual_times").Default(0).Comment("运营手动次数"),
		field.Uint32("operation_auto_times").Default(0).Comment("运营自动次数"),
		field.Uint32("operation_fault_times").Default(0).Comment("运营故障次数"),
		field.Float32("operation_manual_mileage").Default(0).Comment("运营手动里程（米）"),
		field.Float32("operation_auto_mileage").Default(0).Comment("运营自动里程（米）"),
		field.Float32("maintain_lock_duration").Default(0).Comment("运维锁车时长（秒）"),
		field.Float32("maintain_manual_duration").Default(0).Comment("运维手动时长（秒）"),
		field.Float32("maintain_auto_duration").Default(0).Comment("运维自动时长（秒）"),
		field.Float32("maintain_remote_duration").Default(0).Comment("运维远程时长（秒）"),
		field.Float32("maintain_fault_duration").Default(0).Comment("运维故障时长（秒）"),
		field.Uint32("maintain_lock_times").Default(0).Comment("运维锁车次数"),
		field.Uint32("maintain_manual_times").Default(0).Comment("运维手动次数"),
		field.Uint32("maintain_auto_times").Default(0).Comment("运维自动次数"),
		field.Uint32("maintain_remote_times").Default(0).Comment("运维远程次数"),
		field.Uint32("maintain_fault_times").Default(0).Comment("运维故障次数"),
		field.Float32("maintain_manual_mileage").Default(0).Comment("运维手动里程（米）"),
		field.Float32("maintain_auto_mileage").Default(0).Comment("运维自动里程（米）"),
		field.Float32("maintain_remote_mileage").Default(0).Comment("运维远程里程（米）"),

		field.Uint32("alarm_times").Default(0).Comment("告警次数"),

		field.Time("stats_time").Comment("统计的时间"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
	}
}

// Edges of the StatsDailyCar.
func (StatsDailyCar) Edges() []ent.Edge {
	return []ent.Edge{}
}
