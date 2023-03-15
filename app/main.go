package main

import (
	"fmt"

	"github.com/api-sample/app/consts"
	"github.com/api-sample/app/router"
)

func main() {
	e := router.NewRouter()
	fmt.Println(consts.APIPort)
	e.Logger.Fatal(e.Start(":" + consts.APIPort))
}
