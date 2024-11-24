package payments

import (
	"tech-challenge-fase-1/internal/core/dtos"
	"tech-challenge-fase-1/internal/core/repositories"
)

type GetPaymentStatusUseCase struct {
	paymentRepository     repositories.PaymentRepositoryInterface
}

func NewGetPaymentStatusUseCase(
	paymentRepository repositories.PaymentRepositoryInterface,
) *GetPaymentStatusUseCase {
	return &GetPaymentStatusUseCase{
		paymentRepository:     paymentRepository,
	}
}

func (uc *GetPaymentStatusUseCase) Execute(
	orderId string,
) (*dtos.PaymentStatusDTO, error) {
	payment, err := uc.paymentRepository.FindPaymentByOrderID(orderId)
	if err != nil {
		return nil, err
	}
	return dtos.NewPaymentStatusDTOFromEntity(payment), nil
}

