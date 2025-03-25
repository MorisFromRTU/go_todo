package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	var passwordTest string = "somepassword"
	hashedPasswordTest, err := HashPassword(passwordTest)
	assert.NoError(t, err)
	isTestPasswordCorrect := CheckPasswordHash(passwordTest, hashedPasswordTest)
	assert.Equal(t, true, isTestPasswordCorrect)
}
