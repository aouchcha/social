package controller

import (
	"encoding/json"
	"net/http"

	"socialNetwork/internal/entity"
)

func (h *Handler) responseJSON(w http.ResponseWriter, response entity.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Code)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
