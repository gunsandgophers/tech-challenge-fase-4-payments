package repositories

import "tech-challenge-fase-1/internal/core/entities"

type PaymentRepositoryInterface interface {
	Insert(payment *entities.Payment) error
	Update(order *entities.Payment) error
	FindPaymentByOrderID(orderID string) (*entities.Payment, error)
}
