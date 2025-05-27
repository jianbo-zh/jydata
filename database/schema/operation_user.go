package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type OperationUser struct {
	ent.Schema
}

// Fields of the User.
func (OperationUser) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Comment("唯一标识"),
		field.Int("scenic_area_id").Default(0).Comment("景区ID"),
		field.String("username").Comment("性名"),
		field.String("nickname").Default("").Comment("微信昵称"),
		field.String("phone").Unique().Comment("手机号"),
		field.String("password").Sensitive().Comment("密码"),
		field.String("open_id").Default("").Comment("微信openid(唯一标识)"),
		field.String("avatar_url").Default("").Comment("微信头像"),
		field.Int("status").Default(1).Comment("状态：0 - 不可用，1 - 可用"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the User.
func (OperationUser) Edges() []ent.Edge {
	return nil
}
