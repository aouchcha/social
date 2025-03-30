package utils

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
)

func ParseImage(image_url string) (string, int, error) {
	parts := strings.SplitN(image_url, ",", 2)
	if len(parts) != 2 || !strings.HasPrefix(parts[0], "data:image/") {
		return "", http.StatusBadRequest, errors.New("invalid image format")
	}

	mime := parts[0]
	var ext string
	switch {
	case strings.Contains(mime, "png"):
		ext = ".png"
	case strings.Contains(mime, "jpeg") || strings.Contains(mime, "jpg"):
		ext = ".jpeg"
	case strings.Contains(mime, "gif"):
		ext = ".gif"
	default:
		return "", http.StatusBadRequest, errors.New("invalid image format")
	}

	data, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return "", http.StatusBadRequest, errors.New("Failed to decode image: " + err.Error())
	}

	const maxSize = 20 * 1024 * 1024
	if len(data) > maxSize {
		return "", http.StatusBadRequest, errors.New("image size exceeds 20MB")
	}
	imagePath := fmt.Sprintf("uploads/%s%s", uuid.New().String(), ext)
	if err := os.MkdirAll("uploads", 0755); err != nil {
		return "", http.StatusInternalServerError, errors.New("Failed to create uploads directory: " + err.Error())
	}
	if err := os.WriteFile(imagePath, data, 0644); err != nil {
		return "", http.StatusInternalServerError, errors.New("Failed to save image: " + err.Error())
	}

	return imagePath, http.StatusOK, nil

}
