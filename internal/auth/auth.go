package auth

import (
	"context"
	"encoding/base64"
	"errors"
	"net/http"
	"strings"
)

type User struct {
	ID       int
	Password string
}

var users = map[string]User{
	"user1": {ID: 1, Password: "password1"},
	"user2": {ID: 2, Password: "password2"},
}

func Authenticate(w http.ResponseWriter, r *http.Request) (int, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return 0, errors.New("unauthorized")
	}

	if !strings.HasPrefix(authHeader, "Basic ") {
		return 0, errors.New("unauthorized")
	}

	encodedCredentials := strings.TrimPrefix(authHeader, "Basic ")
	decodedCredentials, err := base64.StdEncoding.DecodeString(encodedCredentials)
	if err != nil {
		return 0, errors.New("unauthorized")
	}

	credentials := strings.SplitN(string(decodedCredentials), ":", 2)
	if len(credentials) != 2 {
		return 0, errors.New("unauthorized")
	}

	username := credentials[0]
	password := credentials[1]

	if user, userExists := users[username]; userExists && user.Password == password {
		return user.ID, nil
	}
	return 0, errors.New("unauthorized")
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := Authenticate(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "user_id", userID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
