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
	baseRepository *CacheRepository
	prefix         string
}

func NewCacheUsersRepository(dependencies *CacheUsersRepositoryDependencies) *CacheUsersRepository {
	return &CacheUsersRepository{
		baseRepository: dependencies.CacheRepository,
		prefix:         CacheUsersPrefix,
	}
}

func (r *CacheUsersRepository) GetUser(ctx context.Context, ID int) (*database.User, error) {
	key := r.prefix + strconv.Itoa(ID)
	userString, err := r.baseRepository.Get(ctx, key)

	if err != nil {
		return nil, err
	}

	var user *database.User

	err = json.Unmarshal([]byte(userString), &user)
	return user, err
}

func (r *CacheUsersRepository) SetUser(ctx context.Context, user *database.User) error {
	key := r.prefix + strconv.Itoa(int(user.ID))
	userString, err := json.Marshal(user)

	if err != nil {
		return err
	}
	err = r.baseRepository.Set(ctx, key, string(userString), CacheTTLLow)

	return err
}
