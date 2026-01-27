package ports

import (
	"context"

	"github.com/zmonsoonz/go-test-api/internal/content/domain"
)

type UserService interface {
	Get(ctx context.Context) (domain.User, error)
}

type UserRepository interface {
	Get(ctx context.Context) (int64, error)
}