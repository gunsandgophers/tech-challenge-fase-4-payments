package entities

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreatePayment(t *testing.T) {
	// Arrange
	orderID := uuid.NewString()
	amount := 40.9
	// Act
	payment := CreatePayment(orderID, amount)
	// Assert
	assert.NotEmpty(t, payment.GetID())
	assert.Equal(t, payment.GetOrderID(), orderID)
	assert.Equal(t, payment.GetAmount(), amount)
	assert.Equal(t, payment.GetPaymentStatus(), ORDER_PAYMENT_AWAITING_PAYMENT)
}

func TestRestorePayment(t *testing.T) {

	// Arrange
	id := uuid.NewString()
	orderID := uuid.NewString()
	amount := 40.9
	// Act
	payment := RestorePayment(id, orderID, amount, ORDER_PAYMENT_PAID)
	// Assert
	assert.NotEmpty(t, payment.GetID())
	assert.Equal(t, orderID, payment.GetOrderID())
	assert.Equal(t, amount, payment.GetAmount())
	assert.Equal(t, ORDER_PAYMENT_PAID, payment.GetPaymentStatus())
}

func TestUpdateAmountPayment(t *testing.T) {
	// Arrange
	orderID := uuid.NewString()
	amount := 40.9
	payment := CreatePayment(orderID, amount)
	// Act
	newAmount := 30.5
	payment.UpdateAmount(newAmount)

	// Assert
	assert.NotEmpty(t, payment.GetID())
	assert.Equal(t, orderID, payment.GetOrderID())
	assert.Equal(t, newAmount, payment.GetAmount())
}

func TestPaymentStatus_AwaitingPayment(t *testing.T) {
	// Arrange
	orderID := uuid.NewString()
	amount := 40.9
	payment := CreatePayment(orderID, amount)
	// Act
	payment.AwaitingPayment()
	// Assert
	assert.Equal(t, ORDER_PAYMENT_AWAITING_PAYMENT, payment.GetPaymentStatus())
}

func TestPaymentStatus_PaymentReceived(t *testing.T) {
	// Arrange
	orderID := uuid.NewString()
	amount := 40.9
	payment := CreatePayment(orderID, amount)
	// Act
	payment.PaymentReceived()
	// Assert
	assert.Equal(t, ORDER_PAYMENT_PAID, payment.GetPaymentStatus())
}

func TestPaymentStatus_PaymentRejected(t *testing.T) {
	// Arrange
	orderID := uuid.NewString()
	amount := 40.9
	payment := CreatePayment(orderID, amount)
	// Act
	payment.PaymentRejected()
	// Assert
	assert.Equal(t, ORDER_PAYMENT_REJECTED, payment.GetPaymentStatus())
}

func TestPaymentStatus_String(t *testing.T) {
	categories := []OrderPaymentStatus{
		ORDER_PAYMENT_PENDING,
		ORDER_PAYMENT_PAID,
		ORDER_PAYMENT_REJECTED,
		ORDER_PAYMENT_AWAITING_PAYMENT,
	}

	for _, c := range categories {
		assert.Equal(t, c.String(), string(c))
	}
}
