package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProfitReceiver holds the schema definition for the ProfitReceiver entity.
type ProfitReceiver struct {
	ent.Schema
}

// Fields of the ProfitShareReceiver.
func (ProfitReceiver) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Comment("ID"),
		field.Int("scenic_area_id").Comment("景区ID"),
		field.String("mch_id").Default("").Comment("商户号"),
		field.String("phone").Comment("手机号"),
		field.Int("receiver_type").Comment("接收方类型(1-商户ID,2-个人openid)"),
		field.String("receiver_account").Comment("接收方账号"),
		field.String("receiver_name").Comment("接收方姓名"),
		field.Float("sharing_ratio").Default(0).Comment("分润比例"),
		field.Int("state").Default(1).Comment("状态(1-启用 2-禁用)"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the ProfitShareReceiver.
func (ProfitReceiver) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("background_scenic_area", ScenicArea.Type).
			Ref("profit_receivers").
			Unique().
			Field("scenic_area_id").
			Required(),
	}
}
