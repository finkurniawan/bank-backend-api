package controller_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/finkurniawan/bank-backend-api/api/controller"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/finkurniawan/bank-backend-api/api/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPaymentService struct {
	mock.Mock
}

func (m *MockPaymentService) MakePayment(customerID, merchantID uint, amount float64, currency string) error {
	args := m.Called(customerID, merchantID, amount, currency)
	return args.Error(0)
}

func TestPaymentController_MakePayment(t *testing.T) {
	mockService := new(MockPaymentService)
	paymentController := controller.NewPaymentController(mockService)

	app := fiber.New()
	app.Post("/payment", func(c *fiber.Ctx) error {
		c.Locals("claims", jwt.MapClaims{"iss": "1"})
		return c.Next()
	}, paymentController.MakePayment)

	t.Run("Successful Payment", func(t *testing.T) {
		mockService.On("MakePayment", uint(1), uint(1), float64(100), "USD").Return(nil)

		reqBody, _ := json.Marshal(model.PaymentRequest{MerchantID: 1, Amount: 100, Currency: "USD"})
		req := httptest.NewRequest(http.MethodPost, "/payment", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var successResponse model.SuccessResponse
		json.NewDecoder(resp.Body).Decode(&successResponse)
		assert.Equal(t, "Payment successful", successResponse.Message)
	})

	t.Run("Invalid Payment Details", func(t *testing.T) {
		reqBody, _ := json.Marshal(model.PaymentRequest{MerchantID: 0, Amount: 0, Currency: ""})
		req := httptest.NewRequest(http.MethodPost, "/payment", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		var errorResponse model.ErrorResponse
		json.NewDecoder(resp.Body).Decode(&errorResponse)
		assert.Equal(t, "Invalid payment details", errorResponse.Message)
	})

	t.Run("Payment Service Error", func(t *testing.T) {
		mockService.On("MakePayment", uint(1), uint(1), float64(100), "USD").
			Return(fmt.Errorf("payment service error"))

		reqBody, _ := json.Marshal(model.PaymentRequest{MerchantID: 1, Amount: 100, Currency: "USD"})
		req := httptest.NewRequest(http.MethodPost, "/payment", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

		var errorResponse model.ErrorResponse
		json.NewDecoder(resp.Body).Decode(&errorResponse)
		assert.Equal(t, "payment service error", errorResponse.Message)
	})

	t.Run("Missing Claims", func(t *testing.T) {
		app := fiber.New()
		app.Post("/payment", paymentController.MakePayment)

		reqBody, _ := json.Marshal(model.PaymentRequest{MerchantID: 1, Amount: 100, Currency: "USD"})
		req := httptest.NewRequest(http.MethodPost, "/payment", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

		var errorResponse model.ErrorResponse
		json.NewDecoder(resp.Body).Decode(&errorResponse)
		assert.Equal(t, "Failed to get user ID", errorResponse.Message)
	})
}
