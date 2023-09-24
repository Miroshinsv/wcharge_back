package postgres

import (
	"github.com/Miroshinsv/wcharge_back/pkg/postgres"
)

const _defaultEntityCap = 64

type PostgresRepo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) *PostgresRepo {
	return &PostgresRepo{pg}
}
