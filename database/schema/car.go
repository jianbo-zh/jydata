// ent/schema/car.go

package schema

import (
	"time"

	"github.com/jianbo-zh/jydata/database/mixin"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Car holds the schema definition for the Car entity.
type Car struct {
	ent.Schema
}

// Fields of the Car.
func (Car) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		// field.Int("type").Default(1).Comment("车辆类型（1-共享型 2-运力型）"),
		field.Int("operation_mode").Default(1).Comment("运营模式（1-租车模式 2-公交模式）"),
		field.String("car_name").Comment("车辆名称"),
		field.Int("scenic_area_id").Comment("景区ID"),
		field.String("device_id").Unique().Comment("设备ID"),
		field.Int("model_id").Comment("车辆型号ID"),
		field.Ints("images").Comment("车辆图片URL"),
		field.String("license_plate").Comment("车牌号"),
		field.Int("passengers").Comment("可乘坐人数"),
		field.Int("reserved_seats").Default(0).Comment("保留座位数"),
		field.Time("produce_time").Comment("生产日期"),
		field.Int("power_threshold").Comment("电量阈值"),
		field.String("activate_code").Comment("激活码"),

		field.Int("state").Default(1).Comment("1-未激活、3-运营中"),
		field.Int("use_state").Default(1).Comment("使用状态：1-空闲中、2-预约中、3-租用中、6-调度中"),
		field.Int("driving_state").Default(0).Comment("驾驶状态：见proto"),
		field.Int("emergency_state").Default(0).Comment("紧急状态(0-无紧急、1-紧急呼救)"),

		field.Int("use_order_id").Default(0).Comment("绑定订单ID"),
		field.Int("dispatch_task_id").Default(0).Comment("调度任务ID"),
		field.Int("use_flight_id").Default(0).Comment("班次任务ID"),

		field.Int("bind_order_count").Default(0).Comment("绑定订单数"),
		field.Int("total_order_mileage").Default(0).Comment("累积订单里程（米）"),
		field.Int("total_order_time").Default(0).Comment("累积订单使用时长（秒）"),
		field.Int("total_order_count").Default(0).Comment("累积订单数"),
		field.Int("total_order_amount").Default(0).Comment("累积订单金额"),

		field.Int("power_remaining").Default(0).Comment("剩余电量"),
		field.Int("error_count").Default(0).Comment("异常数量"),
		field.String("error_message").Default("").Comment("当前异常"),

		field.Int("is_deleted").Default(-1).Comment("是否删除（1-已删除、-1-未删除）"),
		field.Int("is_commercial_car").Default(1).Comment("是否商用车辆(收费)"),
		field.Int("is_driving_state_valid").Default(0).Comment("是否驾驶状态可用"),

		field.String("map_version").Default("").Comment("地图版本"),

		field.String("next_map_version").Default("").Comment("当前地图版本类型"),
		field.String("next_map_version_state").Default("").Comment("地图版本类型"),
		field.Int("next_map_version_process").Default(0).Comment("地图版本进度"),

		field.String("gr_auto_version").Default("").Comment("车机版本"),
		field.String("gr_ui_version").Default("").Comment("CarUI版本"),

		field.Int("extend_yokee_id").Optional().Nillable().Comment("Yokee扩展ID"),

		field.Time("alive_time").Optional().Comment("心跳时间"),
		field.Time("register_time").Optional().Comment("激活时间"),
		field.Time("driving_state_time").Optional().Comment("驾驶状态更改时间"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the Car.
func (Car) Edges() []ent.Edge {
	return []ent.Edge{ // 设置关联关系
		edge.From("background_scenic_area", ScenicArea.Type).
			Ref("cars").
			Unique().
			Field("scenic_area_id"). // 通过class_id字段关联class表
			Required(),
		edge.From("cars_models", CarsModels.Type).
			Ref("cars").
			Unique().
			Field("model_id"). // 通过class_id字段关联class表
			Required(),
		// 操作日志
		edge.To("car_operate_logs", CarsOperateLog.Type),
		// 订单
		edge.To("orders", Order.Type),
		// 车辆小时统计
		edge.To("stats_hourly_car", StatsHourlyCar.Type),
		// 调度任务
		edge.To("sche_task", ScheTask.Type),
		// 关联的配置文件
		edge.To("config_files", CarConfig.Type),
		// 车辆绑定的配置
		edge.To("car_configs", CarConfigDownload.Type),
	}
}

// Mixin of the Pet.
func (Car) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.SoftDeleteMixin{},
	}
}
