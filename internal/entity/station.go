package entity

import "time"

type Station struct {
	ID           int       `json:"id"`
	SerialNumber string    `json:"serial_number"`
	Address      *Address  `json:"address"`
	Capacity     int       `json:"capacity"`
	FreeCapacity int       `json:"free_capacity"`
	CreateAt     time.Time `json:"create_at"`
	UpdateAt     time.Time `json:"update_at"`
	DeteteAt     time.Time `json:"delete_at"`
}
