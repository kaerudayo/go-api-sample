package db

import (
	"database/sql"
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

func GenID() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	hash := ulid.MustNew(ulid.Timestamp(t), entropy)

	return hash.String()
}

func NewSQLNullTime(t time.Time) sql.NullTime {
	res := sql.NullTime{
		Time:  t,
		Valid: true,
	}
	if t.IsZero() {
		res.Valid = false
	}
	return res
}
