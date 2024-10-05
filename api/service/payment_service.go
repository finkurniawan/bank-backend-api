package service

import (
"fmt"
"time"

"github.com/finkurniawan/bank-backend-api/api/model"
"github.com/finkurniawan/bank-backend-api/api/repository"
"gorm.io/gorm"
)

type PaymentService interface {
	MakePayment(customerID, merchantID uint, amount float64, currency string) error
}

type paymentService struct {
	customerRepository repository.CustomerRepository
	merchantRepository repository.MerchantRepository
	db                 *gorm.DB
}

func NewPaymentService(customerRepository repository.CustomerRepository, merchantRepository repository.MerchantRepository, db *gorm.DB) PaymentService {
	return &paymentService{
		customerRepository: customerRepository,
		merchantRepository: merchantRepository,
		db:                 db,
	}
}

func (s *paymentService) MakePayment(customerID, merchantID uint, amount float64, currency string) error {
	customer, err := s.customerRepository.FindByID(customerID)
	if err != nil {
		return fmt.Errorf("customer not found")
	}

	merchant, err := s.merchantRepository.FindByID(merchantID)
	if err != nil {
		return fmt.Errorf("merchant not found")
	}

	if customer.Balance < amount {
		return fmt.Errorf("insufficient balance")
	}

	tx := s.db.Begin()
	if tx.Error != nil {
		return fmt.Errorf("failed to begin transaction")
	}

	customer.Balance -= amount
	if err := s.customerRepository.UpdateBalance(customer); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update customer balance: %v", err)
	}

	transaction := model.Payment{
		CustomerID:  customerID,
		MerchantID:  merchantID,
		Amount:      amount,
		Currency:    currency,
		Transaction: fmt.Sprintf("Payment to %s at %s", merchant.Name, time.Now().Format(time.RFC3339)),
	}

	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to record transaction: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}
