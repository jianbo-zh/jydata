package database

type AppPushState int

const (
	AppPushState_Pending = 1
	AppPushState_Success = 2
	AppPushState_Failed  = 3
)
