package identity

import (
	"context"
	"go-react-router/internal/domain/identity/aggregate"
)

// Repository defines persistence operations for User.
type Repository interface {
	// Create persists a new user
	Create(ctx context.Context, user aggregate.User) error

	// FindByEmail returns:
	//   (*User, nil)                -> success
	//   (nil, ErrUserNotFound)      -> user doesn't exist
	//   (nil, err)                  -> infrastructure error
	FindByEmail(ctx context.Context, email string) (*aggregate.User, error)
}
