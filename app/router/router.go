package router

import (
	"github.com/api-sample/app/pkg/handler"
	"github.com/labstack/echo/v4"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	v1 := e.Group("/v1")
	v1.GET("/", handler.Login)
	return e
}
