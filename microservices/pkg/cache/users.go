package cache

import (
	"context"
	"encoding/json"
	"pkg/database"
	"strconv"
)

type CacheUsersRepositoryDependencies struct {
	*CacheRepository
}

type CacheUsersRepository struct {
	*CacheRepository
	Prefix string
}

func NewCacheUsersRepository(dependencies *CacheUsersRepositoryDependencies) *CacheUsersRepository {
	return &CacheUsersRepository{
		CacheRepository: dependencies.CacheRepository,
		Prefix:          CacheUsersPrefix,
	}
}

func (r *CacheUsersRepository) GetUser(ctx context.Context, ID int) (*database.User, error) {
	key := r.Prefix + strconv.Itoa(ID)
	userString, err := r.CacheRepository.Get(ctx, key)

	if err != nil {
		return nil, err
	}

	var user *database.User

	err = json.Unmarshal([]byte(userString), &user)
	return user, err
}

func (r *CacheUsersRepository) SetUser(ctx context.Context, user *database.User) error {
	key := r.Prefix + strconv.Itoa(int(user.ID))
	userString, err := json.Marshal(user)

	if err != nil {
		return err
	}
	err = r.CacheRepository.Set(ctx, key, string(userString), CacheTTLLow)

	return err
}
