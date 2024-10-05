package service_test

//
//import (
//	"fmt"
//	"github.com/finkurniawan/bank-backend-api/api/service"
//	"testing"
//
//	"github.com/finkurniawan/bank-backend-api/api/model"
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/mock"
//	"gorm.io/gorm"
//)
//
//type MockMerchantRepository struct {
//	mock.Mock
//}
//
//func (m *MockMerchantRepository) FindByID(id uint) (*model.Merchant, error) {
//	args := m.Called(id)
//	return args.Get(0).(*model.Merchant), args.Error(1)
//}
//
//type MockDB struct {
//	mock.Mock
//}
//
//func (m *MockDB) Begin() *gorm.DB {
//	args := m.Called()
//	return args.Get(0).(*gorm.DB)
//}
//
//func (m *MockDB) Create(value interface{}) *gorm.DB {
//	args := m.Called(value)
//	return args.Get(0).(*gorm.DB)
//}
//
//func (m *MockDB) Commit() *gorm.DB {
//	args := m.Called()
//	return args.Get(0).(*gorm.DB)
//}
//
//func (m *MockDB) Rollback() *gorm.DB {
//	args := m.Called()
//	return args.Get(0).(*gorm.DB)
//}
//
//func TestPaymentService_MakePayment(t *testing.T) {
//	customer := &model.Customer{
//		ID:       1,
//		Username: "testuser",
//		Balance:  1000,
//	}
//
//	merchant := &model.Merchant{
//		ID:   1,
//		Name: "Test Merchant",
//	}
//
//	mockCustomerRepo := new(MockCustomerRepository)
//	mockCustomerRepo.On("FindByID", uint(1)).Return(customer, nil)
//	mockCustomerRepo.On("UpdateBalance", customer).Return(nil)
//
//	mockMerchantRepo := new(MockMerchantRepository)
//	mockMerchantRepo.On("FindByID", uint(1)).Return(merchant, nil)
//
//	mockDB := new(MockDB)
//	mockDB.On("Begin").Return(mockDB)
//	mockDB.On("Create", mock.Anything).Return(mockDB)
//	mockDB.On("Commit").Return(mockDB)
//
//	paymentService := service.NewPaymentService(mockCustomerRepo, mockMerchantRepo, mockDB)
//
//	t.Run("Successful Payment", func(t *testing.T) {
//		err := paymentService.MakePayment(1, 1, 100, "USD")
//		assert.NoError(t, err)
//	})
//
//	t.Run("Insufficient Balance", func(t *testing.T) {
//		err := paymentService.MakePayment(1, 1, 1500, "USD")
//		assert.Error(t, err)
//		assert.EqualError(t, err, "insufficient balance")
//	})
//
//	t.Run("Customer Not Found", func(t *testing.T) {
//		mockCustomerRepo.On("FindByID", uint(2)).Return(nil, fmt.Errorf("record not found"))
//		err := paymentService.MakePayment(2, 1, 100, "USD")
//		assert.Error(t, err)
//		assert.EqualError(t, err, "customer not found")
//	})
//
//	t.Run("Merchant Not Found", func(t *testing.T) {
//		mockMerchantRepo.On("FindByID", uint(2)).Return(nil, fmt.Errorf("record not found"))
//		err := paymentService.MakePayment(1, 2, 100, "USD")
//		assert.Error(t, err)
//		assert.EqualError(t, err, "merchant not found")
//	})
//}
