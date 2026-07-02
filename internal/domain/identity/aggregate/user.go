package aggregate

import "time"

type User struct {
	ID           string
	Email        string
	PasswordHash string
	FullName     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Registration contains the information required to register a new user.
type Registration struct {
    Email    string `json:"email"`
    Password string `json:"password"`
    FullName string `json:"full_name"`
}
