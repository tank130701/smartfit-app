package models

import "time"

type Workout struct {
	ID        int `json:"id"`
	Date      time.Time
	Title     string
	Exercises []Exercise
}

type Exercise struct {
	Title    string
	Reps     []int
	Calories int `json:"calories"`
}
