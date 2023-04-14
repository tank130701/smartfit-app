package repository

// import "my-app/models"
import (
	"my-app/pkg/postgres"
	"database/sql"
	"my-app/models"
)


type Repository struct {
	UserRepository
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