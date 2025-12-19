package helpers

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/api-service/models"
)

func GetCurrentTime() time.Time {
	return time.Now().UTC()
}

func GetBearerToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header is missing")
	}
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", errors.New("invalid authorization header")
	}
	return parts[1], nil
}

func GetCurrentUserID(tokenString string) (uint, error) {
	token, err := models.ParseToken(tokenString)
	if err != nil {
		return 0, err
	}
	return token.UserID, nil
}

func HandleError(w http.ResponseWriter, err error) {
	log.Println(err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}