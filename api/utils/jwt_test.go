package utils

import (
"os"
"testing"

"github.com/stretchr/testify/assert"
)

func TestJWT(t *testing.T) {
	os.Setenv("JWT_SECRET", "test-secret")

	t.Run("Generate and Decode Token", func(t *testing.T) {
issuer := "1"
token, err := GenerateToken(issuer)
assert.NoError(t, err)
assert.NotEmpty(t, token)

claims, err := DecodeToken(token)
assert.NoError(t, err)
assert.Equal(t, issuer, claims.Issuer)
})

	t.Run("Decode Invalid Token", func(t *testing.T) {
_, err := DecodeToken("invalidtoken")
assert.Error(t, err)
})
}
