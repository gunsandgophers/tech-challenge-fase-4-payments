package controllers

import "fmt"

type ProcessPaymentRequest struct {
	PaymentStatus string `json:"payment_status"`
}

type CreatePaymentRequest struct {
	Amount float64 `json:"amount"`
}

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func errParamCantBeEmpty(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) cant be empty", name, typ)
}

func (r *ProcessPaymentRequest) Validate() error {
	if len(r.PaymentStatus) == 0 {
		return errParamIsRequired("payment_status", "string")
	}
	return nil
}

func (r *CreatePaymentRequest) Validate() error {
	if r.Amount < 0 {
		return errParamIsRequired("amount", "string")
	}
	return nil
}
