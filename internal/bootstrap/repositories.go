package bootstrap

import (
	"github.com/jmoiron/sqlx"
	"github.com/zmonsoonz/go-test-api/internal/content/ports"
	repository "github.com/zmonsoonz/go-test-api/internal/content/storage/postgres"
)

type Repos struct {
	Auth ports.AuthRepository
	Track ports.TrackRepository
}

func InitRepos (db *sqlx.DB) *Repos {
	return &Repos{
		Auth: repository.NewAuthRep(db),
		Track: repository.NewTrackRep(db),
	}
}
