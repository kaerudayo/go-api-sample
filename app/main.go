package main

import (
	"github.com/api-sample/app/cmd/router"
	"github.com/api-sample/app/infra"
	"github.com/api-sample/app/pkg/consts"
	"github.com/api-sample/app/pkg/logger"
)

func main() {
	// loggerの初期化
	zapLogger := logger.Init()
	defer func() {
		if err := zapLogger.Sync(); err != nil {
			panic(err)
		}
	}()
	// routerの初期化
	e := router.NewRouter()
	// dbの初期化
	infra.Init(true)
	// serverの起動
	e.Logger.Fatal(e.Start(":" + consts.APIPort))
}
