package repositories

import (
	"errors"
	"tech-challenge-fase-1/internal/core/entities"
	"tech-challenge-fase-1/internal/tests/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInsertPayment(t *testing.T) {
	conn := mocks.NewMockConnectionDB(t)

	repo := NewOrderRepositoryDB(conn)

	payment := entities.RestorePayment(
		uuid.NewString(), uuid.NewString(),
		13.37, entities.ORDER_PAYMENT_PAID,
	)

	conn.On("Exec", mock.Anything, payment.GetID(), payment.GetOrderID(),
		payment.GetAmount(), payment.GetPaymentStatus().String(),
	).Return(nil).Once()

	err := repo.Insert(payment)

	assert.Nil(t, err)
}

func TestInsertPaymentWithError(t *testing.T) {
	conn := mocks.NewMockConnectionDB(t)

	repo := NewOrderRepositoryDB(conn)

	payment := entities.RestorePayment(
		uuid.NewString(), uuid.NewString(),
		13.37, entities.ORDER_PAYMENT_PAID,
	)

	conn.On("Exec", mock.Anything, payment.GetID(), payment.GetOrderID(),
		payment.GetAmount(), payment.GetPaymentStatus().String(),
	).Return(errors.New("error")).Once()

	err := repo.Insert(payment)

	assert.EqualError(t, err, "error")
}

func TestFindPaymentByOrderID(t *testing.T) {
	conn := mocks.NewMockConnectionDB(t)

	row := mocks.NewMockRowDB(t)

	repo := NewOrderRepositoryDB(conn)

	orderID := uuid.NewString()

	payment := entities.RestorePayment(
		uuid.NewString(), orderID,
		13.37, entities.ORDER_PAYMENT_PAID,
	)

	conn.On("QueryRow", mock.Anything, orderID).Return(row).Once()
	row.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

	_, err := repo.FindPaymentByOrderID(payment.GetOrderID())

	assert.Nil(t, err)
}

func TestFindPaymentByOrderIDWithErrNotFound(t *testing.T) {
	conn := mocks.NewMockConnectionDB(t)

	row := mocks.NewMockRowDB(t)

	repo := NewOrderRepositoryDB(conn)

	orderID := uuid.NewString()

	payment := entities.RestorePayment(
		uuid.NewString(), orderID,
		13.37, entities.ORDER_PAYMENT_PAID,
	)

	conn.On("QueryRow", mock.Anything, orderID).Return(row).Once()
	row.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(errors.New(ErrNotFound)).Once()

	_, err := repo.FindPaymentByOrderID(payment.GetOrderID())

	assert.Nil(t, err)
}

func TestFindPaymentByOrderIDWithError(t *testing.T) {
	conn := mocks.NewMockConnectionDB(t)

	row := mocks.NewMockRowDB(t)

	repo := NewOrderRepositoryDB(conn)

	orderID := uuid.NewString()

	payment := entities.RestorePayment(
		uuid.NewString(), orderID,
		13.37, entities.ORDER_PAYMENT_PAID,
	)

	conn.On("QueryRow", mock.Anything, orderID).Return(row).Once()
	row.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(errors.New("error")).Once()

	_, err := repo.FindPaymentByOrderID(payment.GetOrderID())

	assert.EqualError(t, err, "error")
}

func TestUpdatePayment(t *testing.T) {
	conn := mocks.NewMockConnectionDB(t)

	repo := NewOrderRepositoryDB(conn)

	payment := entities.RestorePayment(
		uuid.NewString(), uuid.NewString(),
		13.37, entities.ORDER_PAYMENT_PAID,
	)

	conn.On("Exec", mock.Anything, payment.GetOrderID(),
		payment.GetAmount(), payment.GetPaymentStatus().String(),
		payment.GetID(),
	).Return(nil).Once()

	err := repo.Update(payment)

	assert.Nil(t, err)
}

func TestUpdatePaymentWithError(t *testing.T) {
	conn := mocks.NewMockConnectionDB(t)

	repo := NewOrderRepositoryDB(conn)

	payment := entities.RestorePayment(
		uuid.NewString(), uuid.NewString(),
		13.37, entities.ORDER_PAYMENT_PAID,
	)

	conn.On("Exec", mock.Anything, payment.GetOrderID(),
		payment.GetAmount(), payment.GetPaymentStatus().String(),
		payment.GetID(),
	).Return(errors.New("error")).Once()

	err := repo.Update(payment)

	assert.EqualError(t, err, "error")
}
