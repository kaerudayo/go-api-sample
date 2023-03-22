package user

import (
	"github.com/api-sample/app/infra"
	"github.com/api-sample/app/usecase/user"
	"github.com/labstack/echo/v4"
)

func FindById(c echo.Context) error {
	body := findByIdReq{
		ID: c.Param("id"),
	}

	input := body.toInput()
	u := user.NewUserUsecase(infra.DB)

	output, res := u.FindById(input)
	if res.IsErr() {
		return c.JSON(res.Code, res.Msg)
	}
	return c.JSON(res.Code, output)
}
