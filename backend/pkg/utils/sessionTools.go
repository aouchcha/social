package utils

import (
	"net/http"

	"github.com/gofrs/uuid"
)

func GenerateToken() (string, int, error) {
	token, err := uuid.NewV7()
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	return token.String(), http.StatusCreated, nil
}
