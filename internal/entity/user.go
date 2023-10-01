package entity

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
	"io"
)

var saltSize = 32

type UserLogin struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID           int                `json:"id"`
	Username     string             `json:"username"`
	Email        string             `json:"email"`
	RoleID       int                `json:"role"`
	AddressID    int                `json:"address_id"`
	Password     string             `json:"password,omitempty"`
	PasswordHash string             `json:"-"`
	PasswordSalt string             `json:"-"`
	SuspendedAt  pgtype.Timestamptz `json:"suspended_at"`
	CreateAt     pgtype.Timestamptz `json:"create_at"`
	UpdateAt     pgtype.Timestamptz `json:"update_at"`
	DeleteAt     pgtype.Timestamptz `json:"delete_at"`
}

func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		salt, err := generateSalt(u.Password)
		if err != nil {
			return err
		}
		enc, err := encryptString(u.Password, salt)
		if err != nil {
			return err
		}
		u.PasswordSalt = salt
		u.PasswordHash = enc
	}

	return nil
}

func (u *User) Sanitize() {
	u.Password = ""
}

func generateSalt(secret string) (string, error) {
	secretB := []byte(secret)
	buf := make([]byte, saltSize, saltSize+sha1.Size)
	_, err := io.ReadFull(rand.Reader, buf)

	if err != nil {
		return "", fmt.Errorf("random read failed: %v", err)
	}

	hash := sha1.New()
	hash.Write(buf)
	hash.Write(secretB)
	return string(hash.Sum(buf)), nil
}

func (u *User) ComparePassword(password string) bool {
	password = password + u.PasswordSalt
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil

}

func encryptString(s string, salt string) (string, error) {
	s = s + salt
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
