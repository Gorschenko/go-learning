package di

import "test/internal/users"

type IStatsRepository interface {
	AddClick(linkId uint)
}

type IUsersRepository interface {
	Create(user *users.User) (*users.User, error)
	FindByEmail(email string) (*users.User, error)
}
