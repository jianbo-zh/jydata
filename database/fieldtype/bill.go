package fieldtype

type BillTxType int

const (
	BillTxType_In  BillTxType = 1 // 收入
	BillTxType_Out BillTxType = 0 // 支出
)

type BillTxChannel int

const (
	BillTxChannel_Deposit BillTxChannel = 1 // 用车押金
	BillTxChannel_Refund  BillTxChannel = 3 // 用车退款
	BillTxChannel_Sharing BillTxChannel = 4 // 微信分账
)
