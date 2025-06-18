// ent/schema/car.go

package schema

import (
	"time"

	"github.com/jianbo-zh/jydata/database/mixin"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Car holds the schema definition for the Car entity.
type SshAccount struct {
	ent.Schema
}

// Fields of the Car.
func (SshAccount) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.String("type").Comment("账号类型（pgyvisitor-蒲公英）"),
		field.String("username").Comment("Ssh账号"),
		field.String("password").Comment("Ssh密码"),
		field.Int("scenic_area_id").Optional().Nillable().Comment("景区ID"),
		field.Int("car_id").Optional().Nillable().Comment("车辆ID"),
		field.Int("state").Default(1).Comment("使用状态（1-空闲中 2-已使用）"),
		field.Time("use_time").Optional().Comment("使用时间"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the Car.
func (SshAccount) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Mixin of the Pet.
func (SshAccount) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.SoftDeleteMixin{},
	}
}
