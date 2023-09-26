package postgres

import (
	"context"
	"fmt"

	"github.com/Miroshinsv/wcharge_back/internal/entity"
	//"github.com/Miroshinsv/wcharge_back/pkg/postgres"
)

// Store -.
func (r *PostgresRepo) CreateUserRepo(u entity.User) error {
	sql, args, err := r.Builder.
		Insert("postgres.public.tbl_users").
		Columns("username, email, role_id").
		Values(u.Username, u.Email, u.RoleID).
		ToSql()
	if err != nil {
		return fmt.Errorf("PostgresRepo - CreateUserRepo - r.Builder: %w", err)
	}
	ctx := context.Background()
	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("PostgresRepo - CreateUserRepo - r.Pool.Exec: %w", err)
	}

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

func (r *PostgresRepo) GetUsersRepo() ([]entity.User, error) {
	sql, _, err := r.Builder.
		Select("id, username, email, role_id, created_at, updated_at, deleted_at").
		From("postgres.public.tbl_users").
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
		err = rows.Scan(&e.ID, &e.Username, &e.Email, &e.RoleID, &e.CreateAt, &e.UpdateAt, &e.DeleteAt)
		if err != nil {
			return nil, fmt.Errorf("UserRepo - GetUsers - rows.Scan: %w", err)
		}
		entities = append(entities, e)
	}

	return entities, nil
}
