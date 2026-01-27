package usecase

import (
	"context"

	"github.com/zmonsoonz/go-test-api/internal/content/ports"
)
type UserService struct {
	repository ports.UserRepository
}

func NewUserService(repository ports.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (u *UserService) Get(ctx context.Context)  {
	if ctx != nil {
		u.repository.Get(ctx)
	}
}