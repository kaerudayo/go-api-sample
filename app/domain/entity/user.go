package entity

import (
	"database/sql"

	"github.com/api-sample/app/domain/model"
	"github.com/api-sample/app/pkg/db"
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

func (e User) Model() model.User {
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

func NewUserEntity(m model.User) User {
	return User{
		ID:        m.ID,
		Name:      m.Name,
		Email:     m.Email,
		Password:  m.Password,
		BirthDay:  db.NewSQLNullTime(m.BirthDay),
		CreatedAt: db.NewSQLNullTime(m.CreatedAt),
		UpdatedAt: db.NewSQLNullTime(m.CreatedAt),
	}
}
