package postgres

import (
	"github.com/Miroshinsv/wcharge_back/pkg/postgres"
)

const _defaultEntityCap = 64

type Repo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) *Repo {
	return &Repo{pg}
}
