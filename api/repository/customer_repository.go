package repository

import (
"github.com/finkurniawan/bank-backend-api/api/model"
"gorm.io/gorm"
)

type CustomerRepository interface {
	FindByUsername(username string) (*model.Customer, error)
	FindByID(id uint) (*model.Customer, error)
	UpdateBalance(customer *model.Customer) error
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) FindByUsername(username string) (*model.Customer, error) {
	var customer model.Customer
	if err := r.db.Where("username = ?", username).First(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *customerRepository) FindByID(id uint) (*model.Customer, error) {
	var customer model.Customer
	if err := r.db.First(&customer, id).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *customerRepository) UpdateBalance(customer *model.Customer) error {
	return r.db.Model(customer).Update("balance", customer.Balance).Error
}
