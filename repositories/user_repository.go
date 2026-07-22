package repositories

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/adsomniac/my-app/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
		SELECT
			id,
			email,
			password_hash,
			role,
			is_active,
			created_at,
			updated_at
		FROM users
		WHERE LOWER(email) = LOWER($1)
		LIMIT 1;
	`

	var user models.User
	err := r.DB.QueryRowContext(ctx, query, strings.TrimSpace(email)).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
