package user

import (
	"net/http"

	"github.com/api-sample/app/pkg/result"
)

func (u UserUsecase) FindById(in FindByIdInput) (FindByIdOutput, result.Responce) {
	user := u.q.UserQuery.FindById(in.ID)

	if !user.Exists() {
		return FindByIdOutput{}, result.NewResponce(
			http.StatusNotFound,
			"ユーザーが見つかりませんでした",
		)
	}

	return FindByIdOutput{
		ID:       user.ID,
		Name:     user.Name,
		BirthDay: user.BirthDay,
	}, result.Success("success")
}
