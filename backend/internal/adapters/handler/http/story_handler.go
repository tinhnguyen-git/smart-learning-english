package http

import (
	"encoding/json"
	"net/http"

	"smart-learning-english/backend/internal/core/ports"
)

type StoryHandler struct {
	service ports.ScraperService
}

func NewStoryHandler(service ports.ScraperService) *StoryHandler {
	return &StoryHandler{service: service}
}

type ScrapeRequest struct {
	URL string `json:"url"`
}

func (h *StoryHandler) Scrape(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ScrapeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.URL == "" {
		http.Error(w, "Missing 'url' field", http.StatusBadRequest)
		return
	}

	story, err := h.service.ScrapeAndSave(r.Context(), req.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(story)
}
