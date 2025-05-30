package fieldstate

type OrderAppealState int

// (1-申述中、2-申诉成功、3-申诉失败、4-已取消)
const (
	OrderAppealState_Pending  = 1
	OrderAppealState_Success  = 2
	OrderAppealState_Failed   = 3
	OrderAppealState_Canceled = 4
)
