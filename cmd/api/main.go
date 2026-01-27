package main

import (
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	bootsrap "github.com/zmonsoonz/go-test-api/internal/bootstrap"
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

	repos := bootsrap.InitRepos(db)
	services := bootsrap.InitServices(repos)
	handlers := bootsrap.InitHandlers(services)
	router := bootsrap.InitRouter(handlers)
	
	s := bootsrap.NewServer()

	if err := s.Run(cfg.Http, router); err != nil {
		log.Fatalf("error occurred: %s", err.Error())
	}

}