package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Car holds the schema definition for the Car entity.
type CarConfig struct {
	ent.Schema
}

// Fields of the Car.
func (CarConfig) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Int("scenic_area_id").Optional().Nillable().Comment("景区ID"),
		field.Int("model_id").Optional().Nillable().Comment("型号ID"),
		field.Int("car_id").Optional().Nillable().Comment("车辆ID"),
		field.String("car_version").Default("").Comment("适配车辆版本，*表示适配所有版本"),
		field.String("name").Comment("配置名称"),
		field.String("remark").Comment("配置备注"),
		field.String("version").Comment("配置版本号"),
		field.String("save_path").Comment("配置存储路径"),
		field.Int("content_type").Default(1).Comment("配置类型（1-基础配置 2-景区配置 3-高精度地图 4-POI配置 5-CarUI）"),
		field.Int("content_field").Default(1).Comment("配置内容字段（1-file_id 2-pbtext）"),
		field.Int("content_file_id").Optional().Nillable().Comment("配置文件ID"),
		field.Text("content_pbtext").Default("").Comment("配置内容"),
		field.String("content_sha1").Comment("内容sha1校验码"),
		field.Ints("resource_file_ids").Optional().Comment("资源清单列表"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
	}
}

// Edges of the Car.
func (CarConfig) Edges() []ent.Edge {
	return []ent.Edge{ // 设置关联关系
		edge.From("background_scenic_area", ScenicArea.Type).
			Ref("config_files").
			Unique().
			Field("scenic_area_id"),
		edge.From("cars_models", CarsModels.Type).
			Ref("config_files").
			Unique().
			Field("model_id"),
		edge.From("car", Car.Type).
			Ref("config_files").
			Unique().
			Field("car_id"),
	}
}
