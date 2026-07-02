package command

import (
	"context"
	"errors"
	"net/mail"
	"strings"

	"go-react-router/internal/domain/identity"
	"go-react-router/internal/domain/identity/valueobject"
	"go-react-router/internal/domain/identity/aggregate"
)

// RegisterUserCommandHandler handles identity registration.
type RegisterUserCommandHandler struct {
	repository identity.Repository
	hasher     valueobject.PasswordHasher
}

// NewRegisterUserCommandHandler creates a new RegisterUserCommandHandler.
func NewRegisterUserCommandHandler(
	repository identity.Repository,
	hasher valueobject.PasswordHasher,
) *RegisterUserCommandHandler {
	return &RegisterUserCommandHandler{
		repository: repository,
		hasher:     hasher,
	}
}

// Execute registers a new identity.
func (h *RegisterUserCommandHandler) Execute(
	ctx context.Context,
	registration aggregate.Registration,
) error {

	if err := h.validate(&registration); err != nil {
		return err
	}

	if err := h.ensureEmailIsAvailable(ctx, registration.Email); err != nil {
		return err
	}

	passwordHash, err := h.hasher.Hash(registration.Password)
	if err != nil {
		return err
	}

	newUser := aggregate.User{
		Email:        registration.Email,
		PasswordHash: passwordHash,
		FullName:     registration.FullName,
	}

	return h.repository.Create(ctx, newUser)
}

// validate normalizes and validates the registration.
func (uc *RegisterUserCommandHandler) validate(
	registration *aggregate.Registration,
) error {

	registration.Email = strings.TrimSpace(registration.Email)
	registration.FullName = strings.TrimSpace(registration.FullName)

	switch {
	case registration.Email == "":
		return identity.ErrEmailRequired

	case registration.Password == "":
		return identity.ErrPasswordRequired

	case registration.FullName == "":
		return identity.ErrFullNameRequired
	}

	if _, err := mail.ParseAddress(registration.Email); err != nil {
		return identity.ErrInvalidEmail
	}

	return nil
}

// ensureEmailIsAvailable checks whether the email is already registered.
func (h *RegisterUserCommandHandler) ensureEmailIsAvailable(
	ctx context.Context,
	email string,
) error {

	_, err := h.repository.FindByEmail(ctx, email)

	switch {
	case err == nil:
		return identity.ErrEmailAlreadyExists

	case errors.Is(err, identity.ErrUserNotFound):
		return nil

	default:
		return err
	}
}
