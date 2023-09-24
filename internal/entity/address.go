package entity

type Address struct {
	ID      int     `json:"id"`
	Country string  `json:"country"`
	City    string  `json:"city"`
	Address string  `json:"address"`
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
}
