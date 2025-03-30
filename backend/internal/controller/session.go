package controller

import (
	"encoding/json"
	"net/http"

	"socialNetwork/internal/entity"
)

func (h *Handler) isValidToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.errorHandler(w, r, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	session := entity.Session{} //==> var session entity.Session
	if err := json.NewDecoder(r.Body).Decode(&session); err != nil {
		h.errorHandler(w, r, http.StatusBadRequest, "Invalid Credentials")
		return
	}

	status, err := h.service.ServerLogout(session)
	if err != nil {
		h.errorHandler(w, r, status, err.Error())
		return
	}

	h.responseJSON(w, entity.Response{Msg: "signUp is successfully", Code: status})
}
