package router

import (
	"github.com/api-sample/app/cmd/handler/user"
	"github.com/labstack/echo/v4"
)

func NewUserRoute(group *echo.Group) {
	g := group.Group("/users")
	g.GET("/:id", user.FindByID)
}
