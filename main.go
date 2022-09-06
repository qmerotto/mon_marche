package main

import (
	"gorm.io/gorm/logger"
	"test_tech/common"
	"test_tech/common/database"
	"test_tech/web_api"
	"test_tech/workers"
)

func init() {
	common.InitializeConfiguration()

	database.Connect(
		common.GetSettings().DatabaseUser,
		common.GetSettings().DatabasePassword,
		common.GetSettings().DatabaseHost,
		common.GetSettings().DatabasePort,
		common.GetSettings().DatabaseName,
		common.GetSettings().DatabasePoolSize,
		logger.Info,
	)
	database.RunMigrations()

	workers.Init()
}

func main() {
	go workers.Dispatch()
	web_api.RunServer()
}
