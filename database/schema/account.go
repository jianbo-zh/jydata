package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().Comment("ID"),
		field.Int("scenic_area_id").Optional().Nillable().Comment("景区id"),
		field.String("name").NotEmpty().MaxLen(30).Comment("账户名称"),
		field.String("account").Unique().NotEmpty().MaxLen(30).Comment("账号"),
		field.String("password").NotEmpty().Comment("密码"),
		field.Int("role_id").Default(0).Comment("角色id"),
		field.Int("status").Default(1).Comment("状态1可用0不可用"),
		field.Bool("is_super").Default(false).Comment("是否超管"),
		field.String("token_hash").Default("").Comment("token MD5值，单点登录"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("background_scenic_area", ScenicArea.Type).
			Ref("accounts").
			Unique().
			Field("scenic_area_id"),
	}
}

// Mixin of the Account.
func (Account) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}

// Hooks of the Account.
func (Account) Hooks() []ent.Hook {
	return []ent.Hook{}
}

// Policies of the Account.
func (Account) Policies() []ent.Policy {
	return []ent.Policy{}
}
