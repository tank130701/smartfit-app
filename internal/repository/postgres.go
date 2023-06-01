package repository

import (
	"github.com/jmoiron/sqlx"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	usersTable           = "users"
	sessionsTable        = "sessions"
	usersDataTable       = "users_data"
	workoutsTable        = "workouts"
	workoutsArchiveTable = "workouts_archive"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresConnection(cfg Config) (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.DBName,
		cfg.SSLMode,
	)
	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
