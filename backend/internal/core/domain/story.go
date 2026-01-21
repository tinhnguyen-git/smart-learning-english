package domain

import "time"

// Story represents a reading material for users
type Story struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	SourceURL string    `json:"source_url"`
	CreatedAt time.Time `json:"created_at"`
}
