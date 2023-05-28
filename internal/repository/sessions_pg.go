package repository

import (
	"fmt"
	"my-app/internal/models"

	"github.com/jmoiron/sqlx"
)

type SessionsPostgresRepository struct {
	db *sqlx.DB
}

func NewSessionsPostgresRepository(db *sqlx.DB) *SessionsPostgresRepository {
	return &SessionsPostgresRepository{
		db: db,
	}
}

func (r *SessionsPostgresRepository) SaveSession(session models.Session) (int64, error) {
	var LastInsertId int64
	fmt.Println("Session in Function SaveSession: ", session)
	query := fmt.Sprintf("INSERT INTO %s (session, user_id, created_at) VALUES ($1, $2, $3) RETURNING id", sessionsTable)
	// res, err := r.db.Exec(query, session.Session, session.UserID, session.CreatedAt)
	err := r.db.QueryRow(
		query, session.Session, session.UserID, session.CreatedAt).Scan(&LastInsertId)
	if err != nil {
		return 0, err
	}
	return int64(LastInsertId), nil
}

func (m *SessionsPostgresRepository) GetSessionByID(id int64) (models.Session, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", sessionsTable)
	row := m.db.QueryRow(query, id)
	session := new(models.Session)

	err := row.Scan(&session.Session, &session.UserID, &session.CreatedAt)

	return *session, err
}

func (r *SessionsPostgresRepository) DeleteSession(id int64) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", sessionsTable)

	_, err := r.db.Exec(query, id)

	return err
}

func (m *SessionsPostgresRepository) GetSessionByToken(sessionToken string) (models.Session, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE session = $1", sessionsTable)
	row := m.db.QueryRow(query, sessionToken)
	session := new(models.Session)

	err := row.Scan(&session.ID, &session.Session, &session.UserID, &session.CreatedAt)

	return *session, err
}
