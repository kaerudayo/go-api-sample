package writer

import (
	"github.com/api-sample/app/domain/entity"
	"github.com/api-sample/app/domain/model"
	commad "github.com/api-sample/app/domain/repository/command"
	"gorm.io/gorm"
)

type UserCommandImpl struct {
	db *gorm.DB
}

func NewUserCommandImpl(db *gorm.DB) commad.UserCommand {
	userRepo := UserCommandImpl{db}
	return &userRepo
}

func (impl UserCommandImpl) Create(m *model.User) error {
	e := entity.NewUserEntity(*m)
	if err := impl.db.Create(&e).Error; err != nil {
		return err
	}
	return nil
}
