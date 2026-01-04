package users

import (
	"errors"
	"pkg/api"
	"pkg/database"
)

func NewUsersService(dependencies *UsersServiceDependencies) *UsersService {
	return &UsersService{
		UsersRepository:      dependencies.UsersRepository,
		CacheUsersRepository: dependencies.CacheUsersRepository,
	}
}

func (s *UsersService) GetOne(filters *UserFilters) (*database.User, error) {
	if filters.ID != 0 {
		cacheUser, _ := s.CacheUsersRepository.GetUser(filters.ID)

		if cacheUser != nil {
			return cacheUser, nil
		}
	}

	user, err := s.UsersRepository.GetOne(filters)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New(api.CodeNotFound)
	}

	s.CacheUsersRepository.SetUser(user)

	return user, nil
}

func (s *UsersService) DeleteOne(filters *UserFilters) (int, error) {
	_, err := s.GetOne(filters)

	if err != nil {
		return 0, err
	}

	count, err := s.UsersRepository.DeleteOne(filters)

	return count, err
}
