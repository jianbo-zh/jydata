package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Car holds the schema definition for the Car entity.
type CarConfigPack struct {
	ent.Schema
}

// Fields of the Car.
func (CarConfigPack) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Int("scenic_area_id").Optional().Nillable().Comment("景区ID"),
		field.Int("model_id").Optional().Nillable().Comment("型号ID"),
		field.Int("car_id").Optional().Nillable().Comment("车辆ID"),
		field.String("car_version").Default("").Comment("适配车辆版本，空字符串表示适配所有版本"),
		field.String("name").Comment("配置名称"),
		field.String("remark").Comment("配置备注"),
		field.String("version").Comment("配置版本号"),
		field.Ints("config_ids").Optional().Comment("配置列表"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
	}
}

// Edges of the Car.
func (CarConfigPack) Edges() []ent.Edge {
	return []ent.Edge{}
}
