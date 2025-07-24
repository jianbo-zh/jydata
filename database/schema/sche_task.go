// ent/schema/car.go

package schema

import (
	"time"

	"github.com/jianbo-zh/jydata/database/schema/types"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Car holds the schema definition for the Car entity.
type ScheTask struct {
	ent.Schema
}

// Fields of the Car.
func (ScheTask) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Int("user_origin").Default(5).Comment("5-userapp 6-devopsapplet"),
		field.Int("user_type").Default(1).Comment("用户类型(1-admin 2-user 3-devops)"),
		field.Int("user_id").Default(0).Comment("用户ID"),
		field.Int("scenic_area_id").Comment("景区ID"),
		field.Int("car_id").Comment("车辆ID"),
		field.String("device_id").Comment("设备ID"),
		field.Int("dest_id").Comment("地点ID"),
		field.Float("dest_lon").Comment("经度(wgs84)"),
		field.Float("dest_lat").Comment("纬度(wgs84)"),
		field.Int("sche_mode").Default(1).Comment("调度模式（1-[用户]自由调度 2-[运营]揽客模式 3-[运营]部署模式）"),
		field.JSON("sche_args", types.ScheArgs{}).Optional().Comment("调度参数"),
		field.Int("state").Comment("当前状态（1-初始化 2-调度中 3-已暂停 5-停滞不前 6-已完成 7-已取消 8-系统终止 9-异常状态）"),
		field.Int("abnormal_state").Default(0).Comment("异常子状态（1-无异常 2-离线 3-道路外 4-有载人或物 5-驾驶状态异常 6-紧急制动中）"),
		field.String("remark").Default("").Comment("备注"),
		field.JSON("routing_path", types.RoutingPath{}).Comment("调度路径"),
		field.Time("restart_sche_time").Optional().Default(time.Time{}).Comment("调度临停后多久重启调度时间"),
		field.Time("end_time").Optional().Comment("结束时间"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the Car.
func (ScheTask) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("car", Car.Type).
			Ref("sche_task").
			Unique().
			Field("car_id").
			Required(),
		edge.To("events", ScheTaskEvent.Type),
	}
}
