package workoutgenerator

import "my-app/internal/models"

type Generator interface {
	GenerateWorkouts(previousWorkouts []models.Workout, userData models.UserData) ([]models.Workout, error)
}
