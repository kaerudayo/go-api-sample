package router

import (
	"github.com/api-sample/app/cmd/handler/user"
	"github.com/labstack/echo/v4"
)

func NewAuthUserRoute(group *echo.Group) {
	group.POST("/signup", user.SignUp)
	group.POST("/signin", user.Signin)
}
