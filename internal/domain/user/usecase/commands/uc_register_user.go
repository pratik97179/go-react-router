package commands

import (
	"context"
	"errors"
	"net/mail"
	"strings"

	"go-react-router/internal/domain/user"
)

// RegisterUserUseCase handles user registration.
type RegisterUserUseCase struct {
	repository user.Repository
	hasher     user.PasswordHasher
}

// NewRegisterUserUseCase creates a new RegisterUserUseCase.
func NewRegisterUserUseCase(
	repository user.Repository,
	hasher user.PasswordHasher,
) *RegisterUserUseCase {
	return &RegisterUserUseCase{
		repository: repository,
		hasher:     hasher,
	}
}

// Execute registers a new user.
func (uc *RegisterUserUseCase) Execute(
	ctx context.Context,
	registration user.Registration,
) error {

	if err := uc.validate(&registration); err != nil {
		return err
	}

	if err := uc.ensureEmailIsAvailable(ctx, registration.Email); err != nil {
		return err
	}

	passwordHash, err := uc.hasher.Hash(registration.Password)
	if err != nil {
		return err
	}

	newUser := user.User{
		Email:        registration.Email,
		PasswordHash: passwordHash,
		FullName:     registration.FullName,
	}

	return uc.repository.Create(ctx, newUser)
}

// validate normalizes and validates the registration.
func (uc *RegisterUserUseCase) validate(
	registration *user.Registration,
) error {

	registration.Email = strings.TrimSpace(registration.Email)
	registration.FullName = strings.TrimSpace(registration.FullName)

	switch {
	case registration.Email == "":
		return user.ErrEmailRequired

	case registration.Password == "":
		return user.ErrPasswordRequired

	case registration.FullName == "":
		return user.ErrFullNameRequired
	}

	if _, err := mail.ParseAddress(registration.Email); err != nil {
		return user.ErrInvalidEmail
	}

	return nil
}

// ensureEmailIsAvailable checks whether the email is already registered.
func (uc *RegisterUserUseCase) ensureEmailIsAvailable(
	ctx context.Context,
	email string,
) error {

	_, err := uc.repository.FindByEmail(ctx, email)

	switch {
	case err == nil:
		return user.ErrEmailAlreadyExists

	case errors.Is(err, user.ErrUserNotFound):
		return nil

	default:
		return err
	}
}
