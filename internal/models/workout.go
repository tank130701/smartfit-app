package models

import (
	"time"

	"github.com/lib/pq"
)

type Workout struct {
	ID               int `json:"id"`
	Date             time.Time
	Title            string
	Exercises        []Exercise
	WorkoutExercises []WorkoutExercise
}

type Exercise struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Sets     int    `json:"sets"`
	Reps     pq.Int64Array  `json:"reps"`
	Weights  pq.Float64Array  `json:"weights"`
	Calories int    `json:"calories"`
}

type WorkoutExercise struct {
	WorkoutID  int
	ExerciseID int
}
