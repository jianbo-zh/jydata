package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// SystemLog holds the schema definition for the SystemLog entity.
type SystemLog struct {
	ent.Schema
}

// Fields of the SystemLogs.
func (SystemLog) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique().Immutable(),
		field.Int("timestamp").Default(int(time.Now().Unix())).Immutable().Comment("日志时间"),
		field.String("action").NotEmpty().Comment("操作类型"),
		field.String("user").NotEmpty().Comment("用户"),
		field.String("scenic_area").NotEmpty().Comment("景区"),
		field.String("source_ip").NotEmpty().Comment("来源IP"),
		field.Enum("login_type").NamedValues("Web", "web", "H5", "h5", "App", "app").Comment("登录端类型"),
		field.Text("content").NotEmpty().Comment("日志内容"),
		field.Enum("status").NamedValues("Success", "success", "Failure", "failure").Comment("状态"),
		field.Text("remarks").Optional().Comment("备注"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
	}
}

// Edges of the SystemLogs.
func (SystemLog) Edges() []ent.Edge {
	return []ent.Edge{}
}
