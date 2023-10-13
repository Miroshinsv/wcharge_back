package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testUser = User{
	Username: "user1",
	Password: "user1",
}

func TestUser_ComparePassword(t *testing.T) {
	testClasses := []struct {
		name    string
		payload map[string]string
		rez     bool
	}{
		{
			name: "valid pass",
			payload: map[string]string{
				"username": "user1",
				"password": "user1",
			},
			rez: true,
		},
		{
			name: "invalid pass",
			payload: map[string]string{
				"username": "user1",
				"password": "invalid",
			},
			rez: false,
		},
	}

	for _, tc := range testClasses {
		t.Run(tc.name, func(t *testing.T) {
			u := testUser
			err := u.BeforeCreate()
			if err != nil {
				t.Errorf("BeforeCreate " + err.Error())
			}
			assert.Equal(t, tc.rez, u.ComparePassword(tc.payload["password"]))
		})
	}
}

func TestUser_BeforeCreate(t *testing.T) {
	testClasses := []struct {
		name                      string
		payload                   User
		createPasswordHashAndSalt bool
	}{
		{
			name: "valid pass",
			payload: User{
				Username: "user1",
				Password: "user1",
			},
			createPasswordHashAndSalt: true,
		},
		{
			name: "invalid pass",
			payload: User{
				Username: "user1",
				Password: "",
			},
			createPasswordHashAndSalt: false,
		},
	}

	for _, tc := range testClasses {
		t.Run(tc.name, func(t *testing.T) {
			u := tc.payload
			err := u.BeforeCreate()
			if err != nil {
				t.Errorf("BeforeCreate " + err.Error())
			}
			assert.Equal(t, tc.createPasswordHashAndSalt, u.PasswordHash != "")
			assert.Equal(t, tc.createPasswordHashAndSalt, u.PasswordSalt != "")
		})
	}
}
