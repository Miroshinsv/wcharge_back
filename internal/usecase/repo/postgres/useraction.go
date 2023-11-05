package postgres

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/Miroshinsv/wcharge_back/internal/entity"
)

// делать проверку на ПОВТОРНУЮ ЗАПИСЬ и ПОВТОРНОЕ УДАЛЕНИЕ

// GetUserPowerbanksRepo All user's powerbanks
func (r *Repo) GetUserPowerbanksRepo(userId int) ([]entity.Powerbank, error) {
	sql, args, err := r.Builder.
		Select("powerbank_id").
		From("postgres.public.tbl_user_powerbank").
		Where(squirrel.Eq{"user_id": userId}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("Repo - GetUserPowerbanksRepo - r.Builder: %w", err)
	}
	ctx := context.Background()
	rows, err := r.Pool.Query(ctx, sql, args)
	if err != nil {
		return nil, fmt.Errorf("Repo - GetUserPowerbanksRepo - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	var powerbanksId []int

	for rows.Next() {
		var i int
		err = rows.Scan(&i)
		if err != nil {
			return nil, fmt.Errorf("Repo - GetUserPowerbanksRepo - rows.Scan: %w", err)
		}
		powerbanksId = append(powerbanksId, i)
	}

	var entities []entity.Powerbank
	for _, powerbankId := range powerbanksId {
		pw, _ := r.GetPowerbankRepo(powerbankId)
		entities = append(entities, pw)
	}

	return entities, nil
}

func (r *Repo) insertStationPowerbank(powerbankId int, stationID int) error {
	sql, args, err := r.Builder.
		Insert("public.tbl_station_powerbank").
		Columns("station_id, powerbank_id").
		Values(stationID, powerbankId).
		ToSql()
	if err != nil {
		return fmt.Errorf("insertStationPowerbank - r.Builder: %w", err)
	}
	ctx := context.Background()
	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("insertStationPowerbank - r.Pool.Exec: %w", err)
	}

	return nil
}

func (r *Repo) deleteStationPowerbank(powerbankId int, stationID int) error {
	sql, args, err := r.Builder.
		Delete("public.tbl_station_powerbank").
		Where(squirrel.And{
			squirrel.Eq{"powerbank_id": powerbankId},
			squirrel.Eq{"station_id": stationID},
		}).
		ToSql()
	if err != nil {
		return fmt.Errorf("deleteStationPowerbank - r.Builder: %w", err)
	}
	ctx := context.Background()
	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("deleteStationPowerbank - r.Pool.Exec: %w", err)
	}

	return nil
}

func (r *Repo) insertUserPowerbank(userId int, powerbankId int, stationID int) error {
	sql, args, err := r.Builder.
		Insert("public.tbl_user_powerbank").
		Columns("user_id, powerbank_id").
		Values(userId, powerbankId, stationID).
		ToSql()
	if err != nil {
		return fmt.Errorf("insertUserPowerbank - r.Builder: %w", err)
	}
	ctx := context.Background()
	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("insertUserPowerbank - r.Pool.Exec: %w", err)
	}

	return nil
}

func (r *Repo) deleteUserPowerbank(userId int, powerbankId int) error {
	sql, args, err := r.Builder.
		Delete("public.tbl_user_powerbank").
		Where(squirrel.And{
			squirrel.Eq{"user_id": userId},
			squirrel.Eq{"powerbank_id": powerbankId},
		}).
		ToSql()
	if err != nil {
		return fmt.Errorf("deleteUserPowerbank - r.Builder: %w", err)
	}
	ctx := context.Background()
	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("deleteUserPowerbank - r.Pool.Exec: %w", err)
	}

	return nil
}

// TakePowerbankRepo take user powerbank
func (r *Repo) TakePowerbankRepo(userId int, powerbankId int, stationID int) error {
	err := r.insertUserPowerbank(userId, powerbankId, stationID)
	if err != nil {
		return fmt.Errorf("BackTakePowerbankRepo - %w", err)
	}
	err = r.deleteStationPowerbank(stationID, powerbankId)
	if err != nil {
		return fmt.Errorf("BackTakePowerbankRepo - %w", err)
	}
	return nil
}

func (r *Repo) PutPowerbankRepo(userId int, powerbankId int, stationID int) error {
	err := r.deleteUserPowerbank(userId, powerbankId)
	if err != nil {
		return fmt.Errorf("BackTakePowerbankRepo - %w", err)
	}
	err = r.insertStationPowerbank(stationID, powerbankId)
	if err != nil {
		return fmt.Errorf("BackTakePowerbankRepo - %w", err)
	}
	return nil
}

func (r *Repo) AddPowerbankToStationRepo(powerbankId int, stationID int) error {
	err := r.insertStationPowerbank(stationID, powerbankId)
	if err != nil {
		return fmt.Errorf("AddPowerbankToStationRepo - %w", err)
	}
	return nil
}
