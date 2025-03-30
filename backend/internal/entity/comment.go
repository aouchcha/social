package entity

import "time"

type Comment struct {
	CommentID uint      `json:"comment_id"`
	PostID    uint      `json:"post_id"`
	UserID    uint      `json:"user_id"`
	UserName  string    `json:"username"`
	Content   string    `json:"content"`
	Likes     uint      `json:"likes"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentReaction struct {
	UserID    uint `json:"user_id"`
	CommentID uint `json:"comment_id"`
}
