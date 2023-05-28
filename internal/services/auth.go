package services

import (
	"crypto/sha1"
	"fmt"
	"my-app/internal/models"
	"my-app/internal/repository"
	"time"

	"github.com/google/uuid"
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
	return s.r.Authorization.CreateUser(*user)
}

func (s *AuthService) GenerateSession(username, password string) (models.Session, int64, error) {
	hash := fmt.Sprintf("%x", generatePasswordHash(password)) 
	user, err := s.r.Authorization.GetUser(username, hash)
	fmt.Println("User: ", user)
	if err != nil {
		return models.Session{}, 0, err
	}
	sessionToken := uuid.NewString()

	newSession := &models.Session{
		Session:   sessionToken,
		UserID:    user.ID,
		CreatedAt: time.Now(),
	}
	fmt.Println("newSession: ",newSession)

	id, err := s.r.Session.SaveSession(*newSession)
	if err != nil {
		return models.Session{}, id ,err
	}
	fmt.Println("SessionID: ",id)
	session, err := s.r.Session.GetSessionByToken(sessionToken)
	if err != nil {
		return models.Session{}, id ,err
	}
	fmt.Println("Session: ",session)
	return session, id, nil	
}

func (s *AuthService) GetSession(sessionToken string) (models.Session, error) {
	
	session, err := s.r.Session.GetSessionByToken(sessionToken)
	if err != nil {
		return models.Session{} ,err
	}
	fmt.Println("Session: ",session)
	return session, nil	
}

func (s *AuthService) DeleteSession(id int64) (error) {
	
	err := s.r.Session.DeleteSession(id)
	if err != nil {
		return err
	}
	return nil	
}


func (s *AuthService) GetUser(id int) (models.User, error) {
	user, err := s.r.Authorization.GetUserByID(id)
	fmt.Println("User: ", user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil	
}

func generatePasswordHash(password string) []byte {
	hash := sha1.New()
	hash.Write([]byte(password))
	return hash.Sum([]byte(salt))
}
