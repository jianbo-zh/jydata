package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type CarConfigStatus struct {
	ent.Schema
}

func (CarConfigStatus) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Int("car_id").Comment("车辆ID"),
		field.String("device_id").Comment("车辆设备ID"),
		field.String("config_path").Comment("配置路径"),
		field.String("version").Comment("配置版本号"),
		field.String("remark").Comment("备注说明"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

func (CarConfigStatus) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (CarConfigStatus) Indexes() []ent.Index {
	return []ent.Index{}
}
