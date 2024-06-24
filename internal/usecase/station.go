package usecase

import (
	"fmt"

	"github.com/Miroshinsv/wcharge_back/internal/entity"
)

func (uc *UseCase) CreateStation(s entity.Station) (*entity.Station, error) {
	//ss, err := uc.postgres.CreateStationRepo(s)
	//if err != nil {
	//	return nil, err //fmt.Errorf("UseCase - CreateStation - uc.repo.CreateStationRepo: %w", err)
	//}

	return nil, nil
}

func (uc *UseCase) UpdateStation(id int, s entity.Station) error {
	//err := uc.postgres.UpdateStationRepo(id, s)
	//if err != nil {
	//	return fmt.Errorf("UseCase - UpdateStation - uc.repo.UpdateStationRepo: %w", err)
	//}

	return nil
}

func (uc *UseCase) DeleteStation(id int) error {
	//err := uc.postgres.DeleteStationRepo(id)
	//if err != nil {
	//	return fmt.Errorf("UseCase - DeleteStation - uc.repo.DeleteStationRepo: %w", err)
	//}

	return nil
}

func (uc *UseCase) GetStation(id int) (*entity.Station, error) {
	station, err := uc.postgres.GetStation(id)
	if err != nil {
		return nil, fmt.Errorf("UseCase - GetStation - uc.repo.GetStationRepo: %w", err)
	}

	return station, nil
}

func (uc *UseCase) GetStations() (*[]entity.Station, error) {
	stations, err := uc.postgres.GetStations()
	if err != nil {
		return nil, fmt.Errorf("UseCase - GetStations - : %w", err)
	}

	return stations, nil
}

func (uc *UseCase) GetAllPowerbanksInStation(id int) ([]entity.Powerbank, error) {
	//powerbanks, err := uc.postgres.GetAllPowerbanksInStationRepo(id)
	//if err != nil {
	//	return nil, fmt.Errorf("UseCase - GetAllPowerbanksInStation - %w", err)
	//}

	return nil, nil
}
