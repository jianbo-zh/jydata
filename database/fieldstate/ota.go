package fieldstate

// OTA版本状态（整体部署状态）
type OtaVersionState int

const (
	OtaVersionState_Pending        OtaVersionState = 1 // 待升级
	OtaVersionState_InProgress     OtaVersionState = 2 // 升级中
	OtaVersionState_PartialFailure OtaVersionState = 3 // 部分失败
	OtaVersionState_Success        OtaVersionState = 4 // 升级成功
	OtaVersionState_Failure        OtaVersionState = 5 // 升级失败
)

// OTA升级状态（单车部署状态）
type OtaDeployState int

const (
	OtaDeployState_Pending        OtaDeployState = 1 // 待升级
	OtaDeployState_InProgress     OtaDeployState = 2 // 升级中
	OtaDeployState_Success        OtaDeployState = 3 // 升级成功
	OtaDeployState_Failure        OtaDeployState = 4 // 升级失败
	OtaDeployState_PartialFailure OtaDeployState = 5 // 部分升级失败
)
