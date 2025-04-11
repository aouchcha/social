package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"socialNetwork/internal/entity"
)

func (h *Handler) GetGroups(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.errorHandler(w, r, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	Groups, status, err := h.service.Group.GetGroups(1)
	if err != nil {
		h.errorHandler(w, r, status, err.Error())
		return
	}
	if err := json.NewEncoder(w).Encode(struct {
		Groups []entity.Groups
	}{
		Groups: Groups,
	}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) CreateGroupe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.errorHandler(w, r, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	var input entity.CreateGroupe
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(input)
	defer r.Body.Close()
	GroupeID, status, err := h.service.Group.CreateGroupe(input)
	if err != nil {
		h.errorHandler(w, r, status, err.Error())
		return
	}
	if err := json.NewEncoder(w).Encode(struct {
		GroupeID   int    `json:"group_id"`
		Message    string `json:"message"`
		StatusCode int    `json:"status"`
	}{
		GroupeID:   GroupeID,
		Message:    "Groupe Created Sucsesfully",
		StatusCode: http.StatusOK,
	}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) GroupeInvite(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.errorHandler(w, r, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	var input entity.GroupInvites
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, "Bad Request")
		return
	}

	defer r.Body.Close()
	status, err := h.service.Group.CheckInvite(input)
	if err != nil {
		fmt.Println("Error ===>", err)
		h.errorHandler(w, r, status, err.Error())
		return
	}
	if err := json.NewEncoder(w).Encode(struct {
		Message    string `json:"message"`
		StatusCode int    `json:"status"`
	}{
		Message:    "The Invite Sent Sucsesfully",
		StatusCode: http.StatusCreated,
	}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) GroupeInviteUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.errorHandler(w, r, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	var input entity.GroupInvites
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, "Bad Request")
		return
	}

	defer r.Body.Close()
	status, err := h.service.Group.UpdateInvite(input)
	if err != nil {
		h.errorHandler(w, r, status, err.Error())
		return
	}
	if err := json.NewEncoder(w).Encode(struct {
		Message    string `json:"message"`
		StatusCode int    `json:"status"`
	}{
		Message:    "The Invite Sent Sucsesfully",
		StatusCode: http.StatusCreated,
	}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) GroupeRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.errorHandler(w, r, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	var input entity.GroupRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, "Bad Request")
		return
	}

	defer r.Body.Close()
	status, err := h.service.Group.CheckRequest(input)
	if err != nil {
		h.errorHandler(w, r, status, err.Error())
		return
	}
	if err := json.NewEncoder(w).Encode(struct {
		Message    string `json:"message"`
		StatusCode int    `json:"status"`
	}{
		Message:    "The Request Sent Sucsesfully",
		StatusCode: http.StatusCreated,
	}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) GroupeRequestUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.errorHandler(w, r, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	var input entity.GroupRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, "Bad Request")
		return
	}

	defer r.Body.Close()
	status, err := h.service.Group.UpdateRequest(input)
	if err != nil {
		h.errorHandler(w, r, status, err.Error())
		return
	}
	if err := json.NewEncoder(w).Encode(struct {
		Message    string `json:"message"`
		StatusCode int    `json:"status"`
	}{
		Message:    "The Request Sent Sucsesfully",
		StatusCode: http.StatusCreated,
	}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.errorHandler(w, r, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	var input entity.CreateEvent
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		fmt.Println(err)
		h.errorHandler(w, r, http.StatusBadRequest, "Bad Request")
		return
	}
	fmt.Println(input)

	defer r.Body.Close()
	status, err := h.service.Group.CreateEvent(input)
	if err != nil {
		h.errorHandler(w, r, status, err.Error())
		return
	}
	if err := json.NewEncoder(w).Encode(struct {
		Message    string `json:"message"`
		StatusCode int    `json:"status"`
	}{
		Message:    "Event Created Sucsesfully",
		StatusCode: http.StatusOK,
	}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) GetAllEvents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.errorHandler(w, r, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	group_id := r.URL.Query().Get("groupid")

	Events, status, err := h.service.Group.GetAllEvents(group_id)
	if err != nil {
		h.errorHandler(w, r, status, err.Error())
		return
	}
	if err := json.NewEncoder(w).Encode(struct {
		Events []entity.CreateEvent
	}{
		Events: Events,
	}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.errorHandler(w, r, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	var input entity.EventUpdate
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, "Bad Request")
		return
	}
	status, err := h.service.Group.UpdateEvent(input)
	if err != nil {
		h.errorHandler(w, r, status, err.Error())
		return
	}
	if err := json.NewEncoder(w).Encode(struct {
		Message    string `json:"message"`
		StatusCode int    `json:"status"`
	}{
		Message:    "Event Updated Sucsesfully",
		StatusCode: http.StatusOK,
	}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) GetPostsOfGroup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.errorHandler(w, r, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	group_id := r.URL.Query().Get("groupid")
	posts, status, err := h.service.Group.GetAllPosts(group_id)
	if err != nil {
		h.errorHandler(w, r, status, err.Error())
		return
	}
	if err := json.NewEncoder(w).Encode(struct {
		Posts []entity.Post
	}{
		Posts: posts,
	}); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}
