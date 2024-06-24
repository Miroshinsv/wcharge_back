package postgres

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/Miroshinsv/wcharge_back/internal/entity"
	"log"
)

func (r *Repo) CreatePowerbank(p entity.Powerbank) (*entity.Powerbank, error) {
	sql, args, err := r.Builder.
		Insert(powerbanksTableName).
		Columns("serial_number, capacity, used").
		Values(p.SerialNumber, p.Capacity, p.Used).
		Suffix("returning id").
		ToSql()
	if err != nil {
		log.Printf("CreatePowerbankRepo - r.Builder: %s", err)
		return nil, err
	}

	row := r.Pool.QueryRow(context.Background(), sql, args...)
	err = row.Scan(
		&p.ID,
	)
	if err != nil {
		log.Printf("PostgresRepo - CreateStationRepo - r.Pool.Exec: %s", err)
		return nil, err
	}

	return &p, nil
}

func (r *Repo) UpdatePowerbank(p entity.Powerbank, id int) (*entity.Powerbank, error) {
	sql, args, err := r.Builder.
		Update(powerbanksTableName).
		Set("used", p.Used).
		Set("serial_number", p.SerialNumber).
		Set("capacity", p.Capacity).
		Set("removed", p.Removed).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		log.Printf("UpdatePowerbankRepo - r.Builder: %s", err)
		return nil, err
	}

	row := r.Pool.QueryRow(context.Background(), sql, args...)
	err = row.Scan(
		&p.ID,
		&p.Capacity,
		&p.UpdateAt,
		&p.Capacity,
		&p.Used,
		&p.Position,
		&p.CreateAt,
		&p.Removed,
	)
	if err != nil {
		log.Printf("UpdatePowerbankRepo - r.Pool.Query: %s", err)
		return nil, err
	}

	return &p, nil
}

func (r *Repo) DeletePowerbank(id int) error {
	sql, args, err := r.Builder.
		Update(powerbanksTableName).
		Set("removed", 1).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		log.Printf("DeletePowerbankRepo - r.Builder: %s", err)
		return err
	}

	_, err = r.Pool.Exec(context.Background(), sql, args...)
	if err != nil {
		log.Printf("DeletePowerbankRepo - r.Pool.Query: %s", err)
		return err
	}
	return nil
}

func (r *Repo) GetRamdomPowerbank() (*entity.Powerbank, error) {
	powerbank := entity.Powerbank{}

	sql, args, err := r.Builder.
		Select("powerbanks.id, serial_number, used, tsp.position").
		From(powerbanksTableName).
		Join("rel__stations__powerbanks rsp on powerbanks.id = rsp.powerbank").
		Where("used = false"). // TODO ???
		OrderBy("random()").
		Limit(1).
		ToSql()
	if err != nil {
		log.Printf("GetRamdomPowerbank - r.Builder: %s", err)
		return nil, err
	}

	row := r.Pool.QueryRow(context.Background(), sql, args...)
	err = row.Scan(
		&powerbank.ID,
		&powerbank.SerialNumber,
		&powerbank.Used,
		&powerbank.Position,
	)
	if err != nil {
		log.Printf("GetRamdomPowerbank - r.Pool.Query: %s", err)
		return nil, err
	}

	return &powerbank, nil
}

func (r *Repo) GetPowerbank(id int) (*entity.Powerbank, error) {
	sql, args, err := r.Builder.
		Select("id, serial_number, capacity, used, removed, created_at, updated_at, deleted_at").
		From(powerbanksTableName).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		log.Printf("GetPowerbankRepo - r.Builder: %s", err)
		return nil, err
	}

	row := r.Pool.QueryRow(context.Background(), sql, args...)
	powerbank := entity.Powerbank{}
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
		log.Printf("GetPowerbankRepo - rows.Scan: %s", err)
		return nil, err
	}

	return &powerbank, nil
}

func (r *Repo) GetPowerbanks() (*[]entity.Powerbank, error) {
	sql, _, err := r.Builder.
		Select("id, serial_number, capacity, used, removed, created_at, updated_at, deleted_at").
		From(powerbanksTableName).
		ToSql()
	if err != nil {
		log.Printf("GetPowerbanksRepo - r.Builder: %s", err)
		return nil, err
	}

	rows, err := r.Pool.Query(context.Background(), sql)
	if err != nil {
		log.Printf("GetPowerbanksRepo - r.Pool.Query: %s", err)
		return nil, err
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
			log.Printf("GetPowerbanksRepo - rows.Scan: %s", err)
			return nil, err
		}

		entities = append(entities, s)
	}

	return &entities, nil
}
