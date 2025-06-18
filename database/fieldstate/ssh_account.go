package fieldstate

type SshAccountState int

// 1-未使用 2-使用中
const (
	SshAccountState_Unuse SshAccountState = 1
	SshAccountState_Using SshAccountState = 2
)
