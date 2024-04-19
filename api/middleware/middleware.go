package middleware

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("secret-key")

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract JWT token from request header
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte("Unauthorized access: Missing token"))
			return
		}
		// Verify JWT token
		if err := verifyToken(tokenString); err != nil {
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte("Unauthorized access: Invalid token"))
			return
		}
		// Proceed to the next handler if token is valid
		next.ServeHTTP(w, r)
	})
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
