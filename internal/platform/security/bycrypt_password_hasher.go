package security

import (
	"golang.org/x/crypto/bcrypt"
)

// BcryptPasswordHasher implements user.PasswordHasher using bcrypt.
type BcryptPasswordHasher struct{}

// NewBcryptPasswordHasher creates a new bcrypt password hasher.
func NewBcryptPasswordHasher() *BcryptPasswordHasher {
	return &BcryptPasswordHasher{}
}

// Hash hashes a plain-text password.
func (h *BcryptPasswordHasher) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// Compare compares a plain-text password against a bcrypt hash.
func (h *BcryptPasswordHasher) Compare(
	password string,
	hash string,
) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)
}
