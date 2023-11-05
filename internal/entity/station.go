package entity

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Station struct {
	ID           int                `json:"id"`
	SerialNumber string             `json:"serial_number"`
	AddressId    int                `json:"address"`
	Capacity     int                `json:"capacity"`
	FreeCapacity int                `json:"free_capacity"`
	Removed      int                `json:"removed"`
	CreateAt     pgtype.Timestamptz `json:"create_at"`
	UpdateAt     pgtype.Timestamptz `json:"update_at"`
	DeleteAt     pgtype.Timestamptz `json:"delete_at"`
}
