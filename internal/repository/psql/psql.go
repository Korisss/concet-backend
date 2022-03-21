package psql

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	usersTable = "users"
)

type Configuration struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewDB(cfg *Configuration) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func LoadConfig() *Configuration {
	config := &Configuration{}

	config.Host = os.Getenv("POSTGRES_HOST")
	config.Port = os.Getenv("POSTGRES_PORT")
	config.Username = os.Getenv("POSTGRES_USERNAME")
	config.Password = os.Getenv("POSTGRES_PASSWORD")
	config.DBName = os.Getenv("POSTGRES_DB_NAME")
	config.SSLMode = os.Getenv("POSTGRES_SSL_MODE")

	return config
}
