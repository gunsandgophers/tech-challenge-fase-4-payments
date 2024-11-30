package mocks

import (
	"tech-challenge-fase-1/internal/core/entities"
	"github.com/stretchr/testify/mock"
)

type PaymentRepositoryMock struct {
	mock.Mock
}

func (r *PaymentRepositoryMock) Insert(payment *entities.Payment) error {
	args := r.Called(payment)
	return args.Error(0)
}

func (r *PaymentRepositoryMock) Update(payment *entities.Payment) error {
	args := r.Called(payment)
	return args.Error(0)
}

func (r *PaymentRepositoryMock) FindPaymentByOrderID(orderID string) (*entities.Payment, error) {
	args := r.Called(orderID)
	return args.Get(0).(*entities.Payment), args.Error(1)
}
