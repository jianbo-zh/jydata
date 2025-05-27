// ent/schema/car.go

package schema

import (
	"time"

	"github.com/jianbo-zh/jydata/database/mixin"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Car holds the schema definition for the Car entity.
type CarExtendYokee struct {
	ent.Schema
}

// Fields of the Car.
func (CarExtendYokee) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Int("car_id").Comment("车辆ID"),
		field.Int("yokee_vehicle_id").Comment("九识分配的车辆ID"),
		field.String("yokee_vehicle_name").Comment("九识分配的车辆名称"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the Car.
func (CarExtendYokee) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Mixin of the Pet.
func (CarExtendYokee) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.SoftDeleteMixin{},
	}
}
