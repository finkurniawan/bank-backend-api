package service_test

import (
	"fmt"
	"github.com/finkurniawan/bank-backend-api/api/service"
	"testing"

	"github.com/finkurniawan/bank-backend-api/api/model"
	_ "github.com/finkurniawan/bank-backend-api/api/repository"
	_ "github.com/finkurniawan/bank-backend-api/api/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

type MockCustomerRepository struct {
	mock.Mock
}

func (m *MockCustomerRepository) FindByUsername(username string) (*model.Customer, error) {
	args := m.Called(username)
	return args.Get(0).(*model.Customer), args.Error(1)
}

func (m *MockCustomerRepository) FindByID(id uint) (*model.Customer, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Customer), args.Error(1)
}

func (m *MockCustomerRepository) UpdateBalance(customer *model.Customer) error {
	args := m.Called(customer)
	return args.Error(0)
}

func TestAuthService_Login(t *testing.T) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)

	customer := &model.Customer{
		ID:       1,
		Username: "testuser",
		Password: string(hashedPassword),
		Balance:  1000,
	}

	mockRepo := new(MockCustomerRepository)
	mockRepo.On("FindByUsername", "testuser").Return(customer, nil)

	authService := service.NewAuthService(mockRepo)

	t.Run("Successful Login", func(t *testing.T) {
		token, err := authService.Login("testuser", "password123")
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
	})

	t.Run("Invalid Password", func(t *testing.T) {
		_, err := authService.Login("testuser", "wrongpassword")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid username or password")
	})

	t.Run("User Not Found", func(t *testing.T) {
		mockRepo.On("FindByUsername", "nonexistentuser").Return(nil, fmt.Errorf("record not found"))
		_, err := authService.Login("nonexistentuser", "password123")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid username or password")
	})
}
