package auth

import (
	"os"
	"time"

	"testing"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

var jwtTestKey = []byte(os.Getenv("JWT_KEY_TEST"))

func TestVerifyToken_ValidToken(t *testing.T) {
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		UserID: 123,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})
	testTokenString, err := testToken.SignedString(jwtTestKey)
	assert.NoError(t, err)

	userID, err := VerifyToken(testTokenString)
	assert.NoError(t, err)
	assert.Equal(t, uint(123), userID)
}
