package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"smart-learning-english/backend/internal/core/domain"
	"smart-learning-english/backend/internal/core/ports"
)

type ScraperService struct {
	repo ports.StoryRepository
}

func NewScraperService(repo ports.StoryRepository) ports.ScraperService {
	return &ScraperService{repo: repo}
}

func (s *ScraperService) ScrapeAndSave(ctx context.Context, url string) (*domain.Story, error) {
	c := colly.NewCollector()

	var title string
	var contentBuilder strings.Builder

	// Attempt to find title
	c.OnHTML("h1", func(e *colly.HTMLElement) {
		if title == "" {
			title = strings.TrimSpace(e.Text)
		}
	})

	// Very basic content extraction: Grab all paragraphs
	// This will likely need refinement for specific domains to avoid footer/nav text
	c.OnHTML("p", func(e *colly.HTMLElement) {
		text := strings.TrimSpace(e.Text)
		if len(text) > 20 { // Filter out very short lines
			contentBuilder.WriteString(text)
			contentBuilder.WriteString("\n\n")
		}
	})

	err := c.Visit(url)
	if err != nil {
		return nil, fmt.Errorf("scraping failed: %w", err)
	}

	content := contentBuilder.String()
	if content == "" {
		return nil, fmt.Errorf("no content found at url")
	}

	if title == "" {
		title = "Untitled Story"
	}

	story := &domain.Story{
		Title:     title,
		Content:   content,
		SourceURL: url,
		CreatedAt: time.Now(),
	}

	// Save to DB
	if err := s.repo.Create(ctx, story); err != nil {
		return nil, fmt.Errorf("failed to save story: %w", err)
	}

	return story, nil
}
