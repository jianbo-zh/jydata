package database

// 配置内容类型
type CarConfigContentType int

const (
	CarConfigContentType_Base       CarConfigContentType = 1 // 基础配置
	CarConfigContentType_ScenicArea CarConfigContentType = 2 // 景区配置
	CarConfigContentType_HDMap      CarConfigContentType = 3 // 高精度地图配置
	CarConfigContentType_POI        CarConfigContentType = 4 // Poi配置
	CarConfigContentType_CarUI      CarConfigContentType = 5 // CarUI配置
)

// 配置内容字段（使用的）
type CarConfigContentField int

const (
	CarConfigContentField_FileID CarConfigContentField = 1 // FileID字段
	CarConfigContentField_PbText CarConfigContentField = 2 // PbText字段
)
