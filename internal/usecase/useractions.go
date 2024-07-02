package usecase

import (
	"context"
	"log"

	"github.com/Miroshinsv/wcharge_back/internal/entity"
)

// GetUserPowerbanks all user's powerbanks
func (uc *UseCase) GetUserPowerbanks(userId int) (*[]entity.Powerbank, error) {
	powerbanks, err := uc.postgres.GetUsersPowerbanks(userId)
	if err != nil {
		log.Printf("Error - UseCase - GetUserPowerbanks - GetUserPowerbanksRepo: %s", err)
		return nil, err
	}
	return powerbanks, nil
}

func (uc *UseCase) TakePowerbank(userId int, stationId int) (*entity.Powerbank, error) {
	st, err := uc.postgres.GetStation(stationId)
	if err != nil {
		log.Printf("Error - UseCase - uc.postgres.GetStationRepo: %s", err)
		return nil, err
	}

	pb, err := uc.postgres.GetRamdomPowerbank()
	if err != nil {
		log.Printf("Error - UseCase - uc.postgres.GetRamdomPowebank: %s", err)
		return nil, err
	}

	_, err = uc.mqtt.PushPowerBank(context.Background(), st, pb)
	if err != nil {
		log.Printf("Error - UseCase - uc.mqtt.PushPowerBank: %s", err)
		return nil, err
	}

	err = uc.postgres.TakePowerbank(userId, pb.ID)
	if err != nil {
		log.Printf("Error - UseCase -TakePowerbank - TakePowerbank: %s", err)
		return nil, err
	}
	return pb, nil
}

func (uc *UseCase) PutPowerbank(userId int, powerbankId int, stationId int, position int) error {
	err := uc.postgres.PutPowerbank(userId, powerbankId, stationId, position)
	if err != nil {
		log.Printf("Error - UseCase - PutPowerbank - ReturnPowerbankRepo: %s", err)
		return err
	}
	// TODO Ещё нужно вернуть через mqtt
	/*
		err = uc.mqtt.ReturnPowerbank(powerbank.SerialNumber, station.SerialNumber)
		if err != nil {
			err2 := uc.postgres.TakePowerbank(user.ID, powerbank.ID, station.ID) // откатить изменения ReturnPowerbankRepo
			if err2 != nil {
				return fmt.Errorf("UseCase - TakePowerbank - TakePowerbank (%s) \n BackTakePowerbankRepo (%s)", err.Error(), err2.Error())
			}
			return fmt.Errorf("UseCase - TakePowerbank - mqtt.TakePowerbank - %s", err.Error())
		}
	*/
	return nil
}

func (uc *UseCase) AddPowerbankToStation(powerbankId int, stationId int, position int) error {
	err := uc.postgres.AddPowerbankToStation(powerbankId, stationId, position)
	if err != nil {
		log.Printf("Error - AddPowerbankToStation - %s", err)
		return err
	}
	// Ещё нужно вернуть через mqtt
	return nil
}
