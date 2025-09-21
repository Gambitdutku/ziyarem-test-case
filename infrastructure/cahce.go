package infrastructure

import (
	"context"
	"encoding/json"
	"time"
	"ziyaremtestcase/domain"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	client *redis.Client
	prefix string
}

func NewRedisCache() *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return &RedisCache{client: rdb, prefix: "sensor"}
}

func (r *RedisCache) Get(ctx context.Context, key string) (*domain.SensorData, error) {
	val, err := r.client.Get(ctx, r.prefix+":"+key).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var sd domain.SensorData
	if err := json.Unmarshal([]byte(val), &sd); err != nil {
		return nil, err
	}
	return &sd, nil
}

func (r *RedisCache) Set(ctx context.Context, key string, data *domain.SensorData, ttl time.Duration) error {
	b, _ := json.Marshal(data)
	return r.client.Set(ctx, r.prefix+":"+key, b, ttl).Err()
}
