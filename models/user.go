package models

import "time"

type User struct {
	ID           int64     `json:"id"`
	Username     string    `json:"username"`
	PasswordHash []byte    `json:"passwordHash"`
	CreatedAt    time.Time `json:"created_at"`
}

type UserData struct {
	ID             int64  `json:"id"`
	Age            int    `json:"age"`
	Sex            bool   `json:"sex"`
	Weight         int    `json:"weight"`
	Height         int    `json:"height"`
	Goal           string `json:"goal"`
	Place          string `json:"place"`
	CaloriesIntake int    `json:"calories"`
}
