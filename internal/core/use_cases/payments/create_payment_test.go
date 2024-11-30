package payments

import (
	"errors"
	"tech-challenge-fase-1/internal/core/dtos"
	"tech-challenge-fase-1/internal/core/entities"
	"tech-challenge-fase-1/internal/tests/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreatePaymentUseCase_NewPayment(t *testing.T) {
	// Arrange
	orderID := uuid.NewString()
	amount := 45.5
	repo := &mocks.PaymentRepositoryMock{}
	var nilPayment *entities.Payment
	repo.On("FindPaymentByOrderID", orderID).Return(nilPayment, nil)
	repo.On("Insert", mock.Anything).Return(nil)
	paymentGateway := &mocks.PaymentGatewayMock{}
	paymentRequestDTO := &dtos.PaymentRequestDTO{
		OrderId:     orderID,
		PaymentLink: "link...",
		Method:      dtos.PIX,
		Amount:      amount,
	}
	paymentGateway.On("Execute", mock.Anything, mock.Anything, mock.Anything).Return(paymentRequestDTO, nil)
	uc := NewCreatePaymentUseCase(repo, paymentGateway)

	// Act
	dto, err := uc.Execute(orderID, amount)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, dto)
	assert.Equal(t, orderID, dto.OrderId)
	assert.NotEmpty(t, dto.PaymentLink)
	assert.NotEmpty(t, dto.Method)
	assert.Equal(t, amount, dto.Amount)
}

func TestCreatePaymentUseCase_ExistPayment(t *testing.T) {
	// Arrange
	orderID := uuid.NewString()
	amount := 45.5
	repo := &mocks.PaymentRepositoryMock{}
	payment := entities.CreatePayment(orderID, amount)
	repo.On("FindPaymentByOrderID", orderID).Return(payment, nil)
	repo.On("Update", mock.Anything).Return(nil)
	paymentGateway := &mocks.PaymentGatewayMock{}
	paymentRequestDTO := &dtos.PaymentRequestDTO{
		OrderId:     orderID,
		PaymentLink: "link...",
		Method:      dtos.PIX,
		Amount:      amount,
	}
	paymentGateway.On("Execute", mock.Anything, mock.Anything, mock.Anything).Return(paymentRequestDTO, nil)
	uc := NewCreatePaymentUseCase(repo, paymentGateway)

	// Act
	dto, err := uc.Execute(orderID, amount)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, dto)
	assert.Equal(t, orderID, dto.OrderId)
	assert.NotEmpty(t, dto.PaymentLink)
	assert.NotEmpty(t, dto.Method)
	assert.Equal(t, amount, dto.Amount)
}

func TestCreatePaymentUseCase_InvalidPaymentStatus(t *testing.T) {
	// Arrange
	orderID := uuid.NewString()
	amount := 45.5
	repo := &mocks.PaymentRepositoryMock{}
	payment := entities.CreatePayment(orderID, amount)
	payment.PaymentReceived()
	repo.On("FindPaymentByOrderID", orderID).Return(payment, nil)
	repo.On("Update", mock.Anything).Return(nil)
	paymentGateway := &mocks.PaymentGatewayMock{}
	paymentRequestDTO := &dtos.PaymentRequestDTO{
		OrderId:     orderID,
		PaymentLink: "link...",
		Method:      dtos.PIX,
		Amount:      amount,
	}
	paymentGateway.On("Execute", mock.Anything, mock.Anything, mock.Anything).Return(paymentRequestDTO, nil)
	uc := NewCreatePaymentUseCase(repo, paymentGateway)

	// Act
	_, err := uc.Execute(orderID, amount)

	// Assert
	assert.NotNil(t, err)
}

func TestCreatePaymentUseCase_FindError(t *testing.T) {
	// Arrange
	orderID := uuid.NewString()
	amount := 45.5
	repo := &mocks.PaymentRepositoryMock{}
	var nilPayment *entities.Payment
	repo.On("FindPaymentByOrderID", orderID).Return(nilPayment, errors.New("Find Error"))
	repo.On("Update", mock.Anything).Return(nil)
	paymentGateway := &mocks.PaymentGatewayMock{}
	paymentRequestDTO := &dtos.PaymentRequestDTO{
		OrderId:     orderID,
		PaymentLink: "link...",
		Method:      dtos.PIX,
		Amount:      amount,
	}
	paymentGateway.On("Execute", mock.Anything, mock.Anything, mock.Anything).Return(paymentRequestDTO, nil)
	uc := NewCreatePaymentUseCase(repo, paymentGateway)

	// Act
	_, err := uc.Execute(orderID, amount)

	// Assert
	assert.NotNil(t, err)
}
