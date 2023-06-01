package services

import (
	"my-app/internal/models"
	"my-app/internal/repository"
)

type Service struct {
	Auth     
	UsersData    
	Workouts 
}

func NewService(r *repository.Repositories) *Service {
	return &Service{
		Auth:     NewAuthService(r),
		UsersData:    NewUsersDataService(r),
		Workouts: NewWorkoutsService(r),
	}
}

type Auth interface {
	CreateUser(username, password string) (int, error)
	GenerateSession(username, password string) (models.Session, int64, error)
	GetSession(sessionToken string) (models.Session, error)
	DeleteSession(id int64) (error)
	GetUser(id int) (models.User, error)
}

type UsersData interface {
	CreateUserData(userId int, newData models.UserData) (int, error)
	GetUserData(id int) (models.UserData, error)
	UpdateUserData(user models.User, newData models.UserData) error
}

type Workouts interface {
	Generate(models.User) error
	GetWorkout(id int) (models.Workout, error)
	InsertWorkout(workout models.Workout) (int64, error)
	GetWorkouts() ([]models.Workout, error)
}
