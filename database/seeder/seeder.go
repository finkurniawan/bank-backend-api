package seeder

import (
	"fmt"
	"log"

	"github.com/finkurniawan/bank-backend-api/api/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedCustomers(db *gorm.DB) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)

	customers := []model.Customer{
		{Username: "user1", Password: string(hashedPassword), Balance: 1000},
		{Username: "user2", Password: string(hashedPassword), Balance: 2500},
		{Username: "user3", Password: string(hashedPassword), Balance: 5000},
	}

	for _, customer := range customers {
		if err := db.Create(&customer).Error; err != nil {
			log.Fatalf("Failed to seed customers: %v", err)
		}
	}

	fmt.Println("Customers seeded successfully!")
}

func SeedMerchants(db *gorm.DB) {
	merchants := []model.Merchant{
		{Name: "Merchant A"},
		{Name: "Merchant B"},
		{Name: "Merchant C"},
	}

	for _, merchant := range merchants {
		if err := db.Create(&merchant).Error; err != nil {
			log.Fatalf("Failed to seed merchants: %v", err)
		}
	}

	fmt.Println("Merchants seeded successfully!")
}
