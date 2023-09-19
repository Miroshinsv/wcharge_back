package repo

import (
	"context"
	"fmt"

	"github.com/Miroshinsv/wcharge_back/internal/entity"
	"github.com/Miroshinsv/wcharge_back/pkg/postgres"
)

const _defaultEntityCap = 64

// UserRepo -.
type UserRepo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

// GetHistory -.
func (r *UserRepo) GetUsers(ctx context.Context) ([]entity.User, error) {
	sql, _, err := r.Builder.
		Select("id, username, email, role").
		From("history").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("TranslationRepo - GetHistory - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("TranslationRepo - GetHistory - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	entities := make([]entity.User, 0, _defaultEntityCap)

	for rows.Next() {
		e := entity.User{}

		err = rows.Scan(&e.ID, &e.Username, &e.Email)
		if err != nil {
			return nil, fmt.Errorf("TranslationRepo - GetHistory - rows.Scan: %w", err)
		}

		entities = append(entities, e)
	}

	return entities, nil
}

// Store -.
func (r *UserRepo) Store(ctx context.Context, u entity.User) error {
	sql, args, err := r.Builder.
		Insert("history").
		Columns("id, username, email").
		Values(u.ID, u.Username, u.Email).
		ToSql()
	if err != nil {
		return fmt.Errorf("userRepo - Store - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("userRepo - Store - r.Pool.Exec: %w", err)
	}

	return nil
}
