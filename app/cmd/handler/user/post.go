package user

import (
	"net/http"

	"github.com/api-sample/app/infra"
	"github.com/api-sample/app/usecase/user"
	"github.com/labstack/echo/v4"
)

func Top(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World!")
}

func SignUp(c echo.Context) error {
	return c.JSON(http.StatusOK, "TODO")
}

func Login(c echo.Context) error {
	body := loginReq{}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, "不正なリクエストです")
	}
	input := body.toInput()
	u := user.NewUsecase(infra.DB)
	output, res := u.Login(input, c)
	if res.IsErr() {
		return c.JSON(res.Code, res.Msg)
	}
	return c.JSON(res.Code, output)

}
