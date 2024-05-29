package entity

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Powerbank struct {
	ID           int                `json:"id"`
	Position     int                `json:"position"`
	SerialNumber string             `json:"serial_number"`
	Capacity     int                `json:"capacity"` // объем заряда
	Used         int                `json:"used"`     // сколько уже использована банка в часах // up юзается или нет
	Removed      int                `json:"removed"`
	CreateAt     pgtype.Timestamptz `json:"create_at"`
	UpdateAt     pgtype.Timestamptz `json:"update_at"`
	DeleteAt     pgtype.Timestamptz `json:"delete_at"`
}
