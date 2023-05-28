package repository

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	"my-app/internal/models"
)

type WorkoutsPostgresRepository struct {
	db *sqlx.DB
}

func NewWorkoutsPostgresRepository(db *sqlx.DB) *WorkoutsPostgresRepository {
	return &WorkoutsPostgresRepository{
		db: db,
	}
}

func (m *WorkoutsPostgresRepository) GetWorkoutByID(id int64) (models.Workout, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", workoutsTable)
	row := m.db.QueryRow(query, id)
	workout := new(models.Workout)
	err := row.Scan(&workout.ID, &workout.Title, &workout.Exercises, &workout.Date)

	return *workout, err
}
