package postgres

import (
	"github.com/Miroshinsv/wcharge_back/internal/entity"
	//"github.com/Miroshinsv/wcharge_back/pkg/postgres"
)

func (r *Repo) CreatePowerbankRepo(p entity.Powerbank) error {

	return nil
}

func (r *Repo) UpdatePowerbankRepo(id int, p entity.Powerbank) error {

	return nil
}

func (r *Repo) DeletePowerbankRepo(id int) error {

	return nil
}

func (r *Repo) GetPowerbankRepo(id int) (entity.Powerbank, error) {

	return entity.Powerbank{}, nil
}

func (r *Repo) GetPowerbanksRepo() ([]entity.Powerbank, error) {

	return nil, nil
}
