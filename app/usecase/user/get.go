package user

import (
	"net/http"

	"github.com/api-sample/app/pkg/logger"
	"github.com/api-sample/app/pkg/result"
	"github.com/labstack/echo/v4"
)

func (u Usecase) FindByID(in FindByIDInput, c echo.Context) (FindByIDOutput, result.Response) {
	user, err := u.q.UserQuery.FindByID(in.ID)
	if err != nil {
		logger.Error(err.Error(), c)
		return FindByIDOutput{}, result.NewInternalServerError("サーバーエラーが発生しました")
	}

	if !user.Exists() {
		return FindByIDOutput{}, result.NewResponce(
			http.StatusNotFound,
			"ユーザーが見つかりませんでした",
		)
	}

	return FindByIDOutput{
		ID:   user.ID,
		Name: user.Name,
		ID:   user.ID,
		Name: user.Name,
	}, result.Success("success")
}
