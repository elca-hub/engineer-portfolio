package main

import (
	"devport/infra"
	"devport/infra/database"
	"devport/infra/router"
	"os"
	"time"
)

func main() {
	app := infra.NewHttpServerConfig().
		Name(os.Getenv("APP_NAME")).
		ContextTimeout(10 * time.Second).
		DbSql(database.InstanceMySQL).
		DbNoSql(database.InstanceRedis).
		WebServerPort(os.Getenv("APP_PORT")).
		WebServer(router.InstanceGin)

	app.Start()
}
