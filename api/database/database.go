package database

import (
	"fmt"
	"github.com/finkurniawan/bank-backend-api/api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := fmt.Sprint(
		"host=127.0.0.1 user=postgres password=postgres dbname=bank_api port=5432 sslmode=disable TimeZone=Asia/Jakarta")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.Customer{}, &model.Merchant{}, &model.Payment{})
	if err != nil {
		return nil, err
	}

	log.Println("Database connection successful")
	return db, nil
}
