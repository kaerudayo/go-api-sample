package command

import "github.com/api-sample/app/domain/model"

type UserRepository interface {
	Create(m *model.User) error
}
