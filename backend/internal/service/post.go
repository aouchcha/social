package service

import (
	"errors"
	"log"
	"net/http"
	"socialNetwork/internal/entity"
	"socialNetwork/internal/repository"
	"socialNetwork/pkg/utils"
	"strconv"
	"strings"
)

type PostService struct {
	postRepo repository.Post
}

func newPostService(postRepo repository.Post) *PostService {
	return &PostService{
		postRepo: postRepo,
	}
}
func (s *PostService) GetAllPosts(limit, offset string) ([]entity.Post, int, error) {
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return nil, http.StatusBadRequest, errors.New("invalid limit")
	}
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		return nil, http.StatusBadRequest, errors.New("invalid offset")
	}
	return s.postRepo.GetAllPosts(limitInt, offsetInt)
}

func (s *PostService) GetPostByID(strPostID string) (entity.Post, int, error) {
	postId, err := strconv.Atoi(strPostID)
	if err != nil || postId <= 0 {
		return entity.Post{}, http.StatusBadRequest, errors.New("invalid post id")
	}

	return s.postRepo.GetPostByID(uint(postId))
}

func (s *PostService) CreatePost(input entity.Post) (uint, int, error) {
	input.UserID = uint(input.UserID)

	if strings.TrimSpace(input.Content) == "" || len(strings.TrimSpace(input.Content)) > 10000 {
		return 0, http.StatusBadRequest, errors.New("size of text must be between 1 and 10000")
	}

	allowedTypes := map[string]bool{"pb": true, "pr": true, "ap": true}
	if !allowedTypes[input.Privacy] {
		return 0, http.StatusBadRequest, errors.New("invalid privacy option")
	}

	if input.Image_url != "" {
		var err error
		var status int
		input.Image_url, status, err = utils.ParseImage(input.Image_url)
		if err != nil || status != http.StatusOK {
			return 0, status, err
		}
	}

	postID, status, err := s.postRepo.CreatePost(input)
	if err != nil {
		log.Println(err)
		return 0, status, err
	}
	return postID, http.StatusOK, nil
}

func (s *PostService) GetAllByUserID(strUserID, limitStr, offsetStr string) ([]entity.Post, int, error) {
	userID, err := strconv.Atoi(strUserID)
	if err != nil || userID <= 0 {
		return nil, http.StatusNotFound, errors.New("invalid user ID")
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return nil, http.StatusBadRequest, errors.New("invalid limit")
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return nil, http.StatusBadRequest, errors.New("invalid offset")
	}
	return s.postRepo.GetAllByUserID(uint(userID), limit, offset)
}

func (s *PostService) AddPostReaction(input entity.PostReaction) (int, error) {
	return s.postRepo.AddPostReaction(input)
}
