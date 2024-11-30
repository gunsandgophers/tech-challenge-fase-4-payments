package payments

import (
	"errors"
	"tech-challenge-fase-1/internal/core/entities"
	"tech-challenge-fase-1/internal/tests/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewGetPaymentStatusUseCase(t *testing.T) {
	// Arrange
	orderID := uuid.NewString()
	amount := 45.5
	repo := &mocks.PaymentRepositoryMock{}
	payment := entities.CreatePayment(orderID, amount)
	repo.On("FindPaymentByOrderID", orderID).Return(payment, nil)
	uc := NewGetPaymentStatusUseCase(repo)

	// Act
	dto, err := uc.Execute(orderID)

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, orderID, dto.OrderId)
	assert.Equal(t, entities.ORDER_PAYMENT_AWAITING_PAYMENT.String(), dto.PaymentStatus)
}

func TestNewGetPaymentStatusUseCase_NotFound(t *testing.T) {
	// Arrange
	orderID := uuid.NewString()
	repo := &mocks.PaymentRepositoryMock{}
	var payment = &entities.Payment{}
	repo.On("FindPaymentByOrderID", orderID).Return(payment, errors.New("Not Found"))
	uc := NewGetPaymentStatusUseCase(repo)

	// Act
	_, err := uc.Execute(orderID)

	//Assert
	assert.NotNil(t, err)
}
