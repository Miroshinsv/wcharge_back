package postgres

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/Miroshinsv/wcharge_back/internal/entity"
	"log"
)

func (r *Repo) GetUserByName(userName string) (*entity.User, error) {
	sql, args, err := r.Builder.
		Select("id, username, email, role, phone, addresses.*, roles.*").
		From(usersTableName).
		Join("addresses on users.address = addresses.id").
		Join("roles on users.role = roles.id").
		Where(squirrel.Eq{"username": userName}).
		ToSql()
	if err != nil {
		log.Printf("Error - Postgres - GetUser - r.Builder: %s", err)
		return nil, err
	}

	row := r.Pool.QueryRow(context.Background(), sql, args...)
	u := entity.User{}
	err = row.Scan(
		&u.ID,
		&u.Username,
		&u.Email,
		&u.Phone,

		&u.AddressFull.ID,
		&u.AddressFull.City,
		&u.AddressFull.Country,
		&u.AddressFull.Address,
		&u.AddressFull.Lng,
		&u.AddressFull.Lat,

		&u.RoleFull.Name,
		&u.RoleFull.ID,
		&u.RoleFull.Privileges,
	)
	if err != nil {
		log.Printf("Error - User - GetUsers - rows.Scan: %s", err)
		return nil, err
	}

	return &u, nil
}

func (r *Repo) CreateUser(u entity.User) (*entity.User, error) {
	err := u.BeforeCreate()
	if err != nil {
		log.Printf("Error - Postgres - CreateUser - u.BeforeCreate: %s", err)
		return nil, err
	}

	sql, args, err := r.Builder.
		Insert(usersTableName).
		Columns("username, email, phone, password_hash, password_salt").
		Values(u.Username, u.Email, u.Phone, u.PasswordHash, u.PasswordSalt).
		ToSql()
	if err != nil {
		log.Printf("Error - Postgres - CreateUser - r.Builder: %s", err)
		return nil, err
	}

	row := r.Pool.QueryRow(context.Background(), sql, args...)
	uu := entity.User{}
	err = row.Scan(
		&uu.ID,
		&uu.Username,
		&uu.Email,
		&uu.Phone,
	)
	if err != nil {
		log.Printf("Error - Postgres - CreateUser - r.Pool.Exec: %s", err)
		return nil, err
	}

	return &uu, nil
}

func (r *Repo) UpdateUser(u entity.User, id int) (*entity.User, error) {
	sql, args, err := r.Builder.
		Update(usersTableName).
		Set("username", u.Username).
		Set("email", u.Email).
		Set("phone", u.Phone).
		Set("address_id", u.Address).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		log.Printf("Error - Postgres - UpdateUser - r.Builder: %s", err)
		return nil, err
	}

	row := r.Pool.QueryRow(context.Background(), sql, args...)
	err = row.Scan(
		&u.ID,
		&u.Username,
		&u.Email,
		&u.Phone,
		&u.Address,
	)
	if err != nil {
		log.Printf("Error - Postgres - UpdateUser - r.Pool.Query: %s", err)
		return nil, err
	}

	return &u, nil
}

func (r *Repo) DeleteUser(id int) error {
	sql, args, err := r.Builder.
		Update(usersTableName).
		Set("removed", true).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		log.Printf("Error - Postgres - UpdateUser - r.Builder: %s", err)
		return err
	}

	_, err = r.Pool.Exec(context.Background(), sql, args...)
	if err != nil {
		log.Printf("Error - Postgres - UpdateUser - r.Pool.Query: %s", err)
		return err
	}

	return nil
}

func (r *Repo) GetUser(id int) (*entity.User, error) {
	sql, args, err := r.Builder.
		Select("id, username, email, role_id, phone, address_id, removed, created_at, updated_at, deleted_at").
		From(usersTableName).
		Join("addresses on users.address = addresses.id").
		Join("roles on users.role = roles.id").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		log.Printf("Error -Postgres - GetUser - r.Builder: %s", err)
		return nil, err
	}

	row := r.Pool.QueryRow(context.Background(), sql, args...)
	u := entity.User{}
	err = row.Scan(
		&u.ID,
		&u.Username,
		&u.Email,
		&u.Phone,

		&u.AddressFull.ID,
		&u.AddressFull.City,
		&u.AddressFull.Country,
		&u.AddressFull.Address,
		&u.AddressFull.Lng,
		&u.AddressFull.Lat,

		&u.RoleFull.Name,
		&u.RoleFull.ID,
		&u.RoleFull.Privileges,
	)
	if err != nil {
		log.Printf("Error - User - GetUsers - rows.Scan: %s", err)
		return nil, err
	}

	return &u, nil
}

func (r *Repo) GetUsers() (*[]entity.User, error) {
	sql, args, err := r.Builder.
		Select("id, username, email, role, phone, addresses.*, roles.*").
		From(usersTableName).
		Join("addresses on users.address = addresses.id").
		Join("roles on users.role = roles.id").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("User - GetUsers - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(context.Background(), sql, args...)
	if err != nil {
		return nil, fmt.Errorf("User - GetUsers - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	entities := make([]entity.User, 0, _defaultEntityCap)
	for rows.Next() {
		u := entity.User{}
		err = rows.Scan(
			&u.ID,
			&u.Username,
			&u.Email,
			&u.Phone,

			&u.AddressFull.ID,
			&u.AddressFull.City,
			&u.AddressFull.Country,
			&u.AddressFull.Address,
			&u.AddressFull.Lng,
			&u.AddressFull.Lat,

			&u.RoleFull.Name,
			&u.RoleFull.ID,
			&u.RoleFull.Privileges,
		)
		if err != nil {
			return nil, fmt.Errorf("User - GetUsers - rows.Scan: %w", err)
		}
		entities = append(entities, u)
	}

	return &entities, nil
}
