package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type CarState struct {
	DrivingState     int       `json:"s,omitempty"`
	DrivingStateTime time.Time `json:"t,omitempty"`
}

/******************* 车辆状态统计 *******************/
func (cache *Cache) SetCarState(ctx context.Context, deviceID string, carState *CarState) error {
	bs, err := json.Marshal(carState)
	if err != nil {
		return fmt.Errorf("json.Marshal error: %w", err)
	}
	return cache.Client().HSet(ctx, "car_state", deviceID, string(bs)).Err()
}

func (cache *Cache) DelCarState(ctx context.Context, deviceID string) error {
	return cache.Client().HDel(ctx, "car_state", deviceID).Err()
}

func (cache *Cache) GetCarState(ctx context.Context, deviceID string) (*CarState, error) {
	bs, err := cache.Client().HGet(ctx, "car_state", deviceID).Bytes()
	if err != nil {
		return nil, fmt.Errorf("rdb.HGet error: %w", err)
	}

	var carState *CarState
	if err = json.Unmarshal(bs, &carState); err != nil {
		return nil, fmt.Errorf("json.Unmarshal error: %w", err)
	}

	return carState, nil
}

func (cache *Cache) GetAllCarState(ctx context.Context) (map[string]*CarState, error) {
	results, err := cache.Client().HGetAll(ctx, "car_state").Result()
	if err != nil {
		return nil, err
	}

	states := make(map[string]*CarState)
	for deviceId, val := range results {
		var carState *CarState
		if err = json.Unmarshal([]byte(val), &carState); err != nil {
			return nil, fmt.Errorf("json.Unmarshal error: %w", err)
		}
		states[deviceId] = carState
	}

	return states, nil
}

/******************* 车辆在线 *******************/

// GetCarAliveTime 获取车辆在线
func (cache *Cache) GetCarAliveTime(ctx context.Context, deviceID string) (time.Time, error) {
	aliveTime, err := cache.Client().Get(ctx, "car_alive_time:"+deviceID).Int64()
	if err != nil {
		return time.Time{}, err
	}
	return time.UnixMilli(aliveTime), nil
}

// SetCarAliveTime 设置车辆在线
func (cache *Cache) SetCarAliveTime(ctx context.Context, deviceID string, aliveTime time.Time) error {
	return cache.Client().Set(ctx, "car_alive_time:"+deviceID, aliveTime.UnixMilli(), 0).Err()
}

// SetCarConnected 设置车辆在线
func (cache *Cache) SetCarConnected(ctx context.Context, deviceID string) error {
	return cache.Client().HSet(ctx, "car_connects", deviceID, true).Err()
}

// SetCarDisconnected 设置车辆在线
func (cache *Cache) SetCarDisconnected(ctx context.Context, deviceID string) error {
	return cache.Client().HSet(ctx, "car_connects", deviceID, false).Err()
}

// GetCarConnected 获取车辆是否在线
func (cache *Cache) CheckCarConnect(ctx context.Context, deviceID string) (bool, error) {
	return cache.Client().HGet(ctx, "car_connects", deviceID).Bool()
}

// ClearCarConnected 设置车辆在线
func (cache *Cache) ClearCarConnect(ctx context.Context) error {
	return cache.Client().Del(ctx, "car_connects").Err()
}

/******************* 车辆数据 *******************/

// SetCarData 设置车辆数据
func (cache *Cache) GetSetCarData(ctx context.Context, deviceID string, key string, val []byte) ([]byte, error) {
	bs, err := cache.Client().HGet(ctx, "car_data:"+deviceID, key).Bytes()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, fmt.Errorf("client.HGet error: %w", err)
	}
	err = cache.Client().HSet(ctx, "car_data:"+deviceID, key, val).Err()
	if err != nil {
		return nil, fmt.Errorf("client.HSet error: %w", err)
	}
	return bs, nil
}

// GetCarData 获取车辆数据
func (cache *Cache) GetCarData(ctx context.Context, deviceID string, key string) ([]byte, error) {
	return cache.Client().HGet(ctx, "car_data:"+deviceID, key).Bytes()
}

// GetCarAllData 获取车辆所有数据
func (cache *Cache) GetCarDatas(ctx context.Context, deviceID string) (map[string][]byte, error) {
	result, err := cache.Client().HGetAll(ctx, "car_data:"+deviceID).Result()
	if err != nil {
		return nil, fmt.Errorf("rdb.HGetAll error: %w", err)
	}

	data := make(map[string][]byte)
	for key, val := range result {
		data[key] = []byte(val)
	}

	return data, nil
}

// DelCarAllData 删除车辆数据
func (cache *Cache) ClearCarDatas(ctx context.Context, deviceID string) error {
	return cache.Client().Del(ctx, "car_data:"+deviceID).Err()
}

/******************* 车辆播报 *******************/
type PoiDis struct {
	ID  int     `json:"id,omitempty"`
	Dis float32 `json:"dis,omitempty"`
}

func (cache *Cache) AddNearbyPOI(ctx context.Context, deviceID string, pois []PoiDis) error {
	bs, _ := json.Marshal(pois)
	return cache.Client().SetEX(ctx, "nearby_pois:"+deviceID, string(bs), time.Minute).Err()
}

func (cache *Cache) GetNearbyPOIs(ctx context.Context, deviceID string) ([]PoiDis, error) {
	var pois []PoiDis
	bs, err := cache.Client().Get(ctx, "nearby_pois:"+deviceID).Bytes()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bs, &pois)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal error: %w", err)
	}

	return pois, nil
}

/******************* 车辆路由规划 *******************/

func (cache *Cache) PublishRoutePath(ctx context.Context, deviceID string, routePath string) error {
	return cache.Client().Publish(ctx, "routepath_"+deviceID, routePath).Err()
}

func (cache *Cache) SubscribeRoutePath(ctx context.Context, deviceID string, timeout time.Duration) (string, error) {
	pubsub := cache.Client().Subscribe(ctx, "routepath_"+deviceID)
	defer pubsub.Close()

	cancelCtx, cancelFn := context.WithTimeout(ctx, timeout)
	defer cancelFn()

	msg, err := pubsub.ReceiveMessage(cancelCtx)
	if err != nil {
		return "", fmt.Errorf("pubsub.ReceiveMessage error: %w", err)
	}

	return msg.Payload, nil
}

/******************* 车辆座位表 *******************/

func (cache *Cache) UpdateCarSeatTable(ctx context.Context, deviceID string, seatTable []bool) error {
	return cache.Client().Set(ctx, "car_seat_table:"+deviceID, seatTable, 0).Err()
}

func (cache *Cache) GetCarSeatTable(ctx context.Context, deviceID string) ([]bool, error) {
	var seatTable []bool
	err := cache.Client().Get(ctx, "car_seat_table:"+deviceID).Scan(&seatTable)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, fmt.Errorf("client.HGet error: %w", err)
	}
	return seatTable, nil
}
