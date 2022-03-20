package repository_mock

import (
	"github.com/Korisss/concet-backend/internal/domain"
	"github.com/Korisss/concet-backend/internal/repository"
)

type DB struct {
	users []domain.User
}

func NewRepositoryMock(db *DB) *repository.Repository {
	return &repository.Repository{
		Authorization: NewAuthMock(db),
	}
}

func NewDBMock() *DB {
	return &DB{
		users: []domain.User{},
	}
}

type AuthMock struct {
	db *DB
}

func NewAuthMock(db *DB) *AuthMock {
	return &AuthMock{db: db}
}

func (r *AuthMock) CreateUser(user domain.User) (int, error) {
	return 0, nil
}

func (r *AuthMock) GetUser(email, password string) (domain.User, error) {
	return domain.User{}, nil
}
