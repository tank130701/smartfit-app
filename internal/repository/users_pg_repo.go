package repository

import (
	"database/sql"
	"fmt"
	"my-app/internal/models"
)

type UsersPostgresRepository struct {
	db *sql.DB
}

func NewUsersPostgresRepository(db *sql.DB) *UsersPostgresRepository {
	return &UsersPostgresRepository{
		db: db,
	}
}

func (m *UsersPostgresRepository) CreateUser(user *models.User) (int, error) {
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

func (m *UsersPostgresRepository) GetUserByID(id int64) (*models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", usersTable)
	row := m.db.QueryRow(query, id)
	user := new(models.User)
	err := row.Scan(&user.ID, &user.Username, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (m *UsersPostgresRepository) GetUserByUsername(username string) (*models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE username = ?", usersTable)
	row := m.db.QueryRow(query, username)
	user := new(models.User)
	err := row.Scan(&user.ID, &user.Username, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UsersPostgresRepository) Update(userID int, newUser models.User) (err error) {
	_, err = r.db.Exec("UPDATE ...") // TODO: query
	return
}
