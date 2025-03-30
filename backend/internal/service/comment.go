package service

import (
	"errors"
	"net/http"
	"socialNetwork/internal/entity"
	"socialNetwork/internal/repository"
	"strings"
)

type CommentService struct {
	commentRepo repository.Comment
}

func newCommentService(commentRepo repository.Comment) *CommentService {
	return &CommentService{commentRepo: commentRepo}
}

func (s *CommentService) CreateComment(input entity.Comment) (int, error) {
	if strings.TrimSpace(input.Content) == "" {
		return http.StatusBadRequest, errors.New("invalid content")
	} else if input.PostID == 0 {
		return http.StatusBadRequest, errors.New("invalid post ID")
	}
	return s.commentRepo.CreateComment(input)
}

func (s *CommentService) AddCommentReaction(input entity.CommentReaction) (int, error) {
	return s.commentRepo.AddCommentReaction(input)
}
