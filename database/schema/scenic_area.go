package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ScenicArea holds the schema definition for the ScenicArea entity.
type ScenicArea struct {
	ent.Schema
}

// Fields of the ScenicArea.
func (ScenicArea) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().Comment("景区id"),
		field.String("code").Comment("景区编号"),
		field.String("name").Comment("景区名称"),
		field.String("address").Comment("景区地址"),
		field.String("manager").Comment("负责人"),
		field.String("phone").Comment("联系电话"),
		field.String("mch_id").Default("").Comment("商户号"),
		field.String("mch_name").Default("").Comment("商户名称"),
		field.Int("timezone").Default(8).Comment("时区"),
		field.Float("wgs_lon").Default(0).Comment("经度(地球坐标系)"),
		field.Float("wgs_lat").Default(0).Comment("纬度(地球坐标系)"),
		field.Float("gcj_lon").Default(0).Comment("经度(火星坐标系)"),
		field.Float("gcj_lat").Default(0).Comment("纬度(火星坐标系)"),
		field.Float("bd_lon").Default(0).Comment("经度(百度坐标系)"),
		field.Float("bd_lat").Default(0).Comment("纬度(百度坐标系)"),
		field.Int("status").Default(1).Comment("景区状态"),
		field.Int("extend_yokee_id").Optional().Nillable().Comment("Yokee扩展ID"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the ScenicArea.
// Edges of the BackgroundScenicArea.
func (ScenicArea) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cars", Car.Type),
		// edge.To("files", File.Type),
		edge.To("accounts", Account.Type).Annotations(entsql.OnDelete(entsql.Restrict)),
		edge.To("pois", Poi.Type),
		edge.To("profit_receivers", ProfitReceiver.Type),
		edge.To("pay_tx_bills", PayTxBill.Type),
		edge.To("car_billing_strategies", BillingStrategy.Type),
		edge.To("map", ScenicAreaMap.Type).Unique(),
		edge.To("map_versions", MapVersion.Type),
		edge.To("users", User.Type),
		edge.To("orders", Order.Type),
		edge.To("car_operate_logs", CarsOperateLog.Type),
		edge.To("stats_hourly_car", StatsHourlyCar.Type),
		edge.To("config_files", CarConfig.Type),
	}
}
