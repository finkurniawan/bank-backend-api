package repository

import (
"testing"

"github.com/finkurniawan/bank-backend-api/api/model"
"gorm.io/driver/postgres"
"gorm.io/gorm"
)

func TestCustomerRepository_FindByUsername(t *testing.T) {
	dsn := "host=localhost user=postgres password=password dbname=bank_db_test port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	db.AutoMigrate(&model.Customer{})

	customer := model.Customer{Username: "testuser", Password: "testpassword", Balance: 1000}
	db.Create(&customer)

	repo := NewCustomerRepository(db)

	retrievedCustomer, err := repo.FindByUsername("testuser")
	if err != nil {
		t.Errorf("FindByUsername failed: %v", err)
	}
	if retrievedCustomer.ID != customer.ID {
		t.Errorf("Expected customer ID %d, got %d", customer.ID, retrievedCustomer.ID)
	}
}

func TestCustomerRepository_FindByID(t *testing.T) {
	dsn := "host=localhost user=postgres password=password dbname=bank_db_test port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	db.AutoMigrate(&model.Customer{})

	customer := model.Customer{Username: "testuser", Password: "testpassword", Balance: 1000}
	db.Create(&customer)

	repo := NewCustomerRepository(db)

	retrievedCustomer, err := repo.FindByID(customer.ID)
	if err != nil {
		t.Errorf("FindByID failed: %v", err)
	}
	if retrievedCustomer.Username != customer.Username {
		t.Errorf("Expected customer username %s, got %s", customer.Username, retrievedCustomer.Username)
	}
}

func TestCustomerRepository_UpdateBalance(t *testing.T) {
	dsn := "host=localhost user=postgres password=password dbname=bank_db_test port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	db.AutoMigrate(&model.Customer{})

	customer := model.Customer{Username: "testuser", Password: "testpassword", Balance: 1000}
	db.Create(&customer)

	repo := NewCustomerRepository(db)

	customer.Balance = 1500
	err = repo.UpdateBalance(&customer)
	if err != nil {
		t.Errorf("UpdateBalance failed: %v", err)
	}

	var updatedCustomer model.Customer
	db.First(&updatedCustomer, customer.ID)
	if updatedCustomer.Balance != 1500 {
		t.Errorf("Expected updated balance to be 1500, got %f", updatedCustomer.Balance)
	}
}
