package postgres

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/Miroshinsv/wcharge_back/internal/entity"
	"log"
)

// TODO

// GetUsersPowerbanks All user's powerbanks
// делать проверку на ПОВТОРНУЮ ЗАПИСЬ и ПОВТОРНОЕ УДАЛЕНИЕ
func (r *Repo) GetUsersPowerbanks(userId int) (*[]entity.Powerbank, error) {
	sql, args, err := r.Builder.
		Select("pt.id, pt.serial_number, rsp.position").
		From(usersPowerbanksTableName + " rup").
		Join(powerbanksTableName + " pt on pt.id = rup.powerbank").
		Join(stationsPowerbanksTableName + " rsp on rsp.powerbank = pt.id").
		Where(squirrel.And{
			squirrel.Eq{"rup.user": userId},
			squirrel.Eq{"pt.used": true},
		}).
		ToSql()
	if err != nil {
		log.Printf("Error - Repo - GetUserPowerbanks - r.Builder: %s", err)
		return nil, err
	}

	rows, err := r.Pool.Query(context.Background(), sql, args...)
	if err != nil {
		log.Printf("Error - Repo - GetUserPowerbanks - r.Pool.Query: %s", err)
		return nil, err
	}
	defer rows.Close()

	var entities []entity.Powerbank
	for rows.Next() {
		powerbank := entity.Powerbank{}
		err = rows.Scan(
			&powerbank.ID,
			&powerbank.SerialNumber,
			&powerbank.Used,
			&powerbank.Position,
		)
		if err != nil {
			log.Printf("Error - Repo - GetUserPowerbanks - rows.Scan: %s", err)
			return nil, err
		}

		entities = append(entities, powerbank)
	}

	return &entities, nil
}

// InsertStationPowerbank TODO return value
func (r *Repo) InsertStationPowerbank(powerbankId int, stationId int, position int) error {
	sql, args, err := r.Builder.
		Insert(stationsPowerbanksTableName).
		Columns("station, powerbank, position").
		Values(stationId, powerbankId, position).
		ToSql()
	if err != nil {
		log.Printf("Error - InsertStationPowerbank - r.Builder: %s", err)
		return err
	}

	_, err = r.Pool.Exec(context.Background(), sql, args...)
	if err != nil {
		log.Printf("Error - InsertStationPowerbank - r.Pool.Exec: %s", err)
		return err
	}

	return nil
}

func (r *Repo) DeleteStationPowerbank(powerbankId int, stationId int) error {
	sql, args, err := r.Builder.
		Delete(stationsPowerbanksTableName).
		Where(squirrel.And{
			squirrel.Eq{"powerbank": powerbankId},
			squirrel.Eq{"station": stationId},
		}).
		ToSql()
	if err != nil {
		log.Printf("Error - DeleteStationPowerbank - r.Builder: %s", err)
		return err
	}

	_, err = r.Pool.Exec(context.Background(), sql, args...)
	if err != nil {
		log.Printf("Error - DeleteStationPowerbank - r.Pool.Exec: %s", err)
		return err
	}

	return nil
}

// InsertUserPowerbank TODO return value
func (r *Repo) InsertUserPowerbank(userId int, powerbankId int, stationId int) error {
	sql, args, err := r.Builder.
		Insert(usersPowerbanksTableName).
		Columns("user, powerbank").
		Values(userId, powerbankId, stationId).
		ToSql()
	if err != nil {
		log.Printf("Error - InsertUserPowerbank - r.Builder: %s", err)
		return err
	}

	_, err = r.Pool.Exec(context.Background(), sql, args...)
	if err != nil {
		log.Printf("Error - InsertUserPowerbank - r.Pool.Exec: %s", err)
		return err
	}

	return nil
}

func (r *Repo) DeleteUserPowerbank(userId int, powerbankId int) error {
	sql, args, err := r.Builder.
		Delete(usersPowerbanksTableName).
		Where(squirrel.And{
			squirrel.Eq{"user": userId},
			squirrel.Eq{"powerbank": powerbankId},
		}).
		ToSql()
	if err != nil {
		log.Printf("Error - DeleteUserPowerbank - r.Builder: %s", err)
		return err
	}

	_, err = r.Pool.Exec(context.Background(), sql, args...)
	if err != nil {
		log.Printf("Error - DeleteUserPowerbank - r.Pool.Exec: %s", err)
		return err
	}

	return nil
}

// TakePowerbank take user powerbank
// TODO return value
func (r *Repo) TakePowerbank(userId int, powerbankId int) error {
	sql, args, err := r.Builder.
		Insert(usersPowerbanksTableName).
		Columns("user, powerbank").
		Values(userId, powerbankId).
		ToSql()
	if err != nil {
		log.Printf("Error - TakePowerbank - r.Builder - 1: %s", err)
		return err
	}

	_, err = r.Pool.Exec(context.Background(), sql, args...)
	if err != nil {
		log.Printf("Error - TakePowerbank - r.Pool.Exec - 1: %s", err)
		return err
	}

	sql, args, err = r.Builder.
		Update("powerbanks").
		Set("used", true).
		Where(squirrel.Eq{"id": powerbankId}).
		ToSql()
	if err != nil {
		log.Printf("Error - TakePowerbank - r.Builder - 2: %s", err)
		return err
	}

	_, err = r.Pool.Exec(context.Background(), sql, args...)
	if err != nil {
		log.Printf("Error - TakePowerbank - r.Pool.Exec - 2: %s", err)
		return err
	}

	return nil
}

// PutPowerbank TODO return value
func (r *Repo) PutPowerbank(userId int, powerbankId int, stationId int, position int) error {
	err := r.DeleteUserPowerbank(userId, powerbankId)
	if err != nil {
		log.Printf("Error - PutPowerbank(BackTakedPowerbank): %s", err)
		return err
	}

	err = r.InsertStationPowerbank(powerbankId, stationId, position)
	if err != nil {
		log.Printf("Error - PutPowerbank(BackTakedPowerbank): %s", err)
		return err
	}

	return nil
}

// AddPowerbankToStation TODO return value
func (r *Repo) AddPowerbankToStation(powerbankId int, stationId int, position int) error {
	err := r.InsertStationPowerbank(powerbankId, stationId, position)
	if err != nil {
		log.Printf("Error - AddPowerbankToStation - %s", err)
		return err
	}

	return nil
}
