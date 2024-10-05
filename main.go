package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"

	_ "github.com/arsmn/fiber-swagger/v2"

	"github.com/finkurniawan/bank-backend-api/api/controller"
	"github.com/finkurniawan/bank-backend-api/api/database"
	"github.com/finkurniawan/bank-backend-api/api/middleware"
	"github.com/finkurniawan/bank-backend-api/api/repository"
	"github.com/finkurniawan/bank-backend-api/api/service"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Session(&gorm.Session{})

	customerRepository := repository.NewCustomerRepository(db)
	merchantRepository := repository.NewMerchantRepository(db)

	authService := service.NewAuthService(customerRepository)
	paymentService := service.NewPaymentService(customerRepository, merchantRepository, db)

	authController := controller.NewAuthController(authService)
	paymentController := controller.NewPaymentController(paymentService)

	app := fiber.New()
	app.Use(cors.New())

	app.Post("/login", authController.Login)
	app.Post("/logout", middleware.Protected(), authController.Logout)

	app.Post("/payment", middleware.Protected(), paymentController.MakePayment)

	app.Get("/swagger/*", swagger.HandlerDefault)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(":" + port))
}
