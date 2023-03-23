package user

import (
	"net/http"

	"github.com/api-sample/app/pkg/result"
)

func (u Usecase) FindByID(in FindByIDInput) (FindByIDOutput, result.Response) {
	user := u.q.UserQuery.FindByID(in.ID)

	if !user.Exists() {
		return FindByIDOutput{}, result.NewResponce(
			http.StatusNotFound,
			"ユーザーが見つかりませんでした",
		)
	}

	return FindByIDOutput{
		ID:       user.ID,
		Name:     user.Name,
		BirthDay: user.BirthDay,
	}, result.Success("success")
}
