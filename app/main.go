package main

import (
	"fmt"
	"net/http"

	"github.com/api-sample/app/consts"
	"github.com/labstack/echo/v4"
)

func main() {
	// インスタンスを作成
	e := echo.New()

	e.GET("/", hello)
	fmt.Println(consts.APIPort)

	e.Logger.Fatal(e.Start(":" + consts.APIPort))
}

// ハンドラーを定義
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
