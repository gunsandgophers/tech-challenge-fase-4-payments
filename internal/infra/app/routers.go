package app

import (
	"tech-challenge-fase-1/internal/infra/controllers"
	httpserver "tech-challenge-fase-1/internal/infra/http"
)

// Registra as rotas dos controllers
func registerRouters(app *APIApp) {
	paymentController := controllers.NewPaymentController(
		app.paymentRepository,
		app.paymentGateway,
	)

	baseUrl := "/api/v1"
	app.httpServer.(httpserver.HTTPRoutes).SetBasePath(baseUrl)

	app.httpServer.(httpserver.HTTPRoutes).GET( "/payment/:order_id", paymentController.GetPaymentStatus)
	app.httpServer.(httpserver.HTTPRoutes).POST( "/payment/:order_id", paymentController.CreatePayment)
	app.httpServer.(httpserver.HTTPRoutes).GET( "/payment/:order_id", paymentController.ProcessPayment)
	app.httpServer.(httpserver.HTTPRoutes).SetSwagger("/swagger/*any")
}
