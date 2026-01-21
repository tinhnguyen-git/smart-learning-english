package ports

import (
	"context"

	"smart-learning-english/backend/internal/core/domain"
)

// RegisterRequest represents the data needed to register a new user
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}

// LoginRequest represents the data needed to login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AuthResponse represents the response after successful auth
type AuthResponse struct {
	Token string `json:"token"`
	User  struct {
		ID       string `json:"id"`
		Email    string `json:"email"`
		FullName string `json:"full_name"`
	} `json:"user"`
}

// AuthService defines the methods for authentication business logic
type AuthService interface {
	Register(ctx context.Context, req RegisterRequest) (*AuthResponse, error)
	Login(ctx context.Context, req LoginRequest) (*AuthResponse, error)
}

// ScraperService defines methods for scraping content
type ScraperService interface {
	ScrapeAndSave(ctx context.Context, url string) (*domain.Story, error)
}

// PaymentService defines methods for processing payments
type PaymentService interface {
	ProcessPayment(ctx context.Context, amount float64, currency string) error
}

// UpgradeService defines methods for upgrading users
type UpgradeService interface {
	UpgradeUser(ctx context.Context, userID string) error
}
