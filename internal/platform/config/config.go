package config

import (
	"os"
	"time"
)

type Config struct {
	Http HttpConfig
	DB   DBConfig
}
type HttpConfig struct {
	Port    string
	ReadTimeout time.Duration
	WriteTimeout time.Duration

}
type DBConfig struct {
	Driver   string
	Host     string
	Port     int
	Username     string
	Password string
	DBName     string
	SSLMode  string
}


func Load() (Config, error) {
	cfg := Config{
		Http: HttpConfig{
			Port:         "8080",
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
		},
		DB: DBConfig{
			Host:    "localhost",
			Port:    5436,
			Username: "postgres",
			Password: os.Getenv("DB_PASSWORD"),
			DBName:     "postgres",
			SSLMode:  "disable",
		},
	}

	return cfg, nil
}
