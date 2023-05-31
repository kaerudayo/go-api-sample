package user

import (
	"github.com/api-sample/app/domain/repository/command"
	"github.com/api-sample/app/domain/repository/query"
)

type U struct {
	Q Q
	C C
}

type C struct {
	UserCommand command.UserCommand
}

type Q struct {
	UserQuery query.UserQuery
}

func NewTestUsecase(u U) Usecase {
	return Usecase{
		c: c{
			UserCommand: u.C.UserCommand,
		},
		q: q{
			UserQuery: u.Q.UserQuery,
		},
	}
}
