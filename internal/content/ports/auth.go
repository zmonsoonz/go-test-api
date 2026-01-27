package ports

import (
	"github.com/zmonsoonz/go-test-api/internal/content/domain"
)

type AuthService interface {
	CreateUser(user domain.User) (int, error)
}

type AuthRepository interface {
	CreateUser(user domain.User) (int, error)
}
