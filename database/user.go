package database

import (
	"context"

	"github.com/jianbo-zh/jydata/database/ent"
)

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

func (db *Database) GetUser(ctx context.Context, id int) (*ent.User, error) {
	return db.MainDB().User.Get(ctx, id)
}
