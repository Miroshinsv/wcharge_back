package postgres

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/Miroshinsv/wcharge_back/internal/entity"
)

var tableAddress = "tbl_addresses"

func (r *Repo) CreateAddressRepo(addr entity.Address) error {
	sql, args, err := r.Builder.
		Insert(tableAddress).
		Columns("country, city, address, lat, lng").
		Values(addr.Country, addr.City, addr.Address, addr.Lat, addr.Lng).
		ToSql()
	if err != nil {
		return fmt.Errorf("PostgresRepo - CreateAddressRepo - r.Builder: %w", err)
	}
	ctx := context.Background()
	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("PostgresRepo - CreateAddressRepo - r.Pool.Exec: %w", err)
	}

	return nil
}

func (r *Repo) UpdateAddressRepo(id int, addr entity.Address) error {
	sql, args, err := r.Builder.
		Update(tableAddress).
		Set("country", addr.Country).
		Set("city", addr.City).
		Set("address", addr.Address).
		Set("lat", addr.Lat).
		Set("lng", addr.Lng).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return fmt.Errorf("PostgresRepo - UpdateAddressRepo - r.Builder: %w", err)
	}
	ctx := context.Background()
	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("PostgresRepo - UpdateAddressRepo - r.Pool.Query: %w", err)
	}
	return nil
}

func (r *Repo) DeleteAddressRepo(id int) error {
	sql, args, err := r.Builder.
		Delete(tableAddress).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return fmt.Errorf("PostgresRepo - DeleteAddressRepo - r.Builder: %w", err)
	}
	ctx := context.Background()
	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("PostgresRepo - DeleteAddressRepo - r.Pool.Query: %w", err)
	}
	return nil
}

func (r *Repo) GetAddressRepo(id int) (entity.Address, error) {
	u := entity.Address{}
	sql, args, err := r.Builder.
		Select("id, country, city, address, lat, lng").
		From(tableAddress).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return u, fmt.Errorf("PostgresRepo - GetAddressRepo - r.Builder: %w", err)
	}
	ctx := context.Background()
	row := r.Pool.QueryRow(ctx, sql, args...)
	if err != nil {
		return u, fmt.Errorf("PostgresRepo - GetAddressRepo - r.Pool.Query: %w", err)
	}

	err = row.Scan(&u.ID, &u.Country, &u.City, &u.Address, &u.Lat, &u.Lng)
	if err != nil {
		return entity.Address{}, fmt.Errorf("AddressRepo - GetAddresss - rows.Scan: %w", err)
	}

	return u, nil
}

func (r *Repo) GetAddressesRepo() ([]entity.Address, error) {
	sql, _, err := r.Builder.
		Select("id, country, city, address, lat, lng").
		From(tableAddress).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("GetAddresses - r.Builder: %w", err)
	}
	ctx := context.Background()         //!!!
	rows, err := r.Pool.Query(ctx, sql) //!!!
	if err != nil {
		return nil, fmt.Errorf("GetAddresses - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	entities := make([]entity.Address, 0, _defaultEntityCap)

	for rows.Next() {
		u := entity.Address{}
		err = rows.Scan(&u.ID, &u.Country, &u.City, &u.Address, &u.Lat, &u.Lng)
		if err != nil {
			return nil, fmt.Errorf("AddressRepo - GetAddresss - rows.Scan: %w", err)
		}
		entities = append(entities, u)
	}

	return entities, nil
}
