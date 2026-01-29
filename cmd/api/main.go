package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/zmonsoonz/go-test-api/internal/bootstrap"
	"github.com/zmonsoonz/go-test-api/internal/platform/config"
	"github.com/zmonsoonz/go-test-api/internal/platform/db/postgres"
)
func main()  {

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("failed to load .env file: %v", err)
	}

	cfg, err := config.Load()
	if err != nil {
		logrus.Fatalf("failed to load config: %v", err)
	}

	db, err := postgres.OpenPostgresDB(cfg.DB)
	if err != nil {
		logrus.Fatalf("failed to connect to database: %v", err)
	}

	repos := bootstrap.InitRepos(db)
	services := bootstrap.InitServices(repos)
	handlers := bootstrap.InitHandlers(services)
	router := bootstrap.InitRouter(handlers, services)
	
	s := bootstrap.NewServer()

	if err := s.Run(cfg.Http, router); err != nil {
		logrus.Fatalf("server stopped with error: %v", err)
	}

}