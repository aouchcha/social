package repository

import (
	"database/sql"
	"net/http"

	"socialNetwork/internal/entity"
)

type SessionRepository struct {
	db *sql.DB
}

func newSessionRepository(db *sql.DB) *SessionRepository {
	return &SessionRepository{db: db}
}

func (r *SessionRepository) IsTokenExist(session entity.Session) (bool, error) {
	var isToken bool
	query := `SELECT EXISTS(SELECT * FROM sessions WHERE token = ? AND user_id = ?)`
	if err := r.db.QueryRow(query, session.Token, session.UserID).Scan(&isToken); err != nil {
		return false, err
	}
	if !isToken {
		return false, nil
	}
	return true, nil
}

func (r *SessionRepository) DeleteSessionByToken(session entity.Session) error {
	query := `DELETE FROM sessions WHERE token = ?;`
	if _, err := r.db.Exec(query, session.Token); err != nil {
		return err
	}
	return nil
}

func (r *SessionRepository) DeleteSessionByUserID(session entity.Session) error {
	query := `DELETE FROM sessions WHERE user_id = ?;`
	if _, err := r.db.Exec(query, session.UserID); err != nil {
		return err
	}
	return nil
}

func (r *SessionRepository) CreatSession(session entity.Session) (int, error) {
	query := `INSERT INTO sessions (user_id, token) VALUES (?, ?)`
	if _, err := r.db.Exec(query, session.UserID, session.Token); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
