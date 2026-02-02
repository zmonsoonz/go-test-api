package ports

import (
	"github.com/zmonsoonz/go-test-api/internal/content/domain"
)

type AuthService interface {
	SignUp(user domain.User) (int, error)
	SignIn(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type AuthRepository interface {
	Create(user domain.User) (int, error)
	GetByUsername(username string) (domain.User, error)
}
