package controller

import (
	"github.com/finkurniawan/bank-backend-api/api/model"
	"github.com/finkurniawan/bank-backend-api/api/service"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	var req model.LoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{Message: "Invalid request"})
	}

	token, err := c.authService.Login(req.Username, req.Password)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{Message: err.Error()})
	}

	return ctx.JSON(model.LoginResponse{Token: token})
}

func (c *AuthController) Logout(ctx *fiber.Ctx) error {
	return ctx.JSON(model.SuccessResponse{Message: "Successfully logged out"})
}
