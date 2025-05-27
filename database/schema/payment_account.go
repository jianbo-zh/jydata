package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ScenicArea holds the schema definition for the ScenicArea entity.
type PaymentAccount struct {
	ent.Schema
}

// Fields of the ScenicArea.
func (PaymentAccount) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().Comment("id"),
		field.String("mch_name").Comment("商户名称"),
		field.String("mch_id").Comment("商户ID"),
		field.String("mch_cert_sn").Comment("商户证书序号"),
		field.String("mch_apiv3_key").Comment("商户APIV3密钥"),
		field.String("mch_private_key").Comment("商户证书私钥"),
		field.Time("create_time").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}

// Edges of the ScenicArea.
// Edges of the BackgroundScenicArea.
func (PaymentAccount) Edges() []ent.Edge {
	return []ent.Edge{}
}
