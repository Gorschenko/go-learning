package cache

import (
	"context"
	"pkg/configs"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	*redis.Client
}

func NewCache(config *configs.Config) (*Cache, error) {
	host := config.Cache.Host
	port := config.Cache.Port
	address := host + ":" + strconv.Itoa(port)

	client := redis.NewClient(&redis.Options{
		Addr: address,
	})

	// Проверка подключения
	ctx := context.Background()
	_, err := client.Ping(ctx).Result()

	if err != nil {
		return nil, err
	}

	return &Cache{client}, nil
}
