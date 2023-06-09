package seed

import (
	"database/sql"

	"github.com/api-sample/app/domain/entity"
	"github.com/api-sample/app/domain/model"
	"github.com/api-sample/app/infra"
)

func DefaultSeed(*sql.DB) {
	u := entity.User{
		ID:       "user_1",
		Name:     "user_1",
		Email:    "user_1@example.com",
		Password: "password",
	}
	u.Password = model.HashPass(u.Email, u.Password)
	if err := infra.DB.Create(&u).Error; err != nil {
		panic(err)
	}

	u = entity.User{
		ID:        "user_2",
		Name:      "user_2",
		Email:     "user_2@example.com",
		Password:  "password",
		CreatedAt: sql.NullTime{},
		UpdatedAt: sql.NullTime{},
	}
	u.Password = model.HashPass(u.Email, u.Password)
	if err := infra.DB.Create(&u).Error; err != nil {
		panic(err)
	}
}
