// ent/schema/car.go

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Car holds the schema definition for the Car entity.
type CarAlarm struct {
	ent.Schema
}

// Fields of the Car.
func (CarAlarm) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.String("device_id").Comment("设备ID"),
		field.Uint32("alarm_id").Comment("告警ID"),
		field.String("module_name").Comment("模块名称"),
		field.Uint32("type").Comment("告警类型（0-产生 1-恢复 2-事件）"),
		field.Uint32("level").Comment("告警等级（0-信息 1-告警 2-错误 3-故障）"),
		field.Uint32("can_ignore").Comment("忽略能力（0-不忽略 1-忽略一次 2-重复忽略）"),
		field.Uint32("effect_state").Comment("影响状态（0-无影响 1-手动 2-自动 3-手自动 4-远程 5-手远程 6-自远 7-手自远）"),
		field.String("desc").Comment("告警描述"),
		field.Uint64("uuid").Comment("告警时间"),
		field.Ints("associated_ids").Comment("关联告警（0-不忽略 1-忽略一次 2-重复忽略）"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the Car.
func (CarAlarm) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Index of the Car.
func (CarAlarm) Index() []ent.Index {
	return []ent.Index{}
}
