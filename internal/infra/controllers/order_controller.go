package controllers

import (
	"net/http"
	"tech-challenge-fase-1/internal/core/events"
	"tech-challenge-fase-1/internal/core/queries"
	"tech-challenge-fase-1/internal/core/repositories"
	"tech-challenge-fase-1/internal/core/services"
	"tech-challenge-fase-1/internal/core/use_cases/orders"
	httpserver "tech-challenge-fase-1/internal/infra/http"
)

type OrderController struct {
	paymentRepository       repositories.PaymentRepositoryInterface
	paymentGateway        services.PaymentGatewayInterface
	commandEventManager   events.ManagerEvent
}

func NewOrderController(
	paymentRepository repositories.PaymentRepositoryInterface,
	paymentGateway services.PaymentGatewayInterface,
	commandEventManager events.ManagerEvent,
) *OrderController {
	return &OrderController{
		paymentRepository:       paymentRepository,
		paymentGateway:        paymentGateway,
		commandEventManager:   commandEventManager,
	}
}

// func (cc *OrderController) Checkout(c httpserver.HTTPContext) {
// 	request := CheckoutRequest{}
// 	c.BindJSON(&request)
// 	if err := request.Validate(); err != nil {
// 		sendError(c, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	checkoutUseCase := orders.NewCheckoutOrderUseCase(
// 		cc.orderRepository,
// 		cc.customerService,
// 		cc.productRepository,
// 		cc.paymentGateway,
// 		cc.commandEventManager,
// 	)
// 	checkout, err := checkoutUseCase.Execute(request.CustomerId, request.ProductsIds)
// 	if err != nil {
// 		sendError(c, http.StatusNotAcceptable, err.Error())
// 		return
// 	}
// 	sendSuccess(c, http.StatusCreated, "checkout-order", checkout)
// }

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
//	@Router			/order/{order_id}/payment-status [get]
func (cc *OrderController) GetPaymentStatus(c httpserver.HTTPContext) {
	orderId := c.Param("order_id")
	getPaymentStatusUC := orders.NewGetPaymentStatusUseCase(cc.paymentGateway)
	paymentStatus, err := getPaymentStatusUC.Execute(orderId)
	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}
	sendSuccess(c, http.StatusOK, "get-payment-status-order", paymentStatus)
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
//	@Router			/order/payment [post]
func (cc *OrderController) Payment(c httpserver.HTTPContext) {
	request := &PaymentRequest{}
	c.BindJSON(request)
	if err := request.Validate(); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}
	paymentUseCase := orders.NewPaymentOrderUseCase(
		cc.orderRepository,
	)
	err := paymentUseCase.Execute(request.OrderId, request.PaymentStatus)
	if err != nil {
		sendError(c, http.StatusNotAcceptable, err.Error())
		return
	}
	sendSuccess(c, http.StatusNoContent, "payment-order", nil)
}
