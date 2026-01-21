package services

import (
	"context"
	"log"

	"smart-learning-english/backend/internal/core/ports"
)

type PaymentService struct{}

func NewPaymentService() ports.PaymentService {
	return &PaymentService{}
}

func (s *PaymentService) ProcessPayment(ctx context.Context, amount float64, currency string) error {
	// Mock implementation: Always succeed
	log.Printf("Processing mock payment of %.2f %s... Success!", amount, currency)
	return nil
}
