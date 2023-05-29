package user

import (
	"github.com/api-sample/app/infra"
	"github.com/api-sample/app/usecase/user"
	"github.com/labstack/echo/v4"
)

/*
*
curl -X GET 'localhost:5002/v1/users/01H1K5Y39BEM66XE6DAA3B16PG' \
-H 'Authorization: Bearer xxxx
*/
func FindByID(c echo.Context) error {
	body := findByIDReq{
		ID: c.Param("id"),
	}
	input := body.toInput()
	u := user.NewUsecase(infra.DB)
	output, res := u.FindByID(input, c)
	if res.IsErr() {
		return c.JSON(res.Code, res)
	}
	return c.JSON(res.Code, output)
}
