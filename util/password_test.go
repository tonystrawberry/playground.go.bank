package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

// TestPassword tests the password hashing and checking.
func TestPassword(t *testing.T) {
	password := RandomString(6)

	hashedPassword1, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)

	match := CheckPassword(password, hashedPassword1)
	require.NoError(t, match)

	nonMatch := CheckPassword(password+"1", hashedPassword1)
	require.Error(t, nonMatch)
	require.Equal(t, nonMatch, bcrypt.ErrMismatchedHashAndPassword)

	hashedPassword2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword2)
	require.NotEqual(t, hashedPassword1, hashedPassword2)
}
