package token

import "time"

// Maker is an interface for creating tokens.
type Maker interface {
	// CreateToken creates a token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	// VerifyToken verifies a token and returns the username
	VerifyToken(token string) (*Payload, error)
}
