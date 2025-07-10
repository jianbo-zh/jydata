package schema

import (
	"time"

	"github.com/jianbo-zh/jydata/database/schema/types"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// CarsFlight 车辆发车班次
type CarsFlight struct {
	ent.Schema
}

// Fields of the CarsFlight.
func (CarsFlight) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Comment("ID"),
		field.Int("scenic_area_id").Comment("景区ID"),
		field.String("flight_no").Comment("班次编号"),
		field.Int("car_id").Comment("车辆ID"),
		field.String("device_id").Default("").Comment("车辆编号"),
		field.String("car_name").Comment("车辆名称"),
		field.Int("route_id").Comment("路线ID"),
		field.String("route_name").Comment("路线名称"),
		field.Int("seats_num").Comment("可售卖座位数"),
		field.Int("state").Default(1).Comment("班次状态(1-待发车 2-路线规划中 3-行驶中 4-到停靠点 5-到目的地 6-已取消)"),
		field.Int("curr_stop_id").Optional().Default(0).Comment("当前或下一个停靠点ID"),
		field.Int("curr_stop_index").Optional().Default(0).Comment("当前或下一个停靠点序号"),
		field.Ints("stop_ids").Optional().Default([]int{}).Comment("站点列表"),
		field.Ints("pass_ids").Optional().Default([]int{}).Comment("过站列表"),
		field.String("remark").Default("").Comment("备注"),
		field.JSON("stop_stock", []types.StopStock{}).Comment("座位库存数量"),
		field.Int("extend_yokee_id").Optional().Nillable().Comment("Yokee扩展ID"),
		field.Time("departure_time").Optional().Nillable().Comment("发车时间"),
		field.Time("finish_time").Optional().Nillable().Comment("完成时间"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

func (CarsFlight) Edges() []ent.Edge {
	return []ent.Edge{}
}
