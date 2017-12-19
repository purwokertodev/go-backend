package middleware

import (
	"encoding/base64"
	"net/http"
	"strings"
)

type config struct {
	username string
	password string
}

func NewConfig(username, password string) *config {
	return &config{username: username, password: password}
}

func BasicAuth(config *config, next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		validate := func(username, password string) bool {
			if config.username == username && config.password == password {
				return true
			}
			return false
		}

		auth := strings.SplitN(req.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			http.Error(res, "authorization failed", http.StatusUnauthorized)
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !validate(pair[0], pair[1]) {
			http.Error(res, "authorization failed", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(res, req)
	})
}
