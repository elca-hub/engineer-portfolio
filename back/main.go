package main

import (
	"devport/infra"
	cronpackage "devport/infra/cron"
	"devport/infra/database"
	"devport/infra/file_uploader"
	"devport/infra/log"
	"devport/infra/router"
	"devport/infra/validation"
	"os"
	"time"
)

func main() {
	batch := infra.NewBatchServerConfig().
		CronText("@every 10m").
		LoggingTool(log.InstanceZap).
		DB(database.InstanceSqlBoilerMySql).
		AppName(os.Getenv("APP_NAME")).
		CtxTimeout(10 * time.Second).
		FileUploader(file_uploader.InstanceMinio).
		Cron(cronpackage.InstanceCron)

	batch.Start()

	app := infra.NewHttpServerConfig().
		Name(os.Getenv("APP_NAME")).
		ContextTimeout(10 * time.Second).
		DbSql(database.InstanceSqlBoilerMySql).
		DbNoSql(database.InstanceRedis).
		Logger(log.InstanceZap).
		Validator(validation.InstanceGoPlayground).
		Email().
		FileUploader(file_uploader.InstanceMinio).
		WebServerPort(os.Getenv("APP_PORT")).
		WebServer(router.InstanceGin)

	app.Start()
}
