package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().Comment("用户id"),
		field.Int("scenic_area_id").Optional().Comment("景区ID"),
		field.Int("origin").Default(1).Comment("用户来源：1-微信用户 2-APP用户"),
		field.Int("user_cls").Default(1).Comment("用户分类：1-普通用户 2-测试用户"),
		field.String("open_id").Unique().Comment("微信openid"),
		field.String("username").Optional().Nillable().Unique().Comment("用户账号"),
		field.String("nickname").Default("").Comment("昵称"),
		field.String("phone").Default("").Comment("手机号"),
		field.String("avatar_url").Default("").Comment("微信头像"),
		field.Uint8("gender").Default(0).Comment("性别"),
		field.String("password").Default("").Comment("密码"),
		field.String("country").Default("").Comment("国家"),
		field.String("province").Default("").Comment("省份"),
		field.String("city").Default("").Comment("城市"),
		field.Int("status").Default(1).Comment("1.显示0隐藏"),
		field.Uint8("is_deleted").Default(0).Comment("是否删除"),
		field.Time("login_time").Optional().Comment("登录时间"),
		field.String("language").Default("").Comment("语言"),
		field.Bool("is_tester").Default(false).Comment("是否测试人员"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("background_scenic_area", ScenicArea.Type).
			Ref("users").
			Unique().
			Field("scenic_area_id"),
	}
}

// Table specifies the existing table name.
func (User) Table() string {
	return "jiuyouzx_user"
}
