package bdd

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"tech-challenge-fase-1/internal/core/dtos"
	"tech-challenge-fase-1/internal/core/entities"
	"tech-challenge-fase-1/internal/infra/app"
	"tech-challenge-fase-1/internal/infra/controllers"
	"tech-challenge-fase-1/internal/tests/fixtures"
	"tech-challenge-fase-1/internal/tests/mocks"
	"testing"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/mock"
)

type appCtxKey struct{}
type dependenciesCtxKey struct{}
type paymentCtxKey struct{}
type orderIDCtxKey struct{}
type responseCtxKey struct{}

type Dependencies struct {
	paymentRepository *mocks.PaymentRepositoryMock
	paymentGateway    *mocks.PaymentGatewayMock
}

func newPaymentRequest(ctx context.Context) (context.Context, error) {
	return context.WithValue(ctx, paymentCtxKey{}, controllers.CreatePaymentRequest{}), nil
}

func defineOrderID(ctx context.Context, orderID string) (context.Context, error) {
	return context.WithValue(ctx, orderIDCtxKey{}, orderID), nil
}

func defineAmount(
	ctx context.Context,
	price string,
) (context.Context, error) {
	request := ctx.Value(paymentCtxKey{}).(controllers.CreatePaymentRequest)
	request.Amount, _ = strconv.ParseFloat(price, 64)
	return context.WithValue(ctx, paymentCtxKey{}, request), nil
}

func sendCreatePaymentRequest(
	ctx context.Context,
) (context.Context, error) {

	request := ctx.Value(paymentCtxKey{}).(controllers.CreatePaymentRequest)
	orderID := ctx.Value(orderIDCtxKey{}).(string)
	dependencies := ctx.Value(dependenciesCtxKey{}).(*Dependencies)

	var nilPayment *entities.Payment
	dependencies.paymentRepository.On("FindPaymentByOrderID", mock.Anything).Return(nilPayment, nil)
	dependencies.paymentRepository.On("Insert", mock.Anything).Return(nil)

	paymentRequestDTO := &dtos.PaymentRequestDTO{
		OrderId:     orderID,
		PaymentLink: "link...",
		Method:      dtos.PIX,
		Amount:      request.Amount,
	}

	dependencies.paymentGateway.On("Execute", mock.Anything, mock.Anything, mock.Anything).Return(
		paymentRequestDTO,
		nil,
	)

	app := ctx.Value(appCtxKey{}).(*app.APIApp)
	w := httptest.NewRecorder()
	body, _ := json.Marshal(request)
	req, _ := http.NewRequest("POST", "/api/v1/payment/"+orderID, bytes.NewReader(body))
	app.HTTPServer().ServeHTTP(w, req)
	return context.WithValue(ctx, responseCtxKey{}, w), nil
}

type ResponseCreatePayment struct {
	Data    dtos.PaymentRequestDTO `json:"data,omitempty"`
	Message string                 `json:"message,omitempty"`
}

func paymentCreatedWith(ctx context.Context, orderID string) error {
	response, _ := ctx.Value(responseCtxKey{}).(*httptest.ResponseRecorder)
	status := response.Result().StatusCode

	if status != http.StatusCreated {
		return errors.New("invalid status")
	}

	r := ResponseCreatePayment{}
	err := json.Unmarshal(response.Body.Bytes(), &r)
	if err != nil {
		return err
	}

	if r.Message != "operation: new-payment successfull" {
		return errors.New("error on new-payment")
	}

	if r.Data.OrderId != orderID {
		return errors.New("order ID invalid")
	}

	if r.Data.PaymentLink == "" {
		return errors.New("link invalid")
	}

	return nil
}

func TestFeatures(t *testing.T) {
	dependencies := &Dependencies{
		paymentRepository: &mocks.PaymentRepositoryMock{},
		paymentGateway:    &mocks.PaymentGatewayMock{},
	}

	app := fixtures.NewAPIAppIntegrationTest(dependencies.paymentRepository, dependencies.paymentGateway)
	ctx := context.WithValue(context.Background(), appCtxKey{}, app)
	ctx = context.WithValue(ctx, dependenciesCtxKey{}, dependencies)
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:         "pretty",
			Paths:          []string{"features"},
			DefaultContext: ctx,
			TestingT:       t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(sc *godog.ScenarioContext) {
	sc.Given(`^That I need to create a new payment via API$`, newPaymentRequest)
	sc.Step(`^have the order ID "([^"]*)"$`, defineOrderID)
	sc.Step(`^amount as "([^"]*)"$`, defineAmount)
	sc.When(`^I send the data$`, sendCreatePaymentRequest)
	sc.Then(
		`^the payment need to be created with link and order ID "([^"]*)"$`,
		paymentCreatedWith,
	)
}
