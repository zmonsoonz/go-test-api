package bootstrap

import (
	"github.com/jmoiron/sqlx"
	"github.com/zmonsoonz/go-test-api/internal/content/ports"
	repository "github.com/zmonsoonz/go-test-api/internal/content/storage/postgres"
)

type Repos struct {
	User ports.UserRepository
	Auth ports.AuthRepository
}

func InitRepos (db *sqlx.DB) *Repos {
	return &Repos{
		Auth: repository.NewAuthRep(db),
	}
}
