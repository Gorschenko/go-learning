package cache

import (
	"context"
	"pkg/configs"
	"pkg/logger"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheRepository struct {
	client *redis.Client
}

func NewCacheRepository(config *configs.Config) (*CacheRepository, error) {
	ctx := context.Background()

	host := config.Cache.Host
	port := config.Cache.Port
	address := host + ":" + strconv.Itoa(port)

	client := redis.NewClient(&redis.Options{
		Addr: address,
	})

	// Проверка подключения
	_, err := client.Ping(ctx).Result()

	if err != nil {
		return nil, err
	}

	cache := CacheRepository{
		client,
	}

	return &cache, nil
}

func (r *CacheRepository) Get(ctx context.Context, key string) (string, error) {
	logger := logger.GetLogger(ctx)
	logger.Debug(
		"CacheRepository",
		"Method", "GET",
		"Key", key,
	)
	result, err := r.client.Get(ctx, key).Result()

	logger.Debug(
		"CacheRepository",
		"Method", "GET",
		"Result", result,
		"Error", err,
	)

	return result, err
}

func (r *CacheRepository) Set(ctx context.Context, key, value string, ttl time.Duration) error {
	logger := logger.GetLogger(ctx)
	logger.Debug(
		"CacheRepository",
		"Method", "SET",
		"Key", key,
		"Value", value,
		"TTL", ttl,
	)
	result, err := r.client.Set(ctx, key, value, ttl).Result()

	logger.Debug(
		"CacheRepository",
		"Method", "SET",
		"Result", result,
		"Error", err,
	)

	return err
}

func (r *CacheRepository) Delete(ctx context.Context, keys ...string) error {
	logger := logger.GetLogger(ctx)
	logger.Debug(
		"CacheRepository",
		"Method", "DEL",
		"Key", keys,
	)

	result, err := r.client.Del(ctx, keys...).Result()

	logger.Debug(
		"CacheRepository",
		"Method", "DEL",
		"Result", result,
		"Error", err,
	)

	return err
}
