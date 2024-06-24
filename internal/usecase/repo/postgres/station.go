package postgres

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/Miroshinsv/wcharge_back/internal/entity"
	"log"
)

func (r *Repo) CreateStation(s entity.Station) (*entity.Station, error) {
	sql, args, err := r.Builder.
		Insert(stationsTableName).
		Columns("serial_number, address, capacity, free_capacity").
		Values(s.SerialNumber, s.Address, s.Capacity, s.FreeCapacity).
		ToSql()
	if err != nil {
		log.Printf("Error - Repo - CreateStation - r.Builder: %w", err)
		return nil, err
	}

	row := r.Pool.QueryRow(context.Background(), sql, args...)
	err = row.Scan(
		&s.ID,
		&s.SerialNumber,
		&s.Address,
		&s.Capacity,
		&s.FreeCapacity,
		&s.CreateAt,
		&s.UpdateAt,
	)
	if err != nil {
		log.Printf("Error - Repo - CreateStation - r.Pool.QueryRow: %s", err)
		return nil, err
	}

	return &s, nil
}

func (r *Repo) UpdateStation(s entity.Station, id int) (*entity.Station, error) {
	sql, args, err := r.Builder.
		Update(stationsTableName).
		Set("address", s.Address).
		Set("free_capacity", s.FreeCapacity).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		log.Printf("Error - Repo - UpdateStation - r.Builder: %s", err)
		return nil, err
	}

	row := r.Pool.QueryRow(context.Background(), sql, args...)
	err = row.Scan(
		&s.ID,
		&s.SerialNumber,
		&s.Address,
		&s.Capacity,
		&s.FreeCapacity,
		&s.CreateAt,
		&s.UpdateAt,
	)
	if err != nil {
		log.Printf("Error - Repo - UpdateStation - r.Pool.QueryRow: %s", err)
		return nil, err
	}
	return &s, nil
}

func (r *Repo) DeleteStationRepo(id int) error {
	sql, args, err := r.Builder.
		Update(stationsTableName).
		Set("removed", true).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		log.Printf("Error - Repo - DeleteStation - r.Builder: %s", err)
		return err
	}

	_, err = r.Pool.Exec(context.Background(), sql, args...)
	if err != nil {
		log.Printf("Error - Repo - DeleteStation - r.Pool.Exec: %s", err)
		return err
	}
	return nil
}

func (r *Repo) GetStation(id int) (*entity.Station, error) {
	sql, args, err := r.Builder.
		Select("stations.id, serial_number, addresses.*").
		From(stationsTableName).
		Join("addresses on stations.address = addresses.id").
		Where(squirrel.And{
			squirrel.Eq{"stations.id": id},
			squirrel.Eq{"removed": false},
		}).
		ToSql()
	if err != nil {
		log.Printf("PostgresRepo - GetStationRepo - r.Builder: %s", err)
		return nil, err
	}

	row := r.Pool.QueryRow(context.Background(), sql, args...)
	if err != nil {
		log.Printf("PostgresRepo - GetUserRepo - r.Pool.Query: %s", err)
		return nil, err
	}
	s := entity.Station{}
	err = row.Scan(
		&s.ID,
		&s.SerialNumber,

		&s.AddressFull.ID,
		&s.AddressFull.City,
		&s.AddressFull.Country,
		&s.AddressFull.Address,
		&s.AddressFull.Lng,
		&s.AddressFull.Lat,
	)
	if err != nil {
		log.Printf("PostgresRepo - GetStationRepo - rows.Scan: %s", err)
		return nil, err
	}

	return &s, nil
}

func (r *Repo) GetStations() (*[]entity.Station, error) {
	sql, args, err := r.Builder.
		Select("stations.id, serial_number, addresses.*").
		From(stationsTableName).
		Join("addresses on stations.address = addresses.id").
		Where(squirrel.And{
			squirrel.Eq{"removed": false},
		}).
		ToSql()
	if err != nil {
		log.Printf("Error - Repo - GetStations - r.Builder: %s", err)
		return nil, err
	}

	rows, err := r.Pool.Query(context.Background(), sql, args...)
	if err != nil {
		log.Printf("Error - Repo - GetStations - r.Pool.Query: %s", err)
		return nil, err
	}
	defer rows.Close()

	entities := make([]entity.Station, 0, _defaultEntityCap)
	for rows.Next() {
		s := entity.Station{}
		err = rows.Scan(
			&s.ID,
			&s.SerialNumber,

			&s.AddressFull.ID,
			&s.AddressFull.City,
			&s.AddressFull.Country,
			&s.AddressFull.Address,
			&s.AddressFull.Lng,
			&s.AddressFull.Lat,
		)
		if err != nil {
			log.Printf("Error - Repo - GetStations - rows.Scan: %s", err)
			return nil, err
		}

		entities = append(entities, s)
	}

	return &entities, nil
}

// TODO
//func (r *Repo) GetAllPowerbanksInStation(stationId int) (*[]entity.Powerbank, error) {
//	subQ, subArgs, err := squirrel.
//		Select("powerbank").
//		From("rel__stations__powerbanks").
//		Where(squirrel.Eq{"station": stationId}).
//		PlaceholderFormat(squirrel.Question).
//		ToSql()
//	if err != nil {
//		return nil, fmt.Errorf("GetAllPowerbanksInStationRepo - subQ: %w", err)
//	}
//	sql, args, err := r.Builder.
//		Select("id, serial_number, capacity, used, removed, created_at, updated_at, deleted_at").
//		From("tbl_powerbanks").
//		Where("id in ("+subQ+")", subArgs...).
//		ToSql()
//	if err != nil {
//		return nil, fmt.Errorf("GetAllPowerbanksInStationRepo - r.Builder: %w", err)
//	}
//	ctx := context.Background()                  //!!!
//	rows, err := r.Pool.Query(ctx, sql, args...) //!!!
//	if err != nil {
//		return nil, fmt.Errorf("GetAllPowerbanksInStationRepo - r.Pool.Query: %w", err)
//	}
//	defer rows.Close()
//
//	entities := make([]entity.Powerbank, 0, _defaultEntityCap)
//
//	for rows.Next() {
//		s := entity.Powerbank{}
//
//		err = rows.Scan(&s.ID, &s.SerialNumber, &s.Capacity, &s.Used, &s.Removed, &s.CreateAt, &s.UpdateAt, &s.DeleteAt)
//		if err != nil {
//			return nil, fmt.Errorf("GetAllPowerbanksInStationRepo - rows.Scan: %w", err)
//		}
//
//		entities = append(entities, s)
//	}
//
//	return &entities, nil
//}
