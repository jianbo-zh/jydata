package database

import (
	"context"

	"github.com/jianbo-zh/jydata/database/ent"
	"github.com/jianbo-zh/jydata/database/ent/operationuser"
)

func (db *Database) GetDevopsUser(ctx context.Context, id int) (*ent.OperationUser, error) {
	return db.MainDB().OperationUser.Get(ctx, id)
}

func (db *Database) GetDevopsUserByOpenID(ctx context.Context, openID string) (*ent.OperationUser, error) {
	return db.MainDB().OperationUser.Query().Where(operationuser.OpenIDEQ(openID)).First(ctx)
}

func (db *Database) GetDevopsUserByPhone(ctx context.Context, phone string) (*ent.OperationUser, error) {
	return db.MainDB().OperationUser.Query().Where(operationuser.PhoneEQ(phone)).First(ctx)
}

func (db *Database) UpdateDevopsUserOpenID(ctx context.Context, id int, openID string) (*ent.OperationUser, error) {
	return db.MainDB().OperationUser.UpdateOneID(id).SetOpenID(openID).Save(ctx)
}

func (db *Database) UpdateDevopsUserPassword(ctx context.Context, id int, password string) (*ent.OperationUser, error) {
	return db.MainDB().OperationUser.UpdateOneID(id).SetPassword(password).Save(ctx)
}
