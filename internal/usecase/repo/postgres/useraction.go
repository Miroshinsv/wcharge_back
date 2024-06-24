package postgres

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/Miroshinsv/wcharge_back/internal/entity"
)

// TODO

// делать проверку на ПОВТОРНУЮ ЗАПИСЬ и ПОВТОРНОЕ УДАЛЕНИЕ
// GetUserPowerbanksRepo All user's powerbanks
func (r *Repo) GetUserPowerbanksRepo(userId int) (*[]entity.Powerbank, error) {
	//sql, args, err := r.Builder.
	//	Select("powerbank_id").
	//	From("tbl_user_powerbank").
	//	Where(squirrel.Eq{"user_id": userId}).
	//	ToSql()
	//if err != nil {
	//	return nil, fmt.Errorf("Repo - GetUserPowerbanksRepo - r.Builder: %w", err)
	//}
	//ctx := context.Background()
	//rows, err := r.Pool.Query(ctx, sql, args)
	//if err != nil {
	//	return nil, fmt.Errorf("Repo - GetUserPowerbanksRepo - r.Pool.Query: %w", err)
	//}
	//defer rows.Close()
	//
	//var powerbanksId []int
	//
	//for rows.Next() {
	//	var i int
	//	err = rows.Scan(&i)
	//	if err != nil {
	//		return nil, fmt.Errorf("Repo - GetUserPowerbanksRepo - rows.Scan: %w", err)
	//	}
	//	powerbanksId = append(powerbanksId, i)
	//}

	var entities []entity.Powerbank
	//for _, powerbankId := range powerbanksId {
	//	pw, _ := r.GetPowerbank(powerbankId)
	//	entities = append(entities, *pw)
	//}

	return &entities, nil
}

func (r *Repo) InsertStationPowerbank(powerbankId int, stationId int, position int) error {
	//sql, args, err := r.Builder.
	//	Insert("public.tbl_station_powerbank").
	//	Columns("station_id, powerbank_id, position").
	//	Values(stationId, powerbankId, position).
	//	ToSql()
	//if err != nil {
	//	return fmt.Errorf("InsertStationPowerbank - r.Builder: %w", err)
	//}
	//ctx := context.Background()
	//_, err = r.Pool.Exec(ctx, sql, args...)
	//if err != nil {
	//	return fmt.Errorf("InsertStationPowerbank - r.Pool.Exec: %w", err)
	//}

	return nil
}

func (r *Repo) deleteStationPowerbank(powerbankId int, stationId int) error {
	//sql, args, err := r.Builder.
	//	Delete("public.tbl_station_powerbank").
	//	Where(squirrel.And{
	//		squirrel.Eq{"powerbank_id": powerbankId},
	//		squirrel.Eq{"station_id": stationID},
	//	}).
	//	ToSql()
	//if err != nil {
	//	return fmt.Errorf("deleteStationPowerbank - r.Builder: %w", err)
	//}
	//ctx := context.Background()
	//_, err = r.Pool.Exec(ctx, sql, args...)
	//if err != nil {
	//	return fmt.Errorf("deleteStationPowerbank - r.Pool.Exec: %w", err)
	//}

	return nil
}

func (r *Repo) insertUserPowerbank(userId int, powerbankId int, stationId int) error {
	//sql, args, err := r.Builder.
	//	Insert("public.tbl_user_powerbank").
	//	Columns("user_id, powerbank_id").
	//	Values(userId, powerbankId, stationID).
	//	ToSql()
	//if err != nil {
	//	return fmt.Errorf("insertUserPowerbank - r.Builder: %w", err)
	//}
	//ctx := context.Background()
	//_, err = r.Pool.Exec(ctx, sql, args...)
	//fmt.Print(args...)
	//if err != nil {
	//	return fmt.Errorf("insertUserPowerbank - r.Pool.Exec: %w", err)
	//}

	return nil
}

func (r *Repo) deleteUserPowerbank(userId int, powerbankId int) error {
	//sql, args, err := r.Builder.
	//	Delete("public.tbl_user_powerbank").
	//	Where(squirrel.And{
	//		squirrel.Eq{"user_id": userId},
	//		squirrel.Eq{"powerbank_id": powerbankId},
	//	}).
	//	ToSql()
	//if err != nil {
	//	return fmt.Errorf("deleteUserPowerbank - r.Builder: %w", err)
	//}
	//ctx := context.Background()
	//_, err = r.Pool.Exec(ctx, sql, args...)
	//if err != nil {
	//	return fmt.Errorf("deleteUserPowerbank - r.Pool.Exec: %w", err)
	//}

	return nil
}

// TakePowerbankRepo take user powerbank
func (r *Repo) TakePowerbank(userId int, powerbankId int) error {

	sql, args, err := r.Builder.
		Insert("rel__users__powerbanks").
		Columns("user, powerbank").
		Values(userId, powerbankId).
		ToSql()
	if err != nil {
		return err
	}

	ctx := context.Background()
	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	sql, args, err = r.Builder.
		Update("powerbanks").
		Set("used", true).
		Where(squirrel.Eq{"id": powerbankId}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) PutPowerbankRepo(userId int, powerbankId int, stationId int, position int) error {
	//err := r.deleteUserPowerbank(userId, powerbankId)
	//if err != nil {
	//	return fmt.Errorf("BackTakePowerbankRepo - %w", err)
	//}
	//err = r.InsertStationPowerbank(powerbankId, stationId, position)
	//if err != nil {
	//	return fmt.Errorf("BackTakePowerbankRepo - %w", err)
	//}
	return nil
}

func (r *Repo) AddPowerbankToStationRepo(powerbankId int, stationId int, position int) error {
	//err := r.InsertStationPowerbank(powerbankId, stationId, position)
	//if err != nil {
	//	return fmt.Errorf("AddPowerbankToStationRepo - %w", err)
	//}
	return nil
}
