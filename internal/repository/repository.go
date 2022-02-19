package repository

import (
	"github.com/Korisss/concet-backend/internal/repository/psql"
	"github.com/Korisss/concet-backend/internal/types"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user types.User) (int, error)
	GetUser(email, password string) (types.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: psql.NewAuthPostgres(db),
	}
}

func NewPostgresDB(cfg psql.Config) (*sqlx.DB, error) {
	return psql.NewDB(cfg)
}
