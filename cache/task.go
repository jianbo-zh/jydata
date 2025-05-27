package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

func (cache *Cache) AddTaskQueue(ctx context.Context, taskID int, ts time.Time) error {
	return cache.Client().ZAdd(ctx, "taskjob", &redis.Z{
		Score:  float64(ts.UnixMilli()),
		Member: taskID,
	}).Err()
}
