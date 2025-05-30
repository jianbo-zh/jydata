package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// BillingStrategy holds the schema definition for the BillingStrategy entity.
type BillingStrategy struct {
	ent.Schema
}

// Fields of the CarBillingStrategy.
func (BillingStrategy) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Comment("ID"),
		field.String("name").Comment("策略名称"),
		field.Int("scenic_area_id").Comment("景区ID"),
		field.Int("model_id").Comment("型号ID"),
		field.Int("main_mode").Default(1).Comment("主要模式（1-按时间 2-按里程 3-按站点）"),
		field.Int("start_time_price").Default(0).Comment("起步价(按时)收费价格(单位：分)"),
		field.Int("start_time_unit").Default(1).Comment("起步价(按时)计量单位（单位：秒）"),
		field.Int("normal_time_price").Default(0).Comment("时间收费价格(单位：分)"),
		field.Int("normal_time_unit").Default(1).Comment("时间计量单位（单位：秒）"),
		field.Int("start_mileage_price").Default(0).Comment("起步价(按里程)收费价格(单位：分)"),
		field.Int("start_mileage_unit").Default(1).Comment("起步价(按里程)计量单位（单位：米）"),
		field.Int("normal_mileage_price").Default(0).Comment("里程收费价格(单位：分)"),
		field.Int("normal_mileage_unit").Default(1).Comment("里程计量单位（单位：米）"),
		field.Int("start_stop_price").Default(0).Comment("起步价(按站)收费价格(单位：分)"),
		field.Int("start_stop_unit").Default(1).Comment("起步价(按站)计量单位（单位：站）"),
		field.Int("normal_stop_price").Default(0).Comment("按站收费价格(单位：分)"),
		field.Int("normal_stop_unit").Default(1).Comment("按站计量单位（单位：站）"),
		field.Int("capped_amount").Default(0).Comment("封顶价格（单位：分）"),
		field.Int("deposit_amount").Default(0).Comment("押金金额（单位：分）"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the CarBillingStrategy.
func (BillingStrategy) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("background_scenic_area", ScenicArea.Type).
			Ref("car_billing_strategies").
			Unique().
			Field("scenic_area_id").
			Required(),
	}
}

func (BillingStrategy) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("scenic_area_id", "model_id").Unique(),
	}
}
