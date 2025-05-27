// ent/schema/car.go

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Car holds the schema definition for the Car entity.
type CarsModelsGroupsParams struct {
	ent.Schema
}

// Fields of the VehicleModel.
func (CarsModelsGroupsParams) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().Comment("ID"),
		field.Int("model_id").Default(0).Comment("型号ID"),
		field.Int("group_id").Default(0).Comment("分组ID"),
		field.String("param_remark").Comment("参数描述"),
		field.String("param_name").Comment("参数名称"),
		field.Enum("param_type").Values("int", "double", "string").Comment("参数类型"),
		field.String("value_range").Default("").Comment("取值范围"),
		field.Int("status").Default(0).Comment("状态：0 - 不可用，1 - 可用"),
		field.Int("is_deleted").Default(0).Comment("是否已删除：0 - 正常，1 - 已删除"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the Car.
func (CarsModelsGroupsParams) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("model", CarsModels.Type).Ref("params").Unique().Field("model_id").Required(),
		edge.From("group", CarsModelsGroups.Type).Ref("params").Unique().Field("group_id").Required(),
	}
}

func (CarsModelsGroupsParams) Index() []ent.Index {
	return []ent.Index{
		index.Fields("group_id", "param_key").Unique(),
	}
}
