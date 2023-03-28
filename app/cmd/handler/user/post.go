package user

import (
	"net/http"

	"github.com/api-sample/app/infra"
	"github.com/api-sample/app/usecase/user"
	"github.com/labstack/echo/v4"
)

// curl -X POST 'localhost:5002/v1/signup' -H 'Content-Type: application/json' -d '{"email":"sample@example.com","password":"sample"}'
func SignUp(c echo.Context) error {
	body := signUpReq{}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, "不正なリクエストです")
	}

	in := body.toInput()
	u := user.NewUsecase(infra.DB)
	res := u.SignUp(in, c)
	if res.IsErr() {
		return c.JSON(res.Code, res)
	}
	return c.JSON(http.StatusOK, res)
}

// curl -X POST 'localhost:5002/v1/signin' -H 'Content-Type: application/json' -d '{"email":"sample@example.com","password":"sample"}'
func Signin(c echo.Context) error {
	body := loginReq{}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, "不正なリクエストです")
	}
	input := body.toInput()
	u := user.NewUsecase(infra.DB)
	output, res := u.Login(input, c)
	if res.IsErr() {
		return c.JSON(res.Code, res)
	}
	return c.JSON(res.Code, output)
}
