package usecase

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/zmonsoonz/go-test-api/internal/content/domain"
	"github.com/zmonsoonz/go-test-api/internal/content/ports"
	"golang.org/x/crypto/bcrypt"
)


type AuthService struct {
	repository ports.AuthRepository
}

type tokenClaims struct {
	jwt.MapClaims
	UserId int `json:"user_id"`
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

func (s *AuthService) GenerateToken(username, password string) (string, error)  {
	user, err := s.repository.GetUser(username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword(
			[]byte(user.Password),
			[]byte(password),
		)
	if err != nil {
		return "", err
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims {
		jwt.MapClaims{
		"ExpiresAt":     time.Now().Add(24 * time.Hour).Unix(),
		"IssuedAt":     time.Now().Unix(), 
	},
		user.Id,
	})

	return token.SignedString([]byte("secret"))
} 