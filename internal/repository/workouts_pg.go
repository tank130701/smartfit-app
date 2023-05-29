package repository

import (
	"fmt"
	"my-app/internal/models"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type WorkoutsPostgresRepository struct {
	db *sqlx.DB
}

func NewWorkoutsPostgresRepository(db *sqlx.DB) *WorkoutsPostgresRepository {
	return &WorkoutsPostgresRepository{
		db: db,
	}
}

func (m *WorkoutsPostgresRepository) InsertWorkout(workout models.Workout) (int,error) {
    tx, err := m.db.Begin()
    if err != nil {
        return 0, err
    }

    // Добавляем запись для таблицы Workout
    var workoutID int
    err = tx.QueryRow(`
        INSERT INTO workouts (date, title)
        VALUES ($1, $2)
        RETURNING id;
    `, workout.Date, workout.Title).Scan(&workoutID)
    if err != nil {
        tx.Rollback()
        return 0, err
    }

    // Добавляем запись для таблицы Exercises
    for _, exercise := range workout.Exercises {
        var exerciseID int
        err = tx.QueryRow(`
            INSERT INTO exercises (title, calories)
            VALUES ($1, $2)
            RETURNING id;
        `, exercise.Title, exercise.Calories).Scan(&exerciseID)
        if err != nil {
            tx.Rollback()
            return 0, err
        }

        // Добавляем записи для таблицы WorkoutExercises
        _, err = tx.Exec(`
            INSERT INTO workout_exercises (workout_id, exercise_id, reps)
            VALUES ($1, $2, $3);
        `, workoutID, exerciseID, pq.Array(exercise.Reps))
        if err != nil {
            tx.Rollback()
            return 0, err
        }
    }

    // Заканчиваем транзакцию
    err = tx.Commit()
    if err != nil {
        return 0, err
    }

    return workoutID, nil
}

func (m *WorkoutsPostgresRepository) GetWorkoutByID(id int) (models.Workout, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", workoutsTable)
	row := m.db.QueryRow(query, id)
	workout := new(models.Workout)
	err := row.Scan(&workout.ID, &workout.Title, &workout.Exercises, &workout.Date, &workout.WorkoutExercises)

	return *workout, err
}
