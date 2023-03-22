package router

import (
	"github.com/api-sample/app/cmd/handler"
	"github.com/labstack/echo/v4"
)

func NewUserRoute(e *echo.Echo) {
	v1 := e.Group("/v1")
	v1.GET("", handler.Top)
	v1.GET("/login", handler.Login)
}
