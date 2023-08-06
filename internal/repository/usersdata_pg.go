package repository

import (
	"fmt"
	"my-app/internal/models"
	"strings"

	"github.com/jmoiron/sqlx"
)

type UsersDataPostgresRepository struct {
	db *sqlx.DB
}

func NewUsersDataPostgresRepository(db *sqlx.DB) *UsersDataPostgresRepository {
	return &UsersDataPostgresRepository{
		db: db,
	}
}

func (r *UsersDataPostgresRepository) Create(userId int, userData models.UserData) (int, error) {
	fmt.Println(userData)
	var LastInsertId int
	query := fmt.Sprintf(
	"INSERT INTO %s (user_id, name, age, sex, weight, height, goal, place, calories_intake) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING user_id", usersDataTable)
	err := r.db.QueryRow(
		query, userId, userData.Name, userData.Age, userData.Sex, userData.Weight, userData.Height, userData.Goal,
		userData.Place, userData.CaloriesIntake).Scan(&LastInsertId)
	if err != nil {
		return 0, err
	}
	return int(LastInsertId), nil
}

func (r *UsersDataPostgresRepository) Get(id int) (models.UserData, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", usersDataTable)
	row := r.db.QueryRow(query, id)
	UsersData := new(models.UserData)

	err := row.Scan(&UsersData.UserID, &UsersData.Name, &UsersData.Age,
		 &UsersData.Sex, &UsersData.Weight, &UsersData.Height, &UsersData.Goal, &UsersData.Place, &UsersData.CaloriesIntake)

	return *UsersData, err
}


func (r *UsersDataPostgresRepository) Update(id int, newData models.UserData) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if newData.Name != "" {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argID))
		args = append(args, newData.Name)
		argID++
	}

	if newData.Age != 0 {
		setValues = append(setValues, fmt.Sprintf("age=$%d", argID))
		args = append(args, newData.Age)
		argID++
	}

	if newData.Sex != "" {
		setValues = append(setValues, fmt.Sprintf("sex=$%d", argID))
		args = append(args, newData.Sex)
		argID++
	}

	if newData.Weight != 0 {
		setValues = append(setValues, fmt.Sprintf("weight=$%d", argID))
		args = append(args, newData.Weight)
		argID++
	}

	if newData.Height != 0 {
		setValues = append(setValues, fmt.Sprintf("height=$%d", argID))
		args = append(args, newData.Height)
		argID++
	}

	if newData.Goal != "" {
		setValues = append(setValues, fmt.Sprintf("goal=$%d", argID))
		args = append(args, newData.Goal)
		argID++
	}

	if newData.Place != "" {
		setValues = append(setValues, fmt.Sprintf("place=$%d", argID))
		args = append(args, newData.Place)
		argID++
	}

	if newData.CaloriesIntake != 0 {
		setValues = append(setValues, fmt.Sprintf("calories_intake=$%d", argID))
		args = append(args, newData.CaloriesIntake)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s SET %s WHERE id=$%d`, usersDataTable, setQuery, argID)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}