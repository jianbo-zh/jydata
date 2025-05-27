// ent/schema/car.go

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Car holds the schema definition for the Car entity.
type CarsModelsGroups struct {
	ent.Schema
}

// Fields of the VehicleModel.
func (CarsModelsGroups) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().Comment("ID"),
		field.Int("model_id").Default(0).Comment("型号ID"),
		field.String("group_remark").Default("").Comment("分组备注"),
		field.String("group_name").Comment("参数名称"),
		field.Int("status").Default(0).Comment("状态：0 - 不可用，1 - 可用"),
		field.Int("is_deleted").Default(0).Comment("是否已删除：0 - 正常，1 - 已删除"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the Car.
func (CarsModelsGroups) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("model", CarsModels.Type).Ref("groups").Unique().Field("model_id").Required(),
		edge.To("params", CarsModelsGroupsParams.Type),
	}
}
