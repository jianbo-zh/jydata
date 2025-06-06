package schema

import (
	"time"

	"github.com/jianbo-zh/jydata/database/mixin"
	"github.com/jianbo-zh/jydata/database/schema/types"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Route 线路
type Route struct {
	ent.Schema
}

// Fields of the User.
func (Route) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Int("scenic_area_id").Comment("景区ID"),
		field.Int("type").Comment("线路类型（1-接驳线路）"),
		field.String("name").Comment("线路名称"),
		field.Ints("poi_ids").Comment("poi点列表"),
		field.JSON("routing_path", &types.RoutingPath{}).Optional().Comment("行驶线路"),
		field.String("remark").Comment("线路备注"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the User.
func (Route) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Mixin of the Pet.
func (Route) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.SoftDeleteMixin{},
	}
}
