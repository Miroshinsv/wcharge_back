package entity

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Powerbank struct {
	ID           int                `json:"id"`
	Position     int                `json:"position"`
	SerialNumber string             `json:"serial_number"`
	Capacity     float64            `json:"capacity"` // объем заряда
	Used         bool               `json:"used"`     // сколько уже использована банка в часах // up юзается или нет
	Removed      bool               `json:"removed"`
	CreateAt     pgtype.Timestamptz `json:"create_at" swaggertype:"string"`
	UpdateAt     pgtype.Timestamptz `json:"update_at" swaggertype:"string"`
	DeleteAt     pgtype.Timestamptz `json:"delete_at" swaggertype:"string"`
}
