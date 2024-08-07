package domain

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	jwt.RegisteredClaims
}

// Valid implements jwt.Claims.
func (c *Claims) Valid() error {
	// Periksa apakah token sudah kadaluarsa
	if c.ExpiresAt != nil && c.ExpiresAt.Before(time.Now()) {
		return errors.New("token is expired")
	}
	return nil
}
