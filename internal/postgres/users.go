package postrgres

import (
	"database/sql"
	"fmt"
	"log"
	"my-app/models"
	"time"
)

type UsersPostgres struct {
	db *sql.DB
}

func NewUsersPostgres(db *sql.DB) *UsersPostgres {
	return &UsersPostgres{
		db: db,
	}
}

func (m *UsersPostgres) SaveUser(user *models.User) (int64, error) {
	var LastInsertId int
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash) VALUES ($1, $2) RETURNING id", usersTable)

	err := m.db.QueryRow(
		query, user.Username, user.PasswordHash).Scan(&LastInsertId)
	if err != nil {
		return 0, err
	}

	return int64(LastInsertId), nil
}

func (m *UsersPostgres) DeleteUser(id int64) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", usersTable)

	row := m.db.QueryRow(query, id)

	return row.Err()
}

func (m *UsersPostgres) GetUserByUsername(username string) (*models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE username = ?", usersTable)

	row := m.db.QueryRow(query, username)

	user := new(models.User)

	err := row.Scan(&user.ID, &user.Username, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (m *UsersPostgres) GetUserByID(id int64) (*models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", usersTable)

	row := m.db.QueryRow(query, id)

	user := new(models.User)

	err := row.Scan(&user.ID, &user.Username, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (m *UsersPostgres) GetUsers() ([]*models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s", usersTable)

	rows, err := m.db.Query(query)

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	if err != nil {
		return nil, err
	}

	users := make([]*models.User, 0, 1)
	var id int64
	var username string
	var passwordHash []byte
	var createdAt time.Time

	for rows.Next() {
		err = rows.Scan(&id, &username, &passwordHash, &createdAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &models.User{
			ID:           id,
			Username:     username,
			PasswordHash: passwordHash,
			CreatedAt:    createdAt,
		})
	}

	return users, nil
}

//Надо разбираться 
func (m *UsersPostgres) UpdateUser(user *models.User) error {
	query := fmt.Sprintf("UPDATE %s SET %s.username = ?, WHERE %s.id = ?", usersTable, usersTable, usersTable)

	_, err := m.db.Exec(query, user.Username, user.ID)
	if err != nil {
		return err
	}

	return nil
}