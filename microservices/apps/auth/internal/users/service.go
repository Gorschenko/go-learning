package users

import (
	"errors"
	"pkg/api"
	users_api "pkg/api/users"
	"pkg/database"
)

func NewUsersService(dependencies *UsersServiceDependencies) *UsersService {
	return &UsersService{
		UsersRepository:      dependencies.UsersRepository,
		CacheUsersRepository: dependencies.CacheUsersRepository,
	}
}

func (s *UsersService) GetOne(filters *users_api.UserFiltersDto) (*database.User, error) {
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

func (s *UsersService) UpdateOne(filters *users_api.UserFiltersDto, update *users_api.UserUpdateDto) (int, error) {
	_, err := s.GetOne(filters)

	if err != nil {
		return 0, err
	}

	count, err := s.UsersRepository.UpdateOne(filters, update)

	return int(count), err
}

func (s *UsersService) DeleteOne(filters *users_api.UserFiltersDto) (int, error) {
	_, err := s.GetOne(filters)

	if err != nil {
		return 0, err
	}

	count, err := s.UsersRepository.DeleteOne(filters)

	return int(count), err
}
