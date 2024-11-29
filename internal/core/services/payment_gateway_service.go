package services

import (
	"tech-challenge-fase-1/internal/core/dtos"
)

type PaymentGatewayInterface interface {
	Execute(orderID string, amount float64, method dtos.MethodType) (*dtos.PaymentRequestDTO, error)
}
