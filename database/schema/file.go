// ent/schema/file.go

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// File holds the schema definition for the File entity.
type File struct {
	ent.Schema
}

// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Comment("ID"),
		field.Int("storage_type").Default(1).Comment("存储类型（1-Local、2-OSS）"),
		field.Int("creator_id").Comment("用户ID"),
		field.Int("scenic_area_id").Comment("景区ID"),
		field.Int("file_category").Comment("文件分类(0-其他,1-POI,2-停车点)"),
		field.Int("file_type").Comment("文件类型(0-其他,1-图片,2-音频,3-视频)"),
		field.String("file_sha1").Comment("文件sha1"),
		field.String("file_path").Comment("文件路径"),
		field.String("mime_type").Comment("文件MIME"),
		field.String("file_suffix").Comment("文件后缀"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the File.
func (File) Edges() []ent.Edge {
	return nil
}
