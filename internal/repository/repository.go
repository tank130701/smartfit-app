package repository

// import "my-app/models"
import (
	"database/sql"
	"my-app/internal/models"
)

type Repositories struct {
	Users     Users
	UsersData UsersData
	Sessions  Session
}

func NewRepository(db *sql.DB) *Repositories {
	return &Repositories{
		Users:     NewUsersPostgresRepository(db),
		UsersData: NewUsersDataPostgresRepository(db),
		Sessions:  NewSessionsPostgresRepository(db),
	}
}

type Users interface {
	CreateUser(user *models.User) (int, error)
	Update(userID int, newUser models.User) error
}
type UsersData interface {
	Create(data models.UserData) (int, error)
	Update(userID int, newData models.UserData) error
}

type Session interface {
	SaveSession(session *models.Session) (int64, error)
	DeleteSession(id int64) error
	GetSessionByToken(sessionToken string) (models.Session, error)
}
