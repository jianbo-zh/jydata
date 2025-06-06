package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/protobuf/proto"

	zzzV1 "github.com/jianbo-zh/jypb/api/zzz/v1"
)

func (cache *Cache) CreateFlightTrace(ctx context.Context, flightId int, coord *zzzV1.Coord) error {
	bs, err := proto.Marshal(coord)
	if err != nil {
		return fmt.Errorf("proto.Marshal error: %w", err)
	}
	if err := cache.Client().RPush(ctx, fmt.Sprintf("flight_trace_%d", flightId), bs).Err(); err != nil {
		return fmt.Errorf("redis.RPush error: %w", err)
	}
	if err := cache.Client().Expire(ctx, fmt.Sprintf("flight_trace_%d", flightId), 12*time.Hour).Err(); err != nil {
		return fmt.Errorf("redis.Expire error: %w", err)
	}
	return nil
}

func (cache *Cache) AppendFlightTrace(ctx context.Context, flightId int, coord *zzzV1.Coord) error {
	bs, err := proto.Marshal(coord)
	if err != nil {
		return fmt.Errorf("proto.Marshal error: %w", err)
	}
	return cache.Client().RPush(ctx, fmt.Sprintf("flight_trace_%d", flightId), bs).Err()
}

func (cache *Cache) GetFlightTraces(ctx context.Context, flightId int) ([]*zzzV1.Coord, error) {
	var bbs [][]byte
	err := cache.Client().LRange(ctx, fmt.Sprintf("flight_trace_%d", flightId), 0, -1).ScanSlice(&bbs)
	if err != nil {
		return nil, fmt.Errorf("redis.ScanSlice error: %w", err)
	}

	var coords []*zzzV1.Coord
	for _, bs := range bbs {
		var item zzzV1.Coord
		err := proto.Unmarshal(bs, &item)
		if err != nil {
			return nil, fmt.Errorf("proto.Unmarshal error: %w", err)
		}
		coords = append(coords, &item)
	}

	return coords, nil
}

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
