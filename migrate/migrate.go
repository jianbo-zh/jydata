package migrate

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/jianbo-zh/jydata/cache"
	"github.com/jianbo-zh/jydata/database"
	"github.com/jianbo-zh/jydata/influxdb"
)

var ProviderSet = wire.NewSet(
	database.NewDatabase,
	cache.NewRedis,
	influxdb.NewInfluxDB,
	NewMigrateServer,
)

type MigrateServer struct {
	db       *database.Database
	cache    *cache.Cache
	influxdb *influxdb.InfluxDB

	log *log.Helper
}

func NewMigrateServer(db *database.Database, cache *cache.Cache, influxdb *influxdb.InfluxDB, logger log.Logger) *MigrateServer {
	return &MigrateServer{
		db:       db,
		cache:    cache,
		influxdb: influxdb,

		log: log.NewHelper(logger),
	}
}

func (c *MigrateServer) Run(ctx context.Context) error {
	err := c.db.MainDB().Schema.Create(ctx)
	if err != nil {
		c.log.Errorf("failed to create database schema: %v", err)
		return fmt.Errorf("failed to create database schema: %w", err)
	}
	return nil
}
