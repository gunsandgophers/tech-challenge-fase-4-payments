package controllers

import (
	"net/http"
	"tech-challenge-fase-1/internal/core/repositories"
	"tech-challenge-fase-1/internal/core/services"
	"tech-challenge-fase-1/internal/core/use_cases/payments"
	httpserver "tech-challenge-fase-1/internal/infra/http"
)

type PaymentController struct {
	paymentRepository repositories.PaymentRepositoryInterface
	paymentGateway    services.PaymentGatewayInterface
}

func NewPaymentController(
	paymentRepository repositories.PaymentRepositoryInterface,
	paymentGateway services.PaymentGatewayInterface,
) *PaymentController {
	return &PaymentController{
		paymentRepository: paymentRepository,
		paymentGateway:    paymentGateway,
	}
}

// PaymentRequest
//
//	@Summary		Request new payment
//	@Description	request a new payment by order_id
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			order_id	path		string	true	"Get Payment Status"
//	@Success		200			{object}	dtos.PaymentRequestDTO
//	@Failure		400			{string}	string	"when bad request"
//	@Failure		406			{string}	string	"when invalid params or invalid object"
//	@Router			/payment/{order_id} [post]
func (ctrl *PaymentController) CreatePayment(c httpserver.HTTPContext) {
	orderId := c.Param("order_id")
	request := &CreatePaymentRequest{}
	c.BindJSON(request)
	if err := request.Validate(); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}
	uc := payments.NewCreatePaymentUseCase(ctrl.paymentRepository, ctrl.paymentGateway)
	paymentDTO, err := uc.Execute(orderId, request.Amount)
	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}
	sendSuccess(c, http.StatusCreated, "new-payment", paymentDTO)
}

// GetPaymentStatus godoc
//
//	@Summary		Get a payment status
//	@Description	get payment status by order_id
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			order_id	path		string	true	"Get Payment Status"
//	@Success		200			{object}	dtos.PaymentStatusDTO
//	@Failure		400			{string}	string	"when bad request"
//	@Failure		406			{string}	string	"when invalid params or invalid object"
//	@Router			/payment/{order_id} [get]
func (ctrl *PaymentController) GetPaymentStatus(c httpserver.HTTPContext) {
	orderId := c.Param("order_id")
	getPaymentStatusUC := payments.NewGetPaymentStatusUseCase(ctrl.paymentRepository)
	paymentStatus, err := getPaymentStatusUC.Execute(orderId)
	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}
	sendSuccess(c, http.StatusOK, "get-payment", paymentStatus)
}

// Payment godoc
//
//	@Summary		Process order payment
//	@Description	process the payment for an order
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			payment	body		PaymentRequest	true	"Payment"
//	@Success		200		{object}	string			""
//	@Failure		400		{string}	string			"when bad request"
//	@Failure		406		{string}	string			"when invalid params or invalid object"
//	@Router			/payment/{order_id} [put]
func (ctrl *PaymentController) ProcessPayment(c httpserver.HTTPContext) {
	orderID := c.Param("order_id")
	request := &ProcessPaymentRequest{}
	c.BindJSON(request)
	if err := request.Validate(); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}
	paymentUseCase := payments.NewPaymentOrderUseCase(
		ctrl.paymentRepository,
	)
	err := paymentUseCase.Execute(orderID, request.PaymentStatus)
	if err != nil {
		sendError(c, http.StatusNotAcceptable, err.Error())
		return
	}
	sendSuccess(c, http.StatusNoContent, "payment-order", nil)
}
