// ent/schema/order.go

package schema

import (
	"time"

	"github.com/jianbo-zh/jydata/database/schema/types"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// 订单分润
type OrderSharing struct {
	ent.Schema
}

func (OrderSharing) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Comment("ID"),
		field.Int("order_id").Comment("订单ID"),
		field.String("sharing_no").Comment("分润单号"),
		field.String("wx_sharing_id").Default("").Comment("分润单号"),
		field.Int("sharing_amount").Default(0).Comment("分润金额"),
		field.JSON("receivers", []types.OrderSharingReceiver{}).Optional().Comment("分润接收者"),
		field.Int("state").Default(1).Comment("分润状态(1-待分润、2-已完成)"),
		field.String("remark").Default("").Comment("备注"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

func (OrderSharing) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).Ref("sharing").Unique().Field("order_id").Required(),
	}
}
