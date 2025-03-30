package controller

import (
	"encoding/json"
	"net/http"
	"socialNetwork/internal/entity"
)

func (h *Handler) createComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.errorHandler(w, r, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	input := entity.Comment{}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}
	if status, err := h.service.Comment.CreateComment(input); err != nil {
		h.errorHandler(w, r, status, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) commentReaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodPut {
		h.errorHandler(w, r, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var input entity.CommentReaction
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}
	if status, err := h.service.Comment.AddCommentReaction(input); err != nil {
		h.errorHandler(w, r, status, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}
