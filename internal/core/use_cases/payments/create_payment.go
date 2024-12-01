package payments

import (
	"tech-challenge-fase-1/internal/core/dtos"
	"tech-challenge-fase-1/internal/core/entities"
	"tech-challenge-fase-1/internal/core/errors"
	"tech-challenge-fase-1/internal/core/repositories"
	"tech-challenge-fase-1/internal/core/services"
)

type CreatePaymentUseCase struct {
	paymentRepository repositories.PaymentRepositoryInterface
	paymentGateway    services.PaymentGatewayInterface
}

func NewCreatePaymentUseCase(
	paymentRepository repositories.PaymentRepositoryInterface,
	paymentGateway services.PaymentGatewayInterface,
) *CreatePaymentUseCase {
	return &CreatePaymentUseCase{
		paymentRepository: paymentRepository,
		paymentGateway: paymentGateway,
	}
}

func (uc *CreatePaymentUseCase) getPayment(orderId string, amount float64) (*entities.Payment, error) {
	payment, err := uc.paymentRepository.FindPaymentByOrderID(orderId)
	if err != nil {
		return nil, err
	}

	if payment == nil {
		payment = entities.CreatePayment(orderId, amount)
		uc.paymentRepository.Insert(payment)
		return payment, nil
	}

	paymentStatus := payment.GetPaymentStatus()
	if paymentStatus != entities.ORDER_PAYMENT_REJECTED && paymentStatus != entities.ORDER_PAYMENT_AWAITING_PAYMENT {
		return nil, errors.ErrInvalidPaymentStatus
	}
	payment.UpdateAmount(amount)
	uc.paymentRepository.Update(payment)
	return payment, nil
}

func (uc *CreatePaymentUseCase) Execute(orderId string, amount float64) (*dtos.PaymentRequestDTO, error) {
	payment, err := uc.getPayment(orderId, amount)
	if err != nil {
		return nil, err
	}
	return uc.paymentGateway.Execute(
		payment.GetOrderID(),
		payment.GetAmount(),
		dtos.PIX,
	)
}
