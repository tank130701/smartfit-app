package models

import "time"

type Workout struct {
	ID        int `json:"id"`
	Date      time.Time
	Title     string
	Exercises []Exercise
	WorkoutExercises []WorkoutExercise
}

type Exercise struct {
	Title    string
	Reps     []int
	Calories int `json:"calories"`
}

type WorkoutExercise struct {
	WorkoutID   int
	ExerciseID  int
}