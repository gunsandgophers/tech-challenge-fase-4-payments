package payments

import (
	"strings"
	"tech-challenge-fase-1/internal/core/entities"
	"tech-challenge-fase-1/internal/core/errors"
	"tech-challenge-fase-1/internal/core/repositories"
)

type PaymentOrderUseCase struct {
	paymentRepository repositories.PaymentRepositoryInterface
}

func NewPaymentOrderUseCase(
	paymentRepository repositories.PaymentRepositoryInterface,
) *PaymentOrderUseCase {
	return &PaymentOrderUseCase{
		paymentRepository: paymentRepository,
	}
}

func (uc *PaymentOrderUseCase) checkValidPaymentStatus(
	paymentStatusString string,
) (entities.OrderPaymentStatus, error) {
	paymentStatus := entities.OrderPaymentStatus(strings.ToUpper(paymentStatusString))
	if paymentStatus != entities.ORDER_PAYMENT_PAID &&
		paymentStatus != entities.ORDER_PAYMENT_REJECTED {
		return paymentStatus, errors.ErrInvalidPaymentStatus
	}

	return paymentStatus, nil
}

func (uc *PaymentOrderUseCase) checkValidOrder(orderID string) (*entities.Payment, error) {
	// DEVERA BUSCAR A ORDER VIA GRPC E VALIDAR
	// order, err := uc.orderRepository.FindOrderByID(orderId)
	// if err != nil {
	// 	return nil, err
	// }
	// if order.GetPreparationStatus() != entities.ORDER_PREPARATION_AWAITING {
	// 	return nil, errors.ErrOrderNotAwaitingPreparation
	// }

	payment, err := uc.paymentRepository.FindPaymentByOrderID(orderID)
	if err != nil {
		return nil, err
	}
	if payment.GetPaymentStatus() != entities.ORDER_PAYMENT_AWAITING_PAYMENT {
		return nil, errors.ErrOrderNotAwaitingPayment
	}
	return payment, nil
}

func (uc *PaymentOrderUseCase) processPayment(
	payment *entities.Payment,
	paymentStatus entities.OrderPaymentStatus,
) error {
	if paymentStatus == entities.ORDER_PAYMENT_PAID {
		payment.PaymentReceived()
	} else {
		payment.PaymentRejected()
	}
	return uc.paymentRepository.Update(payment)
}

func (uc *PaymentOrderUseCase) Execute(
	orderId string,
	paymentStatusString string,
) error {
	paymentStatus, err := uc.checkValidPaymentStatus(paymentStatusString)
	if err != nil {
		return err
	}
	payment, err := uc.checkValidOrder(orderId)
	if err != nil {
		return err
	}
	return uc.processPayment(payment, paymentStatus)
}
