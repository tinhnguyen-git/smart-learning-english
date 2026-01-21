package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"smart-learning-english/backend/internal/core/domain"
	"smart-learning-english/backend/internal/core/ports"
)

type StoryRepository struct {
	db *pgxpool.Pool
}

func NewStoryRepository(db *pgxpool.Pool) ports.StoryRepository {
	return &StoryRepository{db: db}
}

func (r *StoryRepository) Create(ctx context.Context, story *domain.Story) error {
	query := `
		INSERT INTO stories (title, content, source_url, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	err := r.db.QueryRow(ctx, query,
		story.Title,
		story.Content,
		story.SourceURL,
		story.CreatedAt,
	).Scan(&story.ID)

	if err != nil {
		return fmt.Errorf("failed to create story: %w", err)
	}
	return nil
}
