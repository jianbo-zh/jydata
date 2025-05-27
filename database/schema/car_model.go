// ent/schema/car.go

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Car holds the schema definition for the Car entity.
type CarsModels struct {
	ent.Schema
}

// Fields of the VehicleModel.
func (CarsModels) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().Comment("ID"),
		field.String("model_name").Comment("型号名称"),
		field.String("model_remark").Default("").Comment("型号名称(中文)"),
		field.Int("group_count").Default(0).Comment("分组数量"),
		field.Int("status").Default(0).Comment("状态：0 - 不可用，1 - 可用"),
		field.Int("car_incr").Default(0).Comment("车辆数递增值"),
		field.Int("is_deleted").Default(0).Comment("是否已删除：0 - 正常，1 - 已删除"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the Car.
func (CarsModels) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cars", Car.Type),
		edge.To("groups", CarsModelsGroups.Type),
		edge.To("params", CarsModelsGroupsParams.Type),
		edge.To("config_files", CarConfig.Type),
	}
}
