package controller

import (
	"encoding/json"
	"net/http"
	"socialNetwork/internal/entity"
)


func (h *Handler) getALLPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.errorHandler(w, r, http.StatusMethodNotAllowed, "Methode not allowed")
		return
	}
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	posts, status, err := h.service.Post.GetAllPosts(limitStr, offsetStr)
	if err != nil {
		h.errorHandler(w, r, status, err.Error())
		return
	}
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) getPostbyID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.errorHandler(w, r, http.StatusMethodNotAllowed, "Methode not allowed")
		return
	}

	strPostID := r.URL.Path[len("/api/post/"):]
	post, status, err := h.service.Post.GetPostByID(strPostID)
	if err != nil {
		h.errorHandler(w, r, status, err.Error())
		return
	}

	if err := json.NewEncoder(w).Encode(post); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) createPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.errorHandler(w, r, http.StatusMethodNotAllowed, "Methode not allowed")
		return
	}

	var input entity.Post
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	postID, status, err := h.service.Post.CreatePost(input)
	if err != nil {
		h.errorHandler(w, r, status, err.Error())
		return
	}

	if err := json.NewEncoder(w).Encode(map[string]uint{
		"post_id": postID,
	}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) getMyPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.errorHandler(w, r, http.StatusMethodNotAllowed, "Methode not allowed")
		return
	}
	strUserID := r.URL.Path[len("/api/profile/posts/"):]

	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	posts, status, err := h.service.Post.GetAllByUserID(strUserID, limitStr, offsetStr)
	if err != nil {
		h.errorHandler(w, r, status, err.Error())
		return
	}
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) postReaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodPut {
		h.errorHandler(w, r, http.StatusMethodNotAllowed, "Methode not allowed")
		return
	}

	var input entity.PostReaction
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	
	if status, err := h.service.Post.AddPostReaction(input); err != nil {
		h.errorHandler(w, r, status, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}
