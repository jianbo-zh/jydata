package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Car holds the schema definition for the Car entity.
type AppVersion struct {
	ent.Schema
}

// Fields of the Car.
func (AppVersion) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.String("platform").Comment("平台系统"),
		field.String("app_name").Comment("应用名称"),
		field.String("version").Comment("版本号"),
		field.String("content").Comment("升级说明"),
		field.Bool("is_force_upgrade").Comment("是否强制升级"),
		field.Int("file_id").Comment("文件ID"),
		field.Int("state").Comment("发布状态(1-预发布 2-正式发布 3-终止发布)"),
		field.Time("publish_time").Default(time.Now()).Comment("发布时间"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}
