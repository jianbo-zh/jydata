package database

import (
	"context"

	"github.com/jianbo-zh/jydata/database/ent"
	"github.com/jianbo-zh/jydata/database/ent/billingstrategy"
)

func (db *Database) GetChargeStrategyByModelID(ctx context.Context, scenicAreaID int, modelID int) (*ent.BillingStrategy, error) {
	return db.MainDB().BillingStrategy.Query().
		Where(billingstrategy.ScenicAreaIDEQ(scenicAreaID)).
		Where(billingstrategy.ModelIDEQ(modelID)).
		Only(ctx)
}
