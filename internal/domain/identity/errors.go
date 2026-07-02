package identity

import "errors"

var (
	// Repository errors.
	ErrUserNotFound      = errors.New("user not found")
	ErrEmailAlreadyExists = errors.New("email already exists")

	// Validation errors.
	ErrEmailRequired    = errors.New("email is required")
	ErrInvalidEmail     = errors.New("invalid email address")
	ErrPasswordRequired = errors.New("password is required")
	ErrFullNameRequired = errors.New("full name is required")
)
