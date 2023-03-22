package router

import (
	"github.com/api-sample/app/cmd/handler/user"
	"github.com/labstack/echo/v4"
)

func NewUserRoute(e *echo.Echo) {
	v1 := e.Group("/v1")
	v1.GET("", user.Top)
	v1.GET("/signup", user.SignUp)
	v1.GET("/users/:id", user.FindById)
}
