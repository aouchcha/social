package jwt

import (
	"github.com/gofrs/uuid"
)
// duplicated function ===>utils/sessionTools
func GeneratorJWT(val string) (string, error) {
	jwt, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	return jwt.String(), nil
}
