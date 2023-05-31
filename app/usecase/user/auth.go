package user

import (
	"fmt"
	"net/http"

	"github.com/api-sample/app/domain/model"

	"github.com/api-sample/app/pkg/db"
	"github.com/api-sample/app/pkg/logger"
	"github.com/api-sample/app/pkg/result"
	"github.com/labstack/echo/v4"
)

func (u Usecase) SignUp(in SignUpInput, c echo.Context) result.Response {
	user, err := u.q.UserQuery.FindByEmail(in.Email)
	if err != nil {
		logger.Error(err.Error(), c)
		return result.NewInternalServerError("Internal Server Error")
	}

	if user.Exists() {
		return result.NewResponse(
			http.StatusFound,
			"Users of this email are already registered",
		)
	}

	user.Email = in.Email
	user.Password = model.HashPass(in.Email, in.Password)

	if err := u.c.UserCommand.Create(&user); err != nil {
		logger.Error(err.Error(), c)

		return result.NewInternalServerError("Internal Server Error")
	}

	return result.Success("")
}

func (u Usecase) Login(in LoginInput, c echo.Context) (LoginOutput, result.Response) {
	user, err := u.q.UserQuery.FindByEmail(in.Email)
	if err != nil {
		logger.Error(err.Error(), c)
		return LoginOutput{}, result.NewInternalServerError("Internal Server Error")
	}

	if !user.Exists() {
		return LoginOutput{}, result.NewResponse(
			http.StatusNotFound,
			"email or password is incorrect",
		)
	}

	if !user.ValidPass(in.Password) {
		return LoginOutput{}, result.NewResponse(
			http.StatusUnauthorized,
			"email or password is incorrect",
		)
	}

	token, err := db.CreateAndSetToken(user.ID)
	if err != nil {
		logger.Error(err.Error(), c)
		return LoginOutput{},
			result.NewInternalServerError("Internal Server Error")
	}

	return LoginOutput{
		ID:          user.ID,
		AccessToken: fmt.Sprintf("Bearer %s", token),
	}, result.Success("")
}
