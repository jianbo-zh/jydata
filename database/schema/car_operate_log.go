package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CarsOperateLog holds the schema definition for the CarsOperateLog entity.
type CarsOperateLog struct {
	ent.Schema
}

// Fields of the CarsOperateLog.
func (CarsOperateLog) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Comment("ID"),
		field.Int("scenic_area_id").Default(1).Comment("景区ID"),
		field.String("origin").Comment("操作端(用户、运维、后台)"),
		field.Int("origin_uid").Comment("用户id"),
		field.String("origin_ip").Comment("用户ip"),
		field.String("operate_action").Comment("操作动作"),
		field.Int("car_id").Comment("车辆id"),
		field.String("device_id").Comment("设备id"),
		field.Float("car_longitude").Optional().Nillable().Comment("经度"),
		field.Float("car_latitude").Optional().Nillable().Comment("纬度"),
		field.Float("car_mileage").Optional().Nillable().Comment("当前里程"),
		field.Float("car_speed").Optional().Nillable().Comment("当前速度"),
		field.Float("car_power").Optional().Nillable().Comment("当前电量"),
		field.String("car_status").Optional().Nillable().Comment("车辆状态"),
		field.String("operate_result").SchemaType(map[string]string{dialect.MySQL: "varchar(2000)"}).Comment("操作结果"),
		field.Int("operate_state").Default(0).Comment("操作状态(1-失败、2-成功)"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

func (CarsOperateLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("background_scenic_area", ScenicArea.Type).
			Ref("car_operate_logs").
			Unique().
			Field("scenic_area_id").
			Required(),
		edge.From("car", Car.Type).
			Ref("car_operate_logs").
			Unique().
			Field("car_id").
			Required(),
	}
}
