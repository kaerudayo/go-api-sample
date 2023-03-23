package user

import (
	"github.com/api-sample/app/domain/repository/command"
	"github.com/api-sample/app/domain/repository/query"
	"github.com/api-sample/app/infra/reader"
	"github.com/api-sample/app/infra/writer"
	"gorm.io/gorm"
)

type c struct {
	UserCommand command.UserCommand
}

type q struct {
	UserQuery query.UserQuery
}

type Usecase struct {
	c c
	q q
}

func NewUsecase(db *gorm.DB) Usecase {
	return Usecase{
		q: q{
			UserQuery: reader.NewUserQueryImpl(db),
		},
		c: c{
			UserCommand: writer.NewUserCommandImpl(db),
		},
	}
}
