package main

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/encoding/json"
	"github.com/google/uuid"

	_ "go.uber.org/automaxprocs"

	"github.com/jianbo-zh/jylib/bootstrap"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	Version = "1.0.0"
)

func init() {
	bootstrap.RegisterServiceInfo(
		uuid.NewString(),
		"jydata.migrate",
		Version,
		map[string]string{},
	)

	// UseEnumNumbers 枚举类型返回数字
	json.MarshalOptions.UseEnumNumbers = true
}

func main() {
	ctx := context.Background()
	config := bootstrap.LoadConfig()
	logger := bootstrap.NewLoggerProvider(bootstrap.Service, config.Log)

	migrater, cleanup, err := wireApp(config.Infra, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := migrater.Run(ctx); err != nil {
		panic(err)
	}
	fmt.Println("migrate success")
}
