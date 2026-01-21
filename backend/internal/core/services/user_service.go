package services

import (
	"context"
	"fmt"
	"time"

	"smart-learning-english/backend/internal/core/ports"
)

type UserService struct {
	userRepo ports.UserRepository
	payment  ports.PaymentService
}

func NewUserService(userRepo ports.UserRepository, payment ports.PaymentService) ports.UpgradeService {
	return &UserService{
		userRepo: userRepo,
		payment:  payment,
	}
}

func (s *UserService) UpgradeUser(ctx context.Context, userID string) error {
	// 1. Get User
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return err
	}
	if user == nil {
		return fmt.Errorf("user not found")
	}

	if user.IsPremium {
		return fmt.Errorf("user is already premium")
	}

	// 2. Process Payment (Mock cost: $9.99)
	if err := s.payment.ProcessPayment(ctx, 9.99, "USD"); err != nil {
		return fmt.Errorf("payment failed: %w", err)
	}

	// 3. Update User Status
	now := time.Now()
	expiry := now.AddDate(1, 0, 0) // 1 Year subscription
	user.IsPremium = true
	user.PremiumExpiry = &expiry
	user.UpdatedAt = now

	if err := s.userRepo.Update(ctx, user); err != nil {
		return err
	}

	return nil
}
