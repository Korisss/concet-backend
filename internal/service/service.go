package service

import (
	"github.com/Korisss/concet-backend/internal/repository"
	"github.com/Korisss/concet-backend/internal/types"
)

type Authorization interface {
	CreateUser(user types.User) (int, error)
	GenerateToken(email, password string) (int, string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
