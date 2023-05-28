package models

import "time"

type User struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	PasswordHash []byte    `json:"passwordHash"`
	CreatedAt    time.Time `json:"created_at"`
}

type UserData struct {
	UserID         int    `json:"user_id"`
	Name           string `json:"name"`
	Age            int    `json:"age"`
	Sex            string   `json:"sex"`
	Weight         int    `json:"weight"`
	Height         int    `json:"height"`
	Goal           string `json:"goal"`
	Place          string `json:"place"`
	CaloriesIntake int    `json:"calories"`
}
