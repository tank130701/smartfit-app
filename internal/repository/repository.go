package repository

// import "my-app/models"
import (
	"database/sql"
	"my-app/models"
	postrgres "my-app/internal/postgres"
)

type Repository struct {
	UserRepository
	SessionRepository
}

type UserRepository interface {
	SaveUser(user *models.User) (int64, error)
	// DeleteUser(id int64) error
	// GetUserByID(id int64) (*models.User, error)
	GetUserByUsername(nickname string) (*models.User, error)
	// GetUsers() ([]*models.User, error)
	// UpdateUser(user *models.User) error
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UserRepository: postrgres.NewUsersPostgres(db),
	}
}

type SessionRepository interface {
	SaveSession(session *models.Session) (int64, error)
	DeleteSession(id int64) error
	GetSessionByToken(sessionToken string) (*models.Session, error)
	GetSessionByID(id int64) (*models.Session, error)
}
