package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
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
	// Launch headless browser
	l := launcher.New().Headless(true).MustLaunch()
	browser := rod.New().ControlURL(l).MustConnect()
	defer browser.MustClose()

	// Navigate to page and wait for content
	page := browser.MustPage(url)
	page.MustWaitLoad()

	// Wait for dynamic content (JavaScript rendering)
	time.Sleep(3 * time.Second)

	// Extract title
	var title string
	titleEl, err := page.Element("h1")
	if err == nil && titleEl != nil {
		title, _ = titleEl.Text()
	}
	if title == "" {
		title = "Untitled Story"
	}

	// Extract content - try common content selectors
	var contentBuilder strings.Builder
	contentSelectors := []string{
		"div.chapter-content p",      // Common novel site
		"div.content p",              // Generic
		"article p",                  // Blog style
		"div.text-content p",         // Alternative
		"div.story-content p",        // Story sites
		"div#chapter-content p",      // ID based
		"p",                          // Fallback to all paragraphs
	}

	for _, selector := range contentSelectors {
		elements, err := page.Elements(selector)
		if err == nil && len(elements) > 0 {
			for _, el := range elements {
				text, _ := el.Text()
				text = strings.TrimSpace(text)
				if len(text) > 30 { // Filter short/navigation text
					contentBuilder.WriteString(text)
					contentBuilder.WriteString("\n\n")
				}
			}
			if contentBuilder.Len() > 100 {
				break // Found enough content
			}
		}
	}

	content := contentBuilder.String()
	if content == "" {
		// Fallback: get all visible text
		body, err := page.Element("body")
		if err == nil && body != nil {
			content, _ = body.Text()
		}
	}

	if strings.TrimSpace(content) == "" {
		return nil, fmt.Errorf("no content found at url")
	}

	story := &domain.Story{
		Title:     strings.TrimSpace(title),
		Content:   strings.TrimSpace(content),
		SourceURL: url,
		CreatedAt: time.Now(),
	}

	// Save to DB
	if err := s.repo.Create(ctx, story); err != nil {
		return nil, fmt.Errorf("failed to save story: %w", err)
	}

	return story, nil
}
