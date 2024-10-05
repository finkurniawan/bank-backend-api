package middleware

import (
"strconv"

"github.com/gofiber/fiber/v2"
"github.com/finkurniawan/bank-backend-api/api/utils"
)

func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")

		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
		}

		claims, err := utils.DecodeToken(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
		}

		c.Locals("claims", claims)
		return c.Next()
	}
}

func GetUserIDFromClaims(ctx *fiber.Ctx) (uint, error) {
	claims, ok := ctx.Locals("claims").(utils.Claims)
	if !ok {
		return 0, fiber.ErrInternalServerError
	}

	userID, err := strconv.ParseUint(claims.Issuer, 10, 32)
	if err != nil {
		return 0, fiber.ErrInternalServerError
	}

	return uint(userID), nil
}
