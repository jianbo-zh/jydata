package cache

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"

	configV1 "github.com/jianbo-zh/jypb/config/v1"
)

type Cache struct {
	client *redis.Client
}

func (r *Cache) Client() *redis.Client {
	return r.client
}

func NewRedis(conf *configV1.Infra, logger log.Logger) (*Cache, func(), error) {
	if conf.Redis == nil {
		panic("config infra.redis not found")
	}
	// redis
	client := redis.NewClient(&redis.Options{
		Network:      conf.Redis.Network,
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		DB:           int(conf.Redis.Db),
		DialTimeout:  conf.Redis.DialTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
	})
	client.AddHook(redisotel.TracingHook{})

	return &Cache{
			client: client,
		}, func() {
			if err := client.Close(); err != nil {
				log.Error(err)
			}
		}, nil
}
