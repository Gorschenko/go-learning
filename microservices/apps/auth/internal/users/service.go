package users

import (
	"errors"
	"pkg/api"
	"pkg/database"
)

func NewUsersService(dependencies *UsersServiceDependencies) *UsersService {
	return &UsersService{
		UsersRepository: dependencies.UsersRepository,
	}
}

func (s *UsersService) GetOne(filters *GetOneUserFilters) (*database.User, error) {
	user, err := s.UsersRepository.GetOne(filters)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New(api.CodeNotFound)
	}

	return user, nil
}
