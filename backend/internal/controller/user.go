package controller

import (
	"encoding/json"
	"net/http"

	"socialNetwork/internal/entity"
	"socialNetwork/internal/repository"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.errorHandler(w, r, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	UserInfo := entity.UserProfile{}
	if err := json.NewDecoder(r.Body).Decode(&UserInfo); err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, "Invalid Credentials")
		return
	}

	defer r.Body.Close()

	status, err := h.service.ServiceSingUp(UserInfo)
	if err != nil {
		h.errorHandler(w, r, status, err.Error())
		return
	}
	h.responseJSON(w, entity.Response{Msg: "signUp is successfully", Code: status})
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.errorHandler(w, r, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	UserProfile := entity.UserProfile{}
	if err := json.NewDecoder(r.Body).Decode(&UserProfile); err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, "Invalid Credentials")
		return
	}

	defer r.Body.Close()

	session, setupInfo, status, err := h.service.ServiceLogIn(UserProfile)
	if err != nil {
		h.errorHandler(w, r, status, err.Error())
		return
	}
	h.responseJSON(w, entity.Response{Msg: "Logged in successfully", Code: status, Session: session, UserID: uint(setupInfo.Id)})
}

func (h *Handler) signOut(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.errorHandler(w, r, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	session := entity.Session{}
	if err := json.NewDecoder(r.Body).Decode(&session); err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, "Invalid Credentials")
		return
	}

	defer r.Body.Close()

	sessions := repository.SessionRepository{}
	if err := sessions.DeleteSessionByUserID(session); err != nil {
		h.errorHandler(w, r, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	h.responseJSON(w, entity.Response{Msg: "Sign Out in successfully", Code: http.StatusOK})
}
