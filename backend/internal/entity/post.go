package entity

import "time"

type Post struct {
	PostID        uint      `json:"post_id"`
	UserID        uint      `json:"user_id"`
	GroupID       uint      `json:"group_id"`
	UserName      string    `json:"username"`
	Content       string    `json:"content"`
	Image_url     string    `json:"image_url"`
	Privacy       string    `json:"privacy"`
	Likes         uint      `json:"likes"`
	Comments      []Comment `json:"comments"`
	CommentsCount uint      `json:"comments_count"`
	CreatedAt     time.Time `json:"created_at"`
}

type PostReaction struct {
	UserID uint `json:"user_id"`
	PostID uint `json:"post_id"`
}
