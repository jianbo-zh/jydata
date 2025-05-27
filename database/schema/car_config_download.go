package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type CarConfigDownload struct {
	ent.Schema
}

func (CarConfigDownload) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Int64("uuid").Unique(),
		field.Int("car_id").Comment("车辆ID"),
		field.String("device_id").Comment("车辆设备ID"),
		field.String("download_state").Default("").Comment("下载状态"),
		field.Int("download_process").Default(0).Comment("下载进度"),
		field.Ints("config_ids").Comment("配置IDs"),
		field.String("remark").Comment("备注说明"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

func (CarConfigDownload) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (CarConfigDownload) Indexes() []ent.Index {
	return []ent.Index{}
}
