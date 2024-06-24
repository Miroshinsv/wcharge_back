package postgres

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/Miroshinsv/wcharge_back/internal/entity"
	"log"
)

func (r *Repo) CreateAddress(addr entity.Address) (*entity.Address, error) {
	sql, args, err := r.Builder.
		Insert(addressTableName).
		Columns("country, city, address, lat, lng").
		Values(addr.Country, addr.City, addr.Address, addr.Lat, addr.Lng).
		ToSql()
	if err != nil {
		log.Printf("PostgresRepo - CreateAddressRepo - r.Builder: %s", err)
		return nil, err
	}
	row := r.Pool.QueryRow(context.Background(), sql, args...)
	err = row.Scan(
		&addr.ID,
		&addr.Address,
		&addr.City,
		&addr.Lat,
		&addr.Lng,
		&addr.Country,
	)
	if err != nil {
		log.Printf("PostgresRepo - CreateAddressRepo - r.Pool.Exec: %s", err)
		return nil, err
	}

	return &addr, nil
}

func (r *Repo) UpdateAddress(addr entity.Address, id int) (*entity.Address, error) {
	sql, args, err := r.Builder.
		Update(addressTableName).
		Set("country", addr.Country).
		Set("city", addr.City).
		Set("address", addr.Address).
		Set("lat", addr.Lat).
		Set("lng", addr.Lng).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		log.Printf("PostgresRepo - UpdateAddressRepo - r.Builder: %s", err)
		return nil, err
	}
	row := r.Pool.QueryRow(context.Background(), sql, args...)
	err = row.Scan(
		&addr.ID,
		&addr.Address,
		&addr.City,
		&addr.Lat,
		&addr.Lng,
		&addr.Country,
	)
	if err != nil {
		log.Printf("PostgresRepo - UpdateAddressRepo - r.Pool.Query: %s", err)
		return nil, err
	}
	return &addr, nil
}

func (r *Repo) DeleteAddress(id int) error {
	sql, args, err := r.Builder.
		Delete(addressTableName).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		log.Printf("PostgresRepo - DeleteAddressRepo - r.Builder: %s", err)
		return err
	}
	_, err = r.Pool.Exec(context.Background(), sql, args...)
	if err != nil {
		log.Printf("PostgresRepo - DeleteAddressRepo - r.Pool.Query: %s", err)
		return err
	}
	return nil
}

func (r *Repo) GetAddress(id int) (*entity.Address, error) {
	sql, args, err := r.Builder.
		Select("id, country, city, address, lat, lng").
		From(addressTableName).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		log.Printf("PostgresRepo - GetAddressRepo - r.Builder: %s", err)
		return nil, err
	}

	row := r.Pool.QueryRow(context.Background(), sql, args...)
	u := entity.Address{}
	err = row.Scan(
		&u.ID,
		&u.Country,
		&u.City,
		&u.Address,
		&u.Lat,
		&u.Lng,
	)
	if err != nil {
		log.Printf("AddressRepo - GetAddresss - rows.Scan: %ы", err)
		return nil, err
	}

	return &u, nil
}

func (r *Repo) GetAddresses() (*[]entity.Address, error) {
	sql, _, err := r.Builder.
		Select("id, country, city, address, lat, lng").
		From(addressTableName).
		ToSql()
	if err != nil {
		log.Printf("GetAddresses - r.Builder: %s", err)
		return nil, err
	}

	rows, err := r.Pool.Query(context.Background(), sql)
	if err != nil {
		log.Printf("GetAddresses - r.Pool.Query: %s", err)
		return nil, err
	}
	defer rows.Close()

	entities := make([]entity.Address, 0, _defaultEntityCap)
	for rows.Next() {
		u := entity.Address{}
		err = rows.Scan(&u.ID, &u.Country, &u.City, &u.Address, &u.Lat, &u.Lng)
		if err != nil {
			log.Printf("AddressRepo - GetAddresss - rows.Scan: %s", err)
			return nil, err
		}
		entities = append(entities, u)
	}

	return &entities, nil
}
