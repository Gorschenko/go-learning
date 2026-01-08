package users

import (
	"context"
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

func (s *UsersService) GetOne(ctx context.Context, filters *users_api.UserFiltersDto) (*database.User, error) {
	if filters.ID != 0 {
		cacheUser, _ := s.CacheUsersRepository.GetUser(ctx, filters.ID)

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

	s.CacheUsersRepository.SetUser(ctx, user)

	return user, nil
}

/*
Нужно исправить:
Нужно обновить кэш.
*/
func (s *UsersService) UpdateOne(ctx context.Context, filters *users_api.UserFiltersDto, update *users_api.UserUpdateDto) (int, error) {
	_, err := s.GetOne(ctx, filters)

	if err != nil {
		return 0, err
	}

	count, err := s.UsersRepository.UpdateOne(filters, update)

	return int(count), err
}

func (s *UsersService) DeleteOne(ctx context.Context, filters *users_api.UserFiltersDto) (int, error) {
	_, err := s.GetOne(ctx, filters)

	if err != nil {
		return 0, err
	}

	count, err := s.UsersRepository.DeleteOne(filters)

	return int(count), err
}
