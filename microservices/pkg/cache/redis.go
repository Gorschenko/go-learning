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
	Ctx    context.Context
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
		Ctx:    ctx,
	}

	return &cache, nil
}

func (r *CacheRepository) Get(key string) (string, error) {
	return r.Client.Get(r.Ctx, key).Result()
}

func (r *CacheRepository) Set(key, value string, ttl time.Duration) error {
	return r.Client.Set(r.Ctx, key, value, ttl).Err()
}

func (r *CacheRepository) Delete(keys ...string) error {
	return r.Client.Del(r.Ctx, keys...).Err()
}
