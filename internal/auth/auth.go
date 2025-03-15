package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey exacts an API Key from the headers of an HTTP request
// Ex : Authorization: APIKey {insert apikey here}

func GetAPIKey(header http.Header) (string, error) {
	val := header.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentifaction info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of the auth header")
	}
	return vals[1], nil
}
