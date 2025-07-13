package auth

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestCheckHashpassword(t *testing.T) {
	password1 := "correctPassword123!"
	password2 := "differentPassword123!"

	hash1, _ := HashPassword(password1)
	hash2, _ := HashPassword(password2)
	
	// Test: Password provided should match hash
	err := CheckPasswordHash(hash1, password1)
	require.NoError(t, err)

	// Test: Password provided does not match hash
	err = CheckPasswordHash(hash2, password1)
	require.Error(t, err)
}
