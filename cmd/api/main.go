package main

import (
	"log"

	"github.com/joho/godotenv"
	server "github.com/zmonsoonz/go-test-api"
	repository "github.com/zmonsoonz/go-test-api/internal/content/storage/postgres"
	"github.com/zmonsoonz/go-test-api/internal/content/usecase"
	"github.com/zmonsoonz/go-test-api/internal/platform/config"
	"github.com/zmonsoonz/go-test-api/internal/platform/db"
	api "github.com/zmonsoonz/go-test-api/internal/platform/http"
)
func main()  {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error occurred: %s", err.Error())
	}
	
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("error occurred: %s", err.Error())
	}
	db, err := db.OpenPostgresDB(cfg.DB)
	if  err != nil {
		log.Fatalf("error occurred: %s", err.Error())
	}
	userRep := repository.NewUserRep(db)
	userService := usecase.NewUserService(userRep)
	handlers := new(api.Handler.)
	s := new(server.Server)
	if err := s.Run(cfg, handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurred: %s", err.Error())
	}

}