package models

type Workout struct {
	ID      int64  `json:"id"`
	Workout string `json:"workout"`
	Calories int `json:"calories"`
}

type WorkoutsArchive struct {
	ID      int64  `json:"id"`
	Workout Workout `json:"workout"`
}