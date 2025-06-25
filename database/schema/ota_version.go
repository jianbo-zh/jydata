package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/jianbo-zh/jydata/database/mixin"
	"github.com/jianbo-zh/jydata/database/schema/types"
)

// File holds the schema definition for the File entity.
type OtaVersion struct {
	ent.Schema
}

// Fields of the File.
func (OtaVersion) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Comment("ID"),
		field.Int("scenic_area_id").Optional().Nillable().Comment("景区ID"),
		field.Int("model_id").Optional().Nillable().Comment("景区ID"),
		field.String("name").Comment("名称"),
		field.String("version").Comment("版本号"),
		field.JSON("content", types.OtaContent{}).Comment("OTA内容"),
		field.Int("state").Comment("升级状态（1-待升级 2-升级中 3-部分失败 4-升级成功 5-升级失败）"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the File.
func (OtaVersion) Edges() []ent.Edge {
	return nil
}

// Mixin of the Pet.
func (OtaVersion) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.SoftDeleteMixin{},
	}
}
