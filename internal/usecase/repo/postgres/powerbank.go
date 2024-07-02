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
		log.Printf("Error - CreatePowerbank - r.Builder: %s", err)
		return nil, err
	}

	row := r.Pool.QueryRow(context.Background(), sql, args...)
	err = row.Scan(
		&p.ID,
	)
	if err != nil {
		log.Printf("Error - Postgres - CreateStation - r.Pool.Exec: %s", err)
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
		log.Printf("Error - UpdatePowerbank - r.Builder: %s", err)
		return nil, err
	}

	row := r.Pool.QueryRow(context.Background(), sql, args...)
	err = row.Scan(
		&p.ID,
		&p.Capacity,
		&p.UpdateAt,
		&p.Capacity,
	)
	if err != nil {
		log.Printf("Error - UpdatePowerbank - r.Pool.Query: %s", err)
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
		log.Printf("Error - DeletePowerbank - r.Builder: %s", err)
		return err
	}

	_, err = r.Pool.Exec(context.Background(), sql, args...)
	if err != nil {
		log.Printf("Error - DeletePowerbank - r.Pool.Query: %s", err)
		return err
	}
	return nil
}

func (r *Repo) GetRamdomPowerbank() (*entity.Powerbank, error) {

	sql, args, err := r.Builder.
		Select("powerbanks.id, serial_number, used, tsp.position").
		From(powerbanksTableName).
		Join("rel__stations__powerbanks rsp on powerbanks.id = rsp.powerbank").
		Where("used = false"). // TODO ???
		OrderBy("random()").
		Limit(1).
		ToSql()
	if err != nil {
		log.Printf("Error - GetRamdomPowerbank - r.Builder: %s", err)
		return nil, err
	}

	row := r.Pool.QueryRow(context.Background(), sql, args...)
	powerbank := entity.Powerbank{}
	err = row.Scan(
		&powerbank.ID,
		&powerbank.SerialNumber,
		&powerbank.Used,
		&powerbank.Position,
	)
	if err != nil {
		log.Printf("Error - GetRamdomPowerbank - r.Pool.Query: %s", err)
		return nil, err
	}

	return &powerbank, nil
}

func (r *Repo) GetPowerbank(id int) (*entity.Powerbank, error) {
	sql, args, err := r.Builder.
		Select("id, serial_number, capacity, used").
		From(powerbanksTableName).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		log.Printf("Error - GetPowerbank - r.Builder: %s", err)
		return nil, err
	}

	row := r.Pool.QueryRow(context.Background(), sql, args...)
	powerbank := entity.Powerbank{}
	err = row.Scan(
		&powerbank.ID,
		&powerbank.SerialNumber,
		&powerbank.Capacity,
	)
	if err != nil {
		log.Printf("Error - GetPowerbankRepo - rows.Scan: %s", err)
		return nil, err
	}

	return &powerbank, nil
}

func (r *Repo) GetPowerbanks() (*[]entity.Powerbank, error) {
	sql, _, err := r.Builder.
		Select("id, serial_number, capacity, used").
		From(powerbanksTableName).
		Where(squirrel.Eq{"removed": false}).
		ToSql()
	if err != nil {
		log.Printf("Error - GetPowerbanks - r.Builder: %s", err)
		return nil, err
	}

	rows, err := r.Pool.Query(context.Background(), sql)
	if err != nil {
		log.Printf("Error - GetPowerbanks - r.Pool.Query: %s", err)
		return nil, err
	}
	defer rows.Close()

	entities := make([]entity.Powerbank, 0, _defaultEntityCap)
	for rows.Next() {
		p := entity.Powerbank{}
		err = rows.Scan(
			&p.ID,
			&p.SerialNumber,
			&p.Capacity,
		)
		if err != nil {
			log.Printf("Error - GetPowerbankы - rows.Scan: %s", err)
			return nil, err
		}

		entities = append(entities, p)
	}

	return &entities, nil
}

func (r *Repo) GetAllPowerbanksInStation(stationId int) (*[]entity.Powerbank, error) {
	sql, args, err := r.Builder.
		Select("powerbanks.id, serial_number, capacity, used").
		From(powerbanksTableName).
		Join("rel__stations__powerbanks rsp on powerbanks.id = rsp.powerbank").
		Where(squirrel.And{
			squirrel.Eq{"rsp.station": stationId},
			squirrel.Eq{"removed": false},
		}).
		ToSql()
	if err != nil {
		log.Printf("Error - GetAllPowerbanksInStation - r.Builder: %s", err)
		return nil, err
	}

	rows, err := r.Pool.Query(context.Background(), sql, args...)
	if err != nil {
		log.Printf("Error - GetAllPowerbanksInStation - r.Pool.Query: %s", err)
		return nil, err
	}
	defer rows.Close()

	entities := make([]entity.Powerbank, 0, _defaultEntityCap)
	for rows.Next() {
		p := entity.Powerbank{}
		err = rows.Scan(
			&p.ID,
			&p.SerialNumber,
			&p.Capacity,
		)
		if err != nil {
			log.Printf("Error - GetAllPowerbanksInStation - rows.Scan: %s", err)
			return nil, err
		}
		entities = append(entities, p)
	}

	return &entities, nil
}
