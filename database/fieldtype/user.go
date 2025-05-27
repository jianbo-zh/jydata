package fieldtype

// 用户来源
type UserOrigin int

const (
	UserOrigin_WxApplet UserOrigin = 1 // 微信用户
	UserOrigin_App      UserOrigin = 2 // App用户
)

// 用户分类
type UserCls int

const (
	UserCls_Normal UserCls = 1 // 普通用户
	UserCls_Tester UserCls = 2 // 测试用户
)

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
