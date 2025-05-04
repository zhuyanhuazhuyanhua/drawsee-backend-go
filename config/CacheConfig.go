package config

import (
	"context"
	"encoding/json"
	"time"

	"drawsee/constant"

	"github.com/go-redis/redis/v8"
)

// CacheManager 定义缓存管理器
type CacheManager struct {
	client *redis.Client
}

// NewCacheManager 创建新的缓存管理器
func NewCacheManager(redisConnectionFactory *redis.Client) *CacheManager {
	return &CacheManager{
		client: redisConnectionFactory,
	}
}

// Get 获取缓存值
func (cm *CacheManager) Get(ctx context.Context, key string, result interface{}) error {
	val, err := cm.client.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil
		}
		return err
	}
	return json.Unmarshal(val, result)
}

// Set 设置缓存值
func (cm *CacheManager) Set(ctx context.Context, key string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	var expiration time.Duration
	switch key {
	case constant.DASHBOARD_STATISTICS_KEY:
		expiration = time.Hour
	case constant.INVITATION_CODE_PAGE_KEY:
		expiration = 3 * 24 * time.Hour
	default:
		expiration = 24 * time.Hour
	}

	return cm.client.Set(ctx, key, data, expiration).Err()
}

// NewRedisClient 创建新的 Redis 客户端
func NewRedisClient(addr string, password string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
}
