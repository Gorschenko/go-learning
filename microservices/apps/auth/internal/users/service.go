package users

import (
	"errors"
	"pkg/database"
	"pkg/static"
)

func NewUsersService(dependencies *UsersServiceDependencies) *UsersService {
	return &UsersService{
		UsersRepository: dependencies.UsersRepository,
	}
}

func (s *UsersService) GetOne(filters *GetOneUserFilters) (*database.User, error) {
	user, err := s.UsersRepository.GetOne(filters)

	if user == nil {
		return nil, errors.New(static.ErrorUserNotFound)
	}

	return user, err
}
