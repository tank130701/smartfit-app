package models

import (
	"fmt"
	"time"
)

type Session struct {
	ID        int64     `json:"id"`
	Session   string    `json:"session"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (s *Session) IsExpired() error {
	if time.Now().Sub(s.CreatedAt) > time.Hour*24 {
		return fmt.Errorf("session in expired on %v", time.Now().Add(3*time.Hour).Local().Sub(s.CreatedAt.Add(time.Minute)))
	}
	return nil
}
