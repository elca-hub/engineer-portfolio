package main

import (
	"devport/infra"
	"devport/infra/database"
	"devport/infra/file_uploader"
	"devport/infra/log"
	"devport/infra/router"
	"devport/infra/validation"
	"os"
	"time"
)

func main() {
	app := infra.NewHttpServerConfig().
		Name(os.Getenv("APP_NAME")).
		ContextTimeout(10 * time.Second).
		DbSql(database.InstanceMySQL).
		DbNoSql(database.InstanceRedis).
		Logger(log.InstanceZap).
		Validator(validation.InstanceGoPlayground).
		Email().
		FileUploader(file_uploader.InstanceMinio).
		WebServerPort(os.Getenv("APP_PORT")).
		WebServer(router.InstanceGin)

	app.Start()
}
