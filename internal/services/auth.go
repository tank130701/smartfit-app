package services

import (
	"crypto/sha1"
	"my-app/internal/models"
	"my-app/internal/repository"
	"time"
)

const (
	salt = "hjqrhjqw124617ajfhajs"
)

type AuthService struct {
	r *repository.Repositories
}

func NewAuthService(r *repository.Repositories) *AuthService {
	return &AuthService{r: r}
}

func (s *AuthService) CreateUser(username, password string) (int, error) {

	passwordHash := generatePasswordHash(password)
	user := &models.User{
		Username:     username,
		PasswordHash: passwordHash,
		CreatedAt: time.Now(),
		// Data: nil,
	}
	return s.r.Users.CreateUser(user)
}

func (s *AuthService) SignIn(username, password string)  {
	//TODO implement me
	panic("implement me")
}

func generatePasswordHash(password string) []byte {
	hash := sha1.New()
	hash.Write([]byte(password))
	return hash.Sum([]byte(salt))
}
