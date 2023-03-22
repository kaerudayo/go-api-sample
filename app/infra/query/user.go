package infra

import (
	"github.com/api-sample/app/domain/model"
	"github.com/api-sample/app/domain/repository/query"
	"gorm.io/gorm"
)

type UserQueryImpl struct {
	db *gorm.DB
}

func NewUserQueryImpl(db *gorm.DB) query.UserQuery {
	userRepo := UserQueryImpl{db}
	return &userRepo
}

func (impl UserQueryImpl) FindById(id string) model.User {
	var user model.User
	impl.db.Where("id = ?", id).Find(&user)
	return user
}
