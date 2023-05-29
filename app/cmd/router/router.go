package router

import (
	"github.com/api-sample/app/cmd/handler"
	"github.com/api-sample/app/cmd/middleware"
	"github.com/api-sample/app/pkg/logger"
	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.Use(logger.Middleware())
	e.Use(m.Recover())

	v1 := e.Group("/v1")
	NewAuthUserRoute(v1)

	v1.Use(middleware.ValidateToken())

	e.GET("/v1", handler.Top)

	NewUserRoute(v1)
	return e
}
