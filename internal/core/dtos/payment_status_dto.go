package dtos

import "tech-challenge-fase-1/internal/core/entities"

type PaymentStatusDTO struct {
	OrderId       string `json:"order_id"`
	PaymentStatus string `json:"payment_status"`
}

func NewPaymentStatusDTOFromEntity(payment *entities.Payment) *PaymentStatusDTO {
	return &PaymentStatusDTO{
		OrderId:       payment.GetOrderID(),
		PaymentStatus: payment.GetPaymentStatus().String(),
	}
}

type MethodType string

const (
	PIX         MethodType = "PIX"
	CREDIT_CARD MethodType = "CREDIT_CARD"
)

type PaymentRequestDTO struct {
	OrderId     string
	PaymentLink string
	Method      MethodType
	Amount      float64
}
