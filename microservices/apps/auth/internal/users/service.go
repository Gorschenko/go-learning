package users

import (
	"errors"
	"pkg/database"
	pkg_errors "pkg/errors"
)

func NewUsersService(dependencies *UsersServiceDependencies) *UsersService {
	return &UsersService{
		UsersRepository: dependencies.UsersRepository,
	}
}

func (s *UsersService) GetOne(filters *GetOneUserFilters) (*database.User, error) {
	user, err := s.UsersRepository.GetOne(filters)

	if user == nil {
		return nil, errors.New(string(pkg_errors.CodeNotFound))
	}

	return user, err
}
