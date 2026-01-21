package services

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"smart-learning-english/backend/internal/core/domain"
	"smart-learning-english/backend/internal/core/ports"
)

type AuthService struct {
	repo       ports.UserRepository
	jwtSecret  string
	tokenExpiry time.Duration
}

func NewAuthService(repo ports.UserRepository, jwtSecret string) ports.AuthService {
	return &AuthService{
		repo:        repo,
		jwtSecret:   jwtSecret,
		tokenExpiry: 24 * time.Hour,
	}
}

func (s *AuthService) Register(ctx context.Context, req ports.RegisterRequest) (*ports.AuthResponse, error) {
	// 1. Check if user exists
	existingUser, err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// 2. Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 3. Create User
	newUser := &domain.User{
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		FullName:     req.FullName,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := s.repo.Create(ctx, newUser); err != nil {
		return nil, err
	}

	// 4. Generate Token
	token, err := s.generateToken(newUser.ID)
	if err != nil {
		return nil, err
	}

	return &ports.AuthResponse{
		Token: token,
		User: struct {
			ID       string `json:"id"`
			Email    string `json:"email"`
			FullName string `json:"full_name"`
		}{
			ID:       newUser.ID,
			Email:    newUser.Email,
			FullName: newUser.FullName,
		},
	}, nil
}

func (s *AuthService) Login(ctx context.Context, req ports.LoginRequest) (*ports.AuthResponse, error) {
	// 1. Get User
	user, err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid credentials")
	}

	// 2. Compare Password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	// 3. Generate Token
	token, err := s.generateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &ports.AuthResponse{
		Token: token,
		User: struct {
			ID       string `json:"id"`
			Email    string `json:"email"`
			FullName string `json:"full_name"`
		}{
			ID:       user.ID,
			Email:    user.Email,
			FullName: user.FullName,
		},
	}, nil
}

func (s *AuthService) generateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(s.tokenExpiry).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}
