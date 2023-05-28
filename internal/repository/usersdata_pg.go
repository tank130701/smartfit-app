package repository

import (
	"github.com/jmoiron/sqlx"
	"my-app/internal/models"
	"fmt"
)

type UsersDataPostgresRepository struct {
	db *sqlx.DB
}

func NewUsersDataPostgresRepository(db *sqlx.DB) *UsersDataPostgresRepository {
	return &UsersDataPostgresRepository{
		db: db,
	}
}

func (r *UsersDataPostgresRepository) Create(userData models.UserData) (int, error) {
	//TODO implement me
	//panic("implement me")
	var LastInsertId int
	query := fmt.Sprintf(
	"INSERT INTO %s (user_id, sex, age, weight, height, goal, place, calories) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", usersDataTable)
	err := r.db.QueryRow(
		query, userData.UserID, userData.Sex, userData.Age, userData.Weight, userData.Height,
		userData.Place).Scan(&LastInsertId)
	if err != nil {
		return 0, err
	}
	return int(LastInsertId), nil
}

func (r *UsersDataPostgresRepository) Update(userID int, newData models.UserData) error {
	//TODO implement me
	panic("implement me")
}
