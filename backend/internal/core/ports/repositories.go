package ports

import (
	"context"

	"smart-learning-english/backend/internal/core/domain"
)

// UserRepository defines the methods for user data access
type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	GetByID(ctx context.Context, id string) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
}

// StoryRepository defines methods for story data persistence
type StoryRepository interface {
	Create(ctx context.Context, story *domain.Story) error
	// Could add List(), GetByID() etc later
}
