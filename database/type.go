package database

type UserType int

func (u UserType) String() string {
	switch u {
	case UserType_User:
		return "user"
	case UserType_Admin:
		return "admin"
	case UserType_Devops:
		return "devops"
	}
	return "unknown"
}

const (
	UserType_User   UserType = 1 // 普通用户
	UserType_Admin  UserType = 2 // 管理员
	UserType_Devops UserType = 3 // 运维人员
)

type GeneralUser struct {
	Type UserType
	ID   int
}
