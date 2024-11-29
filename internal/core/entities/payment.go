package entities

import (
	"github.com/google/uuid"
)

type (
	OrderPaymentStatus string
)

func (s OrderPaymentStatus) String() string {
	return string(s)
}

const (
	ORDER_PAYMENT_PENDING          OrderPaymentStatus = "PENDING"
	ORDER_PAYMENT_PAID             OrderPaymentStatus = "PAID"
	ORDER_PAYMENT_REJECTED         OrderPaymentStatus = "REJECTED"
	ORDER_PAYMENT_AWAITING_PAYMENT OrderPaymentStatus = "AWAITING_PAYMENT"
)

type Payment struct {
	id            string
	orderID       string
	amount        float64
	paymentStatus OrderPaymentStatus
}

func CreatePayment(orderID string, amount float64) *Payment {
	return RestorePayment(
		uuid.NewString(),
		orderID,
		amount,
		ORDER_PAYMENT_AWAITING_PAYMENT,
	)
}

func RestorePayment(
	id string,
	orderID string,
	amount float64,
	paymentStatus OrderPaymentStatus,
) *Payment {
	return &Payment{
		id:            id,
		orderID:       orderID,
		amount:        amount,
		paymentStatus: paymentStatus,
	}
}

func (p *Payment) GetID() string {
	return p.id
}

func (p *Payment) GetOrderID() string {
	return p.orderID
}

func (p *Payment) GetAmount() float64 {
	return p.amount
}

func (p *Payment) GetPaymentStatus() OrderPaymentStatus {
	return p.paymentStatus
}

func (p *Payment) UpdateAmount(amount float64) {
	p.amount = amount
}

func (p *Payment) AwaitingPayment() {
	p.paymentStatus = ORDER_PAYMENT_AWAITING_PAYMENT
}

func (p *Payment) PaymentReceived() {
	p.paymentStatus = ORDER_PAYMENT_PAID
}

func (p *Payment) PaymentRejected() {
	p.paymentStatus = ORDER_PAYMENT_REJECTED
}
