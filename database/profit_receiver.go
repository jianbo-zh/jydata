package database

import (
	"context"

	"github.com/jianbo-zh/jydata/database/ent"
	"github.com/jianbo-zh/jydata/database/ent/profitreceiver"
)

// 分润状态
type ProfitState int

// 接收者类型
type ReceiverType int

const (
	ProfitState_Enable  ProfitState = 1 // 启用
	ProfitState_Disable ProfitState = 2 // 禁用
)

const (
	ReceiverType_Merchant ReceiverType = 1 // 商户
	ReceiverType_Personal ReceiverType = 2 // 个人
)

func (m *Database) GetEnableProfitReceivers(ctx context.Context, scenicAreaID int) ([]*ent.ProfitReceiver, error) {
	return m.MainDB().ProfitReceiver.Query().
		Where(profitreceiver.ScenicAreaIDEQ(scenicAreaID)).
		Where(profitreceiver.StateEQ(int(ProfitState_Enable))).
		All(ctx)
}
