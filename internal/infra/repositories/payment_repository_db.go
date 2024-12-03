package repositories

import (
	// "encoding/json"
	"tech-challenge-fase-1/internal/core/entities"
	"tech-challenge-fase-1/internal/infra/database"
)

type PaymentRepositoryDB struct {
	conn database.ConnectionDB
}

func NewOrderRepositoryDB(conn database.ConnectionDB) *PaymentRepositoryDB {
	return &PaymentRepositoryDB{conn: conn}
}

func (r *PaymentRepositoryDB) Insert(payment *entities.Payment) error {
	sql := `
	INSERT INTO payments(id, id_order, amount, payment_status)
	VALUES ($1, $2, $3, $4)
	`
	return r.conn.Exec(
		sql,
		payment.GetID(),
		payment.GetOrderID(),
		payment.GetAmount(),
		payment.GetPaymentStatus().String(),
	)
}

func (r *PaymentRepositoryDB) FindPaymentByOrderID(orderId string) (*entities.Payment, error) {
	sql := `
	SELECT
		id,
		id_order,
		amount,
		payment_status
	FROM payments
	WHERE id_order = $1`
	row := r.conn.QueryRow(sql, orderId)
	return r.toEntity(row)
}

func (r *PaymentRepositoryDB) Update(payment *entities.Payment) error {
	sql := `
	UPDATE payments
	SET
		id_order = $1,
		amount = $2,
		payment_status = $3,
	WHERE id = $4
	`
	return r.conn.Exec(
		sql,
		payment.GetOrderID(),
		payment.GetAmount(),
		payment.GetPaymentStatus().String(),
		payment.GetID(),
	)
}

func (r *PaymentRepositoryDB) toEntity(row database.RowDB) (*entities.Payment, error) {
	var id string
	var orderID string
	var amount float64
	var paymentStatus entities.OrderPaymentStatus
	err := row.Scan(&id, &orderID, &amount, &paymentStatus)
	if err != nil {
		if err.Error() == ErrNotFound {
			return nil, ErrOrderNotFound
		}
		return nil, err
	}
	return entities.RestorePayment(
		id,
		orderID,
		amount,
		paymentStatus,
	), nil
}
