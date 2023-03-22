package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Top(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World!")
}

func Login(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
