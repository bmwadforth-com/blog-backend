package util

import (
	"errors"
	"net/http"
	"strings"
)

func GetBearerToken(req *http.Request) (string, error) {
	authHeader := req.Header.Get("Authorization")
	bearer := strings.Split(authHeader, "Bearer ")
	if len(bearer) == 2 {
		return bearer[1], nil
	} else {
		return "", errors.New("auth header is malformed")
	}
}
