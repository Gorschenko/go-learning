package cache

import (
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

func (r *CacheUsersRepository) GetUser(ID int) (*database.User, error) {
	key := r.Prefix + strconv.Itoa(ID)
	userString, err := r.Client.Get(r.Ctx, key).Bytes()

	if err != nil {
		return nil, err
	}

	var user *database.User

	err = json.Unmarshal(userString, &user)
	return user, err
}

func (r *CacheUsersRepository) SetUser(user *database.User) error {
	key := r.Prefix + strconv.Itoa(int(user.ID))
	userString, err := json.Marshal(user)

	if err != nil {
		return err
	}

	err = r.Client.Set(r.Ctx, key, userString, CacheTTLLow).Err()

	return err
}
