package repository


import (
	"github.com/jmoiron/sqlx"
	"my-app/internal/models"
)

type Repositories struct {
	Authorization     
	UsersData 
	Session 
	Workout 
}

func NewRepository(db *sqlx.DB) *Repositories {
	return &Repositories{
		Authorization:     NewAuthPostgresRepository(db),
		UsersData: NewUsersDataPostgresRepository(db),
		Session:  NewSessionsPostgresRepository(db),
		Workout: NewWorkoutsPostgresRepository(db),
	}
}

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
	GetUserByID(id int) (models.User, error)
	UpdateUser(userID int, newUser models.User) error
}
type UsersData interface {
	Create(data models.UserData) (int, error)
	Update(userID int, newData models.UserData) error
}

type Session interface {
	SaveSession(session models.Session) (int64, error)
	DeleteSession(id int64) error
	GetSessionByToken(sessionToken string) (models.Session, error)
	GetSessionByID(id int64) (models.Session, error)
}

type Workout interface {
	GetWorkoutByID(id int64) (models.Workout, error)
}

