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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) Login(username, password string) (string, error) {
	args := m.Called(username, password)
	return args.String(0), args.Error(1)
}

func TestAuthController_Login(t *testing.T) {
	mockService := new(MockAuthService)
	authController := controller.NewAuthController(mockService)

	app := fiber.New()
	app.Post("/login", authController.Login)

	t.Run("Successful Login", func(t *testing.T) {
		mockService.On("Login", "testuser", "password123").Return("testtoken", nil)

		reqBody, _ := json.Marshal(model.LoginRequest{Username: "testuser", Password: "password123"})
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var loginResponse model.LoginResponse
		json.NewDecoder(resp.Body).Decode(&loginResponse)
		assert.Equal(t, "testtoken", loginResponse.Token)
	})

	t.Run("Invalid Credentials", func(t *testing.T) {
		mockService.On("Login", "testuser", "wrongpassword").Return("", fmt.Errorf("invalid username or password"))

		reqBody, _ := json.Marshal(model.LoginRequest{Username: "testuser", Password: "wrongpassword"})
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)

		var errorResponse model.ErrorResponse
		json.NewDecoder(resp.Body).Decode(&errorResponse)
		assert.Equal(t, "invalid username or password", errorResponse.Message)
	})

	t.Run("Invalid Request Body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		var errorResponse model.ErrorResponse
		json.NewDecoder(resp.Body).Decode(&errorResponse)
		assert.Equal(t, "Invalid request", errorResponse.Message)
	})
}

func TestAuthController_Logout(t *testing.T) {
	mockService := new(MockAuthService)
	authController := controller.NewAuthController(mockService)

	app := fiber.New()
	app.Post("/logout", authController.Logout)

	t.Run("Successful Logout", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/logout", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var successResponse model.SuccessResponse
		json.NewDecoder(resp.Body).Decode(&successResponse)
		assert.Equal(t, "Successfully logged out", successResponse.Message)
	})
}
