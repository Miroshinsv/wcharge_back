package usecase

import (
	"fmt"

	"github.com/Miroshinsv/wcharge_back/internal/entity"
)

func (uc *UseCase) CreateStation(s entity.Station) error {
	err := uc.postgres.CreateStationRepo(s)
	if err != nil {
		return fmt.Errorf("UseCase - CreateStation - uc.repo.CreateStationRepo: %w", err)
	}

	return nil
}

func (uc *UseCase) UpdateStation(s entity.Station) error {
	err := uc.postgres.UpdateStationRepo(s)
	if err != nil {
		return fmt.Errorf("UseCase - UpdateStation - uc.repo.UpdateStationRepo: %w", err)
	}

	return nil
}

func (uc *UseCase) DeleteStation(s entity.Station) error {
	err := uc.postgres.DeleteStationRepo(s)
	if err != nil {
		return fmt.Errorf("UseCase - DeleteStation - uc.repo.DeleteStationRepo: %w", err)
	}

	return nil
}

func (uc *UseCase) GetStation(s entity.Station) (entity.Station, error) {
	station, err := uc.postgres.GetStationRepo(s)
	if err != nil {
		return entity.Station{}, fmt.Errorf("UseCase - GetStation - uc.repo.GetStationRepo: %w", err)
	}

	return station, nil
}

func (uc *UseCase) GetStations() ([]entity.Station, error) {
	stations, err := uc.postgres.GetStationsRepo()
	if err != nil {
		return nil, fmt.Errorf("UseCase - GetStations - uc.repo.GetStationsRepo: %w", err)
	}

	return stations, nil
}
