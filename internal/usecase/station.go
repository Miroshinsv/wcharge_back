package usecase

import (
	"log"

	"github.com/Miroshinsv/wcharge_back/internal/entity"
)

func (uc *UseCase) CreateStation(s entity.Station) (*entity.Station, error) {
	ss, err := uc.postgres.CreateStation(s)
	if err != nil {
		log.Printf("Error - UseCase - CreateStation - uc.repo.CreateStation: %s", err)
		return nil, err
	}

	return ss, nil
}

func (uc *UseCase) UpdateStation(id int, s entity.Station) (*entity.Station, error) {
	ss, err := uc.postgres.UpdateStation(s, id)
	if err != nil {
		log.Printf("Error - UseCase - UpdateStation - uc.repo.UpdateStation: %s", err)
		return nil, err
	}

	return ss, nil
}

func (uc *UseCase) DeleteStation(id int) error {
	err := uc.postgres.DeleteStation(id)
	if err != nil {
		log.Printf("Error - UseCase - DeleteStation - uc.repo.DeleteStation: %s", err)
		return err
	}

	return nil
}

func (uc *UseCase) GetStation(id int) (*entity.Station, error) {
	station, err := uc.postgres.GetStation(id)
	if err != nil {
		log.Printf("Error - UseCase - GetStation - uc.repo.GetStation: %s", err)
		return nil, err
	}

	return station, nil
}

func (uc *UseCase) GetStations() (*[]entity.Station, error) {
	stations, err := uc.postgres.GetStations()
	if err != nil {
		log.Printf("Error - UseCase - GetStations - : %s", err)
		return nil, err
	}

	return stations, nil
}
