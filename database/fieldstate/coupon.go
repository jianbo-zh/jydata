package fieldstate

type CouponState int

// 1-未使用、2-已使用、3-已过期
const (
	CouponState_Pending = 1
	CouponState_Used    = 2
	CouponState_Expired = 3
)
