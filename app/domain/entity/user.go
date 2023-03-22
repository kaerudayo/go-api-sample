package entity

import (
	"database/sql"

	"github.com/api-sample/app/domain/model"
)

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	BirthDay  sql.NullTime
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

func (e User) toModel() model.User {
	return model.User{
		ID:        e.ID,
		Name:      e.Name,
		Email:     e.Email,
		Password:  e.Password,
		BirthDay:  e.BirthDay.Time,
		CreatedAt: e.CreatedAt.Time,
		UpdatedAt: e.CreatedAt.Time,
	}
}
