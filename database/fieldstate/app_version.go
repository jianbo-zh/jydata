package fieldstate

type AppVersionState int

const (
	AppVersionState_PreRelease  AppVersionState = 1
	AppVersionState_Release     AppVersionState = 2
	AppVersionState_Termination AppVersionState = 3
)
