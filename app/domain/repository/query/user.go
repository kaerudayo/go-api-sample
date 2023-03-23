package query

import "github.com/api-sample/app/domain/model"

type UserQuery interface {
	FindByID(id string) (model.User, error)
}
