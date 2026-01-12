package usecase

import (
	"github.com/zmonsoonz/go-test-api/internal/content/ports"
)
type UserService struct {
	repository ports.UserRepository
}

func NewUserService(repository ports.UserRepository) *UserService {
	return &UserService{repository: repository}
}