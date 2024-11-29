package app

import (
	"tech-challenge-fase-1/internal/infra/controllers"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Registra as rotas dos controllers
func registerRouters(app *APIApp) {
	paymentController := controllers.NewPaymentController(
		app.paymentRepository,
		app.mercadoPagoGateway,
	)

	baseUrl := "/api/v1"
	app.httpServer.SetBasePath(baseUrl)

	app.httpServer.GET( "/payment/:order_id", paymentController.GetPaymentStatus)
	app.httpServer.POST( "/payment/:order_id", paymentController.CreatePayment)
	app.httpServer.GET( "/payment/:order_id", paymentController.ProcessPayment)
	app.httpServer.SetSwagger("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
