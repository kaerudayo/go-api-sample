package model

import (
	"encoding/hex"
	"time"

	"encoding/hex"
	"time"

	"golang.org/x/crypto/scrypt"
)

type User struct {
	ID       string
	Name     string
	Email    string
	Password string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *User) Exists() bool {
	return m.ID != ""
}

func HashPass(email, pass string) string {
	converted, _ := scrypt.Key([]byte(pass), []byte(email), 16384, 8, 1, 16)
	return hex.EncodeToString(converted[:])
}

func (m *User) ValidPass(pass string) bool {
	return HashPass(m.Email, pass) == m.Password
}
