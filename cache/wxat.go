package cache

import (
	"context"
	"time"
)

// GetWxAccessToken 获取微信AccessToken
func (cache *Cache) GetWxAccessToken(ctx context.Context, appid string) (string, error) {
	return cache.Client().Get(ctx, "wxaccesstoken_"+appid).Result()
}

// CacheWxAccessToken 缓存微信AccessToken
func (cache *Cache) CacheWxAccessToken(ctx context.Context, appid string, accessToken string, expiresIn time.Duration) error {
	return cache.Client().SetEX(ctx, "wxaccesstoken_"+appid, accessToken, expiresIn).Err()
}
