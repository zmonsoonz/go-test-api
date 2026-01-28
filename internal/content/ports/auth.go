package ports

import (
	"github.com/zmonsoonz/go-test-api/internal/content/domain"
)

type AuthService interface {
	CreateUser(user domain.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type AuthRepository interface {
	CreateUser(user domain.User) (int, error)
	GetUser(username string) (domain.User, error)
}
