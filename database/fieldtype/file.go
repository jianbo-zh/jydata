package fieldtype

type StorageType int

const (
	StorageType_Local StorageType = 1 // 本地文件
	StorageType_OSS   StorageType = 2 // 阿里云OSS
)

type FileType int // 文件类型

const (
	FileType_Other    FileType = 0 // 其他
	FileType_Image    FileType = 1 // 图片
	FileType_Audio    FileType = 2 // 音频
	FileType_Video    FileType = 3 // 视频
	FileType_Protobuf FileType = 4 // Protobuf文件
	FileType_Bin      FileType = 5 // Bin文件
	FileType_Apk      FileType = 6 // Apk文件
	FileType_Conf     FileType = 7 // 配置文件
)

func (f FileType) String() string {
	switch f {
	case FileType_Other:
		return "Other"
	case FileType_Image:
		return "Image"
	case FileType_Audio:
		return "Audio"
	case FileType_Video:
		return "Video"
	case FileType_Protobuf:
		return "Protobuf"
	case FileType_Bin:
		return "Bin"
	case FileType_Apk:
		return "Apk"
	case FileType_Conf:
		return "Conf"
	default:
		return "Unknown"
	}
}

type FileCategory int // 文件分类

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

func (f FileCategory) String() string {
	switch f {
	case FileCategory_Other:
		return "Other"
	case FileCategory_POI:
		return "POI"
	case FileCategory_ParkingSpot:
		return "ParkingSpot"
	case FileCategory_Map:
		return "Map"
	case FileCategory_Avatar:
		return "Avatar"
	case FileCategory_Qrcode:
		return "Qrcode"
	case FileCategory_BGM:
		return "BGM"
	case FileCategory_App:
		return "App"
	case FileCategory_Conf:
		return "Conf"
	default:
		return "Unknown"
	}
}
