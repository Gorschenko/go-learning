package auth_test

import (
	"test/internal/auth"
	"test/internal/users"
	"testing"
)

const email = "test@test.ru"

type MockUsersRepository struct{}

func (repository *MockUsersRepository) Create(user *users.User) (*users.User, error) {
	return &users.User{
		Email: email,
	}, nil
}

func (repository *MockUsersRepository) FindByEmail(email string) (*users.User, error) {
	return nil, nil
}

func TestRegisterSuccess(t *testing.T) {

	authService := auth.NewAuthService(&MockUsersRepository{})
	resultedEmail, err := authService.Register(email, "test", "test")
	if err != nil {
		t.Fatal(err)
	}

	if resultedEmail != email {
		t.Fatalf("Expected %s got %s", email, resultedEmail)
	}
}
