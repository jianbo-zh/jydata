//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jianbo-zh/jydata/migrate"

	confV1 "github.com/jianbo-zh/jypb/config/v1"
)

// wireApp init kratos application.
func wireApp(*confV1.Infra, log.Logger) (*migrate.MigrateServer, func(), error) {
	panic(wire.Build(migrate.ProviderSet))
}
