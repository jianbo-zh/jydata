package database

import (
	"context"
	"fmt"

	"github.com/jianbo-zh/jydata/database/ent"
	"github.com/jianbo-zh/jydata/database/ent/mapversion"

	"entgo.io/ent/dialect/sql"
)

func (m *Database) CreateMapVersion(ctx context.Context, req *ent.MapVersion) (*ent.MapVersion, error) {
	return m.MainDB().MapVersion.Create().
		SetScenicAreaID(req.ScenicAreaID).
		SetVersion(req.Version).
		SetFileID(req.FileID).
		SetFilePath(req.FilePath).
		SetName(req.Name).
		SetRemark(req.Remark).
		Save(ctx)
}

func (m *Database) GetMapVersion(ctx context.Context, scenicAreaId int, version string) (*ent.MapVersion, error) {
	return m.MainDB().MapVersion.Query().
		Where(mapversion.ScenicAreaIDEQ(scenicAreaId)).
		Where(mapversion.VersionEQ(version)).
		First(ctx)
}

func (m *Database) GetMapVersionNewest(ctx context.Context, scenicAreaId int) (*ent.MapVersion, error) {
	return m.MainDB().MapVersion.Query().
		Where(mapversion.ScenicAreaIDEQ(scenicAreaId)).
		Order(mapversion.ByID(sql.OrderDesc())).
		First(ctx)
}

func (m *Database) GetMapVersions(ctx context.Context, scenicAreaIDs []int, name string, offset int, limit int) (int, []*ent.MapVersion, error) {

	query := m.MainDB().MapVersion.Query()
	if len(scenicAreaIDs) > 0 {
		query.Where(mapversion.ScenicAreaIDIn(scenicAreaIDs...))
	}

	if name != "" {
		query.Where(mapversion.NameContains(name))
	}

	count, err := query.Clone().Count(ctx)
	if err != nil {
		return 0, nil, fmt.Errorf("query.Count error: %w", err)
	}

	list, err := query.Order(mapversion.ByID(sql.OrderDesc())).Offset(offset).Limit(limit).All(ctx)
	if err != nil {
		return 0, nil, fmt.Errorf("query.All error: %w", err)
	}

	return count, list, nil
}
