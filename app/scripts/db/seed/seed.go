package seed

import (
	"database/sql"
	"time"

	"github.com/api-sample/app/domain/entity"
	"github.com/api-sample/app/infra"
)

func DefaultSeed(*sql.DB) {
	u := entity.User{
		ID:       "user_1",
		Name:     "user_1",
		Email:    "user_1@example.com",
		Password: "user_1_pass",
		CreatedAt: sql.NullTime{
			Time:  time.Time{},
			Valid: false,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Time{},
			Valid: false,
		},
	}
	if err := infra.DB.Create(&u).Error; err != nil {
		panic(err)
	}

	u = entity.User{
		ID:        "user_2",
		Name:      "user_2",
		Email:     "user_2@example.com",
		Password:  "user_2_pass",
		CreatedAt: sql.NullTime{},
		UpdatedAt: sql.NullTime{},
	}
	if err := infra.DB.Create(&u).Error; err != nil {
		panic(err)
	}
}
