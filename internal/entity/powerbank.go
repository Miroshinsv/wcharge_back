package entity

import "time"

type Powerbank struct {
	ID           int       `json:"id"`
	SerialNumber string    `json:"serial_number"`
	Capacity     int       `json:"capacity"` // объем заряда
	Used         int       `json:"used"`     // сколько уже использована банка в часах
	CreateAt     time.Time `json:"create_at"`
	UpdateAt     time.Time `json:"update_at"`
	DeleteAt     time.Time `json:"delete_at"`
}
