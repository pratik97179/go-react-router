package database

import (
	"context"
	"errors"
	"strings"
	"time"

	"go-react-router/internal/domain/identity"
	"go-react-router/internal/domain/identity/aggregate"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// UserRepository implements identity.Repository using PostgreSQL.
type UserRepository struct {
	db *pgxpool.Pool
}

// userRecord represents the persisted form of a User.
type userRecord struct {
	ID           string    `db:"id"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	FullName     string    `db:"full_name"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

// NewUserRepository creates a new UserRepository.
func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func toRecord(u aggregate.User) userRecord {
	return userRecord{
		ID:           u.ID,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		FullName:     u.FullName,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}

func toDomain(r userRecord) aggregate.User {
	return aggregate.User{
		ID:           r.ID,
		Email:        r.Email,
		PasswordHash: r.PasswordHash,
		FullName:     r.FullName,
		CreatedAt:    r.CreatedAt,
		UpdatedAt:    r.UpdatedAt,
	}
}

// Create persists a new identity.
func (r *UserRepository) Create(
	ctx context.Context,
	u aggregate.User,
) error {

	record := toRecord(u)

	err := r.db.QueryRow(
		ctx,
		`
		INSERT INTO users (
			email,
			password_hash,
			full_name
		)
		VALUES ($1, $2, $3)
		RETURNING
			id,
			email,
			password_hash,
			full_name,
			created_at,
			updated_at
		`,
		record.Email,
		record.PasswordHash,
		record.FullName,
	).Scan(
		&record.ID,
		&record.Email,
		&record.PasswordHash,
		&record.FullName,
		&record.CreatedAt,
		&record.UpdatedAt,
	)

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) &&
			pgErr.Code == "23505" &&
			strings.Contains(pgErr.ConstraintName, "email") {

			return identity.ErrEmailAlreadyExists
		}

		return err
	}

	return nil
}

// FindByEmail returns a identity by email.
func (r *UserRepository) FindByEmail(
	ctx context.Context,
	email string,
) (*aggregate.User, error) {

	var record userRecord

	err := r.db.QueryRow(
		ctx,
		`
		SELECT
			id,
			email,
			password_hash,
			full_name,
			created_at,
			updated_at
		FROM users
		WHERE email = $1
		`,
		email,
	).Scan(
		&record.ID,
		&record.Email,
		&record.PasswordHash,
		&record.FullName,
		&record.CreatedAt,
		&record.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, identity.ErrUserNotFound
		}

		return nil, err
	}

	domainUser := toDomain(record)

	return &domainUser, nil
}
