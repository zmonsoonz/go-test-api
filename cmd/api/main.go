package main

import (
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/zmonsoonz/go-test-api/internal/bootstrap"
	"github.com/zmonsoonz/go-test-api/internal/platform/config"
	"github.com/zmonsoonz/go-test-api/internal/platform/db/postgres"
)
func main()  {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error occurred: %s", err.Error())
	}
	
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("error occurred: %s", err.Error())
	}
	db, err := postgres.OpenPostgresDB(cfg.DB)
	if  err != nil {
		log.Fatalf("error occurred: %s", err.Error())
	}

	repos := bootstrap.InitRepos(db)
	services := bootstrap.InitServices(repos)
	handlers := bootstrap.InitHandlers(services)
	router := bootstrap.InitRouter(handlers)
	
	s := bootstrap.NewServer()

	if err := s.Run(cfg.Http, router); err != nil {
		log.Fatalf("error occurred: %s", err.Error())
	}

}