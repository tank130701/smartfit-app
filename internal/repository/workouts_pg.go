package repository

import (
	"database/sql"
	"my-app/internal/models"
	"fmt"
)

type WorkoutsPostgresRepository struct {
	db *sql.DB
}

func NewWorkoutsPostgresRepository(db *sql.DB) *WorkoutsPostgresRepository {
	return &WorkoutsPostgresRepository{
		db: db,
	}
}


func (m *UsersPostgresRepository) GetWorkoutByID(id int64) (*models.Workout, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", workoutsTable)
	row := m.db.QueryRow(query, id)
	workout := new(models.Workout)
	err := row.Scan(&workout.ID, &workout.Title, &workout.Exercises, &workout.Date)
	if err != nil {
		return nil, err
	}
	return workout, nil
}