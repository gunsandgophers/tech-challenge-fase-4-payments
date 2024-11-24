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

func (uc *PaymentOrderUseCase) checkValidOrder(orderId string) (*entities.Payment, error) {
	// DEVERA BUSCAR A ORDER VIA GRPC E VALIDAR
	// order, err := uc.orderRepository.FindOrderByID(orderId)
	// if err != nil {
	// 	return nil, err
	// }
	// if order.GetPreparationStatus() != entities.ORDER_PREPARATION_AWAITING {
	// 	return nil, errors.ErrOrderNotAwaitingPreparation
	// }
	// DEVERA BUSCAR O PAYMENT E VALIDAR
	// if order.GetPaymentStatus() != entities.ORDER_PAYMENT_AWAITING_PAYMENT {
	// 	return nil, errors.ErrOrderNotAwaitingPayment
	// }
	// return order, nil
	return nil, nil
}

func (uc *PaymentOrderUseCase) processPayment(
	order *entities.Payment,
	paymentStatus entities.OrderPaymentStatus,
) error {
	// DEVE ATUALIZAR O PAGAMENTO
	// if paymentStatus == entities.ORDER_PAYMENT_PAID {
	// 	order.PaymentReceived()
	// } else {
	// 	order.PaymentRejected()
	// }
	// return uc.orderRepository.Update(order)
	return nil
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
