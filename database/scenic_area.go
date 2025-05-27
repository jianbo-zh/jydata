package database

import (
	"context"

	"github.com/jianbo-zh/jydata/database/ent"
)

func (m *Database) GetScenicArea(ctx context.Context, id int) (*ent.ScenicArea, error) {
	return m.MainDB().ScenicArea.Get(ctx, id)
}
