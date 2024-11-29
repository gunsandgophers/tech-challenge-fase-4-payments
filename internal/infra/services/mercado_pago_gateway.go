package services

import (
	"tech-challenge-fase-1/internal/core/dtos"
	"tech-challenge-fase-1/internal/infra/events"
)

type MercadoPagoGateway struct {
	eventManager *events.EventManager
}

func NewMercadoPagoGateway(eventManager *events.EventManager) *MercadoPagoGateway {
	return &MercadoPagoGateway{
		eventManager: eventManager,
	}
}

func (m *MercadoPagoGateway) Execute(orderID string, amount float64, method dtos.MethodType) (*dtos.PaymentRequestDTO, error) {
	link := "https://www.pngall.com/wp-content/uploads/2/QR-Code-PNG-Images.png"
	return &dtos.PaymentRequestDTO{
		OrderId:     orderID,
		PaymentLink: link,
		Method:      method,
		Amount:      amount,
	}, nil
}
