package query

import "github.com/api-sample/app/domain/model"

// make mockgen SOURCE=user.go
type UserQuery interface {
	FindByID(id string) (model.User, error)
	FindByEmail(id string) (model.User, error)
}
