package repository

import (
	"database/sql"
	"errors"
	"net/http"
	"socialNetwork/internal/entity"
)

type CommentRepository struct {
	db *sql.DB
}

func newCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) CreateComment(input entity.Comment) (int, error) {
	query := `INSERT INTO comments(post_id, user_id, content) VALUES($1, $2, $3)`
	prep, err := r.db.Prepare(query)
	if err != nil {
		return http.StatusInternalServerError, errors.New("error preparing query" + err.Error())
	}
	defer prep.Close()
	if _, err := prep.Exec(input.PostID, input.UserID, input.Content); err != nil {
		return http.StatusBadRequest, errors.New("error creating comment" + err.Error())
	}
	return http.StatusOK, nil
}

func (r *CommentRepository) AddCommentReaction(input entity.CommentReaction) (int, error) {
	query := "SELECT is_like FROM likes_comment WHERE user_id = $1 and comment_id = $2;"
	prep, err := r.db.Prepare(query)
	if err != nil {
		return http.StatusInternalServerError, errors.New("error preparing query" + err.Error())
	}
	defer prep.Close()
	var like uint
	if err := prep.QueryRow(input.UserID, input.CommentID).Scan(&like); err != nil {
		return http.StatusInternalServerError, errors.New("error scanning query result" + err.Error())
	}
	if like == 1 {
		query := "DELETE FROM likes_comment WHERE user_id = $1 and comment_id = $2;"
		prep, err := r.db.Prepare(query)
		if err != nil {
			return http.StatusInternalServerError, errors.New("error preparing query" + err.Error())
		}
		defer prep.Close()
		if _, err := prep.Exec(input.UserID, input.CommentID); err != nil {
			return http.StatusInternalServerError, errors.New("error deleting like" + err.Error())
		}
	} else {
		query := "UPDATE likes_comment SET is_like = $1 WHERE user_id = $2 and comment_id = $3;"
		prep, err := r.db.Prepare(query)
		if err != nil {
			return http.StatusInternalServerError, errors.New("error preparing query" + err.Error())
		}
		defer prep.Close()
		if _, err := prep.Exec(1, input.UserID, input.CommentID); err != nil {
			return http.StatusInternalServerError, errors.New("error adding reaction" + err.Error())
		}
	}

	return http.StatusOK, nil
}
