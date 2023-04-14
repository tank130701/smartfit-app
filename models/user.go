package models

import "time"

type User struct {
	ID           int64     `json:"id"`
	Username     string    `json:"username"`
	Age          int       `json:"age"`
	Sex          bool      `json:"sex"`
	Weight       int       `json:"weight"`
	PasswordHash []byte    `json:"passwordHash"`
	CreatedAt    time.Time `json:"created_at"`
}
