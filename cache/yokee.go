package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type YokeeOnlineCars struct {
	DeviceIds  []string
	UpdateTime time.Time
}

/******************* Yokee车辆在线记录 *******************/
func (cache *Cache) SetYokeeOnlineCars(ctx context.Context, data *YokeeOnlineCars) error {
	bs, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("json.Marshal error: %w", err)
	}
	return cache.Client().Set(ctx, "yokee_online_cars", string(bs), 0).Err()
}

func (cache *Cache) GetYokeeOnlineCars(ctx context.Context) (*YokeeOnlineCars, error) {
	bs, err := cache.Client().Get(ctx, "yokee_online_cars").Bytes()
	if err != nil {
		return nil, fmt.Errorf("rdb.HGet error: %w", err)
	}

	var data YokeeOnlineCars
	if err = json.Unmarshal(bs, &data); err != nil {
		return nil, fmt.Errorf("json.Unmarshal error: %w", err)
	}

	return &data, nil
}

/******************* Yokee车辆详情记录 *******************/
type Goal struct {
	StopId    int
	StopName  string
	GoalIndex int
}
type YokeeCarDetail struct {
	Id                 int
	Name               string
	StationId          int
	LonWgs84           float64
	LatWgs84           float64
	LonGcj02           float64
	LatGcj02           float64
	Online             bool
	BatteryCharging    bool
	BatteryPower       float64
	BusinessStatus     string
	BusinessStatusName string
	DispatchId         int64
	CurrentGoalIndex   int32
	Goals              []Goal
	NoLoadEndurance    float64
	FullLoadEndurance  float64
	MaxSpeed           float64
	UpdateTime         time.Time
}

func (cache *Cache) SetYokeeCarDetail(ctx context.Context, deviceID string, carDetail *YokeeCarDetail) error {
	bs, err := json.Marshal(carDetail)
	if err != nil {
		return fmt.Errorf("json.Marshal error: %w", err)
	}
	return cache.Client().HSet(ctx, "yokee_cars_detail", deviceID, string(bs)).Err()
}

func (cache *Cache) GetYokeeCarDetail(ctx context.Context, deviceID string) (*YokeeCarDetail, error) {
	bs, err := cache.Client().HGet(ctx, "yokee_cars_detail", deviceID).Bytes()
	if err != nil {
		return nil, fmt.Errorf("rdb.HGet error: %w", err)
	}

	var carDetail *YokeeCarDetail
	if err = json.Unmarshal(bs, &carDetail); err != nil {
		return nil, fmt.Errorf("json.Unmarshal error: %w", err)
	}

	return carDetail, nil
}
