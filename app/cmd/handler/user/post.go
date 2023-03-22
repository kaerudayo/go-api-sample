package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Top(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World!")
}

func SignUp(c echo.Context) error {
	return c.JSON(http.StatusOK, "TODO")
}
