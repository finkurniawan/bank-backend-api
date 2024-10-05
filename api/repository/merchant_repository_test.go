package repository

import (
"testing"

"github.com/finkurniawan/bank-backend-api/api/model"
"gorm.io/driver/postgres"
"gorm.io/gorm"
)

func TestMerchantRepository_FindByID(t *testing.T) {
	dsn := "host=localhost user=postgres password=password dbname=bank_db_test port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	db.AutoMigrate(&model.Merchant{})

	merchant := model.Merchant{Name: "Test Merchant"}
	db.Create(&merchant)

	repo := NewMerchantRepository(db)

	retrievedMerchant, err := repo.FindByID(merchant.ID)
	if err != nil {
		t.Errorf("FindByID failed: %v", err)
	}
	if retrievedMerchant.Name != merchant.Name {
		t.Errorf("Expected merchant name %s, got %s", merchant.Name, retrievedMerchant.Name)
	}
}
