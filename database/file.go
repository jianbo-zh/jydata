package database

import (
	"context"

	"github.com/jianbo-zh/jydata/database/ent"
	"github.com/jianbo-zh/jydata/database/ent/file"
)

// CreateFile 添加文件
func (db *Database) CreateFile(ctx context.Context, file *ent.File) (*ent.File, error) {
	return db.MainDB().File.Create().
		SetCreatorID(file.CreatorID).
		SetScenicAreaID(file.ScenicAreaID).
		SetFileType(file.FileType).
		SetFileCategory(file.FileCategory).
		SetFilePath(file.FilePath).
		SetFileSha1(file.FileSha1).
		SetFileSuffix(file.FileSuffix).
		SetMimeType(file.MimeType).
		Save(ctx)
}

func (db *Database) GetFileByID(ctx context.Context, id int) (*ent.File, error) {
	return db.MainDB().File.Get(ctx, id)
}

func (db *Database) GetFilesByIDs(ctx context.Context, ids []int) ([]*ent.File, error) {
	return db.MainDB().File.Query().Where(file.IDIn(ids...)).All(ctx)
}
