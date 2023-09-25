package postgres

import (
	"context"

	"github.com/Miroshinsv/wcharge_back/internal/entity"
	//"github.com/Miroshinsv/wcharge_back/pkg/postgres"
)

func (r *PostgresRepo) CreatePowerbank(ctx context.Context, p entity.Powerbank) error {

	return nil
}

func (r *PostgresRepo) UpdatePowerbank(ctx context.Context, p entity.Powerbank) error {

	return nil
}

func (r *PostgresRepo) DeletePowerbank(ctx context.Context, p entity.Powerbank) error {

	return nil
}

func (r *PostgresRepo) GetPowerbank(ctx context.Context, p entity.Powerbank) (entity.Powerbank, error) {

	return entity.Powerbank{}, nil
}

func (r *PostgresRepo) GetPowerbanks(ctx context.Context) ([]entity.Powerbank, error) {

	return nil, nil
}
