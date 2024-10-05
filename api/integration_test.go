package api

import (
	"bytes"
	"encoding/json"
	"github.com/finkurniawan/bank-backend-api/database/seeder"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/finkurniawan/bank-backend-api/api/controller"
	"github.com/finkurniawan/bank-backend-api/api/database"
	"github.com/finkurniawan/bank-backend-api/api/middleware"
	"github.com/finkurniawan/bank-backend-api/api/model"
	"github.com/finkurniawan/bank-backend-api/api/repository"
	"github.com/finkurniawan/bank-backend-api/api/service"
	"github.com/stretchr/testify/assert"
)

func TestAPIIntegration(t *testing.T) {
	db, err := database.ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Session(&gorm.Session{})

	seeder.SeedCustomers(db)
	seeder.SeedMerchants(db)

	customerRepository := repository.NewCustomerRepository(db)
	merchantRepository := repository.NewMerchantRepository(db)

	authService := service.NewAuthService(customerRepository)
	paymentService := service.NewPaymentService(customerRepository, merchantRepository, db)

	authController := controller.NewAuthController(authService)
	paymentController := controller.NewPaymentController(paymentService)

	app := fiber.New()
	app.Use(cors.New())

	app.Post("/login", authController.Login)
	app.Post("/logout", middleware.Protected(), authController.Logout)

	app.Post("/payment", middleware.Protected(), paymentController.MakePayment)

	t.Run("Login and Make Payment", func(t *testing.T) {
		loginReqBody, _ := json.Marshal(model.LoginRequest{Username: "user1", Password: "password123"})
		loginReq := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(loginReqBody))
		loginReq.Header.Set("Content-Type", "application/json")

		loginResp, err := app.Test(loginReq)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, loginResp.StatusCode)

		var loginResponse model.LoginResponse
		json.NewDecoder(loginResp.Body).Decode(&loginResponse)
		token := loginResponse.Token

		paymentReqBody, _ := json.Marshal(model.PaymentRequest{MerchantID: 1, Amount: 100, Currency: "USD"})
		paymentReq := httptest.NewRequest(http.MethodPost, "/payment", bytes.NewBuffer(paymentReqBody))
		paymentReq.Header.Set("Content-Type", "application/json")
		paymentReq.Header.Set("Authorization", token)

		paymentResp, err := app.Test(paymentReq)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, paymentResp.StatusCode)
	})
}
