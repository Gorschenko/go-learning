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

func (s *UsersService) GetOne(filters *UserFilters) (*database.User, error) {
	user, err := s.UsersRepository.GetOne(filters)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New(api.CodeNotFound)
	}

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
