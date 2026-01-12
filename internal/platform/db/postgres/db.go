package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/zmonsoonz/go-test-api/internal/platform/config"
)


func OpenPostgresDB(cfg config.DBConfig) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode)
	db, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}