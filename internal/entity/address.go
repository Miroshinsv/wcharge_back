package entity

type Address struct {
	ID      int     `json:"id"`
	Country int     `json:"country"`
	City    int     `json:"city"`
	Address string  `json:"address"`
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
}
