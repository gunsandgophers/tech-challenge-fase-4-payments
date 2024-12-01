package payments

import (
	"errors"
	"tech-challenge-fase-1/internal/core/entities"
	"tech-challenge-fase-1/internal/tests/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPaymentOrderUseCase_Paid(t *testing.T) {
	// Arrange
	orderID := uuid.NewString()
	amount := 45.5
	repo := &mocks.PaymentRepositoryMock{}
	payment := entities.CreatePayment(orderID, amount)
	repo.On("FindPaymentByOrderID", orderID).Return(payment, nil)
	repo.On("Update", mock.Anything).Return(nil)
	uc := NewPaymentOrderUseCase(repo)

	// Act
	err := uc.Execute(orderID, entities.ORDER_PAYMENT_PAID.String())

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, entities.ORDER_PAYMENT_PAID, payment.GetPaymentStatus())
}

func TestPaymentOrderUseCase_Rejected(t *testing.T) {
	// Arrange
	orderID := uuid.NewString()
	amount := 45.5
	repo := &mocks.PaymentRepositoryMock{}
	payment := entities.CreatePayment(orderID, amount)
	repo.On("FindPaymentByOrderID", orderID).Return(payment, nil)
	repo.On("Update", mock.Anything).Return(nil)
	uc := NewPaymentOrderUseCase(repo)

	// Act
	err := uc.Execute(orderID, entities.ORDER_PAYMENT_REJECTED.String())

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, entities.ORDER_PAYMENT_REJECTED, payment.GetPaymentStatus())
}

func TestPaymentOrderUseCase_InvalidPaymentStatus(t *testing.T) {
	// Arrange
	orderID := uuid.NewString()
	amount := 45.5
	repo := &mocks.PaymentRepositoryMock{}
	payment := entities.CreatePayment(orderID, amount)
	repo.On("FindPaymentByOrderID", orderID).Return(payment, nil)
	repo.On("Update", mock.Anything).Return(nil)
	uc := NewPaymentOrderUseCase(repo)

	// Act
	err := uc.Execute(orderID, entities.ORDER_PAYMENT_AWAITING_PAYMENT.String())

	// Assert
	assert.NotNil(t, err)
}

func TestPaymentOrderUseCase_NotFound(t *testing.T) {
	// Arrange
	orderID := uuid.NewString()
	repo := &mocks.PaymentRepositoryMock{}
	var payment *entities.Payment
	repo.On("FindPaymentByOrderID", orderID).Return(payment, errors.New("Not Found"))
	repo.On("Update", mock.Anything).Return(nil)
	uc := NewPaymentOrderUseCase(repo)

	// Act
	err := uc.Execute(orderID, entities.ORDER_PAYMENT_PAID.String())

	// Assert
	assert.NotNil(t, err)
}

func TestPaymentOrderUseCase_InvalidPayment(t *testing.T) {
	// Arrange
	orderID := uuid.NewString()
	amount := 45.5
	repo := &mocks.PaymentRepositoryMock{}
	payment := entities.CreatePayment(orderID, amount)
	payment.PaymentReceived()
	repo.On("FindPaymentByOrderID", orderID).Return(payment, nil)
	repo.On("Update", mock.Anything).Return(nil)
	uc := NewPaymentOrderUseCase(repo)

	// Act
	err := uc.Execute(orderID, entities.ORDER_PAYMENT_PAID.String())

	// Assert
	assert.NotNil(t, err)
}
