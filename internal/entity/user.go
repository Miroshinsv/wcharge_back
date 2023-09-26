package entity

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type User struct {
	ID       int                `json:"id"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
	RoleID   int                `json:"role"`
	CreateAt time.Time          `json:"create_at"`
	UpdateAt time.Time          `json:"update_at"`
	DeleteAt pgtype.Timestamptz `json:"delete_at"`
}
