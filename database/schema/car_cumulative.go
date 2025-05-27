// ent/schema/car.go

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Car holds the schema definition for the Car entity.
type CarCumulative struct {
	ent.Schema
}

// Fields of the Car.
func (CarCumulative) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Int("car_id").Unique().Comment("车辆ID"),
		field.String("device_id").Unique().Comment("车辆编号"),
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
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the Car.
func (CarCumulative) Edges() []ent.Edge {
	return []ent.Edge{}
}
