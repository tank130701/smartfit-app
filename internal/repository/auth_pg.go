package repository

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	"my-app/internal/models"
)

type AuthPostgresRepository struct {
	db *sqlx.DB
}

func NewAuthPostgresRepository(db *sqlx.DB) *AuthPostgresRepository {
	return &AuthPostgresRepository{
		db: db,
	}
}

func (m *AuthPostgresRepository) CreateUser(user models.User) (int, error) {
	var LastInsertId int
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash, created_at) VALUES ($1, $2, $3) RETURNING id", usersTable)
	hash := fmt.Sprintf("%x", user.PasswordHash)
	err := m.db.QueryRow(
		query, user.Username, hash, user.CreatedAt).Scan(&LastInsertId)
	if err != nil {
		return 0, err
	}
	return int(LastInsertId), nil
}

func (m *AuthPostgresRepository) GetUser(username, password string) (models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	row := m.db.QueryRow(query, username, password)
	user := models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}


func (m *AuthPostgresRepository) GetUserByID(id int) (models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", usersTable)
	row := m.db.QueryRow(query, id)
	user := models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (m *AuthPostgresRepository) GetUserByUsername(username string) (*models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE username = $1", usersTable)
	row := m.db.QueryRow(query, username)
	user := new(models.User)
	err := row.Scan(&user.ID, &user.Username, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *AuthPostgresRepository) UpdateUser(userID int, newUser models.User) (err error) {
	_, err = r.db.Exec("UPDATE ...") // TODO: query
	return
}
