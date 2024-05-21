package postgres

import (
	"context"
	"fmt"
	"github.com/Miroshinsv/wcharge_back/internal/entity"
)

func (r *Repo) GetRoleRepo(id int) (entity.Role, error) {
	u := entity.Role{
		ID: id,
	}
	sql, args, err := r.Builder.
		Select("id, role_name, role_privileges").
		From("tbl_role").
		Where("tbl_role.id = ?", id).
		ToSql()
	if err != nil {
		return u, fmt.Errorf("Repo - GetRoleRepo - r.Builder: %w", err)
	}
	ctx := context.Background()
	row := r.Pool.QueryRow(ctx, sql, args...)
	if err != nil {
		return u, fmt.Errorf("Repo - GetRoleRepo - r.Pool.Query: %w", err)
	}

	err = row.Scan(&u.ID, &u.RoleName, &u.RolePrivileges)
	if err != nil {
		return u, fmt.Errorf("Repo - GetRoleRepo - rows.Scan: %w", err)
	}

	return u, nil
}

func (r *Repo) GetRolesRepo() ([]entity.Role, error) {
	sql, args, err := r.Builder.
		Select("id, role_name, role_privileges").
		From("postgres.public.tbl_role").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("Repo - GetRolesRepo - r.Builder: %w", err)
	}
	ctx := context.Background()
	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("Repo - GetRolesRepo - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	roles := make([]entity.Role, 0, _defaultEntityCap)

	for rows.Next() {
		role := entity.Role{}
		err = rows.Scan(&role.ID, &role.RoleName, &role.RolePrivileges)
		if err != nil {
			return nil, fmt.Errorf("Repo - GetRolesRepo - rows.Scan: %w", err)
		}
		roles = append(roles, role)
	}

	return roles, nil
}
