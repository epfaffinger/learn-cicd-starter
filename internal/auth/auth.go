package auth

import (
	"errors"
	"net/http"
	"strings"
)

var ErrNoAuthHeaderIncluded = errors.New("no authorization header included")

// GetAPIKey -
// go
func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("X-Api-Key")
	if authHeader == "" {
		return "", ErrNoAuthHeaderIncluded
	}
	splitAuth := strings.SplitN(authHeader, " ", 2)
	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" || strings.TrimSpace(splitAuth[1]) == "" {
		return "", errors.New("malformed authorization header")
	}
	return strings.TrimSpace(splitAuth[1]), nil
}
