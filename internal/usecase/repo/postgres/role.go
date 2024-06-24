package postgres

import (
	"context"
	"github.com/Miroshinsv/wcharge_back/internal/entity"
	"log"
)

func (r *Repo) GetRole(id int) (*entity.Role, error) {
	sql, args, err := r.Builder.
		Select("id, name, privileges").
		From(rolesTableName).
		Where("role.id = ?", id).
		ToSql()
	if err != nil {
		log.Printf("Repo - GetRole - r.Builder: %s", err)
		return nil, err
	}

	row := r.Pool.QueryRow(context.Background(), sql, args...)
	u := entity.Role{}
	err = row.Scan(
		&u.ID,
		&u.Name,
		&u.Privileges,
	)
	if err != nil {
		log.Printf("Repo - GetRole - rows.Scan: %s", err)
		return nil, err
	}

	return &u, nil
}

func (r *Repo) GetRoles() (*[]entity.Role, error) {
	sql, args, err := r.Builder.
		Select("id, name, privileges").
		From("postgres.public.role").
		ToSql()
	if err != nil {
		log.Printf("Repo - GetRolesRepo - r.Builder: %s", err)
		return nil, err
	}

	rows, err := r.Pool.Query(context.Background(), sql, args...)
	if err != nil {
		log.Printf("Repo - GetRolesRepo - r.Pool.Query: %s", err)
		return nil, err
	}
	defer rows.Close()

	roles := make([]entity.Role, 0, _defaultEntityCap)
	for rows.Next() {
		role := entity.Role{}
		err = rows.Scan(
			&role.ID,
			&role.Name,
			&role.Privileges,
		)
		if err != nil {
			log.Printf("Repo - GetRolesRepo - rows.Scan: %s", err)
			return nil, err
		}
		roles = append(roles, role)
	}

	return &roles, nil
}
