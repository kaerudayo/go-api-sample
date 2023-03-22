package user

import (
	"github.com/api-sample/app/domain/repository/query"
	infra "github.com/api-sample/app/infra/query"
	"gorm.io/gorm"
)

type c struct {
}

type q struct {
	UserQuery query.UserQuery
}

type UserUsecase struct {
	c c
	q q
}

func NewUserUsecase(db *gorm.DB) UserUsecase {
	return UserUsecase{
		q: q{
			infra.NewUserQueryImpl(db),
		},
	}
}
