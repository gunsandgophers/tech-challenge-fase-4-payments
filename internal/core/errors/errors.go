package errors

import "errors"

var (
	ErrOrderNotAwaitingPayment = errors.New("Order not awaiting payment")
	ErrInvalidPaymentStatus = errors.New("Invalid Payment Status")
)
