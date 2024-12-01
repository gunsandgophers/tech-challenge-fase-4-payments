package fixtures

import (
	"tech-challenge-fase-1/internal/infra/app"
	httpserver "tech-challenge-fase-1/internal/infra/http"
	"tech-challenge-fase-1/internal/tests/mocks"
)


func NewAPIAppTest() *app.APIApp {
	httpServer := httpserver.NewGinHTTPServerAdapter()
	paymentRepository := &mocks.PaymentRepositoryMock{}
	paymentGateway := &mocks.PaymentGatewayMock{}

	return app.NewAPIApp(httpServer, paymentRepository, paymentGateway)
}
