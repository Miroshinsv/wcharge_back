package usecase

import (
	"github.com/Miroshinsv/wcharge_back/internal/entity"
)

func (uc *UseCase) CreatePowerbank(p entity.Powerbank) (*entity.Powerbank, error) {
	//pp, err := uc.postgres.CreatePowerbankRepo(p)
	//if err != nil {
	//	return nil, err //fmt.Errorf("UseCase - CreateStation - uc.repo.CreateStationRepo: %w", err)
	//}

	return nil, nil
}

func (uc *UseCase) UpdatePowerbank(id int, p entity.Powerbank) error {
	//err := uc.postgres.UpdatePowerbankRepo(id, p)
	//if err != nil {
	//	return fmt.Errorf("UseCase - UpdateStation - uc.repo.UpdateStationRepo: %w", err)
	//}
	//
	return nil
}

func (uc *UseCase) DeletePowerbank(id int) error {
	//err := uc.postgres.DeletePowerbankRepo(id)
	//if err != nil {
	//	return fmt.Errorf("UseCase - DeleteStation - uc.repo.DeleteStationRepo: %w", err)
	//}

	return nil
}

func (uc *UseCase) GetPowerbank(id int) (entity.Powerbank, error) {
	//powerbank, err := uc.postgres.GetPowerbankRepo(id)
	//if err != nil {
	//	return entity.Powerbank{}, fmt.Errorf("GetPowerbank - uc.repo.GetStationRepo: %w", err)
	//}

	return entity.Powerbank{}, nil
}

func (uc *UseCase) GetPowerbanks() ([]entity.Powerbank, error) {
	//powerbanks, err := uc.postgres.GetPowerbanksRepo()
	//if err != nil {
	//	return nil, fmt.Errorf("UseCase - GetStations - uc.repo.GetStationsRepo: %w", err)
	//}

	return nil, nil
}
