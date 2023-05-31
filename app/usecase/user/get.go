package user

import (
	"net/http"

	"github.com/api-sample/app/pkg/logger"
	"github.com/api-sample/app/pkg/result"
	"github.com/labstack/echo/v4"
)

func (u Usecase) FindByID(in FindByIDInput, c echo.Context) (FindByIDOutput, result.Response) {
	user, err := u.q.UserQuery.FindByID(in.UserID)
	if err != nil {
		logger.Error(err.Error(), c)
		return FindByIDOutput{}, result.NewInternalServerError("Internal Server Error")
	}

	if !user.Exists() {
		return FindByIDOutput{}, result.NewResponse(
			http.StatusNotFound,
			"User not found",
		)
	}

	return FindByIDOutput{
		ID:   user.ID,
		Name: user.Name,
	}, result.Success("")
}
