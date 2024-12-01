package services

import (
	"tech-challenge-fase-1/internal/core/dtos"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestMercadoPagoGatewayExecute(t *testing.T) {

	gateway := NewMercadoPagoGateway()

	orderID := uuid.NewString()
	amount := 13.37
	method := dtos.CREDIT_CARD

	response, err := gateway.Execute(orderID, amount, method)

	assert.Nil(t, err)
	assert.Equal(t, orderID, response.OrderId)
}
