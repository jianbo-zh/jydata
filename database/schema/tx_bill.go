package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PayTxBill holds the schema definition for the PayTxBill entity.
type PayTxBill struct {
	ent.Schema
}

// Fields of the PayTxLog.
func (PayTxBill) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Comment("ID"),
		field.String("tx_id").Comment("支付流水号"),
		field.Int("order_id").Comment("订单ID"),
		field.String("order_no").Comment("订单号"),
		field.Int("scenic_area_id").Comment("景区ID"),
		field.String("scenic_area_name").Comment("景区名称"),
		field.Int("tx_type").Comment("类型（1-收入 0-支出）"),
		field.Int("tx_channel").Comment("渠道（1-用车押金、2-用车费用、3-用车退款 4-微信分账）"),
		field.String("tx_account").Comment("对方账号"),
		field.String("tx_account_name").Comment("对方姓名"),
		field.Int("tx_amount").Comment("交易金额(单位: 分)"),
		field.Time("tx_time").Comment("交易时间"),
		field.String("remark").Comment("备注"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the PayTxLog.
func (PayTxBill) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("background_scenic_area", ScenicArea.Type).
			Ref("pay_tx_bills").
			Unique().
			Field("scenic_area_id").
			Required(),
	}
}
