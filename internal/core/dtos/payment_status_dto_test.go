package dtos

import (
	"tech-challenge-fase-1/internal/core/entities"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewPaymentStatusFromEntity(t *testing.T) {
	//Arrange
	payment := entities.CreatePayment(
		uuid.NewString(),
		40.9,
	)
	//Act
	dto := NewPaymentStatusDTOFromEntity(payment)
	//Assert
	assert.NotNil(t, dto)
	assert.Equal(t, payment.GetOrderID(), dto.OrderId)
	assert.Equal(t, payment.GetPaymentStatus().String(), dto.PaymentStatus)
}
