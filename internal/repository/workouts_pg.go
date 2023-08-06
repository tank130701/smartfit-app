package repository

import (
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

func (r *WorkoutsPostgresRepository) InsertWorkout(workout models.Workout) (int64, error) {
	tx := r.db.MustBegin()
	defer tx.Rollback()

	// Вставляем данные в таблицу workouts
	var workoutID int64
	err := tx.QueryRow(`
            INSERT INTO workouts (date, title)
            VALUES ($1, $2)
            RETURNING id`, workout.Date, workout.Title).Scan(&workoutID)
	if err != nil {
		return 0, err
	}

	// Вставляем данные в таблицу exercises и workout_exercises
	for _, exercise := range workout.Exercises {
		var exerciseID int64
		err = tx.QueryRow(`
                INSERT INTO exercises (title, sets, reps, weights, calories)
                VALUES ($1, $2, $3, $4, $5)
                RETURNING id`, exercise.Title, exercise.Sets, pq.Array(exercise.Reps), pq.Array(exercise.Weights), exercise.Calories).Scan(&exerciseID)
		if err != nil {
			return 0, err
		}

		// Связываем упражнение с тренировкой в таблице workout_exercises
		_, err = tx.Exec(`
                INSERT INTO workout_exercises (workout_id, exercise_id)
                VALUES ($1, $2)`, workoutID, exerciseID)
		if err != nil {
			return 0, err
		}
	}

	// Если все прошло успешно, фиксируем транзакцию
	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return workoutID, nil
}


func (r *WorkoutsPostgresRepository) GetWorkoutByID(id int) (models.Workout, error) {
	var w models.Workout
    err := r.db.Get(&w, `
        SELECT *
        FROM workouts
        WHERE id = $1`, id)
    if err != nil {
        return models.Workout{}, err
    }

    // Запрашиваем данные упражнений, связанных с этой тренировкой
    // exercises := []*models.Exercise{}
    var exercises []models.Exercise
    err = r.db.Select(&exercises, `
        SELECT e.*
        FROM exercises e
        JOIN workout_exercises we ON we.exercise_id = e.id
        WHERE we.workout_id = $1`, id)
    if err != nil {
        return models.Workout{}, err
    }

    // Присваиваем полученные данные упражнений соответствующему полю в тренировке
    w.Exercises = exercises

    return w, nil
}

func (r *WorkoutsPostgresRepository) GetWorkouts() ([]models.Workout, error) {
	var workouts []models.Workout
    err := r.db.Select(&workouts, `
        SELECT *
        FROM workouts
        ORDER BY date DESC
        LIMIT 3;`)
    if err != nil {
        return nil, err
    }
    
    // Получаем данные упражнений, связанных с каждой тренировкой
    for i, workout := range workouts {
        exercises := []models.Exercise{}
        err = r.db.Select(&exercises, `
            SELECT e.*
            FROM exercises e
            JOIN workout_exercises we ON we.exercise_id = e.id
            WHERE we.workout_id = $1`, workout.ID)
        if err != nil {
            return nil, err
        }
        workouts[i].Exercises = exercises
    }

    return workouts, nil
}

