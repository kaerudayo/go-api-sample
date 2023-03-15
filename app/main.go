package main

import (
	"github.com/api-sample/app/consts"
	"github.com/api-sample/app/router"
)

func main() {
	e := router.NewRouter()

	e.Logger.Fatal(e.Start(":" + consts.APIPort))
}
