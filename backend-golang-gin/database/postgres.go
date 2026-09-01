package database

import (
	"time"

	"backend-golang-gin/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgres(cfg config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", cfg.DatabaseURL)
	if err != nil { return nil, err }
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(30 * time.Minute)
	return db, nil
}
