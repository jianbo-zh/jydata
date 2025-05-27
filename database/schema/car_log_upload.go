package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type CarLogUpload struct {
	ent.Schema
}

func (CarLogUpload) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Int64("uuid").Unique(),
		field.Int("car_id").Comment("车辆ID"),
		field.String("device_id").Comment("车辆设备ID"),
		field.String("upload_state").Default("").Comment("上传状态"),
		field.Int("upload_process").Default(0).Comment("上传进度"),
		field.String("download_url").Comment("下载地址"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

func (CarLogUpload) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (CarLogUpload) Indexes() []ent.Index {
	return []ent.Index{}
}
