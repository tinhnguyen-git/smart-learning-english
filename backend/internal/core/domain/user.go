package domain

import "time"

// User represents a user in the system
type User struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"` // Never return password hash in JSON
	FullName      string    `json:"full_name"`
	IsPremium     bool      `json:"is_premium"`
	PremiumExpiry *time.Time `json:"premium_expiry,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
