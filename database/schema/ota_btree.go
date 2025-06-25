package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/jianbo-zh/jydata/database/mixin"
)

// File holds the schema definition for the File entity.
type OtaBtree struct {
	ent.Schema
}

// Fields of the File.
func (OtaBtree) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Comment("ID"),
		field.Int("file_id").Comment("文件ID"),
		field.String("name").Comment("名称"),
		field.String("remark").Default("").Comment("备注"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the File.
func (OtaBtree) Edges() []ent.Edge {
	return nil
}

// Mixin of the Pet.
func (OtaBtree) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.SoftDeleteMixin{},
	}
}
