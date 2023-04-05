package model

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	BirthDay  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *User) Exists() bool {
	return m.ID != ""
}

func HashPass(email, pass string) string {
	r := sha256.Sum256([]byte(email))
	mac := hmac.New(sha256.New, r[:])
	mac.Write([]byte(pass))
	return hex.EncodeToString(mac.Sum(nil))
}

func (m *User) ValidPass(pass string) bool {
	return HashPass(m.Email, pass) == m.Password
}
