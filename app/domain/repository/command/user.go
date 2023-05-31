package command

import "github.com/api-sample/app/domain/model"

// make mockgen/writer SOURCE=user.go
type UserCommand interface {
	Create(m *model.User) error
}
