package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"socialNetwork/internal/entity"
)

type PostRepository struct {
	db *sql.DB
}

func newPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) GetAllPosts(limit, offset int) ([]entity.Post, int, error) {
	query := `
	SELECT *
	FROM all_posts_with_details
	LIMIT $1 OFFSET $2;
	`

	prep, err := r.db.Prepare(query)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("Error Preparing query: " + err.Error())
	}
	defer prep.Close()

	posts := []entity.Post{}
	rows, err := prep.Query(limit, offset)
	if err != nil {
		return nil, http.StatusBadRequest, errors.New("Error executing query: " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		post := entity.Post{}
		if err := rows.Scan(
			&post.PostID,
			&post.UserID,
			&post.GroupID,
			&post.UserName,
			&post.Content,
			&post.Image_url,
			&post.Privacy,
			&post.Likes,
			&post.CommentsCount,
			&post.CreatedAt,
		); err != nil {
			return nil, http.StatusInternalServerError, errors.New("Error scanning row: " + err.Error())
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, http.StatusInternalServerError, errors.New("Error scanning rows: " + err.Error())
	}

	if len(posts) == 0 {
		return nil, http.StatusNoContent, fmt.Errorf("no posts found")
	}

	return posts, http.StatusOK, nil
}

func (r *PostRepository) GetPostByID(postID uint) (entity.Post, int, error) {
	var post entity.Post
	query := `
	SELECT * FROM all_posts_with_details 
	WHERE post_id = ?;
	`
	prep, err := r.db.Prepare(query)
	if err != nil {
		return post, http.StatusInternalServerError, errors.New("Error Preparing query: " + err.Error())
	}

	if err := prep.QueryRow(postID).Scan(
		&post.PostID,
		&post.UserID,
		&post.UserName,
		&post.Content,
		&post.Likes,
		&post.CommentsCount,
		&post.CreatedAt,
	); err != nil {
		return post, http.StatusNotFound, errors.New("Error scanning row: " + err.Error())
	}

	comments, status, err := r.getCommentsByPostID(postID)
	if err != nil {
		return post, status, errors.New("Error getting comments: " + err.Error())
	}

	post.Comments = comments
	return post, http.StatusOK, nil
}

func (r *PostRepository) CreatePost(input entity.Post) (uint, int, error) {
	query := `
	INSERT INTO posts (content, image_url, privacy, user_id, group_id) 
	VALUES (?, ?, ?, ?, ?)
	RETURNING post_id;
	`
	prep, err := r.db.Prepare(query)
	if err != nil {
		return 0, http.StatusInternalServerError, errors.New("Error Preparing query: " + err.Error())
	}

	defer prep.Close()
	var id uint
	if err = prep.QueryRow(
		input.Content,
		input.Image_url,
		input.Privacy,
		input.UserID,
		input.GroupID,
	).Scan(&id); err != nil {
		return 0, http.StatusBadRequest, errors.New("Error scanning row: " + err.Error())
	}
	return id, http.StatusOK, nil
}

func (r *PostRepository) GetAllByUserID(userID uint, limit, offset int) ([]entity.Post, int, error) {
	posts := []entity.Post{}
	query := `
	SELECT * FROM all_posts_with_details 
	WHERE author_id = ?
	LIMIT $1 OFFSET $2;
	`
	prep, err := r.db.Prepare(query)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("Error Preparing query: " + err.Error())
	}
	defer prep.Close()
	rows, err := prep.Query(userID, limit, offset)
	if err != nil {
		return nil, http.StatusBadRequest, errors.New("Error querying database: " + err.Error())
	}
	for rows.Next() {
		post := entity.Post{}
		if err := rows.Scan(
			&post.PostID,
			&post.UserID,
			&post.UserName,
			&post.Content,
			&post.Likes,
			&post.CommentsCount,
			&post.CreatedAt,
		); err != nil {
			return nil, http.StatusInternalServerError, errors.New("Error scanning row: " + err.Error())
		}
		posts = append(posts, post)
	}
	return posts, http.StatusOK, nil
}

func (r *PostRepository) AddPostReaction(input entity.PostReaction) (int, error) {
	query := "SELECT is_like FROM likes_post WHERE user_id = $1 and post_id = $2;"
	prep, err := r.db.Prepare(query)
	if err != nil {
		return http.StatusInternalServerError, errors.New("error preparing query" + err.Error())
	}
	defer prep.Close()
	var like uint
	if err := prep.QueryRow(input.UserID, input.PostID).Scan(&like); err != nil {
		return http.StatusInternalServerError, errors.New("error scanning row: " + err.Error())
	}
	if like == 1 {
		query := "DELETE FROM likes_post WHERE user_id = $1 and post_id = $2;"
		prep, err := r.db.Prepare(query)
		if err != nil {
			return http.StatusInternalServerError, errors.New("error preparing query" + err.Error())
		}
		defer prep.Close()
		if _, err := prep.Exec(input.UserID, input.PostID); err != nil {
			return http.StatusInternalServerError, errors.New("error deleting like" + err.Error())
		}
	} else {
		query := "UPDATE likes_post SET is_like = $1 WHERE user_id = $2 and post_id = $3;"
		prep, err := r.db.Prepare(query)
		if err != nil {
			return http.StatusInternalServerError, errors.New("error preparing query" + err.Error())
		}
		defer prep.Close()
		if _, err := prep.Exec(1, input.UserID, input.PostID); err != nil {
			return http.StatusInternalServerError, errors.New("error adding reaction" + err.Error())
		}
	}

	return http.StatusOK, nil
}

func (r *PostRepository) getCommentsByPostID(postID uint) ([]entity.Comment, int, error) {
	query := `
	SELECT 
    c.comment_id,
    c.user_id,
    u.username AS user_name,
    c.content AS comment_content,
		(SELECT COUNT(like_id) 
				FROM likes_comment lc 
				WHERE lc.comment_id = c.comment_id 
				AND lc.is_like = 1) AS likes_count,
		c.created_at
	FROM 
		comments c
		LEFT JOIN users u ON c.user_id = u.user_id
	WHERE 
		c.post_id = ?
	ORDER BY 
    c.created_at DESC;
	`
	prep, err := r.db.Prepare(query)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("Error Preparing query: " + err.Error())
	}
	defer prep.Close()

	rows, err := prep.Query(postID)
	if err != nil {
		return nil, http.StatusBadRequest, errors.New("Error querying database: " + err.Error())
	}
	comments := []entity.Comment{}

	for rows.Next() {
		comment := entity.Comment{}
		if err := rows.Scan(
			&comment.CommentID,
			&comment.UserID,
			&comment.UserName,
			&comment.Content,
			&comment.Likes,
			&comment.CreatedAt,
		); err != nil {
			return nil, http.StatusInternalServerError, errors.New("Error scanning row: " + err.Error())
		}

		comments = append(comments, comment)
	}
	return comments, http.StatusOK, nil
}
