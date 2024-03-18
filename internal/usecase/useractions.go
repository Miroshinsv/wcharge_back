package usecase

import (
	"context"
	"fmt"

	"github.com/Miroshinsv/wcharge_back/internal/entity"
)

// GetUserPowerbanks all user's powerbanks
func (uc *UseCase) GetUserPowerbanks(userId int) ([]entity.Powerbank, error) {
	powerbanks, err := uc.postgres.GetUserPowerbanksRepo(userId)
	if err != nil {
		return nil, fmt.Errorf("UseCase - GetUserPowerbanks - GetUserPowerbanksRepo - " + err.Error())
	}
	return powerbanks, nil
}

func (uc *UseCase) TakePowerbank(userId int, powerbankId int, stationId int) (bool, error) {

	// Ещё нужно выдать через mqtt
	/*
		err = uc.mqtt.TakePowerbank(powerbank.SerialNumber, station.SerialNumber)
		if err != nil {
			err2 := uc.postgres.ReturnPowerbankRepo(user.ID, powerbank.ID, station.ID) // откатить изменения TakePowerbankRepo
			if err2 != nil {
				return fmt.Errorf("UseCase - TakePowerbank - TakePowerbankRepo (%s) \n BackTakePowerbankRepo (%s)", err.Error(), err2.Error())
			}
			return fmt.Errorf("UseCase - TakePowerbank - mqtt.TakePowerbank - %s", err.Error())
		}
	*/
	ctx := context.Background()

	st, err := uc.postgres.GetStationRepo(stationId)
	if err != nil {
		return false, fmt.Errorf("UseCase - uc.postgres.GetStationRepo - %s", err.Error())
	}

	pb, err := uc.postgres.GetPowerbankRepo(powerbankId)
	if err != nil {
		return false, fmt.Errorf("UseCase - uc.postgres.GetPowerbankRepo - %s", err.Error())
	}

	rez, err := uc.mqtt.PushPowerBank(ctx, &st, &pb)
	if err != nil {
		return false, fmt.Errorf("UseCase - uc.mqtt.PushPowerBank - %s", err.Error())
	}

	err = uc.postgres.TakePowerbankRepo(userId, powerbankId, stationId)
	if err != nil {
		return false, fmt.Errorf("UseCase - TakePowerbank - TakePowerbankRepo - %s", err.Error())
	}
	return rez, nil
}

func (uc *UseCase) PutPowerbank(userId int, powerbankId int, stationId int) error {
	err := uc.postgres.PutPowerbankRepo(userId, powerbankId, stationId)
	if err != nil {
		return fmt.Errorf("UseCase - PutPowerbank - ReturnPowerbankRepo - %s", err.Error())
	}
	// Ещё нужно вернуть через mqtt
	/*
		err = uc.mqtt.ReturnPowerbank(powerbank.SerialNumber, station.SerialNumber)
		if err != nil {
			err2 := uc.postgres.TakePowerbankRepo(user.ID, powerbank.ID, station.ID) // откатить изменения ReturnPowerbankRepo
			if err2 != nil {
				return fmt.Errorf("UseCase - TakePowerbank - TakePowerbankRepo (%s) \n BackTakePowerbankRepo (%s)", err.Error(), err2.Error())
			}
			return fmt.Errorf("UseCase - TakePowerbank - mqtt.TakePowerbank - %s", err.Error())
		}
	*/
	return nil
}

func (uc *UseCase) AddPowerbankToStation(powerbankId int, stationId int) error {
	err := uc.postgres.AddPowerbankToStationRepo(powerbankId, stationId)
	if err != nil {
		return fmt.Errorf("AddPowerbankToStation - %w", err)
	}
	// Ещё нужно вернуть через mqtt
	return nil
}
