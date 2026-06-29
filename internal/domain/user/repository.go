package user

import "context"

// Repository defines persistence operations for User.
type Repository interface {
	// Create persists a new user
	Create(ctx context.Context, user User) error

	// FindByEmail returns:
	//   (*User, nil)                -> success
	//   (nil, ErrUserNotFound)      -> user doesn't exist
	//   (nil, err)                  -> infrastructure error
	FindByEmail(ctx context.Context, email string) (*User, error)
}
