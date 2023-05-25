package repository

import (
	"database/sql"
	"fmt"
	"my-app/internal/models"
)

type SessionsPostgresRepository struct {
	db *sql.DB
}

func (r *SessionsPostgresRepository) SaveSession(session *models.Session) (int64, error) {
	//TODO implement me
	// panic("implement me")
	query := fmt.Sprintf("INSERT INTO %s (session, user_id) VALUES ($1, $2)", sessionsTable)

	res, err := r.db.Exec(query, session.Session, session.UserID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *SessionsPostgresRepository) DeleteSession(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (r *SessionsPostgresRepository) GetSessionByToken(sessionToken string) (models.Session, error) {
	//TODO implement me
	panic("implement me")
}

func NewSessionsPostgresRepository(db *sql.DB) *SessionsPostgresRepository {
	return &SessionsPostgresRepository{
		db: db,
	}
}
