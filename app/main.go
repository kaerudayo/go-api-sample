package main

import (
	"github.com/api-sample/app/pkg/consts"
	"github.com/api-sample/app/pkg/db"
	"github.com/api-sample/app/pkg/logger"
	"github.com/api-sample/app/router"
)

func main() {
	// loggerの初期化
	zapLogger := logger.Init()
	defer func() {
		if err := zapLogger.Sync(); err != nil {
			panic(err)
		}
	}()
	e := router.NewRouter()
	db.Init(true)
	e.Logger.Fatal(e.Start(":" + consts.APIPort))
}
