package entity

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
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
	ID             int    `json:"id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	RoleID         int    `json:"role"`
	RoleName       string `json:"role_name"`
	RolePrivileges int    `json:"-"`
	AddressID      int    `json:"address_id"`
	Password       string `json:"password,omitempty"`
	PasswordHash   string `json:"-"`
	PasswordSalt   string `json:"-"`
	Removed        int    `json:"removed"`
	//SuspendedAt    pgtype.Timestamptz `json:"suspended_at"`
	CreateAt pgtype.Timestamptz `json:"create_at"`
	UpdateAt pgtype.Timestamptz `json:"update_at"`
	DeleteAt pgtype.Timestamptz `json:"delete_at"`
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
	salt := base64.StdEncoding.EncodeToString(hash.Sum(buf))
	return salt, nil
}

func (u *User) ComparePassword(password string) bool {
	bSalt, _ := base64.StdEncoding.DecodeString(u.PasswordSalt)
	bPassword := append([]byte(password), bSalt...)
	bPasswordHash, _ := base64.StdEncoding.DecodeString(u.PasswordHash)
	err := bcrypt.CompareHashAndPassword(bPasswordHash, bPassword)
	return err == nil

}

func encryptString(s string, salt string) (string, error) {
	bSalt, _ := base64.StdEncoding.DecodeString(salt)
	ss := append([]byte(s), bSalt...)
	b, err := bcrypt.GenerateFromPassword(ss, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	pass := base64.StdEncoding.EncodeToString(b)
	return pass, nil
}
