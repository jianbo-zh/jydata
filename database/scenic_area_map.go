package database

import (
	"context"

	"github.com/jianbo-zh/jydata/database/ent"
	"github.com/jianbo-zh/jydata/database/ent/scenicareamap"
)

func (m *Database) GetScenicAreaMap(ctx context.Context, scenicAreaID int) (*ent.ScenicAreaMap, error) {
	return m.MainDB().ScenicAreaMap.Query().
		Where(scenicareamap.ScenicAreaIDEQ(scenicAreaID)).
		First(ctx)
}
