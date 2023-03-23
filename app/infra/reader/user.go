package reader

import (
	"github.com/api-sample/app/domain/entity"
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

func (impl UserQueryImpl) FindByID(id string) (model.User, error) {
	var user entity.User
	if err := impl.db.Where("id = ?", id).Find(&user).Error; err != nil {
		return model.User{}, err
	}
	return user.Model(), nil
}
