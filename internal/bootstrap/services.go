package bootstrap

import (
	"github.com/zmonsoonz/go-test-api/internal/content/ports"
	"github.com/zmonsoonz/go-test-api/internal/content/usecase"
)

type Services struct {
	Auth ports.AuthService
}

func InitServices (repos *Repos) *Services {
	return &Services{
		Auth: usecase.NewAuthService(repos.Auth),
	}
}
