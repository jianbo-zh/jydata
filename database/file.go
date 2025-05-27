package database

import (
	"context"

	"github.com/jianbo-zh/jydata/database/ent"
	"github.com/jianbo-zh/jydata/database/ent/file"
)

type FileType int     // 文件类型
type FileCategory int // 文件分类

const (
	FileType_Other    FileType = 0 // 其他
	FileType_Image    FileType = 1 // 图片
	FileType_Audio    FileType = 2 // 音频
	FileType_Video    FileType = 3 // 视频
	FileType_Protobuf FileType = 4 // Protobuf文件
	FileType_Bin      FileType = 5 // Bin文件
	FileType_Apk      FileType = 6 // Apk文件
	FileType_Conf     FileType = 6 // 配置文件
)

const (
	FileCategory_Other       FileCategory = 0 // 其他
	FileCategory_POI         FileCategory = 1 // poi点
	FileCategory_ParkingSpot FileCategory = 2 // 停车点
	FileCategory_Map         FileCategory = 3 // 地图文件
	FileCategory_Avatar      FileCategory = 4 // 头像文件
	FileCategory_Qrcode      FileCategory = 5 // 二维码图片
	FileCategory_BGM         FileCategory = 6 // 背景音乐
	FileCategory_App         FileCategory = 7 // App安装包
	FileCategory_Conf        FileCategory = 8 // 配置文件
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
