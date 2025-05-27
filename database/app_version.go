package database

type Platform string
type AppName string
type AppVersionState int

const (
	Platform_Android Platform = "android"
	Platform_IOS     Platform = "ios"
)

const (
	AppName_GolfApp AppName = "golf_app"
)

const (
	AppVersionState_PreRelease  AppVersionState = 1
	AppVersionState_Release     AppVersionState = 2
	AppVersionState_Termination AppVersionState = 3
)
