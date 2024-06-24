package usecase

import (
	"fmt"

	"github.com/Miroshinsv/wcharge_back/internal/entity"
)

// GetUserPowerbanks all user's powerbanks
func (uc *UseCase) GetUserPowerbanks(userId int) (*[]entity.Powerbank, error) {
	//powerbanks, err := uc.postgres.GetUserPowerbanksRepo(userId)
	//if err != nil {
	//	return nil, fmt.Errorf("UseCase - GetUserPowerbanks - GetUserPowerbanksRepo - " + err.Error())
	//}
	return nil, nil
}

func (uc *UseCase) TakePowerbank(userId int, stationId int) (*entity.Powerbank, error) {

	// Ещё нужно выдать через mqtt
	/*
		err = uc.mqtt.TakePowerbank(powerbank.SerialNumber, station.SerialNumber)
		if err != nil {
			err2 := uc.postgres.ReturnPowerbankRepo(user.ID, powerbank.ID, station.ID) // откатить изменения TakePowerbank
			if err2 != nil {
				return fmt.Errorf("UseCase - TakePowerbank - TakePowerbank (%s) \n BackTakePowerbankRepo (%s)", err.Error(), err2.Error())
			}
			return fmt.Errorf("UseCase - TakePowerbank - mqtt.TakePowerbank - %s", err.Error())
		}
	*/
	//ctx := context.Background()

	//st, err := uc.postgres.GetStationRepo(stationId)
	//if err != nil {
	//	fmt.Printf("UseCase - uc.postgres.GetStationRepo - %s", err.Error())
	//	return nil, err
	//}

	//pb, err := uc.postgres.GetRamdomPowebank()
	//if err != nil {
	//	fmt.Printf("UseCase - uc.postgres.GetRamdomPowebank - %s", err.Error())
	//	return nil, err
	//}

	//_, err = uc.mqtt.PushPowerBank(ctx, &st, pb)
	//if err != nil {
	//	fmt.Printf("UseCase - uc.mqtt.PushPowerBank - %s", err.Error())
	//	return nil, err
	//}

	//err = uc.postgres.TakePowerbank(userId, pb.ID, stationId)
	//if err != nil {
	//	fmt.Printf("UseCase -TakePowerbank - TakePowerbank  - %s", err.Error())
	//	return nil, err
	//}
	return nil, nil
}

func (uc *UseCase) PutPowerbank(userId int, powerbankId int, stationId int, position int) error {
	err := uc.postgres.PutPowerbankRepo(userId, powerbankId, stationId, position)
	if err != nil {
		return fmt.Errorf("UseCase - PutPowerbank - ReturnPowerbankRepo - %s", err.Error())
	}
	// Ещё нужно вернуть через mqtt
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
	err := uc.postgres.AddPowerbankToStationRepo(powerbankId, stationId, position)
	if err != nil {
		return fmt.Errorf("AddPowerbankToStation - %w", err)
	}
	// Ещё нужно вернуть через mqtt
	return nil
}
