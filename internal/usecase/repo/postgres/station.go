package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/Miroshinsv/wcharge_back/internal/entity"
	//"github.com/Miroshinsv/wcharge_back/pkg/postgres"
)

func (r *Repo) CreateStationRepo(s entity.Station) error {
	sql, args, err := r.Builder.
		Insert("public.tbl_stations").
		Columns("serial_number, address_id, capacity, free_capacity").
		Values(s.SerialNumber, s.AddressId, s.Capacity, s.FreeCapacity).
		ToSql()
	if err != nil {
		return fmt.Errorf("PostgresRepo - CreateStationRepo - r.Builder: %w", err)
	}
	ctx := context.Background()
	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("PostgresRepo - CreateStationRepo - r.Pool.Exec: %w", err)
	}

	return nil
}

func (r *Repo) UpdateStationRepo(id int, s entity.Station) error {
	sql, args, err := r.Builder.
		Update("postgres.public.tbl_stations").
		Set("serial_number", s.SerialNumber).
		Set("address_id", s.AddressId).
		Set("capacity", s.Capacity).
		Set("free_capacity", s.FreeCapacity).
		Set("updated_at", time.Now()).
		Where("postgres.public.tbl_users.id = ?", id).
		ToSql()
	if err != nil {
		return fmt.Errorf("PostgresRepo - UpdateStationRepo - r.Builder: %w", err)
	}
	ctx := context.Background()
	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("PostgresRepo - UpdateStationRepo - r.Pool.Query: %w", err)
	}
	return nil
}

func (r *Repo) DeleteStationRepo(id int) error {
	tm := time.Now().GoString()
	sql, args, err := r.Builder.
		Update("postgres.public.tbl_stations").
		Set("updated_at", tm).
		Set("deleted_at", tm).
		Where("postgres.public.tbl_users.id = ?", id).
		ToSql()
	if err != nil {
		return fmt.Errorf("PostgresRepo - DeleteStationRepo - r.Builder: %w", err)
	}
	ctx := context.Background()
	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("PostgresRepo - DeleteStationRepo - r.Pool.Query: %w", err)
	}
	return nil
}

func (r *Repo) GetStationRepo(id int) (entity.Station, error) {
	s := entity.Station{}
	sql, args, err := r.Builder.
		Select("id, serial_number, address_id, capacity, free_capacity, created_at, updated_at, deleted_at").
		From("postgres.public.tbl_users").
		Where("postgres.public.tbl_users.id = ?", id).
		ToSql()
	if err != nil {
		return s, fmt.Errorf("PostgresRepo - GetStationRepo - r.Builder: %w", err)
	}
	ctx := context.Background()
	row := r.Pool.QueryRow(ctx, sql, args...)
	if err != nil {
		return s, fmt.Errorf("PostgresRepo - GetUserRepo - r.Pool.Query: %w", err)
	}

	err = row.Scan(&s.ID, &s.SerialNumber, &s.AddressId, &s.Capacity, &s.FreeCapacity, &s.CreateAt, &s.UpdateAt, &s.DeleteAt)
	if err != nil {
		return s, fmt.Errorf("PostgresRepo - GetStationRepo - rows.Scan: %w", err)
	}

	return s, nil
}

func (r *Repo) GetStationsRepo() ([]entity.Station, error) {
	sql, _, err := r.Builder.
		Select("id, serial_number, address_id, capacity, free_capacity, created_at, updated_at, deleted_at").
		From("postgres.public.tbl_users").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("Repo - GetStationsRepo - r.Builder: %w", err)
	}
	ctx := context.Background()         //!!!
	rows, err := r.Pool.Query(ctx, sql) //!!!
	if err != nil {
		return nil, fmt.Errorf("Repo - GetStationsRepo - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	entities := make([]entity.Station, 0, _defaultEntityCap)

	for rows.Next() {
		s := entity.Station{}

		err = rows.Scan(&s.ID, &s.SerialNumber, &s.AddressId, &s.Capacity, &s.FreeCapacity, &s.CreateAt, &s.UpdateAt, &s.DeleteAt)
		if err != nil {
			return nil, fmt.Errorf("Repo - GetStationsRepo - rows.Scan: %w", err)
		}

		entities = append(entities, s)
	}

	return entities, nil
}
