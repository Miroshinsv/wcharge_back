package usecase

import (
	"fmt"

	"github.com/Miroshinsv/wcharge_back/internal/entity"
)

func (uc *UseCase) CreatePowerbank(p entity.Powerbank) error {
	err := uc.postgres.CreatePowerbankRepo(p)
	if err != nil {
		return fmt.Errorf("UseCase - CreateStation - uc.repo.CreateStationRepo: %w", err)
	}

	return nil
}

func (uc *UseCase) UpdatePowerbank(p entity.Powerbank) error {
	err := uc.postgres.UpdatePowerbankRepo(p)
	if err != nil {
		return fmt.Errorf("UseCase - UpdateStation - uc.repo.UpdateStationRepo: %w", err)
	}

	return nil
}

func (uc *UseCase) DeletePowerbank(p entity.Powerbank) error {
	err := uc.postgres.DeletePowerbankRepo(p)
	if err != nil {
		return fmt.Errorf("UseCase - DeleteStation - uc.repo.DeleteStationRepo: %w", err)
	}

	return nil
}

func (uc *UseCase) GetPowerbank(p entity.Powerbank) (entity.Powerbank, error) {
	powerbank, err := uc.postgres.GetPowerbankRepo(p)
	if err != nil {
		return entity.Powerbank{}, fmt.Errorf("UseCase - GetStation - uc.repo.GetStationRepo: %w", err)
	}

	return powerbank, nil
}

func (uc *UseCase) GetPowerbanks() ([]entity.Powerbank, error) {
	powerbanks, err := uc.postgres.GetPowerbanksRepo()
	if err != nil {
		return nil, fmt.Errorf("UseCase - GetStations - uc.repo.GetStationsRepo: %w", err)
	}

	return powerbanks, nil
}
