package fixtures

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"tech-challenge-fase-1/internal/core/entities"
	"tech-challenge-fase-1/internal/tests/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetPaymentStatus(t *testing.T) {

	repo := &mocks.PaymentRepositoryMock{}
	gateway := &mocks.PaymentGatewayMock{}

	app := NewAPIAppIntegrationTest(repo, gateway)

	orderID := uuid.NewString()

	payment := entities.RestorePayment(uuid.NewString(), orderID, 13.37, entities.ORDER_PAYMENT_PAID)

	repo.On("FindPaymentByOrderID", orderID).Return(payment, nil).Once()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/api/v1/payment/"+orderID, strings.NewReader(""))
	app.HTTPServer().ServeHTTP(w, req)

	var response map[string]interface{}

	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Nil(t, err)
	assert.Equal(t, w.Result().StatusCode, 200)

	assert.Equal(t, response["message"], "operation: get-payment successfull")
}

func TestGetPaymentStatusWithError(t *testing.T) {

	repo := &mocks.PaymentRepositoryMock{}
	gateway := &mocks.PaymentGatewayMock{}

	app := NewAPIAppIntegrationTest(repo, gateway)

	orderID := uuid.NewString()

	repo.On("FindPaymentByOrderID", orderID).Return(&entities.Payment{}, errors.New("error")).Once()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/api/v1/payment/"+orderID, strings.NewReader(""))
	app.HTTPServer().ServeHTTP(w, req)

	var response map[string]interface{}

	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Nil(t, err)
	assert.Equal(t, w.Result().StatusCode, 500)

	assert.Equal(t, response["message"], "error")
}

func TestProcessPayment(t *testing.T) {

	repo := &mocks.PaymentRepositoryMock{}
	gateway := &mocks.PaymentGatewayMock{}

	app := NewAPIAppIntegrationTest(repo, gateway)

	orderID := uuid.NewString()

	payment := entities.RestorePayment(uuid.NewString(), orderID, 13.37, entities.ORDER_PAYMENT_AWAITING_PAYMENT)

	repo.On("FindPaymentByOrderID", orderID).Return(payment, nil).Once()
	repo.On("Update", mock.Anything).Return(nil).Once()

	w := httptest.NewRecorder()

	request, _ := json.Marshal(map[string]string{
		"payment_status": entities.ORDER_PAYMENT_PAID.String(),
	})

	req, err := http.NewRequest("POST", "/api/v1/payment/"+orderID+"/process", bytes.NewReader(request))
	app.HTTPServer().ServeHTTP(w, req)

	var response map[string]interface{}

	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Nil(t, err)
	assert.Equal(t, w.Result().StatusCode, 204)

	assert.Nil(t, response)
}

func TestProcessPaymentWithError(t *testing.T) {

	repo := &mocks.PaymentRepositoryMock{}
	gateway := &mocks.PaymentGatewayMock{}

	app := NewAPIAppIntegrationTest(repo, gateway)

	orderID := uuid.NewString()

	payment := entities.RestorePayment(uuid.NewString(), orderID, 13.37, entities.ORDER_PAYMENT_AWAITING_PAYMENT)

	repo.On("FindPaymentByOrderID", orderID).Return(payment, nil).Once()
	repo.On("Update", mock.Anything).Return(errors.New("error")).Once()

	w := httptest.NewRecorder()

	request, _ := json.Marshal(map[string]string{
		"payment_status": entities.ORDER_PAYMENT_PAID.String(),
	})

	req, err := http.NewRequest("POST", "/api/v1/payment/"+orderID+"/process", bytes.NewReader(request))
	app.HTTPServer().ServeHTTP(w, req)

	var response map[string]interface{}

	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Nil(t, err)
	assert.Equal(t, w.Result().StatusCode, 406)

	assert.Equal(t, response["message"], "error")
}
