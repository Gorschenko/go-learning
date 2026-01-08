package cache

import (
	"context"
	"pkg/configs"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheRepository struct {
	Client *redis.Client
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
		Client: client,
	}

	return &cache, nil
}

func (r *CacheRepository) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

func (r *CacheRepository) Set(ctx context.Context, key, value string, ttl time.Duration) error {
	return r.Client.Set(ctx, key, value, ttl).Err()
}

func (r *CacheRepository) Delete(ctx context.Context, keys ...string) error {
	return r.Client.Del(ctx, keys...).Err()
}
