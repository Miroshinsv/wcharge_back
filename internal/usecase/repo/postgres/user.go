package postgres

import (
	"context"
	"fmt"

	"github.com/Miroshinsv/wcharge_back/internal/entity"
	//"github.com/Miroshinsv/wcharge_back/pkg/postgres"
)

// Store -.
func (r *PostgresRepo) CreateUserRepo(u entity.User) error {

	return nil
}

func (r *PostgresRepo) UpdateUserRepo(u entity.User) error {

	return nil
}

func (r *PostgresRepo) DeleteUserRepo(u entity.User) error {

	return nil
}

func (r *PostgresRepo) GetUserRepo(u entity.User) (entity.User, error) {

	return entity.User{}, nil
}

// GetHistory -.
func (r *PostgresRepo) GetUsersRepo() ([]entity.User, error) {
	sql, _, err := r.Builder.
		Select("id, username, email, role").
		From("users").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("UserRepo - GetUsers - r.Builder: %w", err)
	}
	ctx := context.Background()         //!!!
	rows, err := r.Pool.Query(ctx, sql) //!!!
	if err != nil {
		return nil, fmt.Errorf("UserRepo - GetUsers - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	entities := make([]entity.User, 0, _defaultEntityCap)

	for rows.Next() {
		e := entity.User{}

		err = rows.Scan(&e.ID, &e.Username, &e.Email)
		if err != nil {
			return nil, fmt.Errorf("UserRepo - GetUsers - rows.Scan: %w", err)
		}

		entities = append(entities, e)
	}

	return entities, nil
}
