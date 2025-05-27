package database

import (
	"context"

	"github.com/jianbo-zh/jydata/database/ent"
	"github.com/jianbo-zh/jydata/database/ent/profitreceiver"
	"github.com/jianbo-zh/jydata/database/fieldstate"
)

func (m *Database) GetEnableProfitReceivers(ctx context.Context, scenicAreaID int) ([]*ent.ProfitReceiver, error) {
	return m.MainDB().ProfitReceiver.Query().
		Where(profitreceiver.ScenicAreaIDEQ(scenicAreaID)).
		Where(profitreceiver.StateEQ(int(fieldstate.ProfitState_Enable))).
		All(ctx)
}
