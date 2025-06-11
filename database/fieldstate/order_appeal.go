package fieldstate

type OrderAppealState int

// (1-待审核、2-申诉成功、3-申诉失败、4-已取消、5-待退款)
const (
	OrderAppealState_PendingReview = 1
	OrderAppealState_Success       = 2
	OrderAppealState_Failed        = 3
	OrderAppealState_Canceled      = 4
	OrderAppealState_PendingRefund = 5
)
