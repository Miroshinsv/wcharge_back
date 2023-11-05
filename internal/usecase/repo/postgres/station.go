package postgres

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"time"

	"github.com/Miroshinsv/wcharge_back/internal/entity"
	//"github.com/Miroshinsv/wcharge_back/pkg/postgres"
)

var tableStation = "tbl_stations"

func (r *Repo) CreateStationRepo(s entity.Station) error {
	sql, args, err := r.Builder.
		Insert(tableStation).
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
		Update(tableStation).
		Set("address_id", s.AddressId).
		Set("free_capacity", s.FreeCapacity).
		Set("updated_at", time.Now()).
		Where(squirrel.Eq{"id": id}).
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
	sql, args, err := r.Builder.
		Update(tableStation).
		Set("removed", 1).
		Where(squirrel.Eq{"id": id}).
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
		Select("id, serial_number, address_id, capacity, free_capacity, removed,created_at, updated_at, deleted_at").
		From(tableStation).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return s, fmt.Errorf("PostgresRepo - GetStationRepo - r.Builder: %w", err)
	}
	ctx := context.Background()
	row := r.Pool.QueryRow(ctx, sql, args...)
	if err != nil {
		return s, fmt.Errorf("PostgresRepo - GetUserRepo - r.Pool.Query: %w", err)
	}

	err = row.Scan(&s.ID, &s.SerialNumber, &s.AddressId, &s.Capacity, &s.FreeCapacity, &s.Removed, &s.CreateAt, &s.UpdateAt, &s.DeleteAt)
	if err != nil {
		return s, fmt.Errorf("PostgresRepo - GetStationRepo - rows.Scan: %w", err)
	}

	return s, nil
}

func (r *Repo) GetStationsRepo() ([]entity.Station, error) {
	sql, _, err := r.Builder.
		Select("id, serial_number, address_id, capacity, free_capacity, removed, created_at, updated_at, deleted_at").
		From(tableStation).
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

		err = rows.Scan(&s.ID, &s.SerialNumber, &s.AddressId, &s.Capacity, &s.FreeCapacity, &s.Removed, &s.CreateAt, &s.UpdateAt, &s.DeleteAt)
		if err != nil {
			return nil, fmt.Errorf("Repo - GetStationsRepo - rows.Scan: %w", err)
		}

		entities = append(entities, s)
	}

	return entities, nil
}

func (r *Repo) GetAllPowerbanksInStationRepo(stationId int) ([]entity.Powerbank, error) {
	subQ, subArgs, err := squirrel.
		Select("powerbank_id").
		From("tbl_station_powerbank").
		Where(squirrel.Eq{"station_id": stationId}).
		PlaceholderFormat(squirrel.Question).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("GetAllPowerbanksInStationRepo - subQ: %w", err)
	}
	sql, args, err := r.Builder.
		Select("id, serial_number, capacity, used, removed, created_at, updated_at, deleted_at").
		From("tbl_powerbanks").
		Where("id in ("+subQ+")", subArgs...).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("GetAllPowerbanksInStationRepo - r.Builder: %w", err)
	}
	ctx := context.Background()                  //!!!
	rows, err := r.Pool.Query(ctx, sql, args...) //!!!
	if err != nil {
		return nil, fmt.Errorf("GetAllPowerbanksInStationRepo - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	entities := make([]entity.Powerbank, 0, _defaultEntityCap)

	for rows.Next() {
		s := entity.Powerbank{}

		err = rows.Scan(&s.ID, &s.SerialNumber, &s.Capacity, &s.Used, &s.Removed, &s.CreateAt, &s.UpdateAt, &s.DeleteAt)
		if err != nil {
			return nil, fmt.Errorf("GetAllPowerbanksInStationRepo - rows.Scan: %w", err)
		}

		entities = append(entities, s)
	}

	return entities, nil
}
