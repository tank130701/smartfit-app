package services

import (
	"my-app/internal/models"
	"my-app/internal/repository"
)

type Service struct {
	Auth     Auth
	Users    Users
	Workouts Workouts
}

func NewService(r *repository.Repositories) *Service {
	return &Service{
		Auth:     NewAuthService(r),
		Users:    NewUsersService(r),
		Workouts: NewWorkoutsService(r),
	}
}

type Auth interface {
	CreateUser(username, password string) (int, error)
	SignIn(username, password string) (models.Session, error)
}

type Users interface {
	Get(models.Session) (models.User, error)
	EditData(user models.User, newData models.UserData) error
}

type Workouts interface {
	Generate(models.User) error
	Get(models.User) ([]models.Workout, error)
}
