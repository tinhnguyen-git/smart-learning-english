package main

import (
	"fmt"
	"log"
	"net/http"

	handler "smart-learning-english/backend/internal/adapters/handler/http"
	"smart-learning-english/backend/internal/adapters/repository/postgres"
	"smart-learning-english/backend/internal/core/services"
	"smart-learning-english/backend/pkg/config"
)

func main() {
	// 1. Load Configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("Warning: Could not load config: %v", err)
		// For now, we continue or fail depending on strictness.
		// In a real app, we might fail hard if DB is required.
		// log.Fatal(err)
	}

	// 2. Connect to Database (Handle errors more strictly for this part)
	if cfg == nil {
		log.Fatalf("Config load failed")
	}

	pool, err := postgres.NewConnection(cfg)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer pool.Close()
	log.Println("Connected to database successfully!")

	// 3. Initialize Core (Clean Arch Wiring)
	userRepo := postgres.NewUserRepository(pool)
	storyRepo := postgres.NewStoryRepository(pool)

	authService := services.NewAuthService(userRepo, cfg.JWTSecret)
	scraperService := services.NewScraperService(storyRepo)
	paymentService := services.NewPaymentService()
	userService := services.NewUserService(userRepo, paymentService)

	authHandler := handler.NewAuthHandler(authService)
	storyHandler := handler.NewStoryHandler(scraperService)
	userHandler := handler.NewUserHandler(userService, cfg.JWTSecret)

	// 4. Routes
	// 4. Routes
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to English Learning SuperApp Backend!")
	})
	mux.HandleFunc("/auth/register", authHandler.Register)
	mux.HandleFunc("/auth/login", authHandler.Login)
	mux.HandleFunc("/stories/scrape", storyHandler.Scrape)
	mux.HandleFunc("/users/upgrade", userHandler.Upgrade)

	port := ":8080"
	log.Printf("Server starting on port %s...", port)
	if err := http.ListenAndServe(port, enableCors(mux)); err != nil {
		log.Fatal(err)
	}
}

func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}
