package main

import (
	"log"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
	bootsrap "github.com/zmonsoonz/go-test-api/internal/bootstrap"
	repository "github.com/zmonsoonz/go-test-api/internal/content/storage/postgres"
	"github.com/zmonsoonz/go-test-api/internal/content/usecase"
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
	userRep := repository.NewUserRep(db)
	userService := usecase.NewUserService(userRep)
	handlers := new(bootsrap.Router)
	s := new(bootsrap.Server)
	if err := s.Run(cfg.Http, handlers.InitRoutes(bootsrap.Entities{User: userService})); err != nil {
		log.Fatalf("error occurred: %s", err.Error())
	}

}