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

func (m *WorkoutsPostgresRepository) CreateWorkout(workout models.Workout) error {
    tx, err := m.db.Beginx()
    if err != nil {
        return err
    }

    // Вставляем данные в таблицу workouts
    result := tx.MustExec(`INSERT INTO workouts (date, title) VALUES ($1, $2) RETURNING id`, workout.Date, workout.Title)

    // Получаем ID последней вставленной записи
    workout.ID, err = result.LastInsertId()
    if err != nil {
        tx.Rollback()
        return err
    }

    // Вставляем данные в таблицу exercises и workout_exercises
    for _, exercise := range workout.Exercises {
        // Вставляем данные в таблицу exercises
        result = tx.MustExec(`INSERT INTO exercises (title, calories) VALUES ($1, $2) RETURNING id`, exercise.Title, exercise.Calories)

        // Получаем ID последней вставленной записи
        exercise.EID, err = result.LastInsertId()
        if err != nil {
            tx.Rollback()
            return err
        }

        // Вставляем данные в таблицу workout_exercises
        we := &WorkoutExercise{
            WorkoutID:  workout.ID,
            ExerciseID: exercise.EID,
            Reps:       exercise.Reps,
        }
        tx.MustExec(`INSERT INTO workout_exercises (workout_id, exercise_id, reps) VALUES ($1, $2, $3)`, we.WorkoutID, we.ExerciseID, we.Reps)
    }

    if err := tx.Commit(); err != nil {
        tx.Rollback()
        return err
    }

    return nil
}

func (m *WorkoutsPostgresRepository) GetWorkoutByID(id int) (models.Workout, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", workoutsTable)
	row := m.db.QueryRow(query, id)
	workout := new(models.Workout)
	err := row.Scan(&workout.ID, &workout.Title, &workout.Exercises, &workout.Date)

	return *workout, err
}
