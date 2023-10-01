package postgres

import (
	"context"
	"fmt"
	"github.com/Miroshinsv/wcharge_back/internal/entity"
)

func (r *Repo) GetRoleRepo(u entity.Role) (entity.Role, error) {
	sql, _, err := r.Builder.
		Select("id, role_name, priv").
		From("postgres.public.tbl_role").
		Where("postgres.public.tbl_role.id = ?", u.ID).
		ToSql()
	if err != nil {
		return u, fmt.Errorf("UserRepo - GetUsers - r.Builder: %w", err)
	}
	ctx := context.Background()         //!!!
	rows, err := r.Pool.Query(ctx, sql) //!!!
	if err != nil {
		return u, fmt.Errorf("UserRepo - GetUsers - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&u.ID, &u.RoleName, &u.Privileges)
		if err != nil {
			return u, fmt.Errorf("UserRepo - GetUsers - rows.Scan: %w", err)
		}
	}

	return u, nil
}
