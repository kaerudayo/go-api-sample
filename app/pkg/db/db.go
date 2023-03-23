package db

import (
	"database/sql"
	"time"
)

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
