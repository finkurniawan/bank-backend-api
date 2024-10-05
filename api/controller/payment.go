package controller

import (
	"github.com/finkurniawan/bank-backend-api/api/middleware"
	"github.com/finkurniawan/bank-backend-api/api/model"
	"github.com/finkurniawan/bank-backend-api/api/service"
	"github.com/gofiber/fiber/v2"
)

type PaymentController struct {
	paymentService service.PaymentService
}

func NewPaymentController(paymentService service.PaymentService) *PaymentController {
	return &PaymentController{paymentService: paymentService}
}

func (c *PaymentController) MakePayment(ctx *fiber.Ctx) error {
	var req model.PaymentRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{Message: "Invalid request"})
	}

	if req.MerchantID == 0 || req.Amount <= 0 || req.Currency == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{Message: "Invalid payment details"})
	}

	userID, err := middleware.GetUserIDFromClaims(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{Message: "Failed to get user ID"})
	}

	err = c.paymentService.MakePayment(userID, req.MerchantID, req.Amount, req.Currency)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{Message: err.Error()})
	}

	return ctx.JSON(model.SuccessResponse{Message: "Payment successful"})
}
