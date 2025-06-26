package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/jianbo-zh/jydata/database/mixin"
	"github.com/jianbo-zh/jydata/database/schema/types"
)

// File holds the schema definition for the File entity.
type OtaDeploy struct {
	ent.Schema
}

// Fields of the File.
func (OtaDeploy) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Comment("ID"),
		field.Int64("uuid").Unique().Comment("UUID"),
		field.Int("car_id").Comment("车辆ID"),
		field.String("car_name").Comment("车辆名称"),
		field.String("device_id").Comment("设备ID"),
		field.Int("ota_version_id").Comment("OTA版本ID"),
		field.String("ota_version_name").Comment("OTA版本名称"),
		field.String("ota_version_number").Comment("OTA版本号"),
		field.Int("state").Default(1).Comment("升级状态(1-升级中 2-升级成功 3-升级失败)"),
		field.String("errmsg").Default("").Comment("升级失败描述"),
		field.JSON("process", types.OtaProcess{}).Optional().Comment("OTA升级进度"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the File.
func (OtaDeploy) Edges() []ent.Edge {
	return nil
}

func (OtaDeploy) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("car_id", "ota_version_id").Unique(),
	}
}

// Mixin of the Pet.
func (OtaDeploy) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.SoftDeleteMixin{},
	}
}
