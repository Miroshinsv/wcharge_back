package postgres

import (
	"context"
	"fmt"

	"github.com/Miroshinsv/wcharge_back/internal/entity"
	//"github.com/Miroshinsv/wcharge_back/pkg/postgres"
)

func (r *PostgresRepo) CreateStationRepo(u entity.Station) error {

	return nil
}

func (r *PostgresRepo) UpdateStationRepo(u entity.Station) error {

	return nil
}

func (r *PostgresRepo) DeleteStationRepo(u entity.Station) error {

	return nil
}

func (r *PostgresRepo) GetStationRepo(u entity.Station) (entity.Station, error) {

	return entity.Station{}, nil
}

func (r *PostgresRepo) GetStationsRepo() ([]entity.Station, error) {
	sql, _, err := r.Builder.
		Select("id, serial_number, address_id, capacity, free_capacity").
		From("stations").
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

	entities := make([]entity.Station, 0, _defaultEntityCap)

	for rows.Next() {
		e := entity.Station{}

		err = rows.Scan(&e.ID, &e.SerialNumber, &e.Address.ID, &e.Capacity, &e.FreeCapacity)
		if err != nil {
			return nil, fmt.Errorf("UserRepo - GetUsers - rows.Scan: %w", err)
		}

		entities = append(entities, e)
	}

	return entities, nil
}
