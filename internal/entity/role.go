package entity

type Role struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Privileges int    `json:"privileges"`
}
