package usecase

import (
	"fmt"
	"github.com/Miroshinsv/wcharge_back/internal/entity"
	"log"
)

func (uc *UseCase) CreatePowerbank(p entity.Powerbank) (*entity.Powerbank, error) {
	pb, err := uc.postgres.CreatePowerbank(p)
	if err != nil {
		log.Printf("Error - UseCase - CreateStation - uc.repo.CreateStation: %s", err)
		return nil, err
	}

	return pb, nil
}

func (uc *UseCase) UpdatePowerbank(id int, p entity.Powerbank) (*entity.Powerbank, error) {
	pb, err := uc.postgres.UpdatePowerbank(p, id)
	if err != nil {
		log.Printf("Error - UseCase - UpdateStation - uc.repo.UpdateStation: %s", err)
		return nil, err
	}

	return pb, nil
}

func (uc *UseCase) DeletePowerbank(id int) error {
	err := uc.postgres.DeletePowerbank(id)
	if err != nil {
		log.Printf("Error - UseCase - DeleteStation - uc.repo.DeleteStation: %s", err)
		return fmt.Errorf("UseCase - DeleteStation - uc.repo.DeleteStation: %w", err)
	}

	return nil
}

func (uc *UseCase) GetPowerbank(id int) (*entity.Powerbank, error) {
	powerbank, err := uc.postgres.GetPowerbank(id)
	if err != nil {
		log.Printf("Error - GetPowerbank - uc.repo.GetStation: %s", err)
		return nil, fmt.Errorf("GetPowerbank - uc.repo.GetStation: %w", err)
	}

	return powerbank, nil
}

func (uc *UseCase) GetPowerbanks() (*[]entity.Powerbank, error) {
	powerbanks, err := uc.postgres.GetPowerbanks()
	if err != nil {
		log.Printf("Error - UseCase - GetStations - uc.repo.GetStations: %s", err)
		return nil, fmt.Errorf("UseCase - GetStations - uc.repo.GetStations: %w", err)
	}

	return powerbanks, nil
}

func (uc *UseCase) GetAllPowerbanksInStation(id int) (*[]entity.Powerbank, error) {
	powerbanks, err := uc.postgres.GetAllPowerbanksInStation(id)
	if err != nil {
		log.Printf("Error - UseCase - GetAllPowerbanksInStation - %s", err)
		return nil, err
	}

	return powerbanks, nil
}
