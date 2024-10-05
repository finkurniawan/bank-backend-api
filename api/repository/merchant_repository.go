package repository

import (
"github.com/finkurniawan/bank-backend-api/api/model"
"gorm.io/gorm"
)

type MerchantRepository interface {
	FindByID(id uint) (*model.Merchant, error)
}

type merchantRepository struct {
	db *gorm.DB
}

func NewMerchantRepository(db *gorm.DB) MerchantRepository {
	return &merchantRepository{db: db}
}

func (r *merchantRepository) FindByID(id uint) (*model.Merchant, error) {
	var merchant model.Merchant
	if err := r.db.First(&merchant, id).Error; err != nil {
		return nil, err
	}
	return &merchant, nil
}
