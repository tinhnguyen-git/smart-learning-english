package http

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"smart-learning-english/backend/internal/core/ports"
)

type UserHandler struct {
	upgradeService ports.UpgradeService
	jwtSecret      string
}

func NewUserHandler(upgradeService ports.UpgradeService, jwtSecret string) *UserHandler {
	return &UserHandler{
		upgradeService: upgradeService,
		jwtSecret:      jwtSecret,
	}
}

func (h *UserHandler) Upgrade(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract Token
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing authorization header", http.StatusUnauthorized)
		return
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Parse Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(h.jwtSecret), nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		http.Error(w, "Invalid token claims", http.StatusUnauthorized)
		return
	}

	userID, ok := claims["sub"].(string)
	if !ok {
		http.Error(w, "Invalid user ID in token", http.StatusUnauthorized)
		return
	}

	// Perform Upgrade
	if err := h.upgradeService.UpgradeUser(r.Context(), userID); err != nil {
		// Differentiate between "user already premium" vs "payment failed" vs "db error" ideally.
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User upgraded to premium successfully",
	})
}
