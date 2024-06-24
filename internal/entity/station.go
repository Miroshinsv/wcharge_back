package entity

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Station struct {
	ID           int                 `json:"id"`
	SerialNumber string              `json:"serial_number"`
	Address      int                 `json:"address,omitempty"`
	Capacity     float64             `json:"capacity,omitempty"`
	FreeCapacity float64             `json:"free_capacity,omitempty"`
	Removed      bool                `json:"removed,omitempty"`
	CreateAt     *pgtype.Timestamptz `json:"create_at,omitempty" swaggertype:"string"`
	UpdateAt     *pgtype.Timestamptz `json:"update_at,omitempty" swaggertype:"string"`
	DeleteAt     *pgtype.Timestamptz `json:"delete_at,omitempty" swaggertype:"string"`

	AddressFull Address     `json:"address_full,omitempty"`
	Powerbanks  []Powerbank `json:"powerbanks,omitempty"`
}
