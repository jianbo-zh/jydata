package database

import (
	"context"

	"github.com/jianbo-zh/jydata/database/ent"
)

func (db *Database) GetUser(ctx context.Context, id int) (*ent.User, error) {
	return db.MainDB().User.Get(ctx, id)
}
