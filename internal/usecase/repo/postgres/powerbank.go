package postgres

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/Miroshinsv/wcharge_back/internal/entity"
)

func (r *Repo) CreatePowerbankRepo(p entity.Powerbank) error {
	sql, args, err := r.Builder.
		Insert("tbl_powerbanks").
		Columns("serial_number, capacity, used").
		Values(p.SerialNumber, p.Capacity, p.Used).
		ToSql()
	if err != nil {
		return fmt.Errorf("CreatePowerbankRepo - r.Builder: %w", err)
	}
	ctx := context.Background()
	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("CreatePowerbankRepo - r.Pool.Exec: %w", err)
	}

	return nil
}

func (r *Repo) UpdatePowerbankRepo(id int, p entity.Powerbank) error {
	sql, args, err := r.Builder.
		Update("postgres.public.tbl_powerbanks").
		Set("used", p.Used).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return fmt.Errorf("UpdatePowerbankRepo - r.Builder: %w", err)
	}
	ctx := context.Background()
	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("UpdatePowerbankRepo - r.Pool.Query: %w", err)
	}
	return nil
}

func (r *Repo) DeletePowerbankRepo(id int) error {
	sql, args, err := r.Builder.
		Update("tbl_powerbanks").
		Set("removed", 1).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return fmt.Errorf("DeletePowerbankRepo - r.Builder: %w", err)
	}
	ctx := context.Background()
	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("DeletePowerbankRepo - r.Pool.Query: %w", err)
	}
	return nil
}

func (r *Repo) GetRamdomPowebank() (*entity.Powerbank, error) {
	powerbank := entity.Powerbank{}

	sql, args, err := r.Builder.
		Select("tbl_powerbanks.id, serial_number, used, tsp.position").
		From("tbl_powerbanks").
		Join("tbl_station_powerbank tsp on tbl_powerbanks.id = tsp.powerbank_id").
		Where("used = 0").
		OrderBy("random()").
		Limit(1).
		ToSql()

	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	row := r.Pool.QueryRow(ctx, sql, args...)

	err = row.Scan(
		&powerbank.ID,
		&powerbank.SerialNumber,
		//&powerbank.Capacity,
		&powerbank.Used,
		//&powerbank.Removed,
		//&powerbank.CreateAt,
		//&powerbank.UpdateAt,
		//&powerbank.DeleteAt,
		&powerbank.Position,
	)

	if err != nil {
		return nil, err
	}

	return &powerbank, nil
}

func (r *Repo) GetPowerbankRepo(id int) (entity.Powerbank, error) {
	powerbank := entity.Powerbank{}
	sql, args, err := r.Builder.
		Select("id, serial_number, capacity, used, removed, created_at, updated_at, deleted_at").
		From("tbl_powerbanks").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return powerbank, fmt.Errorf("GetPowerbankRepo - r.Builder: %w", err)
	}
	ctx := context.Background()
	row := r.Pool.QueryRow(ctx, sql, args...)
	if err != nil {
		return powerbank, fmt.Errorf("GetPowerbankRepo - r.Pool.Query: %w", err)
	}

	err = row.Scan(
		&powerbank.ID,
		&powerbank.SerialNumber,
		&powerbank.Capacity,
		&powerbank.Used,
		&powerbank.Removed,
		&powerbank.CreateAt,
		&powerbank.UpdateAt,
		&powerbank.DeleteAt,
	)

	if err != nil {
		return powerbank, fmt.Errorf("GetPowerbankRepo - rows.Scan: %w", err)
	}

	return powerbank, nil
}

func (r *Repo) GetPowerbanksRepo() ([]entity.Powerbank, error) {
	sql, _, err := r.Builder.
		Select("id, serial_number, capacity, used, removed, created_at, updated_at, deleted_at").
		From("tbl_powerbanks").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("GetPowerbanksRepo - r.Builder: %w", err)
	}
	ctx := context.Background()         //!!!
	rows, err := r.Pool.Query(ctx, sql) //!!!
	if err != nil {
		return nil, fmt.Errorf("GetPowerbanksRepo - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	entities := make([]entity.Powerbank, 0, _defaultEntityCap)

	for rows.Next() {
		s := entity.Powerbank{}

		err = rows.Scan(
			&s.ID,
			&s.SerialNumber,
			&s.Capacity,
			&s.Used,
			&s.Removed,
			&s.CreateAt,
			&s.UpdateAt,
			&s.DeleteAt,
		)
		if err != nil {
			return nil, fmt.Errorf("GetPowerbanksRepo - rows.Scan: %w", err)
		}

		entities = append(entities, s)
	}

	return entities, nil
}
