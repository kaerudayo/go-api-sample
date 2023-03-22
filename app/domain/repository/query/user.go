package query

import "github.com/api-sample/app/domain/model"

type UserQuery interface {
	FindById(id string) (user model.User)
}
