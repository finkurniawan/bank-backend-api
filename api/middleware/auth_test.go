package middleware

import (
"net/http"
"net/http/httptest"
"testing"

"github.com/gofiber/fiber/v2"
"github.com/finkurniawan/bank-backend-api/api/utils"
"github.com/golang-jwt/jwt/v5"
"github.com/stretchr/testify/assert"
)

func TestProtected(t *testing.T) {
	app := fiber.New()
	app.Get("/protected", Protected(), func(c *fiber.Ctx) error {
		return c.SendString("Protected route")
	})

	t.Run("Valid Token", func(t *testing.T) {
token, _ := utils.GenerateToken("1")

req := httptest.NewRequest(http.MethodGet, "/protected", nil)
req.Header.Set("Authorization", token)
resp, err := app.Test(req)

assert.NoError(t, err)
assert.Equal(t, http.StatusOK, resp.StatusCode)
})

	t.Run("Invalid Token", func(t *testing.T) {
req := httptest.NewRequest(http.MethodGet, "/protected", nil)
req.Header.Set("Authorization", "invalidtoken")
resp, err := app.Test(req)

assert.NoError(t, err)
assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
})

	t.Run("Missing Token", func(t *testing.T) {
req := httptest.NewRequest(http.MethodGet, "/protected", nil)
resp, err := app.Test(req)

assert.NoError(t, err)
assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
})
}

func TestGetUserIDFromClaims(t *testing.T) {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
claims := jwt.MapClaims{"iss": "1"}
c.Locals("claims", claims)

userID, err := GetUserIDFromClaims(c)
assert.NoError(t, err)
assert.Equal(t, uint(1), userID)

return c.SendStatus(http.StatusOK)
})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
