package command

import "github.com/api-sample/app/domain/model"

type UserCommand interface {
	Create(m *model.User) error
}
