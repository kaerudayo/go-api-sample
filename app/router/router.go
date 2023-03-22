package router

import (
	"github.com/api-sample/app/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.Use(logger.LoggerMiddleware())
	e.Use(middleware.Recover())
	NewUserRoute(e)
	return e
}
