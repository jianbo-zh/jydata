package database

import (
	"context"
	"fmt"

	"github.com/jianbo-zh/jydata/database/ent"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"

	_ "github.com/jianbo-zh/jydata/database/ent/runtime"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	configV1 "github.com/jianbo-zh/jypb/config/v1"
)

//go:generate go run entgo.io/ent/cmd/ent generate --target ./ent --feature intercept,schema/snapshot --template ./tmpl ./schema

type Database struct {
	maindb   *ent.Client
	maindbTx *ent.Tx
}

func (m *Database) MainDB() *ent.Client {
	if m.maindbTx != nil {
		return m.maindbTx.Client()
	}
	return m.maindb
}

func (m *Database) Begin() (*Database, error) {
	return m.BeginWithContext(context.Background())
}

func (m *Database) BeginWithContext(ctx context.Context) (*Database, error) {
	tx, err := m.maindb.Tx(ctx)
	if err != nil {
		return nil, err
	}
	return &Database{
		maindb:   m.maindb,
		maindbTx: tx,
	}, nil
}

func (m *Database) Rollback() error {
	if m.maindbTx != nil {
		return m.maindbTx.Rollback()
	}
	return fmt.Errorf("not tx context")
}

func (m *Database) Commit() error {
	if m.maindbTx != nil {
		return m.maindbTx.Commit()
	}
	return fmt.Errorf("not tx context")
}

func NewDatabase(conf *configV1.Infra, logger log.Logger) (*Database, func(), error) {
	if conf.Database == nil {
		panic("config infra.database not found")
	}

	log := log.NewHelper(logger)
	drv, err := sql.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	sqlDrv := dialect.DebugWithContext(drv, func(ctx context.Context, i ...interface{}) {
		// log.WithContext(ctx).Debug(i...)
	})
	client := ent.NewClient(ent.Driver(sqlDrv))
	if err != nil {
		log.Errorf("failed opening connection to sqlite: %v", err)
		return nil, nil, err
	}

	return &Database{
			maindb: client,
		}, func() {
			if err := client.Close(); err != nil {
				log.Error(err)
			}
		}, nil
}
