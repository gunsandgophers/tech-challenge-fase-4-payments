package mocks

import (
	"tech-challenge-fase-1/internal/core/dtos"

	"github.com/stretchr/testify/mock"
)

type PaymentGatewayMock struct {
	mock.Mock
}

func (r *PaymentGatewayMock) Execute(
	orderID string,
	amount float64,
	method dtos.MethodType,
) (*dtos.PaymentRequestDTO, error) {
	args := r.Called(orderID, amount, method)
	return args.Get(0).(*dtos.PaymentRequestDTO), args.Error(1)
}
