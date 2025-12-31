package helpers

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	id, err := uuid.NewUUID()
	if err != nil {
		log.Println(err)
		return ""
	}
	return id.String()
}

func GenerateRandomString(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("length must be greater than 0")
	}
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func HashString(input string) string {
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}

func IsValidHTTPHeader(header string) bool {
	for _, c := range header {
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '-' || c == '_' || c == ' ' || c == ':') {
			return false
		}
	}
	return true
}

func GetHTTPStatusCodeString(code int) string {
	switch code {
	case http.StatusOK:
		return "OK"
	case http.StatusNotFound:
		return "Not Found"
	case http.StatusInternalServerError:
		return "Internal Server Error"
	default:
		return fmt.Sprintf("Unknown Status Code: %d", code)
	}
}

func IsTimeout(err error) bool {
	if err, ok := err.(net.Error); ok && err.Timeout() {
		return true
	}
	return false
}

func GetElapsedTime(start time.Time) time.Duration {
	return time.Since(start)
}

func SanitizeString(input string) string {
	return strings.TrimSpace(input)
}

func ValidateEmail(email string) bool {
	if len(email) < 3 {
		return false
	}
	if !strings.Contains(email, "@") {
		return false
	}
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	if len(parts[0]) < 1 || len(parts[1]) < 3 {
		return false
	}
	return true
}