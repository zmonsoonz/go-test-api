package usecase

import (
	"time"

	"github.com/zmonsoonz/go-test-api/internal/content/domain"
	"github.com/zmonsoonz/go-test-api/internal/content/ports"
	"golang.org/x/crypto/bcrypt"
)


type AuthService struct {
	repository ports.AuthRepository
}

func NewAuthService(repository ports.AuthRepository) *AuthService {
	return &AuthService{repository: repository}
}

func (s *AuthService) CreateUser(user domain.User) (int, error)  {
	user.Password = generatePasswordHash(user.Password)
	user.CreatedAt = time.Now()
	return s.repository.CreateUser(user)
} 

func generatePasswordHash(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}