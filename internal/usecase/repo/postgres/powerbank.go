package postgres

import (
	"github.com/Miroshinsv/wcharge_back/internal/entity"
	//"github.com/Miroshinsv/wcharge_back/pkg/postgres"
)

func (r *PostgresRepo) CreatePowerbankRepo(p entity.Powerbank) error {

	return nil
}

func (r *PostgresRepo) UpdatePowerbankRepo(p entity.Powerbank) error {

	return nil
}

func (r *PostgresRepo) DeletePowerbankRepo(p entity.Powerbank) error {

	return nil
}

func (r *PostgresRepo) GetPowerbankRepo(p entity.Powerbank) (entity.Powerbank, error) {

	return entity.Powerbank{}, nil
}

func (r *PostgresRepo) GetPowerbanksRepo() ([]entity.Powerbank, error) {

	return nil, nil
}
